package out_test

import (
	"bufio"
	"github.com/AndreaGhizzoni/zenium/out"
	"github.com/AndreaGhizzoni/zenium/samples"
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

// utility method to log file path
func logPath(t *testing.T, p string) {
	t.Logf("trying path: %s", p)
	abs, err := filepath.Abs(p)
	t.Logf("abs: %s %v", abs, err)
	dir, file := filepath.Split(abs)
	t.Logf("dir, file: %s %s", dir, file)
}

// this function tests the correct behavior of out.NewPrinter method with
// correct inputs
func TestPrinter_NewPrinter(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	// list of paths
	var paths = []string{
		tD + "text.out",
		tD + "some/folder/../.text.out",
		tD + "some/folder/../text.out",
		"", // write to stdout
	}

	for _, p := range paths {
		logPath(t, p)

		if _, err := out.NewPrinter(p); err != nil {
			t.Fatalf("out.NewPrinter(%s) must not fail.", p)
		}

		t.Logf("Ok, can open %s", p)
	}
}

// this function tests the correct behavior of out.NewPrinter method with wrong
// inputs
func TestPrinter_NewPrinter_Arguments(t *testing.T) {
	// list of wrong paths
	var paths = []string{
		"/text.out",
		"/../text.out",
		"/.text.out",
		"/../.text.out",
	}

	for _, p := range paths {
		logPath(t, p)

		if _, err := out.NewPrinter(p); err == nil {
			t.Fatalf("out.NewPrinter(%s) must fail.", p)
		}

		t.Logf("Ok, can't open %s", p)
	}
}

// this function tests the correct behavior of out.Printer.WriteSlice method
// with correct inputs
func TestPrinter_WriteSlice(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path  string
		slice []int64
	}{
		{tD + "text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{tD + "text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{tD + "text.out", []int64{0000000, 00000000}},
		{tD + "text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{tD + "text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},
		{tD + "text.out", []int64{}},

		{tD + ".text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{tD + ".text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{tD + ".text.out", []int64{0000000, 00000000}},
		{tD + ".text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{tD + ".text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},
		{tD + ".text.out", []int64{}},

		{tD + "dir/.text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{tD + "dir/.text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{tD + "dir/.text.out", []int64{0000000, 00000000}},
		{tD + "dir/.text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{tD + "dir/.text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},
		{tD + "dir/.text.out", []int64{}},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)
		logPath(t, tt.path)

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
		if fileStat.Name() != filepath.Base(tt.path) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), tt.path)
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
		scanner.Scan()
		sliceLen, err := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, err)

		// read int64 from file
		sliceFromFile := make([]int64, sliceLen)
		if sliceLen != 0 {
			i := 0
			for scanner.Scan() {
				x, err := strconv.ParseInt(scanner.Text(), 10, 64)
				failIf(t, err)
				sliceFromFile[i] = x
				i += 1
			}
		}

		// check if while read there war an error
		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}

		// check lengths of sliceFromFile
		t.Logf("len(sliceFromFile)=%d", len(sliceFromFile))
		t.Log(sliceFromFile)
		if len(sliceFromFile) != len(tt.slice) {
			t.Fatalf("len(sliceFromFile) != len(tt.slice) : %d, %d",
				len(sliceFromFile),
				len(tt.slice))
		}

		// check if every elements from matrix's file is equal to input matrix
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

// this function tests the correct behavior of out.Printer.WriteSlice method
// with wrong inputs
func TestPrinter_WriteSlice_Arguments(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var nilSlice []int64
	var tableTest = []struct {
		path  string
		slice []int64
	}{
		{tD + "text.out", nilSlice},
		{tD + ".text.out", nilSlice},
		{tD + "dir/.text.out", nilSlice},
		{"", nilSlice},
	}

	for _, tt := range tableTest {
		logPath(t, tt.path)

		printer, err := out.NewPrinter(tt.path)
		failIf(t, err)

		if err := printer.WriteSlice(tt.slice); err == nil {
			t.Fatalf("printer.WriteSlice(%v) must fail", tt.slice)
		}
	}
}

// this function tests the correct behavior of out.Printer.WriteMatrix method
// with correct inputs
func TestPrinter_WriteMatrix(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path   string
		matrix [][]int64
	}{
		{tD + "text.out", [][]int64{{1, 1, 2, 3}, {5, 8, 13, 21}}},
		{tD + "text.out", [][]int64{{0, 0, 0, 0}, {0, 0, 0, 0}}},
		{tD + "text.out", [][]int64{{0000000}, {00000000}}},
		{tD + "text.out", [][]int64{{-1, -1, -2, -3}, {-5, -8, -13, -21}}},
		{tD + "text.out", [][]int64{{1, 1, -2, 3}, {-5, 8, -13, 21}}},
		{tD + "text.out", [][]int64{{}}},

		{tD + ".text.out", [][]int64{{1, 1, 2, 3}, {5, 8, 13, 21}}},
		{tD + ".text.out", [][]int64{{0, 0, 0, 0}, {0, 0, 0, 0}}},
		{tD + ".text.out", [][]int64{{0000000}, {00000000}}},
		{tD + ".text.out", [][]int64{{-1, -1, -2, -3}, {-5, -8, -13, -21}}},
		{tD + ".text.out", [][]int64{{1, 1, -2, 3}, {-5, 8, -13, 21}}},
		{tD + ".text.out", [][]int64{{}}},

		{tD + "dir/.text.out", [][]int64{{1, 1, 2, 3}, {5, 8, 13, 21}}},
		{tD + "dir/.text.out", [][]int64{{0, 0, 0, 0}, {0, 0, 0, 0, 0}}},
		{tD + "dir/.text.out", [][]int64{{0000000}, {00000000}}},
		{tD + "dir/.text.out", [][]int64{{-1, -1, -2, -3}, {-5, -8, -13, -21}}},
		{tD + "dir/.text.out", [][]int64{{1, 1, -2, 3}, {-5, 8, -13, 21}}},
		{tD + "dir/.text.out", [][]int64{{}}},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)
		logPath(t, tt.path)

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
		if fileStat.Name() != filepath.Base(tt.path) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), tt.path)
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

// this function tests the correct behavior of out.Printer.WriteMatrix method
// with wrong inputs
func TestPrinter_WriteMatrix_Arguments(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var nilMatrix [][]int64
	var tableTest = []struct {
		path   string
		matrix [][]int64
	}{
		{tD + "text.out", nilMatrix},
		{tD + "text.out", [][]int64{}},

		{tD + ".text.out", nilMatrix},
		{tD + ".text.out", [][]int64{}},

		{tD + "dir/.text.out", nilMatrix},
		{tD + "dir/text.out", [][]int64{}},

		{"", nilMatrix},
		{"", [][]int64{}},
	}

	for _, tt := range tableTest {
		logPath(t, tt.path)

		printer, err := out.NewPrinter(tt.path)
		failIf(t, err)

		if err := printer.WriteMatrix(tt.matrix); err == nil {
			t.Fatalf("printer.WriteSlice(%v) must fail", tt.matrix)
		}
	}
}

// this function tests the correct behavior of out.Printer.WriteBound method
// with correct inputs
func TestPrinter_WriteBound(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path  string
		bound *samples.Bound
	}{
		{tD + "text.out", samples.NewBound(1, 3)},
		{tD + "text.out", samples.NewBound(1, 6)},
		{tD + "text.out", samples.NewBound(-1, 2)},
		{tD + "text.out", samples.NewBound(-100, 6)},

		{tD + ".text.out", samples.NewBound(1, 3)},
		{tD + ".text.out", samples.NewBound(1, 6)},
		{tD + ".text.out", samples.NewBound(-1, 2)},
		{tD + ".text.out", samples.NewBound(-100, 6)},

		{tD + "dir/.text.out", samples.NewBound(1, 3)},
		{tD + "dir/.text.out", samples.NewBound(1, 6)},
		{tD + "dir/.text.out", samples.NewBound(-1, 2)},
		{tD + "dir/.text.out", samples.NewBound(-100, 6)},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)
		logPath(t, tt.path)

		printer, err := out.NewPrinter(tt.path)
		failIf(t, err)
		// try to write
		if err := printer.WriteBound(*tt.bound); err != nil {
			t.Fatal(err)
		}

		// get the stat from file already written
		fileStat, err := os.Stat(tt.path)
		failIf(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(tt.path) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), tt.path)
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

		// read the first int64 that represents the number of bounds
		scanner.Scan()
		numberOfBounds, errR := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, errR)
		t.Logf("Bounds to read from file: %d", numberOfBounds)

		if numberOfBounds != 1 {
			t.Fatal("Number of bounds into file != 1")
		}

		// read lower bound elements as int64 from file
		scanner.Scan()
		bLow, errBLow := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, errBLow)
		t.Logf("Lower bound to read from file: %d", bLow)
		if bLow != tt.bound.Low() {
			t.Fatalf("Lower bound from file != input lower bound: %d != %d",
				bLow, tt.bound.Low())
		}

		// read upper bound elements as int64 from file
		scanner.Scan()
		bUp, errBUp := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, errBUp)
		t.Logf("Upper bound to read from file: %d", bUp)
		if bUp != tt.bound.Up() {
			t.Fatalf("Upper bound from file != input upper bound: %d != %d",
				bUp, tt.bound.Up())
		}
	}
}

// this function tests the correct behavior of out.Printer.WriteBound method
// with wrong inputs
func TestPrinter_WriteBounds(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path   string
		bounds []samples.Bound
	}{
		{tD + "text.out", []samples.Bound{
			*samples.NewBound(1, 3),
			*samples.NewBound(1, 6),
			*samples.NewBound(-1, 2),
			*samples.NewBound(-100, 6),
		}},
		{tD + "text.out", []samples.Bound{}},

		{tD + ".text.out", []samples.Bound{
			*samples.NewBound(1, 3),
			*samples.NewBound(1, 6),
			*samples.NewBound(-1, 2),
			*samples.NewBound(-100, 6),
		}},
		{tD + ".text.out", []samples.Bound{}},

		{tD + "dir/.text.out", []samples.Bound{
			*samples.NewBound(1, 3),
			*samples.NewBound(1, 6),
			*samples.NewBound(-1, 2),
			*samples.NewBound(-100, 6),
		}},
		{tD + "dir/.text.out", []samples.Bound{}},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)
		logPath(t, tt.path)

		printer, err := out.NewPrinter(tt.path)
		failIf(t, err)
		// try to write
		if err := printer.WriteBounds(tt.bounds); err != nil {
			t.Fatal(err)
		}

		// get the stat from file already written
		fileStat, err := os.Stat(tt.path)
		failIf(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(tt.path) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), tt.path)
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

		// read the first int64 that represents the number of bounds
		scanner.Scan()
		numberOfBounds, errR := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, errR)
		t.Logf("Bounds to read from file: %d", numberOfBounds)

		if numberOfBounds != int64(len(tt.bounds)) {
			t.Fatal("Number of bounds into file != 1")
		}

		var i int64 = 0
		for ; i < numberOfBounds; i++ {
			// read lower bound elements as int64 from file
			scanner.Scan()
			bLow, errBLow := strconv.ParseInt(scanner.Text(), 10, 64)
			failIf(t, errBLow)
			t.Logf("Lower bound to read from file: %d", bLow)
			if bLow != tt.bounds[i].Low() {
				t.Fatalf("Lower bound from file != input lower bound: %d != %d",
					bLow, tt.bounds[i].Low())
			}

			// read upper bound elements as int64 from file
			scanner.Scan()
			bUp, errBUp := strconv.ParseInt(scanner.Text(), 10, 64)
			failIf(t, errBUp)
			t.Logf("Upper bound to read from file: %d", bUp)
			if bUp != tt.bounds[i].Up() {
				t.Fatalf("Upper bound from file != input upper bound: %d != %d",
					bUp, tt.bounds[i].Up())
			}
		}
	}
}

func TestPrinter_WriteBounds_Arguments(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path   string
		bounds []samples.Bound
	}{
		{tD + "text.out", nil},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)
		logPath(t, tt.path)

		printer, err := out.NewPrinter(tt.path)
		failIf(t, err)

		// try to write
		if err := printer.WriteBounds(tt.bounds); err == nil {
			t.Fatalf("With bound: %v printer.WriteBounds must fail", tt.bounds)
		}
		t.Logf("printer.WriteBounds(%v) failed -> OK.", tt.bounds)
	}
}
