package users

import (
	"encoding/json"
	"fmt"

	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/liviudnicoara/egopay/pkg/filesystem"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Save(User) error
	Get(ID uuid.UUID) (User, error)
	GetAll() ([]User, error)
}

type userRepository struct {
	storageFilePath string
	m               sync.RWMutex
}

func NewUserRepository(storagePath string) UserRepository {
	_ = filesystem.CreateDirectoryIfNotExists(storagePath)
	storageFilePath := filepath.Join(storagePath, "users.json")

	if _, err := os.Stat(storageFilePath); err != nil {
		users := map[uuid.UUID]User{}
		usersFile, _ := json.Marshal(users)
		_ = os.WriteFile(storageFilePath, usersFile, 0644)
	}

	return &userRepository{
		storageFilePath: storageFilePath,
	}
}

func (r *userRepository) Save(user User) error {
	r.m.Lock()
	defer r.m.Unlock()
	usersFile, err := os.ReadFile(r.storageFilePath)
	if err != nil {
		return errors.WithMessage(err, "could not access users")
	}

	var users map[uuid.UUID]User
	err = json.Unmarshal(usersFile, &users)
	if err != nil {
		return errors.WithMessage(err, "could not read users")
	}

	users[user.ID] = user
	usersFile, err = json.Marshal(users)
	if err != nil {
		return errors.WithMessage(err, "could not save user: %s")
	}

	err = os.WriteFile(r.storageFilePath, usersFile, 0644)
	if err != nil {
		return errors.WithMessage(err, "could not save user: %s")
	}

	return nil
}

func (r *userRepository) Get(ID uuid.UUID) (User, error) {
	r.m.RLock()
	defer r.m.RLocker().Unlock()

	usersFile, err := os.ReadFile(r.storageFilePath)
	if err != nil {
		return User{}, errors.WithMessage(err, "could not access users")
	}

	var users map[uuid.UUID]User
	err = json.Unmarshal(usersFile, &users)
	if err != nil {
		return User{}, errors.WithMessage(err, "could not read users")
	}

	user, found := users[ID]
	if !found {
		return User{}, errors.New(fmt.Sprintf("did not found user: %s", ID.String()))
	}

	return user, nil
}

func (r *userRepository) GetAll() ([]User, error) {
	r.m.RLock()
	defer r.m.RLocker().Unlock()

	usersFile, err := os.ReadFile(r.storageFilePath)
	if err != nil {
		return nil, errors.WithMessage(err, "could not access users")
	}

	var users map[uuid.UUID]User
	err = json.Unmarshal(usersFile, &users)
	if err != nil {
		return nil, errors.WithMessage(err, "could not read users")
	}

	result := make([]User, len(users))
	i := 0
	for _, u := range users {
		result[i] = u
		i++
	}

	return result, nil
}
