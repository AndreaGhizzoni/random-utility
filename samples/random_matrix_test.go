package samples

import "testing"

func TestCreateRandomMatrix(t *testing.T) {
	var rows int64 = 10
	var cols int64 = 10
	var min int64 = -100
	var max int64 = 100

	testRandomMatrix(rows, cols, min, max, t)
}

func TestCreatePositiveRandomMatrix(t *testing.T) {
	var rows int64 = 10
	var cols int64 = 10
	var min int64 = 1
	var max int64 = 100

	testRandomMatrix(rows, cols, min, max, t)
}

func TestCreateNegativeRandomMatrix(t *testing.T) {
	var rows int64 = 10
	var cols int64 = 10
	var min int64 = -100
	var max int64 = -1

	testRandomMatrix(rows, cols, min, max, t)
}

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
