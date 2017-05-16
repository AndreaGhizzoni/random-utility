package samples

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// TODO add doc
type SGenerator struct{}

var zero = big.NewInt(0)
var one = big.NewInt(1)

// TODO add doc
func NewSecureGenerator() *SGenerator {
	return &SGenerator{}
}

// utility method to check if big.Int given given is <= 0.
// dim is the dimension to check and msg is a string to put before the error
// message string
func isLessThenZero(dim *big.Int, msgIfTrue string) error {
	if dim.Cmp(zero) == -1 || dim.Cmp(zero) == 0 { // dim <= zero
		return fmt.Errorf(msgIfTrue+" given is invalid: %v <= 0.", dim)
	}
	return nil
}

// utility method to check if min >= max.
func checkBounds(min, max *big.Int) error {
	if min.Cmp(max) == 1 || min.Cmp(max) == 0 { // min >= max
		return fmt.Errorf("Bounds malformed: (min) %v >= %v (max)", min, max)
	}
	return nil
}

// TODO add doc
func (g *SGenerator) generateInt(min, max *big.Int) (*big.Int, error) {
	r, err := rand.Int(rand.Reader, max.Sub(max, min))
	if err != nil {
		return nil, err
	}
	return r.Add(r, min), nil

}

// TODO add doc
func (g *SGenerator) Int(min, max *big.Int) (*big.Int, error) {
	if err := checkBounds(min, max); err != nil {
		return nil, err
	}
	return g.generateInt(min, max)
}

// This function generate a slice of len length, with random numbers X where
// min <= X < max.
// If len <= 0 or min > max return a error.
func (g *SGenerator) Slice_(len, min, max *big.Int) ([]*big.Int, error) {
	if err := isLessThenZero(len, "Slice length"); err != nil {
		return nil, err
	}

	if err := checkBounds(min, max); err != nil {
		return nil, err
	}

	perm := []*big.Int{}
	i := big.NewInt(0)
	for ; i.Cmp(len) == -1; i.Add(i, one) {
		if r, err := g.generateInt(min, max); err != nil {
			return nil, err
		} else {
			perm = append(perm, r)
		}
	}

	return perm, nil
}

/*
// This function generate a matrix with r rows and c cols. The number X in it
// are min <= X < max.
// If r <= 0 or c <= 0 or min > max, this function return an error
func (g *Generator) Matrix(r, c, min, max int64) ([][]int64, error) {
	// check rows dimension
	if err := checkDimension(r, "Matrix rows"); err != nil {
		return nil, err
	}

	// check columns too because, in case of error in matrix generation, the
	// following error message will be displayed: "Slice length given ...", and
	// is this method do not generate a slice, but a matrix.
	if err := checkDimension(c, "Matrix columns"); err != nil {
		return nil, err
	}

	matrix := make([][]int64, r)
	var i int64 = 0
	for ; i < r; i++ {
		perm, e := g.Slice(c, min, max)
		if e != nil {
			return nil, e
		}
		matrix[i] = perm
	}

	return matrix, nil
}

// This function generate a random bound of fixed length. min and max are the
// minimum and the maximum bounds that the bound will be generated of length
// width.
func (g *Generator) Bound(min, max, width int64) (*Bound, error) {
	if err := checkDimension(width, "Bound with"); err != nil {
		return nil, err
	}

	if err := checkBound(min, max); err != nil {
		return nil, err
	}

	bLow, bUp := g.generateInt(min, max), int64(0)
	if bLow+width > max {
		bUp = max
		bLow = bUp - width
	} else {
		bUp = bLow + width
	}
	return NewBound(bLow, bUp), nil
}
*/
