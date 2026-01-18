package amount

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestFromDecimalStrictScale(t *testing.T) {
	cases := []struct {
		name        string
		input       decimal.Decimal
		want        Amount
		wantErr     bool
	}{
		{
			name:    "exact two decimals",
			input:   decimal.New(2000, -2), // 20.00
			want:    Amount(2000),
			wantErr: false,
		},
		{
			name:    "more than two decimals",
			input:   decimal.New(1234, -3), // 1.234
			wantErr: true,
		},
		{
			name:    "less than two decimals",
			input:   decimal.New(1234, -1), // 123.4
			wantErr: true,
		},
		{
			name:    "negative value",
			input:   decimal.New(-9876, -2), // -98.76
			want:    Amount(-9876),
			wantErr: false,
		},
	}

	for _, tc := range cases {
		if got, err := FromDecimal(tc.input); (err != nil) != tc.wantErr {
			t.Fatalf("%s: expected err=%v, got err=%v", tc.name, tc.wantErr, err)
		} else if err == nil && got != tc.want {
			t.Fatalf("%s: expected %d, got %d", tc.name, tc.want, got)
		}
	}
}

func TestFromString(t *testing.T) {
	if got, err := FromString("20.00"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	} else if got != Amount(2000) {
		t.Fatalf("expected 2000, got %d", got)
	}

	if _, err := FromString("20.0"); err == nil {
		t.Fatalf("expected error for non-2-decimal input")
	}
}

func TestFromFloat(t *testing.T) {
	got, err := FromFloat(12.34)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != Amount(1234) {
		t.Fatalf("expected 1234, got %d", got)
	}
}

func TestFormattingAndConversion(t *testing.T) {
	amount := Amount(2000)
	if amount.String() != "2000" {
		t.Fatalf("expected string 2000, got %q", amount.String())
	}

	dec := amount.AsDecimal()
	if !dec.Equal(decimal.New(2000, -2)) {
		t.Fatalf("expected 20.00, got %s", dec.String())
	}

	if amount.AsFloat() != 20.0 {
		t.Fatalf("expected 20, got %f", amount.AsFloat())
	}

	amount = Amount(1)
	dec = amount.AsDecimal()
	if !dec.Equal(decimal.New(1, -2)) {
		t.Fatalf("expected 0.01, got %s", dec.String())
	}
}
