package samples

import (
	"math/rand"
	"testing"
)

// This function test the creation of random positive int64 with correct
// argument
func TestGetPositiveInt64(t *testing.T) {
	var min int64 = 10
	var max int64 = 100

	testSingleRandom(rand.New(NewTimeSeed()), min, max, t)
}

// This function test the creation of random negative int64 with correct
// argument
func TestGetNegativeInt64(t *testing.T) {
	var min int64 = -100
	var max int64 = -1

	testSingleRandom(rand.New(NewTimeSeed()), min, max, t)
}

// utility functions that test a single random number generated.
func testSingleRandom(r *rand.Rand, min, max int64, t *testing.T) {
	random, err := getInt64(r, min, max)
	if err != nil {
		t.Fatalf("getInt64() returned an error: %v", err)
	}
	t.Logf("Int64[%d,%d] = %d ", min, max, random)

	// random number must be in range min <= X < max
	if random <= min || random > max {
		t.Fatalf("bounds not respected: min=%d, max=%d", min, max)
	}
}

// This function test the creation of positive random int64 with incremental
// range.
func TestIncrementalPositiveInt64(t *testing.T) {
	var min int64 = 1
	var max int64 = 10
	var factor int64 = 5
	times := 20

	doIncrementalTest(min, max, factor, times, true, t)
}

// This function test the creation of negative random int64 with incremental
// range.
func TestIncrementalNegativeInt64(t *testing.T) {
	var min int64 = -10
	var max int64 = -1
	var factor int64 = 5
	times := 20

	doIncrementalTest(min, max, factor, times, false, t)
}

// utility function that execute an incremental test.
func doIncrementalTest(min, max, factor int64, times int,
	workWithPositive bool, t *testing.T) {

	r := rand.New(NewTimeSeed())
	for i := 0; i < times; i++ {
		testSingleRandom(r, min, max, t)

		if workWithPositive {
			max *= factor
		} else {
			min *= factor
		}
	}
}

// This function test the creation of random int64 with incorrect input data.
func TestArgumentsInt64(t *testing.T) {
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

// benchmark random int64 generator
func BenchmarkInt64(b *testing.B) {
	r := rand.New(NewTimeSeed())
	for i := 0; i < b.N; i++ {
		getInt64(r, 10, 100)
	}
}
