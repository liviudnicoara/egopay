package users

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/liviudnicoara/egopay/internal/app/accounts"
	"github.com/liviudnicoara/egopay/internal/app/price"
	"github.com/liviudnicoara/egopay/internal/domain"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type UserService interface {
	Register(ctx context.Context, userName string, email string, password string) error
	Login(ctx context.Context, userName string, password string) (*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	AddAccount(ctx context.Context, userID uuid.UUID, password string) (string, error)
	GetBalance(ctx context.Context, userID uuid.UUID, address string) (balanceFiat string, balanceETH string, err error)
}

type userService struct {
	userRepository UserRepository
	accountService accounts.AccountService
	priceService   price.PriceService
}

func NewUserService(userRepository UserRepository, accountService accounts.AccountService, priceService price.PriceService) UserService {
	return &userService{
		userRepository: userRepository,
		accountService: accountService,
		priceService:   priceService,
	}
}

func (s *userService) Register(ctx context.Context, userName string, email string, password string) error {
	user := User{
		ID:       uuid.New(),
		Username: strings.ToLower(userName),
		Email:    email,
	}

	hashedPassword, err := user.HashPassword(password)
	if err != nil {
		return errors.WithMessagef(err, "could not hash password: %s", password)
	}

	user.Password = user.EncodePassword(hashedPassword)

	err = s.userRepository.Save(ctx, user)

	if err != nil {
		return errors.WithMessagef(err, "could not save user: %s", userName)
	}

	return nil
}

func (s *userService) Login(ctx context.Context, userName string, password string) (*User, error) {
	return nil, nil
}

func (s *userService) GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	u, err := s.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*User, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if strings.EqualFold(u.Email, email) {
			return &u, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("user with email %s was not found", email))
}

func (s *userService) AddAccount(ctx context.Context, userID uuid.UUID, password string) (string, error) {
	eg := errgroup.Group{}
	var user User
	var accountAddress string

	eg.Go(func() error {
		ad, err := s.accountService.CreateAccount(password)
		accountAddress = ad

		if err != nil {
			return errors.WithMessagef(err, "could not create account for user: %s", userID)
		}

		return nil
	})

	eg.Go(func() error {
		u, err := s.userRepository.Get(ctx, userID)
		user = u

		if err != nil {
			return errors.WithMessagef(err, "could not get user: %s", userID)
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return "", nil
	}

	user.AccountAddresses = append(user.AccountAddresses, accountAddress)
	err := s.userRepository.Save(ctx, user)
	if err != nil {
		return "", errors.WithMessagef(err, "could not save account for user: %s", userID)
	}

	return accountAddress, nil
}

func (s *userService) GetBalance(ctx context.Context, userID uuid.UUID, address string) (balanceFiat string, balanceETH string, err error) {
	user, err := s.userRepository.Get(ctx, userID)

	if err != nil {
		return "", "", errors.WithMessagef(err, "could not get user: %s", userID)
	}

	if !user.HasAccount(address) {
		return "", "", errors.New(fmt.Sprintf("account %s not found for userID %s", address, userID.String()))
	}

	balance, err := s.accountService.GetBalance(ctx, address)

	if err != nil {
		return "", "", errors.WithMessagef(err, "could not get balance for account %s for user: %s", address, userID.String())
	}

	ethBalance := domain.ETH(balance)
	price := s.priceService.GetPrice()
	fiatBalance := domain.USD{Fiat: domain.NewFiatFromFloat(ethBalance.ToFloat64() * price.ToFloat64())}

	return fiatBalance.String(), ethBalance.String(), nil
}
