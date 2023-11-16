package bills

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/liviudnicoara/egopay/internal/app/accounts"
	"github.com/liviudnicoara/egopay/internal/app/contracts"
	"github.com/liviudnicoara/egopay/internal/domain"

	"github.com/pkg/errors"
)

type BillService interface {
	Split(ctx context.Context, from string, payers []string, amount domain.USD, password string) (deployedContractAddress string, txHashHex string, err error)
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

func (s *billService) Split(ctx context.Context, owner string, payers []string, amount domain.USD, password string) (deployedContractAddress string, txHashHex string, err error) {
	account, err := s.accountService.GetAccount(owner, password)
	if err != nil {
		return "", "", errors.WithMessagef(err, "could not get account info for %s", owner)
	}

	nounce, err := s.client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(account.PrivateKey.PublicKey))
	if err != nil {
		return "", "", errors.WithMessagef(err, "could not get pending nounce for %s", owner)
	}

	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", "", errors.WithMessagef(err, "could not get suggested price for %s", owner)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, s.chainID)
	auth.Nonce = big.NewInt(int64(nounce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)

	payerAddreses := make([]common.Address, len(payers))
	for i, p := range payers {
		payerAddreses[i] = common.HexToAddress(p)
	}

	contractAmount := big.NewInt(amount.ToInt64())
	address, tx, _, err := contracts.DeploySplitBill(auth, s.client, payerAddreses, contractAmount)
	if err != nil {
		return "", "", errors.WithMessagef(err, "could not deploy contract for %s", owner)
	}

	return address.String(), tx.Hash().Hex(), nil
}
