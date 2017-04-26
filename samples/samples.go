package samples

import (
	"fmt"
	"math/rand"
)

// This is the generator of random numbers.
type Generator struct {
	r *rand.Rand
}

// This method returns a new instance of Generator type.
// This generator is initialized with current time as seed.
func NewGenerator() *Generator {
	return &Generator{
		r: rand.New(NewTimeSeed()),
	}
}

// utility method to check if dimension given is <= 0
// dim is the dimension to check and msg is a string to put before the error
// message string
func checkDimension(dim int64, msg string) error {
	if dim <= 0 {
		return fmt.Errorf(msg+" given is invalid: %d <= 0.", dim)
	}
	return nil
}

// utility method to check if min > max
func checkBound(min, max int64) error {
	if min > max {
		return fmt.Errorf("Bounds malformed: (min) %d > %d (max)", min, max)
	}
	return nil
}

// Int64 generate a random number between min <= x < max. err != nil if
// min > max.
func (g *Generator) Int64(min, max int64) (int64, error) {
	if err := checkBound(min, max); err != nil {
		return -1, err
	}
	return g.r.Int63n(max-min) + min, nil
}

// This function generate a slice of len length, with random numbers X where
// min <= X < max.
// If len <= 0 or min > max return a error.
func (g *Generator) Slice(len, min, max int64) ([]int64, error) {
	if err := checkDimension(len, "Slice length"); err != nil {
		return nil, err
	}

	if err := checkBound(min, max); err != nil {
		return nil, err
	}

	perm := make([]int64, len)
	var i int64 = 0
	for ; i < len; i++ {
		intRandom, e := g.Int64(min, max)
		if e != nil {
			return nil, e
		}
		perm[i] = intRandom
	}

	return perm, nil
}

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

// TODO add doc
func (g *Generator) Bound(min, max, width int64) (int64, int64) {
	return -1, -1
}

// TODO add doc
func (g *Generator) OrderedSlice(len, min, max int64) ([]int64, error) {
	return nil, nil
}
