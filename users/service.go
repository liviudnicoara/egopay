package users

import (
	"strings"

	"github.com/google/uuid"
	"github.com/liviudnicoara/egopay/accounts"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

type UserService interface {
	Register(userName string, password string) (User, error)
	Login(userName string, password string) (User, error)
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

func (s *userService) Register(userName string, password string) (User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return User{}, errors.WithMessagef(err, "could not hash password: %s", password)
	}

	user := User{
		ID:       uuid.New(),
		UserName: strings.ToLower(userName),
		Password: hashedPassword,
	}

	err = s.userRepository.Save(user)

	if err != nil {
		return User{}, errors.WithMessagef(err, "could not save user: %s", userName)
	}

	return user, nil
}

func (s *userService) Login(userName string, password string) (User, error) {
	return User{}, nil
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

func hashPassword(password string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	hashedPassword := string(salt)
	return hashedPassword, nil
}
