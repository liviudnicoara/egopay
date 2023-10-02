package convert

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/params"
)

func ConvertWEItoETH(amount *big.Int) big.Float {
	// 1 ETH = 1^18 WEI
	weiAmount := new(big.Float)
	weiAmount.SetString(amount.String())

	ethAmount := new(big.Float).Quo(weiAmount, big.NewFloat(math.Pow10(18)))

	return *ethAmount
}

func ConvertETHtoWEI(amount *big.Float) *big.Int {
	// 1 ETH = 1^18 WEI
	truncInt, _ := amount.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))

	fracStr := strings.Split(fmt.Sprintf("%.18f", amount), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)

	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}
