package accounts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/liviudnicoara/egopay/pkg/convert"
	"github.com/pkg/errors"
)

type AccountService interface {
	CreateAccount(password string) (string, error)
	GetAccount(address string, password string) (Account, error)
	GetBalance(ctx context.Context, address string) (big.Float, error)
}

type accountService struct {
	accountRepository AccountRepository
	client            *ethclient.Client
}

func NewAccountService(accountRepository AccountRepository, client *ethclient.Client) AccountService {
	return &accountService{
		accountRepository: accountRepository,
		client:            client,
	}
}

func (s *accountService) CreateAccount(password string) (string, error) {
	return s.accountRepository.CreateAccount(password)
}

func (s *accountService) GetAccount(address string, password string) (Account, error) {
	return s.accountRepository.GetAccount(address, password)
}

func (s *accountService) GetBalance(ctx context.Context, address string) (big.Float, error) {
	addr := common.HexToAddress(address)
	balance, err := s.client.BalanceAt(ctx, addr, nil)

	if err != nil {
		return big.Float{}, errors.WithMessagef(err, "could not retirve balance for address %s", address)
	}

	return convert.ConvertWEItoETH(balance), nil
}
