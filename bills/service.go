package bills

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/liviudnicoara/egopay/accounts"
	"github.com/liviudnicoara/egopay/contracts"
	"github.com/pkg/errors"
)

type BillService interface {
	Split(ctx context.Context, from string, amount int, password string) (deployedContractAddress common.Address, txHashHex string, err error)
}

type billService struct {
	accountService accounts.AccountService
	client         *ethclient.Client
	chainID        *big.Int
}

func NewBillService(accountService accounts.AccountService, client *ethclient.Client) BillService {
	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		panic("could not retrive chain id of network")
	}

	return &billService{
		accountService: accountService,
		client:         client,
		chainID:        chainID,
	}
}

func (s *billService) Split(ctx context.Context, from string, amount int, password string) (deployedContractAddress common.Address, txHashHex string, err error) {
	account, err := s.accountService.GetAccount(from, password)
	if err != nil {
		return common.Address{}, "", errors.WithMessagef(err, "could not get account info for %s", from)
	}

	nounce, err := s.client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(account.PrivateKey.PublicKey))
	if err != nil {
		return common.Address{}, "", errors.WithMessagef(err, "could not get pending nounce for %s", from)
	}

	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.Address{}, "", errors.WithMessagef(err, "could not get suggested price for %s", from)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, s.chainID)
	auth.Nonce = big.NewInt(int64(nounce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)

	payers := []common.Address{crypto.PubkeyToAddress(account.PrivateKey.PublicKey)}
	contractAmount := big.NewInt(int64(amount))
	address, tx, _, err := contracts.DeploySplitBill(auth, s.client, payers, contractAmount)
	if err != nil {
		return common.Address{}, "", errors.WithMessagef(err, "could not deploy contract for %s", from)
	}

	return address, tx.Hash().Hex(), nil
}
