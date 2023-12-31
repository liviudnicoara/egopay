package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/liviudnicoara/egopay/internal/app/users"
	"github.com/liviudnicoara/egopay/pkg/errors"
)

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user
// @ID sign-up
// @Tags user
// @Accept json
// @Produce json
// @Param user body UserRegisterRequest true "User info for registration"
// @Success 201 {object} UserResponse
// @Failure 400 {object} errors.Error
// @Failure 404 {objects} errors.Error
// @Failure 500 {objects} errors.Error
// @Router /api/user/signup [post]
func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u users.User
	req := &UserRegisterRequest{}
	if err := req.bind(c, &u, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}
	if err := h.userService.Register(c.Context(), u.Username, u.Email, u.Password); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newUserResponse(&u))
}

// Login godoc
// @Summary Login for existing user
// @Description Login for existing user
// @ID login
// @Tags user
// @Accept json
// @Produce json
// @Param user body UserLoginRequest true "Credentials to use"
// @Success 200 {object} UserResponse
// @Failure 400 {object} errors.Error
// @Failure 401 {object} errors.Error
// @Failure 422 {object} errors.Error
// @Failure 404 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Router /api/user/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	req := &UserLoginRequest{}
	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(errors.NewError(err))
	}

	u, err := h.userService.GetByEmail(c.Context(), req.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.NewError(err))
	}

	if u == nil {
		return c.Status(http.StatusForbidden).JSON(errors.AccessForbidden())
	}

	if !u.CheckPassword(req.Password) {
		return c.Status(http.StatusForbidden).JSON(errors.AccessForbidden())
	}

	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Gets the currently logged-in user
// @ID current-user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse
// @Failure 400 {object} errors.Error
// @Failure 401 {object} errors.Error
// @Failure 422 {object} errors.Error
// @Failure 404 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Security ApiKeyAuth
// @Router /api/user [get]
func (h *Handler) CurrentUser(c *fiber.Ctx) error {
	u, err := h.userService.GetByID(c.Context(), userIDFromToken(c))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.NewError(err))
	}
	if u == nil {
		return c.Status(http.StatusNotFound).JSON(errors.NotFound())
	}
	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

func userIDFromToken(c *fiber.Ctx) uuid.UUID {
	var user *jwt.Token
	l := c.Locals("user")
	if l == nil {
		return uuid.Nil
	}
	user = l.(*jwt.Token)
	id := uuid.MustParse(((user.Claims.(jwt.MapClaims)["id"]).(string)))
	return id
}
