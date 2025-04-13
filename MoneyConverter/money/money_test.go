package money

import "testing"

func mustParseCurrency(t *testing.T, code string) Currency {
	t.Helper()
	currency, err := ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency %s", code)
	}
	return currency
}
func mustParseAmount(t *testing.T, value string, code string) Amount {
	t.Helper()

	n, err := ParseDecimal(value)
	if err != nil {
		t.Fatalf("invalid number %s", value)
	}

	currency, err := ParseCurrency(code)
	if err != nil {
		t.Fatalf("invalid currency code %s", code)
	}

	amount, err := NewAmount(n, currency)
	if err != nil {
		t.Fatalf("cannot create amount with value %v and currency code %s", n, code)
	}

	return amount
}
