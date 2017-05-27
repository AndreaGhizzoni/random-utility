package util

import (
	"errors"
	"math/big"
	"github.com/AndreaGhizzoni/zenium/structures"
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

// TODO add description
func FromStringToBigInt(aString string, prefixError string) (*big.Int, error) {
	bigInt := big.NewInt(1)
	if aString == "" {
		return bigInt, nil
	}
	if _, success := bigInt.SetString(aString, baseTen); !success {
		return nil, errors.New(prefixError + ": Fail to cast to big.Int from " +
			"string.")
	}
	return bigInt, nil
}

// TODO add description
func CountSliceIfNotNil(slice []*big.Int) (*big.Int, error) {
	var count = big.NewInt(0)
	for _, element := range slice {
		if element == nil {
			return nil, errors.New("Nil found in slice")
		} else {
			count.Add(count, One)
		}
	}
	return count, nil
}

// TODO add description
func CountBoundsIfNotNil(bounds []*structures.Bound) (*big.Int, error) {
	var count = big.NewInt(0)
	for _, element := range bounds {
		if element == nil {
			return nil, errors.New("Nil found in slice")
		} else {
			count.Add(count, One)
		}
	}
	return count, nil
}

// TODO add description
func CountMatrixIfNotNil(matrix [][]*big.Int) (*big.Int, *big.Int, error) {
	var rows = big.NewInt(0)
    var maxColumns = big.NewInt(0)
	for _, row := range matrix {
		rowLength, err := CountSliceIfNotNil(row)
		if err != nil {
			return nil, nil, err
		}else{
			if rowLength.Cmp(maxColumns) == 1{ // rowLength > maxColumns
				maxColumns = rowLength
			}
			rows.Add(rows, One)
		}
	}
	return rows, maxColumns, nil
}
