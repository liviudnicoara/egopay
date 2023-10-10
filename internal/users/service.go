package users

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/liviudnicoara/egopay/internal/accounts"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type UserService interface {
	All() ([]User, error)
	Register(userName string, email string, password string) error
	Login(userName string, password string) (*User, error)
	GetByID(userID uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	AddAccount(userID uuid.UUID, password string) (string, error)
}

type userService struct {
	userRepository UserRepository
	accountService accounts.AccountService
}

func NewUserService(userRepository UserRepository, accountService accounts.AccountService) UserService {
	return &userService{
		userRepository: userRepository,
		accountService: accountService,
	}
}

func (s *userService) Register(userName string, email string, password string) error {
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

	err = s.userRepository.Save(user)

	if err != nil {
		return errors.WithMessagef(err, "could not save user: %s", userName)
	}

	return nil
}

func (s *userService) Login(userName string, password string) (*User, error) {
	return nil, nil
}

func (s *userService) All() ([]User, error) {
	u, err := s.userRepository.GetAll()

	return u, err
}

func (s *userService) GetByID(userID uuid.UUID) (*User, error) {
	u, err := s.userRepository.Get(userID)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *userService) GetByEmail(email string) (*User, error) {
	users, err := s.userRepository.GetAll()
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

func (s *userService) AddAccount(userID uuid.UUID, password string) (string, error) {
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
		u, err := s.userRepository.Get(userID)
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
	err := s.userRepository.Save(user)
	if err != nil {
		return "", errors.WithMessagef(err, "could not save account for user: %s", userID)
	}

	return accountAddress, nil
}
