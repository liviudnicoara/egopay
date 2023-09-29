package accounts

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type AccountService interface {
	CreateAccount(password string) (string, error)
	GetAccount(address string, password string) (Account, error)
	GetBalance(ctx context.Context, address string) (big.Float, error)
}

type accountService struct {
	accountRepository AccountRepository
	ethClient         *ethclient.Client
}

func NewAccountService(accountRepository AccountRepository, ethClient *ethclient.Client) AccountService {
	return &accountService{
		accountRepository: accountRepository,
		ethClient:         ethClient,
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
	balance, err := s.ethClient.BalanceAt(ctx, addr, nil)

	if err != nil {
		return big.Float{}, errors.WithMessagef(err, "could not retirve balance for address %s", address)
	}

	return convertWEItoETH(balance), nil
}

func convertWEItoETH(amount *big.Int) big.Float {
	// 1 ETH = 1^18 WEI
	weiAmount := new(big.Float)
	weiAmount.SetString(amount.String())
	ethAmount := new(big.Float).Quo(weiAmount, big.NewFloat(math.Pow10(18)))
	return *ethAmount
}
