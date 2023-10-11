package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/internal/domain"
	"github.com/liviudnicoara/egopay/pkg/errors"
)

// MakeTransfer godoc
// @Summary Creates a new transfer
// @Description Creates a new transfer
// @ID create-transfer
// @Tags transfers
// @Accept json
// @Produce json
// @Param transfer body CreateTransferRequest true "Info for creating transfer"
// @Success 201 {object} CreateTransferResponse
// @Failure 400 {object} errors.Error
// @Failure 404 {objects} errors.Error
// @Failure 500 {objects} errors.Error
// @Security ApiKeyAuth
// @Router /api/transfers/make [post]
func (h *Handler) MakeTransfer(c *fiber.Ctx) error {
	req := &CreateTransferRequest{}
	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	err := h.transferService.Transfer(c.Context(), req.FromAddress, req.ToAddress, domain.NewUSDFromFloat(req.Amount), req.Password)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newCreateTransferResponse(err == nil))
}
