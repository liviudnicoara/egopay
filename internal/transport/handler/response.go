package handler

import (
	_ "github.com/gofiber/fiber/v2"
	"github.com/liviudnicoara/egopay/internal/app/users"
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

type AccountBalanceResponse struct {
	Address      string `json:"address"`
	BalanceFiat  string `json:"balanceFiat"`
	BalanceToken string `json:"balanceToken"`
}

func newAccountBalanceResponse(address string, balanceFiat string, balanceToken string) *AccountBalanceResponse {
	r := new(AccountBalanceResponse)
	r.Address = address
	r.BalanceFiat = balanceFiat
	r.BalanceToken = balanceToken
	return r
}

type CreateAccountResponse struct {
	Address string `json:"address"`
}

func newCreateAccountResponse(address string) *CreateAccountResponse {
	r := new(CreateAccountResponse)
	r.Address = address
	return r
}

type CreateBillResponse struct {
	BillAddress     string `json:"billAddress"`
	TransactionHash string `json:"transactionHash"`
	Success         bool   `json:"success"`
}

func newCreateBillResponse(billAddress, tx string, isSuccess bool) *CreateBillResponse {
	r := new(CreateBillResponse)
	r.BillAddress = billAddress
	r.TransactionHash = tx
	r.Success = isSuccess
	return r
}

type CreateTransferResponse struct {
	Success bool `json:"success"`
}

func newCreateTransferResponse(isSuccess bool) *CreateBillResponse {
	r := new(CreateBillResponse)
	r.Success = isSuccess
	return r
}
