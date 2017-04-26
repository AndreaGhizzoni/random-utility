package samples

import (
	"math/rand"
	"testing"
)

// This function test the creation of random int64 with correct input data.
func TestGetInt64(t *testing.T) {
	var tableTest = []struct {
		min, max int64
	}{
		{0, 10},
		{0, 100},
		{0, 1000},
		{0, 10000},
		{-10, 0},
		{-100, 0},
		{-1000, 0},
		{-10000, 0},
		{-10, 10},
		{-10, 100},
		{-1000, 1000},
		{-10000, 10000},
		{-100000, 100000},
		{-1000000, 1000000},
	}

	r := rand.New(NewTimeSeed())
	for _, tt := range tableTest {
		random, err := getInt64(r, tt.min, tt.max)
		if err != nil {
			t.Fatalf("getInt64() returned an error: %v", err)
		}

		// random number must be in range tt.min <= X < tt.max
		if random < tt.min || random >= tt.max {
			t.Fatalf("bounds not respected: min=%d, max=%d", tt.min, tt.max)
		}

		t.Logf("getInt64(%d,%d) = %d ", tt.min, tt.max, random)
	}
}

// This function test the creation of random int64 with wrong input data.
func TestGetInt64_Arguments(t *testing.T) {
	// this must fail with nil random generator
	if _, err := getInt64(nil, 10, 100); err == nil {
		t.Fatal("With nil as Rand struct, getInt64() needs to return " +
			"(_, nil)")
	}

	// this must fail with min > max
	if _, err := getInt64(rand.New(NewTimeSeed()), 10, 5); err == nil {
		t.Fatal("With min > max as arguments, getInt64() needs to " +
			"return (_, nil)")
	}
}


