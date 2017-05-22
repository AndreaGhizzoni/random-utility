package samples

import (
	"crypto/rand"
	"github.com/AndreaGhizzoni/zenium/structures"
	"github.com/AndreaGhizzoni/zenium/util"
	"math/big"
)

// TODO add doc
type SGenerator struct{}

// TODO add doc
func NewSecureGenerator() *SGenerator {
	return &SGenerator{}
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
	for ; i.Cmp(len) == -1; i.Add(i, util.One) {
		if r, err := generateInt(min, max); err != nil {
			return nil, err
		} else {
			perm = append(perm, r)
		}
	}

	return perm, nil
}

// TODO add description
func generateBound(min, max, width *big.Int) (*structures.Bound, error) {
	bLow, err := generateInt(min, max)
	if err != nil {
		return nil, err
	}
	bUp := big.NewInt(0)

	bLowPlusWidth := big.NewInt(0).Add(bLow, width)
	if bLowPlusWidth.Cmp(max) == 1 { // bLow+width > max
		bUp = max
		bLow.Sub(bUp, width) //bLow = bUp - width
	} else {
		bUp.Add(bLow, width)
	}
	return structures.NewBound(bLow, bUp), nil
}

// TODO add doc
func (g *SGenerator) Int(min, max *big.Int) (*big.Int, error) {
	if err := util.CheckBounds(min, max); err != nil {
		return nil, err
	}
	return generateInt(min, max)
}

// This function generate a slice of len length, with random numbers X where
// min <= X < max.
// If len <= 0 or min > max return a error.
func (g *SGenerator) Slice(len, min, max *big.Int) ([]*big.Int, error) {
	if err := util.IsLessThenOne(len, "Slice length"); err != nil {
		return nil, err
	}

	if err := util.CheckBounds(min, max); err != nil {
		return nil, err
	}

	return generateSlice(len, min, max)
}

// This function generate a matrix with r rows and c columns. The numbers in it
// are min <= X < max.
// If r <= 0 or c <= 0 or min => max, this function return an error.
func (g *SGenerator) Matrix(r, c, min, max *big.Int) ([][]*big.Int, error) {
	// check rows dimension
	if err := util.IsLessThenOne(r, "Matrix rows"); err != nil {
		return nil, err
	}

	// check columns too because, in case of error in matrix generation, the
	// following error message will be displayed: "Slice length given ...", and
	// this method do not generate any slice, but a matrix.
	if err := util.IsLessThenOne(c, "Matrix columns"); err != nil {
		return nil, err
	}

	// check ranges
	if err := util.CheckBounds(min, max); err != nil {
		return nil, err
	}

	matrix := [][]*big.Int{}
	i := big.NewInt(0)
	for ; i.Cmp(r) == -1; i.Add(i, util.One) {
		perm, e := generateSlice(c, min, max)
		if e != nil {
			return nil, e
		}
		matrix = append(matrix, perm)
	}

	return matrix, nil
}

// This function generate a slice of random structure.Bound.
// min and max represents the lower and the upprer bounds of element's slice.
// width is the fixed with of all the bounds.
// amount is the number of bounds that will be generated.
func (g *SGenerator) Bounds(min, max, width, amount *big.Int) ([]*structures.Bound, error) {
	if err := util.IsLessThenOne(width, "Bound width"); err != nil {
		return nil, err
	}

	if err := util.CheckBounds(min, max); err != nil {
		return nil, err
	}

	if err := util.IsWidthContainedInBounds(min, max, width); err != nil {
		return nil, err
	}

	boundSlice := []*structures.Bound{}
	i := big.NewInt(0)
	for ; i.Cmp(amount) == -1; i.Add(i, util.One) {
		if bound, err := generateBound(min, max, width); err != nil {
			return nil, err
		} else {
			boundSlice = append(boundSlice, bound)
		}
	}
	return boundSlice, nil
}
