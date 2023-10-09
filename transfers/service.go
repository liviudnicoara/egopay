package transfers

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/liviudnicoara/egopay/accounts"
	"github.com/liviudnicoara/egopay/pkg/convert"
	"github.com/pkg/errors"
)

const GAS_LIMIT = 21000

type TransferService interface {
	Transfer(ctx context.Context, from string, to string, amount *big.Float, password string) error
}

type transferService struct {
	accountService accounts.AccountService
	client         *ethclient.Client
	chainID        *big.Int
}

func NewTransferService(accountService accounts.AccountService, client *ethclient.Client) TransferService {
	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		panic("could not retrive chain id of network")
	}

	return &transferService{
		accountService: accountService,
		client:         client,
		chainID:        chainID,
	}
}

func (s *transferService) Transfer(ctx context.Context, from string, to string, amount *big.Float, password string) error {
	fromAddr := common.HexToAddress(from)
	toAddr := common.HexToAddress(to)

	account, err := s.accountService.GetAccount(from, password)
	if err != nil {
		return errors.WithMessagef(err, "could not get account info for %s", from)
	}

	nonce, err := s.client.NonceAt(ctx, fromAddr, nil)
	if err != nil {
		return errors.WithMessagef(err, "could not get nounce for address %s", from)
	}

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return errors.WithMessage(err, "could not get gas price")
	}

	tx := types.NewTransaction(nonce, toAddr, convert.ConvertETHtoWEI(amount), GAS_LIMIT, gasPrice, nil)

	tx, err = types.SignTx(tx, types.NewEIP155Signer(s.chainID), account.PrivateKey)
	if err != nil {
		return errors.WithMessagef(err, "could not sign transaction from address %s to address %s", from, to)
	}

	err = s.client.SendTransaction(ctx, tx)
	if err != nil {
		return errors.WithMessagef(err, "could not send transaction from address %s to address %s", from, to)
	}

	return nil
}
