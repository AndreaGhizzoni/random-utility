package out_test

import (
	"bufio"
	"github.com/AndreaGhizzoni/zenium/out"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// utility method
func failIf(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// TODO description
func TestNewPrinter(t *testing.T) {
	// TODO
}

// this function tests the correct behavior of out.NewPrinter method with wrong
// inputs
func TestNewPrinter_Arguments(t *testing.T) {
	// list of wrong paths
	var paths = []string{
		"/text.out",
		"/../text.out",
		"/.text.out",
		"/../.text.out",
	}

	for _, p := range paths {
		t.Logf("trying path: %s", p)

		abs, err := filepath.Abs(p)
		t.Logf("abs: %s %v", abs, err)
		dir, file1 := filepath.Split(abs)
		t.Logf("dir, file: %s %s", dir, file1)

		if _, err := out.NewPrinter(p); err == nil {
			t.Fatalf("out.NewPrinter(%s) must fail.", dir+file1)
		}

		t.Logf("Ok, can't open %s", dir+file1)
	}
}

// this function tests the correct behavior of out.Printer.WriteSlice method
// with correct inputs
func TestPrinter_WriteSlice(t *testing.T) {
	tD := "_test"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path  string
		slice []int64
	}{
		{tD + "/text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{tD + "/text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{tD + "/text.out", []int64{0000000, 00000000}},
		{tD + "/text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{tD + "/text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},

		{tD + "/.text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{tD + "/.text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{tD + "/.text.out", []int64{0000000, 00000000}},
		{tD + "/.text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{tD + "/.text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},

		{tD + "/dir/.text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{tD + "/dir/.text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{tD + "/dir/.text.out", []int64{0000000, 00000000}},
		{tD + "/dir/.text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{tD + "/dir/.text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)

		t.Logf("trying path: %s", tt.path)
		abs, err := filepath.Abs(tt.path)
		t.Logf("abs: %s %v", abs, err)
		dir, f := filepath.Split(abs)
		t.Logf("dir, file: %s %s", dir, f)

		printer, err := out.NewPrinter(tt.path)
		failIf(t, err)

		// try to write
		if err := printer.WriteSlice(tt.slice); err != nil {
			t.Fatal(err)
		}

		// get the stat from file already written
		fileStat, err := os.Stat(tt.path)
		failIf(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(f) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), f)
		}

		// checking file size
		if fileStat.Size() == 0 {
			t.Fatal("file already written has size == 0")
		}

		// open new file and check if slice in it is equal to the slice that
		// I have.
		file, err := os.Open(tt.path)
		failIf(t, err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		sliceLen, err := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, err)

		// read int64 from file
		sliceFromFile := make([]int64, sliceLen)
		i := 0
		for scanner.Scan() {
			x, err := strconv.ParseInt(scanner.Text(), 10, 64)
			failIf(t, err)
			sliceFromFile[i] = x
			i += 1
		}

		// check if while read there war an error
		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}

		// check lengths of sliceFromFile
		t.Logf("len(sliceFromFile)=%d", len(sliceFromFile))
		t.Log(sliceFromFile)
		if len(sliceFromFile) != len(tt.slice) {
			t.Fatalf("len(outRead) != len(tt.slice) : %d, %d",
				len(sliceFromFile),
				len(tt.slice))
		}

		// check if every elements from slice's file is equal to input slice
		for i, e := range sliceFromFile {
			if e != tt.slice[i] {
				t.Fatalf("Element from file %d != element from in slice %d",
					e, tt.path[i])
			}
		}

		// closing  file. not used defer because I'm in a loop.
		file.Close()
	}
}

// TODO add docs
func TestPrinter_WriteSlice_Arguments(t *testing.T) {
	// TODO
}

// this function tests the correct behavior of out.Printer.WriteMatrix method
// with correct inputs
func TestPrinter_WriteMatrix(t *testing.T) {
	tD := "_test"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path   string
		matrix [][]int64
	}{
		{tD + "/text.out", [][]int64{{1, 1, 2, 3}, {5, 8, 13, 21}}},
		{tD + "/text.out", [][]int64{{0, 0, 0, 0}, {0, 0, 0, 0}}},
		{tD + "/text.out", [][]int64{{0000000}, {00000000}}},
		{tD + "/text.out", [][]int64{{-1, -1, -2, -3}, {-5, -8, -13, -21}}},
		{tD + "/text.out", [][]int64{{1, 1, -2, 3}, {-5, 8, -13, 21}}},

		{tD + "/.text.out", [][]int64{{1, 1, 2, 3}, {5, 8, 13, 21}}},
		{tD + "/.text.out", [][]int64{{0, 0, 0, 0}, {0, 0, 0, 0}}},
		{tD + "/.text.out", [][]int64{{0000000}, {00000000}}},
		{tD + "/.text.out", [][]int64{{-1, -1, -2, -3}, {-5, -8, -13, -21}}},
		{tD + "/.text.out", [][]int64{{1, 1, -2, 3}, {-5, 8, -13, 21}}},

		{tD + "/dir/.text.out", [][]int64{{1, 1, 2, 3}, {5, 8, 13, 21}}},
		{tD + "/dir/.text.out", [][]int64{{0, 0, 0, 0}, {0, 0, 0, 0, 0}}},
		{tD + "/dir/.text.out", [][]int64{{0000000}, {00000000}}},
		{tD + "/dir/.text.out", [][]int64{{-1, -1, -2, -3}, {-5, -8, -13, -21}}},
		{tD + "/dir/.text.out", [][]int64{{1, 1, -2, 3}, {-5, 8, -13, 21}}},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)

		t.Logf("trying path: %s", tt.path)
		abs, err := filepath.Abs(tt.path)
		t.Logf("abs: %s %v", abs, err)
		dir, f := filepath.Split(abs)
		t.Logf("dir, file: %s %s", dir, f)

		printer, err := out.NewPrinter(tt.path)
		failIf(t, err)

		// try to write
		if err := printer.WriteMatrix(tt.matrix); err != nil {
			t.Fatal(err)
		}

		// get the stat from file already written
		fileStat, err := os.Stat(tt.path)
		failIf(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(f) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), f)
		}

		// checking file size
		if fileStat.Size() == 0 {
			t.Fatal("file already written has size == 0")
		}

		// open new file and check if matrix in it is equal to the matrix that
		// I have.
		file, err := os.Open(tt.path)
		failIf(t, err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		// read the first int64 that represents the matrix rows
		scanner.Scan()
		matrixRows, errR := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, errR)
		t.Logf("Rows read from file: %d", matrixRows)
		// read the second int64 that represents the matrix cols
		scanner.Scan()
		matrixCols, errC := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, errC)
		t.Logf("Cols read from file: %d", matrixCols)

		// read matrix elements as int64 from file
		matrixFromFile := make([][]int64, matrixRows)
		var i, j int64 = 0, 0
		for ; i < matrixRows; i++ {
			matrixFromFile[i] = make([]int64, matrixCols)
			for ; j < matrixCols; j++ {
				scanner.Scan()
				x, err := strconv.ParseInt(scanner.Text(), 10, 64)
				failIf(t, err)
				matrixFromFile[i][j] = x
			}
			j = 0
		}

		// check if while read there war an error
		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}

		// check lengths of matrixFromFile
		t.Logf("# of rows: len(matrixFromFile)=%d", len(matrixFromFile))
		t.Log(matrixFromFile)
		if len(matrixFromFile) != len(tt.matrix) {
			t.Fatalf("len(matrixFromFile) != len(tt.matrix) : %d, %d",
				len(matrixFromFile),
				len(tt.matrix))
		}

		// check if every elements from matrix's file is equal to input matrix
		for i, row := range matrixFromFile {
			for j, e := range row {
				if e != tt.matrix[i][j] {
					t.Fatalf("Element from file %d != element from given "+
						"matrix %d", e, tt.matrix[i][j])
				}
			}
		}

		// closing  file. not used defer because I'm in a loop.
		file.Close()
	}
}

// TODO add docs
func TestPrinter_WriteMatrix_Arguments(t *testing.T) {
	// TODO
}
