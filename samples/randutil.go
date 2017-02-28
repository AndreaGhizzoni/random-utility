package samples

import (
	"fmt"
	"math/rand"
	"github.com/iris-contrib/errors"
)

// Int64 generate a random number (min >= x > max) with *rand.Rand given.
// if r is nil or min > max, this function return an error.
func getInt64(r *rand.Rand, min, max int64) (int64, error) {
	if r == nil {
		return -1, errors.New("Random object given is nil")
	}

	if min > max {
		return -1, fmt.Errorf("Bounds error: %d > %d", min, max)
	}

	return r.Int63n(max-min) + min, nil
}
