package randutil

import "math/rand"

func Int64(r *rand.Rand, min, max int64) (int64, error) {
	// TODO do checks
	return r.Int63n(max-min) + min, nil
}
