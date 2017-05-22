package util

import (
	"errors"
	"math/big"
)

// One represents the integer 1 as *big.Int
var One = big.NewInt(1)

// IsLessThenOne check if dim < 1. If is true create a new error with msgIfTrue
// as error's message prefix, otherwise return nil.
func IsLessThenOne(dim *big.Int, msgIfTrue string) error {
	if dim.Cmp(One) == -1 { // dim < 1
		return errors.New(msgIfTrue + " given is invalid: " + dim.String() +
			" < 1")
	}
	return nil
}

// CheckBounds check if min >= max. If is true create a new error with
// appropriate description, otherwise return nil.
func CheckBounds(min, max *big.Int) error {
	if min.Cmp(max) == 1 || min.Cmp(max) == 0 { // min >= max
		return errors.New("Bounds malformed: (min) " + min.String() +
			" >= " + max.String() + " (max)")
	}
	return nil
}

// IsWidthContainedInBounds check if argument width given is greater then
// |max-min|. If is true create a new error with appropriate description,
// otherwise return nil.
func IsWidthContainedInBounds(min, max, width *big.Int) error {
	maxMinusMin := big.NewInt(0).Sub(max, min)
	maxMinusMin.Abs(maxMinusMin)
	if width.Cmp(maxMinusMin) == 1 { // width > maxMinusMin
		return errors.New("Width given is greater than interval selected: " +
			width.String() + " > " + maxMinusMin.String())
	}
	return nil
}
