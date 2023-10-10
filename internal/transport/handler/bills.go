package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/internal/domain"
	"github.com/liviudnicoara/egopay/pkg/errors"
)

// // GetBills godoc
// // @Summary Get previous bills
// // @Description Get previous bills
// // @ID get-bills
// // @Tags bills
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} BillsResponse
// // @Failure 400 {object} utils.Error
// // @Failure 500 {object} utils.Error
// // @Router /api/accounts/{address}/balance [get]
// func (h *Handler) GetBills(c *fiber.Ctx) error {
// 	address := c.Params("address")
// 	balance, err := h.userService.GetBalance(c.Context(), userIDFromToken(c), address)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(errors.NewError(err))
// 	}

// 	return c.Status(http.StatusOK).JSON(newAccountBalanceResponse(address, balance))
// }

// CreateBill godoc
// @Summary Creates a new bill
// @Description Creates a new bill
// @ID create-bill
// @Tags bills
// @Accept json
// @Produce json
// @Param bill body CreateBillRequest true "Info for creating bill"
// @Success 201 {object} CreateBillResponse
// @Failure 400 {object} errors.Error
// @Failure 404 {objects} errors.Error
// @Failure 500 {objects} errors.Error
// @Router /api/bills/create [post]
func (h *Handler) CreateBill(c *fiber.Ctx) error {
	req := &CreateBillRequest{}
	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	billAddres, tx, err := h.billService.Split(c.Context(), req.Address, req.Payers, domain.NewUSDFromFloat(req.Amount), req.Password)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newCreateBillResponse(billAddres, tx, err == nil))
}
