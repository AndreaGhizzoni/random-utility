package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"testing"
)

// This function test the creation of random int64 with correct input data.
func TestGenerator_Int64(t *testing.T) {
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

	gen := samples.NewGenerator()
	for _, tt := range tableTest {
		random, err := gen.Int64(tt.min, tt.max)
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
func TestGenerator_Int64_Arguments(t *testing.T) {
	gen := samples.NewGenerator()

	// this must fail with min > max
	min, max := int64(10), int64(5)
	if _, err := gen.Int64(min, max); err == nil {
		t.Fatalf("gen.Int64(%d, %d) must fail", min, max)
	}
}

// This function tests the creation of random slice with correct input data.
func TestGenerator_Slice(t *testing.T) {
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

// This function tests the creation of random slice with wrong input data.
func TestGenerator_Slice_Arguments(t *testing.T) {
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

// This function tests the creation of random matrix with correct input data.
func TestGenerator_Matrix(t *testing.T) {
	var tableTest = []struct {
		r   int64
		c   int64
		min int64
		max int64
	}{
		{1, 1, 1, 10},
		{10, 10, 1, 10},
		{20, 20, 1, 10},
		{1, 1, -10, -1},
		{10, 10, -10, -1},
		{20, 20, -10, -1},
		{1, 1, -100, 100},
		{10, 10, -1000, 1000},
		{20, 20, -10000, 10000},
		{20, 20, -100000, 100000},
		{20, 20, -1000000, 1000000},
		{20, 20, -10000000, 10000000},
		{1, 20, -10000, 10000},
		{20, 1, -100000, 100000},
		{20, 5, -1000000, 1000000},
		{5, 20, -10000000, 10000000},
	}

	gen := samples.NewGenerator()
	for _, tt := range tableTest {
		matrix, err := gen.Matrix(tt.r, tt.c, tt.min, tt.max)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("random matrix generated: %v", matrix)

		// check rows of generated matrix
		actualRows := int64(len(matrix))
		if actualRows != tt.r {
			t.Fatalf("rows length mismatch: actual %d != expected %d",
				actualRows, tt.r)
		}

		// check cols of generated matrix
		actualCols := int64(len(matrix[0]))
		if actualCols != tt.c {
			t.Fatalf("cols length mismatch: actual %d != expected %d",
				actualCols, tt.c)
		}

		for i, slice := range matrix {
			for j, e := range slice {
				if e < tt.min || e >= tt.max {
					t.Fatalf("matrix[%d][%d]=%d is out of bound (%d, %d)",
						i, j, e, tt.min, tt.max)
				}
			}
		}
	}
}

// This function tests the creation of random matrix with wrong input data.
func TestGenerator_Matrix_Arguments(t *testing.T) {
	gen := samples.NewGenerator()

	// this must fail, rows < 0
	if _, err := gen.Matrix(-1, 1, 1, 10); err == nil {
		t.Error("With negative rows, Matrix() needs to return (_, nil)")
	}

	// this must fail, cols < 0
	if _, err := gen.Matrix(1, -1, 1, 10); err == nil {
		t.Error("With negative cols, Matrix() needs to return (_, nil)")
	}

	// this must fail, min > max
	if _, err := gen.Matrix(1, 1, 10, 1); err == nil {
		t.Error("With min > max, Matrix() needs to return (_, nil)")
	}
}

func TestGenerator_Bound(t *testing.T) {
	var tableTest = []struct {
		min int64
		max int64
	}{
		{1, 10},
		{1, 100},
		{1, 1000},
		{1, 10000},
		{-10, 1},
		{-100, 1},
		{-1000, 1},
		{-10000, 1},
		{-10, 10},
		{-100, 100},
		{-1000, 1000},
		{-10000, 10000},
	}
	var width int64 = 1
	var maxWidth int64 = 9

	gen := samples.NewGenerator()
	for ; width <= maxWidth; width++ {
		for _, tt := range tableTest {
			blow, bup, err := gen.Bound(tt.min, tt.max, width)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("bound generated: [%d, %d] [%d]", blow, bup, width)

			if blow < tt.min {
				t.Fatalf("generated blow is less then min: %d < %d", blow,
					tt.min)
			}
			if bup > tt.max {
				t.Fatalf("generated bup is greater then max: %d > %d", blow,
					tt.max)
			}
			if bup-blow != width {
				t.Fatalf("bound width is not equal as requested: %d != %d",
					bup-blow, width)
			}
		}
	}

}
