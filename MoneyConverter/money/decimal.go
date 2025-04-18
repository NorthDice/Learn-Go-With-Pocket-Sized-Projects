package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Decimal represents a decimal number
type Decimal struct {
	// subunit is the smallest unit of the currency
	subunits  int64
	precision byte
}

var (
	// ErrInvalidDecimal is returned if the decimal is malformed.
	ErrInvalidDecimal = Error("unable to convert the decimal")

	// ErrTooLarge is returned if the quantity is too large - this would cause floating point precision errors.
	ErrTooLarge = Error("quantity over 10^12 is too large")
)

// ParseDecimal converts a string into its Decimal representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if subunits > 1e12 {
		return Decimal{}, ErrTooLarge
	}

	precision := byte(len(fracPart))
	dec := Decimal{subunits: subunits, precision: precision}

	// Let's clean the representation a bit. Remove trailing zeroes.
	dec.simplify()

	return dec, nil
}
func (d *Decimal) String() string {
	if d.precision == 0 {
		return fmt.Sprintf("%d", d.subunits)
	}

	centsPerUnit := pow10(d.precision)
	frac := d.subunits % centsPerUnit
	integer := d.subunits / centsPerUnit

	decimalFormat := "%d.%0" + strconv.Itoa(int(d.precision)) + "d"

	return fmt.Sprintf(decimalFormat, integer, frac)
}

// simplifies removes trailing zeroes - as long as they're on the right side of the decimal separator.
func (d *Decimal) simplify() {
	// Using %10 returns the last digit in base 10 of a number.
	// If the precision is positive, that digit belongs to the right side of the decimal separator.
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}

// pow10 returns is a quick way to calculate 10^n.
// It's optimized for small n, and uses a simple loop for larger n.
func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}
