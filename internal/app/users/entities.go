package users

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               uuid.UUID
	Username         string
	Email            string
	Image            *string
	Password         string
	AccountAddresses []string
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) EncodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func (u *User) DecodePassword(password string) []byte {
	dec, _ := base64.StdEncoding.DecodeString(password)
	return dec
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword(u.DecodePassword(u.Password), []byte(plain))
	return err == nil
}

func (u *User) HasAccount(address string) bool {
	for _, a := range u.AccountAddresses {
		if strings.EqualFold(address, a) {
			return true
		}
	}

	return false
}
