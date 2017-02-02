package randutil

import "math/rand"

func RandomInt64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func RandomFloat64(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}
