package samples

import (
	"errors"
	"math/rand"
)

// Int64 generate a random number (min >= x > max) with *rand.Rand given.
// if r is nil or min > max, this function return an error.
func getInt64(r *rand.Rand, min, max int64) (int64, error) {
	if r == nil {
		return -1, errors.New("Random object given is nil")
	}

	if err := checkBound(min, max); err != nil {
		return -1, err
	}

	return r.Int63n(max-min) + min, nil
}
