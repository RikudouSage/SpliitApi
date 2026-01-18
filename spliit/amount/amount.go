package amount

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

type Amount int64

func FromDecimal(val decimal.Decimal) (Amount, error) {
	if val.Exponent() != -2 {
		return 0, fmt.Errorf("amount must have exactly 2 decimal places, got exponent %d", val.Exponent())
	}

	return Amount(val.CoefficientInt64()), nil
}

func FromString(val string) (Amount, error) {
	dec, err := decimal.NewFromString(val)
	if err != nil {
		return 0, err
	}

	return FromDecimal(dec)
}

func FromFloat(val float64) (Amount, error) {
	return FromDecimal(decimal.NewFromFloat(val))
}

func (receiver Amount) String() string {
	return strconv.FormatInt(int64(receiver), 10)
}

func (receiver Amount) AsDecimal() decimal.Decimal {
	return decimal.New(int64(receiver), -2)
}

func (receiver Amount) AsFloat() float64 {
	return receiver.AsDecimal().InexactFloat64()
}
