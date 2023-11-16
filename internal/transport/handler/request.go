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

type CreateBillRequest struct {
	Address  string   `json:"address" validate:"required"`
	Payers   []string `json:"payers" validate:"required"`
	Amount   float64  `json:"amount" validate:"required"`
	Password string   `json:"password" validate:"required"`
}

func (r *CreateBillRequest) bind(c *fiber.Ctx, v *Validator) error {

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}

type CreateTransferRequest struct {
	FromAddress string  `json:"fromAddress" validate:"required"`
	ToAddress   string  `json:"toAddress" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	Password    string  `json:"password" validate:"required"`
}

func (r *CreateTransferRequest) bind(c *fiber.Ctx, v *Validator) error {

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}
