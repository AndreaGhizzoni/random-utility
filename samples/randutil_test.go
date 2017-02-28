package samples

import (
	"math/rand"
	"testing"
)

// This function test the creation of random int64 with correct argument
func TestGetInt64(t *testing.T) {
	var min int64 = 10
	var max int64 = 100

	r := rand.New(NewTimeSeed())
	random, err := getInt64(r,min, max)

	t.Logf("Int64[%d,%d] = %d ", min, max, random)

	if err != nil {
		t.Fatalf("randutil.Int64() returned an error: %v", err)
	}
	// random number must be in range min <= X < max
	if random <= min || random > max {
		t.Fatalf("bounds not respected: min=%d, max=%d", min, max)
	}
}

// This function test the creation of random int64 with incremental range.
func TestIncrementalInt64(t *testing.T) {
	var min int64 = 1
	var max int64 = 10
	var factor int64 = 5
	times := 20

	r := rand.New(NewTimeSeed())
	random, err := getInt64(r, min, max)

	for i := 0; i < times; i++ {
		t.Logf("Int64[%d,%d] = %d ", min, max, random)

		if err != nil {
			t.Fatalf("randutil.Int64() returned an error: %v", err)
		}
		// random number must be in range min <= X < max
		if random <= min || random > max {
			t.Fatalf("bounds not respected: min=%d, max=%d",
				min, max)
		}

		max *= factor
		random, err = getInt64(r, min, max)
	}
}

// This function test the creation of random int64 with incorrect input data.
func TestArgumentsInt64(t *testing.T) {
	// this must fail with nil random generator
	if _, err := getInt64(nil, 10, 100); err == nil {
		t.Fatal("With nil as Rand struct, Int64() needs to return " +
			"(_, nil)")
	}

	// this must fail with min > max
	if _, err := getInt64(rand.New(NewTimeSeed()), 10, 5); err == nil {
		t.Fatal("With min > max as arguments, Int64() needs to return " +
			"(_, nil)")
	}
}

// benchmark random int64 generator
func BenchmarkInt64(b *testing.B) {
	r := rand.New(NewTimeSeed())
	for i := 0; i < b.N; i++ {
		getInt64(r, 10, 100)
	}
}
