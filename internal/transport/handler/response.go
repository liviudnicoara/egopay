package handler

import (
	_ "github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/internal/users"
	"github.com/liviudnicoara/egopay/pkg/jwt"
)

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func newUserResponse(u *users.User) *UserResponse {
	r := new(UserResponse)
	r.Username = u.Username
	r.Email = u.Email
	r.Image = u.Image
	r.Token = jwt.GenerateJWT(u.ID)
	return r
}
