package samples

import (
	"testing"
)

// This function tests a creation of random slice with correct input data.
func TestCreateRandomSlice(t *testing.T) {
	var actualLength int64 = 10
	var min int64 = 10
	var max int64 = 100

	gen := New()
	randomSlice, err := gen.Slice(actualLength, min, max)
	t.Logf("random slice generated: %v", randomSlice)

	if err != nil {
		t.Fatal(err)
	}

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
	_, err1 := gen.Slice(-1, 1, 10)
	if err1 == nil {
		t.Fatal("With negative length, Slice() needs to return (_, nil)")
	}

	// this must fail: min > max
	_, err2 := gen.Slice(10, 10, 1)
	if err2 == nil {
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
