package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/internal/app/users"
)

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json: "email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

func (r *UserRegisterRequest) bind(c *fiber.Ctx, u *users.User, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	u.Username = r.Username
	u.Email = r.Email
	u.Password = r.Password

	return nil
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *UserLoginRequest) bind(c *fiber.Ctx, v *Validator) error {

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}

type CreateAccountRequest struct {
	Password string `json:"password" validate:"required"`
}

func (r *CreateAccountRequest) bind(c *fiber.Ctx, v *Validator) error {

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}
