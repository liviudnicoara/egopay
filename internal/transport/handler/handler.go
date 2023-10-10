package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/internal/users"
	"github.com/liviudnicoara/egopay/pkg/jwt"
)

type Handler struct {
	userService users.UserService
	validator   *Validator

	jwtMiddleware func(*fiber.Ctx) error
}

func NewHandler(us users.UserService) *Handler {
	v := NewValidator()

	jwtMW := jwtware.New(
		jwtware.Config{
			SigningKey: jwt.JWTSecret,
			AuthScheme: "Token",
		})

	return &Handler{
		userService:   us,
		validator:     v,
		jwtMiddleware: jwtMW,
	}
}

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")

	h.registerUser(&v1)
}

func (h *Handler) registerUser(v *fiber.Router) {
	users := (*v).Group("/users")
	users.Get("/all", h.All)
	users.Post("/signup", h.SignUp)
	users.Post("/login", h.Login)
	users.Get("", h.jwtMiddleware, h.CurrentUser)
}
