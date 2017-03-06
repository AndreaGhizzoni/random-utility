package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"testing"
)

func TestCreationRandomSlice(t *testing.T) {
	var tableTest = []struct {
		length int64
		min    int64
		max    int64
	}{
		{1, 1, 10},
		{10, 1, 10},
		{20, 1, 10},
		{1, -10, -1},
		{10, -10, -1},
		{20, -10, -1},
		{1, -100, 100},
		{10, -1000, 1000},
		{20, -10000, 10000},
		{20, -100000, 100000},
		{20, -1000000, 1000000},
		{20, -10000000, 10000000},
	}

	gen := samples.NewGenerator()
	for _, tt := range tableTest {
		rSlice, err := gen.Slice(tt.length, tt.min, tt.max)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("rslice generated: %v", rSlice)

		// check slice length
		actual := int64(len(rSlice))
		if tt.length != actual {
			t.Fatalf("length mismatch: actual %d != expected %d", actual,
				tt.length)
		}

		// check if the generate random data are in bounds
		for i, e := range rSlice {
			if e < tt.min || e >= tt.max {
				t.Fatalf("randomSlice[%d]=%d is out of bound (%d, %d)",
					i, e, tt.min, tt.max)
			}
		}
	}
}

// This function tests the creation of random slice with incorrect input data.
func TestArgumentSlice(t *testing.T) {
	gen := samples.NewGenerator()

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
	gen := samples.NewGenerator()
	for i := 0; i < b.N; i++ {
		gen.Slice(100, 1, 10)
	}
}
