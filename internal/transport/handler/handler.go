package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/internal/app/bills"
	"github.com/liviudnicoara/egopay/internal/app/transfers"
	"github.com/liviudnicoara/egopay/internal/app/users"
	"github.com/liviudnicoara/egopay/pkg/jwt"
)

type Handler struct {
	userService     users.UserService
	billService     bills.BillService
	transferService transfers.TransferService
	validator       *Validator

	jwtMiddleware func(*fiber.Ctx) error
}

func NewHandler(us users.UserService, bs bills.BillService, ts transfers.TransferService) *Handler {
	v := NewValidator()

	jwtMW := jwtware.New(
		jwtware.Config{
			SigningKey: jwt.JWTSecret,
			AuthScheme: "Token",
		})

	return &Handler{
		userService:     us,
		billService:     bs,
		transferService: ts,
		validator:       v,
		jwtMiddleware:   jwtMW,
	}
}

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")

	h.registerUsers(&v1)
	h.registerAccounts(&v1)
	h.registerBills(&v1)
	h.registerTransfers(&v1)
}

func (h *Handler) registerUsers(v *fiber.Router) {
	users := (*v).Group("/user")
	users.Post("/signup", h.SignUp)
	users.Post("/login", h.Login)
	users.Get("", h.jwtMiddleware, h.CurrentUser)
}

func (h *Handler) registerAccounts(v *fiber.Router) {
	accounts := (*v).Group("/accounts", h.jwtMiddleware)
	accounts.Get("/:address/balance", h.GetAccountBalance)
	accounts.Post("/create", h.CreateAccount)
}

func (h *Handler) registerBills(v *fiber.Router) {
	bills := (*v).Group("/bills", h.jwtMiddleware)
	// users.Get("", h.GetBills)
	bills.Post("/create", h.CreateBill)
}

func (h *Handler) registerTransfers(v *fiber.Router) {
	transfers := (*v).Group("/transfers", h.jwtMiddleware)
	transfers.Post("/make", h.MakeTransfer)
}
