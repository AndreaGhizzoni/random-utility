package samples

import (
	"crypto/rand"
	"github.com/AndreaGhizzoni/zenium/structures"
	"github.com/AndreaGhizzoni/zenium/util"
	"math/big"
)

// This is the generator of random numbers.
type Generator struct {
	min, max *big.Int
}

// NewGenerator returns a new instance of samples.Generator type.
// This generator uses crypto/rand to generate *big.Int numbers between min
// and max. error is returned if min > max or if min == nil || max == nil.
func NewGenerator(min, max *big.Int) (*Generator, error) {
	if err := util.CheckBoundsIfNotNil(min, max); err != nil {
		return nil, err
	}
	return &Generator{min, max}, nil
}

func (g *Generator) generateInt() (*big.Int, error) {
	width := big.NewInt(0).Sub(g.max, g.min)
	randomInWidth, err := rand.Int(rand.Reader, width)
	if err != nil {
		return nil, err
	}

	return randomInWidth.Add(randomInWidth, g.min), nil
}

func (g *Generator) generateSlice(len *big.Int) ([]*big.Int, error) {
	randomSlice := []*big.Int{}
	for i := big.NewInt(0); i.Cmp(len) == -1; i.Add(i, util.One) {
		if random, err := g.generateInt(); err != nil {
			return nil, err
		} else {
			randomSlice = append(randomSlice, random)
		}
	}

	return randomSlice, nil
}

func (g *Generator) generateBound(width *big.Int) (*structures.Bound, error) {
	lowerBound, err := g.generateInt()
	if err != nil {
		return nil, err
	}
	upperBound := big.NewInt(0)

	lowerBoundPlusWidth := big.NewInt(0).Add(lowerBound, width)
	if lowerBoundPlusWidth.Cmp(g.max) == 1 { // lowerBound + width > max
		upperBound = g.max
		lowerBound.Sub(upperBound, width) // lowerBound = upperBound - width
	} else {
		upperBound.Add(lowerBound, width) // upperBound = lowerBound + width
	}
	return structures.NewBound(lowerBound, upperBound), nil
}

// Int generate a random *big.Int according to samples.Generator instanced.
// error is returned if generation fails.
func (g *Generator) Int() (*big.Int, error) {
	return g.generateInt()
}

// Slice generate a slice of length len. error is returned if len == nil or
// if single *big.Int generation fails.
func (g *Generator) Slice(len *big.Int) ([]*big.Int, error) {
	if err := util.IsNilOrLessThenOne(len, "Slice length"); err != nil {
		return nil, err
	}

	return g.generateSlice(len)
}

// Matrix generate a matrix with rows and columns given according to
// samples.Generator instanced. error is returned if: rows == nil,
// columns == nil, rows >= 1, columns >= 1 or if single *bit.Int generation
// fails.
func (g *Generator) Matrix(rows, columns *big.Int) ([][]*big.Int, error) {
	if err := util.IsNilOrLessThenOne(rows, "Matrix rows"); err != nil {
		return nil, err
	}

	if err := util.IsNilOrLessThenOne(columns, "Matrix columns"); err != nil {
		return nil, err
	}

	matrix := [][]*big.Int{}
	for i := big.NewInt(0); i.Cmp(rows) == -1; i.Add(i, util.One) {
		random, err := g.generateSlice(columns)
		if err != nil {
			return nil, err
		}
		matrix = append(matrix, random)
	}

	return matrix, nil
}

// This function generate a slice of random structures.Bound. width is the
// fixed with of all the bounds. amount is the number of bounds that will be
// generated. error is returned if: width == nil, width >= 1, width can not
// be placed between min and max or if single *bit.Int generation fails.
func (g *Generator) Bounds(width, amount *big.Int) ([]*structures.Bound, error) {
	if err := util.IsNilOrLessThenOne(width, "Bound width"); err != nil {
		return nil, err
	}

	err := util.IsWidthContainedInBounds(g.min, g.max, width)
	if err != nil {
		return nil, err
	}

	bounds := []*structures.Bound{}
	for i := big.NewInt(0); i.Cmp(amount) == -1; i.Add(i, util.One) {
		if bound, err := g.generateBound(width); err != nil {
			return nil, err
		} else {
			bounds = append(bounds, bound)
		}
	}
	return bounds, nil
}
