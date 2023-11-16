package domain

import (
	"fmt"
	"math/big"
)

type ETH big.Float

func NewETH(value big.Float) ETH {
	return ETH(value)
}

func (eth ETH) ToFloat64() float64 {
	f, _ := (*big.Float)(&eth).Float64()
	return f
}

func (eth ETH) String() string {
	return fmt.Sprintf("%.4f ETH", eth.ToFloat64())
}

var multiplier = big.NewFloat(100)

type USD int64

func NewUSDFromFloat(f float64) USD {
	return USD((f * 100) + 0.5)
}

func NewUSDFromBigFloat(f big.Float) USD {
	v, _ := new(big.Float).Mul(&f, multiplier).Float64()

	return USD(v + 0.5)
}

func ToUSD(f float64) USD {
	return NewUSDFromFloat(f)
}

func (m USD) ToFloat64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

func (m USD) ToInt64() int64 {
	return int64(m)
}

func (m USD) ToBigFloat() *big.Float {
	return big.NewFloat(m.ToFloat64())
}

func (m USD) String() string {
	return fmt.Sprintf("$ %.2f", m.ToFloat64())
}
