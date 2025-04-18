package money

// Amount represents a monetary amount
type Amount struct {
	quantity Decimal
	currency Currency
}

// ErrTooPrecise is returned if the quantity is too precise.
const ErrTooPrecise = Error("quantity is too precise")

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	switch {
	case quantity.precision > currency.precision:
		return Amount{}, ErrTooPrecise
	case quantity.precision < currency.precision:
		quantity.subunits *= pow10(currency.precision - quantity.precision)
		quantity.precision = currency.precision

	}

	return Amount{quantity: quantity, currency: currency}, nil
}

func (a Amount) String() string {
	return a.quantity.String() + " " + a.currency.code
}

func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > 1e12:
		return ErrTooLarge

	case a.quantity.precision > a.currency.precision:

		return ErrTooPrecise
	}

	return nil
}
