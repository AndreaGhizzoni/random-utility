package samples

import (
	"crypto/rand"
	"github.com/AndreaGhizzoni/zenium/structures"
	"github.com/AndreaGhizzoni/zenium/util"
	"math/big"
)

// TODO add doc
type SGenerator struct {
	min, max *big.Int
}

// TODO add doc
func NewSecureGenerator(min, max *big.Int) (*SGenerator, error) {
	if err := util.CheckBounds(min, max); err != nil {
		return nil, err
	}
	return &SGenerator{min, max}, nil
}

// TODO add description
func (this *SGenerator) generateInt() (*big.Int, error) {
	width := big.NewInt(0).Sub(this.max, this.min)
	randomInWidth, err := rand.Int(rand.Reader, width)
	if err != nil {
		return nil, err
	}

	return randomInWidth.Add(randomInWidth, this.min), nil
}

// TODO add description
func (this *SGenerator) generateSlice(len *big.Int) ([]*big.Int, error) {
	randomSlice := []*big.Int{}
	for i := big.NewInt(0); i.Cmp(len) == -1; i.Add(i, util.One) {
		if random, err := this.generateInt(); err != nil {
			return nil, err
		} else {
			randomSlice = append(randomSlice, random)
		}
	}

	return randomSlice, nil
}

// TODO add description
func (this *SGenerator) generateBound(width *big.Int) (*structures.Bound, error) {
	lowerBound, err := this.generateInt()
	if err != nil {
		return nil, err
	}
	upperBound := big.NewInt(0)

	lowerBoundPlusWidth := big.NewInt(0).Add(lowerBound, width)
	if lowerBoundPlusWidth.Cmp(this.max) == 1 { // lowerBound + width > max
		upperBound = this.max
		lowerBound.Sub(upperBound, width) // lowerBound = upperBound - width
	} else {
		upperBound.Add(lowerBound, width) // upperBound = lowerBound + width
	}
	return structures.NewBound(lowerBound, upperBound), nil
}

// TODO add doc
func (this *SGenerator) Int() (*big.Int, error) {
	return this.generateInt()
}

// TODO check this comments
// This function generate a slice of len length, with random numbers X where
// min <= X < max.
// If len <= 0 or min > max return a error.
func (this *SGenerator) Slice(len *big.Int) ([]*big.Int, error) {
	if err := util.IsLessThenOne(len, "Slice length"); err != nil {
		return nil, err
	}

	return this.generateSlice(len)
}

// TODO check this comments
// This function generate a matrix with r rows and c columns. The numbers in it
// are min <= X < max.
// If r <= 0 or c <= 0 or min => max, this function return an error.
func (this *SGenerator) Matrix(rows, columns *big.Int) ([][]*big.Int, error) {
	if err := util.IsLessThenOne(rows, "Matrix rows"); err != nil {
		return nil, err
	}

	if err := util.IsLessThenOne(columns, "Matrix columns"); err != nil {
		return nil, err
	}

	matrix := [][]*big.Int{}
	for i := big.NewInt(0); i.Cmp(rows) == -1; i.Add(i, util.One) {
		random, err := this.generateSlice(columns)
		if err != nil {
			return nil, err
		}
		matrix = append(matrix, random)
	}

	return matrix, nil
}

// TODO check this comments
// This function generate a slice of random structure.Bound.
// min and max represents the lower and the upprer bounds of element's slice.
// width is the fixed with of all the bounds.
// amount is the number of bounds that will be generated.
func (this *SGenerator) Bounds(width, amount *big.Int) ([]*structures.Bound, error) {
	if err := util.IsLessThenOne(width, "Bound width"); err != nil {
		return nil, err
	}

    err := util.IsWidthContainedInBounds(this.min, this.max, width)
	if err!= nil {
		return nil, err
	}

	bounds := []*structures.Bound{}
	for i := big.NewInt(0); i.Cmp(amount) == -1; i.Add(i, util.One) {
		if bound, err := this.generateBound(width); err != nil {
			return nil, err
		} else {
			bounds = append(bounds, bound)
		}
	}
	return bounds, nil
}
