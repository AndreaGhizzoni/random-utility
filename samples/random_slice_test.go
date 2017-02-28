package samples

import (
	"testing"
)

// This function tests a creation of random slice of with correct input data.
func TestCreateRandomSlice(t *testing.T) {
	var actualLength int64 = 10
	var min int64 = -100
	var max int64 = 100

	testCreateRandomSlice(min, max, actualLength, t)
}

// This function tests a creation of random slice of positive numbers with
// correct input data.
func TestCreatePositiveRandomSlice(t *testing.T) {
	var actualLength int64 = 10
	var min int64 = 10
	var max int64 = 100

	testCreateRandomSlice(min, max, actualLength, t)
}

// This function tests a creation of random slice of negative numbers with
// correct input data.
func TestCreateNegativeRandomSlice(t *testing.T) {
	var actualLength int64 = 10
	var min int64 = -100
	var max int64 = -1

	testCreateRandomSlice(min, max, actualLength, t)
}

// utility functions to create and check a generated slice.
func testCreateRandomSlice(min, max, actualLength int64, t *testing.T) {
	gen := New()
	randomSlice, err := gen.Slice(actualLength, min, max)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("random slice generated: %v", randomSlice)

	// check slice length
	lenRandomSlice := int64(len(randomSlice))
	if lenRandomSlice != actualLength {
		t.Fatalf("length mismatch: %d != %d",
			lenRandomSlice, actualLength)
	}

	// check if the generate random data are in bounds
	for i, e := range randomSlice {
		if e < min || e >= max {
			t.Fatalf("randomSlice[%d]=%d is out of bound (%d, %d)",
				i, e, min, max)
		}
	}
}

// This function tests the creation of random slice with incorrect input data.
func TestArgumentSlice(t *testing.T) {
	gen := New()

	// this must fail: length < 0
	if _, err := gen.Slice(-1, 1, 10); err == nil {
		t.Fatal("With negative length, Slice() needs to return (_, nil)")
	}

	// this must fail: min > max
	if _, err := gen.Slice(10, 10, 1); err == nil {
		t.Fatal("With min > max as argument, Slice() needs to return " +
			"(_, nil)")
	}
}

// benchmark random slice generator
func BenchmarkGenerateSlice(b *testing.B) {
	gen := New()
	for i := 0; i < b.N; i++ {
		gen.Slice(100, 1, 10)
	}
}
