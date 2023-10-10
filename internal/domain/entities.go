package domain

import (
	"fmt"
	"math/big"
)

var multiplier = big.NewFloat(100)

type Fiat int64

func NewFiatFromFloat(f float64) Fiat {
	return Fiat((f * 100) + 0.5)
}

func NewFiatFromBigFloat(f big.Float) Fiat {
	v, _ := new(big.Float).Mul(&f, multiplier).Float64()

	return Fiat(v + 0.5)
}

func ToFiat(f float64) Fiat {
	return NewFiatFromFloat(f)
}

func (f Fiat) ToFloat64() float64 {
	x := float64(f)
	x = x / 100
	return x
}

func (f Fiat) ToInt64() int64 {
	return int64(f)
}

func (f Fiat) ToBigFloat() *big.Float {
	return big.NewFloat(f.ToFloat64())
}

type USD struct {
	Fiat
}

func NewUSDFromFloat(f float64) USD {
	return USD{NewFiatFromFloat(f)}
}

func NewUSDFromBigFloat(f big.Float) USD {
	return USD{NewFiatFromBigFloat(f)}
}

func (m USD) String() string {
	return fmt.Sprintf("$ %.2f", m.ToFloat64())
}
