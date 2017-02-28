package samples

import (
	"fmt"
	"math/rand"

	"github.com/AndreaGhizzoni/zenium/randutil"
)

// TODO add doc
type Generator struct {
	r *rand.Rand
}

// TODO add doc
func New() *Generator {
	return &Generator{
		r: rand.New(randutil.NewTimeSeed()),
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
func (g *Generator) Slice(len int64, min, max int64) ([]int64, error) {
	err := checkLen(len)
	if err != nil {
		return nil, err
	}

	err = checkBound(min, max)
	if err != nil {
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

// TODO add doc
func (g *Generator) Matrix(rows, cols int64, min, max int64) ([][]int64, error) {
	return nil, nil
}

// TODO add doc
func (g *Generator) Bound(min, max int64, width int64) (int64, int64) {
	return -1, -1
}

// TODO add doc
func (g *Generator) OrderedSlice(len int64, min, max int64) ([]int64, error) {
	return nil, nil
}
