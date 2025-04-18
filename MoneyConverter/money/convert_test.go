package money

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   Amount
		to       Currency
		validate func(t *testing.T, got Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: mustParseAmount(t, "34.98", "USD"),
			to:     mustParseCurrency(t, "EUR"),
			validate: func(t *testing.T, got Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %v", err.Error())
				}
				expected := mustParseAmount(t, "69.96", "EUR")
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %v, got %v", expected, got)
				}

			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
	}
}
