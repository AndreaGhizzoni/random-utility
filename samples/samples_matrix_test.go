package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"testing"
)

func TestCreationMatrix(t *testing.T) {
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

// This function tests the creation of random matrix with incorrect input data.
func TestArgumentsMatrix(t *testing.T) {
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

// benchmark random matrix generator
func BenchmarkGenerateMatrix(b *testing.B) {
	gen := samples.NewGenerator()
	for i := 0; i < b.N; i++ {
		gen.Matrix(10, 10, -10, 10)
	}
}
