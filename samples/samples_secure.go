package samples

import (
	"crypto/rand"
	"fmt"
	"github.com/AndreaGhizzoni/zenium/structures"
	"math/big"
)

// TODO add doc
type SGenerator struct{}

var one = big.NewInt(1)

// TODO add doc
func NewSecureGenerator() *SGenerator {
	return &SGenerator{}
}

// dim is the dimension to check if it's < 1 and msg is a string to put before
// the error message if check is true.
func isLessThenOne(dim *big.Int, msgIfTrue string) error {
	if dim.Cmp(one) == -1 { // dim < 1
		return fmt.Errorf(msgIfTrue+" given is invalid: %v < 1.", dim)
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

// TODO add description
func generateInt(min, max *big.Int) (*big.Int, error) {
	r, err := rand.Int(rand.Reader, max.Sub(max, min))
	if err != nil {
		return nil, err
	}
	return r.Add(r, min), nil

}

// TODO add description
func generateSlice(len, min, max *big.Int) ([]*big.Int, error) {
	perm := []*big.Int{}
	i := big.NewInt(0)
	for ; i.Cmp(len) == -1; i.Add(i, one) {
		if r, err := generateInt(min, max); err != nil {
			return nil, err
		} else {
			perm = append(perm, r)
		}
	}

	return perm, nil
}

// TODO add doc
func (g *SGenerator) Int(min, max *big.Int) (*big.Int, error) {
	if err := checkBounds(min, max); err != nil {
		return nil, err
	}
	return generateInt(min, max)
}

// This function generate a slice of len length, with random numbers X where
// min <= X < max.
// If len <= 0 or min > max return a error.
func (g *SGenerator) Slice(len, min, max *big.Int) ([]*big.Int, error) {
	if err := isLessThenOne(len, "Slice length"); err != nil {
		return nil, err
	}

	if err := checkBounds(min, max); err != nil {
		return nil, err
	}

	return generateSlice(len, min, max)
}

// This function generate a matrix with r rows and c columns. The numbers in it
// are min <= X < max.
// If r <= 0 or c <= 0 or min => max, this function return an error.
func (g *SGenerator) Matrix(r, c, min, max *big.Int) ([][]*big.Int, error) {
	// check rows dimension
	if err := isLessThenOne(r, "Matrix rows"); err != nil {
		return nil, err
	}

	// check columns too because, in case of error in matrix generation, the
	// following error message will be displayed: "Slice length given ...", and
	// this method do not generate any slice, but a matrix.
	if err := isLessThenOne(c, "Matrix columns"); err != nil {
		return nil, err
	}

	// check ranges
	if err := checkBounds(min, max); err != nil {
		return nil, err
	}

	matrix := [][]*big.Int{}
	i := big.NewInt(0)
	for ; i.Cmp(r) == -1; i.Add(i, one) {
		perm, e := generateSlice(c, min, max)
		if e != nil {
			return nil, e
		}
		matrix = append(matrix, perm)
	}

	return matrix, nil
}

// This function generate a random bound of fixed length. min and max are the
// minimum and the maximum bounds that the bound will be generated of length
// width.
func (g *SGenerator) Bound(min, max, width *big.Int) (*structures.Bound, error) {
	if err := isLessThenOne(width, "Bound with"); err != nil {
		return nil, err
	}

	if err := checkBound(min, max); err != nil {
		return nil, err
	}

	bLow, err := generateInt(min, max)
	if err != nil {
		return nil, err
	}
	bUp := big.NewInt(0)

	i := big.NewInt(0).Add(bLow, width)
	if i.Cmp(max) == 1 { // bLow+width > max
		bUp = max
		bLow.Sub(bUp, width) //bLow = bUp - width
	} else {
		bUp.Add(bLow, width)
	}
	return structures.NewBound(bLow, bUp), nil
}
