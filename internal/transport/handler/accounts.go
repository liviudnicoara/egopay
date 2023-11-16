package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/pkg/errors"
)

// GetAccountBalance godoc
// @Summary Get the account balance
// @Description Get the account balance
// @ID get-account-balance
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param address path string true "address of the account balance to get"
// @Success 200 {object} AccountBalanceResponse
// @Failure 400 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Security ApiKeyAuth
// @Router /api/accounts/{address}/balance [get]
func (h *Handler) GetAccountBalance(c *fiber.Ctx) error {
	address := c.Params("address")
	balanceFiat, balanceETH, err := h.userService.GetBalance(c.Context(), userIDFromToken(c), address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newAccountBalanceResponse(address, balanceFiat, balanceETH))
}

// CreateAccount godoc
// @Summary Creates a new account
// @Description Creates a new account
// @ID create-account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body CreateAccountRequest true "Info for creating account"
// @Success 201 {object} CreateAccountResponse
// @Failure 400 {object} errors.Error
// @Failure 404 {objects} errors.Error
// @Failure 500 {objects} errors.Error
// @Security ApiKeyAuth
// @Router /api/accounts/create [post]
func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	req := &CreateAccountRequest{}
	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	address, err := h.userService.AddAccount(c.Context(), userIDFromToken(c), req.Password)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newCreateAccountResponse(address))
}
