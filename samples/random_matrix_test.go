package samples

import "testing"

// This function tests a creation of random matrix of with correct input data.
func TestCreateRandomMatrix(t *testing.T) {
	var rows int64 = 10
	var cols int64 = 10
	var min int64 = -100
	var max int64 = 100

	testRandomMatrix(rows, cols, min, max, t)
}

// This function tests a creation of random matrix of positive numbers with
// correct input data.
func TestCreatePositiveRandomMatrix(t *testing.T) {
	var rows int64 = 10
	var cols int64 = 10
	var min int64 = 1
	var max int64 = 100

	testRandomMatrix(rows, cols, min, max, t)
}

// This function tests a creation of random matrix of negative numbers with
// correct input data.
func TestCreateNegativeRandomMatrix(t *testing.T) {
	var rows int64 = 10
	var cols int64 = 10
	var min int64 = -100
	var max int64 = -1

	testRandomMatrix(rows, cols, min, max, t)
}

// utility functions to create and check a generated matrix.
func testRandomMatrix(r, c, min, max int64, t *testing.T) {
	gen := New()
	matrix, err := gen.Matrix(r, c, min, max)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("random matrix generated: %v", matrix)

	// check rows of generated matrix
	rows := int64(len(matrix))
	if rows != r {
		t.Fatalf("rows length mismatch: %d != %d", rows, r)
	}

	// check cols of generated matrix
	cols := int64(len(matrix[0]))
	if cols != c {
		t.Fatalf("cols length mismatch: %d != %d", cols, c)
	}

	for i, slice := range matrix {
		for j, e := range slice {
			if e < min || e >= max {
				t.Fatalf("matrix[%d][%d]=%d is out of bound (%d, %d)",
					i, j, e, min, max)
			}
		}
	}
}

// This function tests the creation of random matrix with incorrect input data.
func TestArgumentMatrix(t *testing.T) {
	gen := New()

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
	gen := New()
	for i := 0; i < b.N; i++ {
		gen.Matrix(10, 10, -10, 10)
	}
}
