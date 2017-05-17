package util

import (
	"errors"
	"math/big"
)

var One = big.NewInt(1)

// dim is the dimension to check if it's < 1 and msg is a string to put before
// the error message if check is true.
func IsLessThenOne(dim *big.Int, msgIfTrue string) error {
	if dim.Cmp(One) == -1 { // dim < 1
		errors.New(msgIfTrue + " given is invalid: " + dim.String() + " < 1")
	}
	return nil
}

// utility method to check if min >= max.
func CheckBounds(min, max *big.Int) error {
	if min.Cmp(max) == 1 || min.Cmp(max) == 0 { // min >= max
		return errors.New("Bounds malformed: (min) " + min.String() +
			" >= " + max.String() + " (max)")
	}
	return nil
}
