package randutil

import (
	"math/rand"
	"testing"
)

func BenchmarkInt64(b *testing.B) {
	r := rand.New(NewTimeSeed())
	for i := 0; i < b.N; i++ {
		Int64(r, 10, 100)
	}
}

func TestArgumentsInt64(t *testing.T) {
	_, err := Int64(nil, 10, 100)
	if err == nil {
		t.Fatalf("With nil as Rand struct, Int64() needs to return (_, nil)")
	}

	_, err = Int64(rand.New(NewTimeSeed()), 10, 5)
	if err == nil {
		t.Fatalf("With min > max as arguments, Int64() needs to return (_, nil)")
	}
}

func TestGetInt64(t *testing.T) {
	var min int64 = 10
	var max int64 = 100

	r := rand.New(NewTimeSeed())
	random, err := Int64(r, min, max)

	t.Logf("Int64[%d,%d] = %d ", min, max, random)

	if err != nil {
		t.Fatalf("randutil.Int64() returned an error: %v", err)
	}
	if random < min || random > max {
		t.Fatalf("bounds not respected: min=%d, max=%d", min, max)
	}
}

func TestIncrementalInt64(t *testing.T) {
	var min int64 = 1
	var max int64 = 10
	var factor int64 = 5
	times := 20

	r := rand.New(NewTimeSeed())
	random, err := Int64(r, min, max)

	for i := 0; i < times; i++ {

		t.Logf("Int64[%d,%d] = %d ", min, max, random)

		if err != nil {
			t.Fatalf("randutil.Int64() returned an error: %v", err)
		}
		if random < min || random > max {
			t.Fatalf("bounds not respected: min=%d, max=%d", min, max)
		}

		max *= factor
		random, err = Int64(r, min, max)
	}
}
