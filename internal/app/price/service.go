package price

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/liviudnicoara/egopay/internal/app/contracts"
	"github.com/liviudnicoara/egopay/internal/domain"
)

type PriceService interface {
	Start()
	GetPrice() domain.USD
	Stop()
}

type priceService struct {
	client         *ethclient.Client
	mu             sync.RWMutex
	priceFeedProxy *contracts.AggregatorV3Interface
	fetchInterval  time.Duration
	divisor        *big.Int
	currentPrice   domain.USD
	done           chan struct{}
}

func NewPriceService(client *ethclient.Client, feedContractAddress string, fetchInterval time.Duration) PriceService {
	ok := isContractAddress(feedContractAddress, client)
	if !ok {
		log.Fatalf("address %s is not a contract address\n", feedContractAddress)
	}

	priceFeedProxyAddress := common.HexToAddress(feedContractAddress)
	priceFeedProxy, err := contracts.NewAggregatorV3Interface(priceFeedProxyAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := priceFeedProxy.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Compute a big.int which is 10**decimals.
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	return &priceService{
		client:         client,
		priceFeedProxy: priceFeedProxy,
		fetchInterval:  fetchInterval,
		divisor:        divisor,
	}
}

func (s *priceService) Start() {
	s.done = make(chan struct{})
	s.updateCurrentPrice()

	go func(done <-chan struct{}, interval time.Duration) {
		t := time.NewTicker(interval)
		for {
			select {
			case <-t.C:
				s.updateCurrentPrice()

			case <-done:
				return
			}
		}
	}(s.done, s.fetchInterval)
}

func (s *priceService) GetPrice() domain.USD {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.currentPrice
}

func (s *priceService) Stop() {
	close(s.done)
}

func (s *priceService) updateCurrentPrice() {
	roundData, err := s.priceFeedProxy.LatestRoundData(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}

	floatRoundData := divideBigInt(roundData.Answer, s.divisor)

	s.mu.Lock()
	s.currentPrice = domain.NewUSDFromBigFloat(*floatRoundData)
	s.mu.Unlock()
}

func isContractAddress(addr string, client *ethclient.Client) bool {
	if len(addr) == 0 {
		log.Fatal("feedAddress is empty.")
	}

	// Ensure it is an Ethereum address: 0x followed by 40 hexadecimal characters.
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(addr) {
		log.Fatalf("address %s non valid\n", addr)
	}

	// Ensure it is a contract address.
	address := common.HexToAddress(addr)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	return isContract
}

func divideBigInt(num1 *big.Int, num2 *big.Int) *big.Float {
	if num2.BitLen() == 0 {
		log.Fatal("cannot divide by zero.")
	}
	num1BigFloat := new(big.Float).SetInt(num1)
	num2BigFloat := new(big.Float).SetInt(num2)
	result := new(big.Float).Quo(num1BigFloat, num2BigFloat)
	return result
}
