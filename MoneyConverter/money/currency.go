package money

// Currency represents a currency code
type Currency struct {
	code      string
	precision byte
}

// ExchangeRate represents the exchange rate between two currencies
type ExchangeRate Decimal

// ErrInvalidCurrencyCode is returned if the currency code is malformed.
const ErrInvalidCurrencyCode = Error("invalid currency code")

// ParseCurrency converts a string into its Currency representation and may return ErrInvalidCurrencyCode.
func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}
