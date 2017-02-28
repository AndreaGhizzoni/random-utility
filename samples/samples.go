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
func New() *Generator {
	return &Generator{
		r: rand.New(NewTimeSeed()),
	}
}

// utility method to check if len <= 0
func checkLen(len int64) error {
	if len <= 0 {
		return fmt.Errorf("length (%d) must be grater then zero.", len)
	}
	return nil
}

// utility method to check if min > max
func checkBound(min, max int64) error {
	if min > max {
		return fmt.Errorf("bounds malformed: %d > %d", min, max)
	}
	return nil
}

// This function generate a slice of len length, with random numbers X where
// min <= X < max.
// If len <= 0 or min > max return a error.
func (g *Generator) Slice(len, min, max int64) ([]int64, error) {
	if err := checkLen(len); err != nil {
		return nil, err
	}

	if err := checkBound(min, max); err != nil {
		return nil, err
	}

	perm := make([]int64, len)
	var i int64 = 0
	for ; i < len; i++ {
		intRandom, e := getInt64(g.r, min, max)
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
	if err := checkLen(r); err != nil {
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
