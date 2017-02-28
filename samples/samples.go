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

func checkLen(len int64) error {
	if len <= 0 {
		return fmt.Errorf("length (%d) must be grater then zero.", len)
	}
	return nil
}

func checkBound(min, max int64) error {
	if min > max {
		return fmt.Errorf("bounds malformed: %d > %d", min, max)
	}
	return nil
}

// TODO add doc
func (s *Generator) Slice(len int64, min, max int64) ([]int64, error) {
	err := checkLen(len)
	if err != nil {
		return nil, err
	}

	err = checkBound(min, max)
	if err != nil {
		return nil, err
	}

	r := rand.New(randutil.NewTimeSeed())
	perm := make([]int64, len)
	var i int64 = 0
	for ; i < len; i++ {
		intRandom, e := randutil.Int64(r, min, max)
		if e != nil {
			return nil, e
		}
		perm[i] = intRandom
	}

	return perm, nil
}

// TODO add doc
func (s *Generator) Matrix(rows, cols int64, min, max int64) ([][]int64, error) {
	return nil, nil
}

// TODO add doc
func (s *Generator) Bound(min, max int64, width int64) (int64, int64) {
	return -1, -1
}

// TODO add doc
func (s *Generator) OrderedSlice(len int64, min, max int64) ([]int64, error) {
	return nil, nil
}
