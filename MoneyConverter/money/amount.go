package money

// Amount represents a monetary amount
type Amount struct {
	quantity Decimal
	currency Currency
}

// ErrTooPrecise is returned if the quantity is too precise.
const ErrTooPrecise = Error("quantity is too precise")

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		// In order to avoid converting 0.00001 cent
		return Amount{}, ErrTooPrecise
	}
	quantity.precision = currency.precision

	return Amount{quantity: quantity, currency: currency}, nil
}
