package util

import (
	"errors"
	"github.com/AndreaGhizzoni/zenium/structures"
	"math/big"
)

// One represents the integer 1 as *big.Int
var One = big.NewInt(1)

var baseTen = 10

// IsNilOrLessThenOne check if dim != nil and dim < 1. If is true create a new
// error with msgIfTrue as error's message prefix, otherwise return nil.
func IsNilOrLessThenOne(dim *big.Int, msgIfTrue string) error {
	if dim == nil {
		return errors.New("Dimension given can not be nil.")
	}

	if dim.Cmp(One) == -1 { // dim < 1
		return errors.New(msgIfTrue + " given is invalid: " + dim.String() +
			" < 1")
	}
	return nil
}

// CheckBoundsIfNotNil check min != nil && max != nil and thane check if
// min >= max. If is true create a new error with appropriate description,
// otherwise return nil.
func CheckBoundsIfNotNil(min, max *big.Int) error {
	if min == nil {
		return errors.New("minimum given can not be nil.")
	}
	if max == nil {
		return errors.New("maximum given can not be nil.")
	}

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

// FromStringToBigInt converts aString given into the corresponding *big.Int
// in base 10. error is return if: aString is not a valid number.
func FromStringToBigInt(aString string, prefixError string) (*big.Int, error) {
	stringIntoBigInt := big.NewInt(1)
	if aString == "" { // safe condition
		return stringIntoBigInt, nil
	}
	if _, success := stringIntoBigInt.SetString(aString, baseTen); !success {
		return nil, errors.New(prefixError + ": Fail to cast from string to " +
			"big.Int.")
	}
	return stringIntoBigInt, nil
}

// CountSliceIfNotNil counts the *big.Int in the slice given and returns it
// as big.Int. Counts goes well if every *big.Int in slice is not nil, otherwise
// returns an error.
func CountSliceIfNotNil(slice []*big.Int) (*big.Int, error) {
	counter := big.NewInt(0)
	for _, element := range slice {
		if element == nil {
			return nil, errors.New("Nil element found in slice.")
		}
		counter.Add(counter, One)
	}
	return counter, nil
}

// CountBoundsIfNotNil counts the *structure.Bound the slice given and returns
// it as big.Int. Counts goes well if every *big.Int in slice is not nil,
// otherwise returns an error.
func CountBoundsIfNotNil(bounds []*structures.Bound) (*big.Int, error) {
	counter := big.NewInt(0)
	for _, element := range bounds {
		if element == nil {
			return nil, errors.New("Nil element found in slice.")
		}
		counter.Add(counter, One)
	}
	return counter, nil
}

// CountMatrixIfNotNil return the rows and columns of given matrix of *big.Int
// as *big.Int. Counts goes well if every *big.Int in matrix is not nil,
// otherwise return an error.
func CountMatrixIfNotNil(matrix [][]*big.Int) (rows, columns *big.Int, err error) {
	rows = big.NewInt(0)
	columns = big.NewInt(0)
	for _, row := range matrix {
		rowLength, err := CountSliceIfNotNil(row)
		if err != nil {
			return nil, nil, err
		}
		if rowLength.Cmp(columns) == 1 { // rowLength > maxColumns
			columns = rowLength
		}
		rows.Add(rows, One)
	}
	return rows, columns, nil
}
