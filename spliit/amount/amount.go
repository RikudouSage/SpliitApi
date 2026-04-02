package amount

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// Amount stores a decimal monetary value with two fractional digits.
type Amount int64

// FromDecimal converts a decimal with exactly two fractional digits into an Amount.
func FromDecimal(val decimal.Decimal) (Amount, error) {
	if val.Exponent() != -2 {
		return 0, fmt.Errorf("amount must have exactly 2 decimal places, got exponent %d", val.Exponent())
	}

	return Amount(val.CoefficientInt64()), nil
}

// FromString parses a decimal string into an Amount.
func FromString(val string) (Amount, error) {
	dec, err := decimal.NewFromString(val)
	if err != nil {
		return 0, err
	}

	return FromDecimal(dec)
}

// FromFloat converts a floating-point value into an Amount.
func FromFloat(val float64) (Amount, error) {
	return FromDecimal(decimal.NewFromFloat(val))
}

// String returns the raw fixed-point integer representation.
func (receiver Amount) String() string {
	return strconv.FormatInt(int64(receiver), 10)
}

// AsDecimal converts the amount to a decimal.Decimal with two fractional digits.
func (receiver Amount) AsDecimal() decimal.Decimal {
	return decimal.New(int64(receiver), -2)
}

// AsFloat converts the amount to a float64.
func (receiver Amount) AsFloat() float64 {
	return receiver.AsDecimal().InexactFloat64()
}
