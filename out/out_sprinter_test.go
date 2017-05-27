package out_test

import (
	"bufio"
	"github.com/AndreaGhizzoni/zenium/out"
	"github.com/AndreaGhizzoni/zenium/samples"
	"github.com/AndreaGhizzoni/zenium/structures"
	"github.com/AndreaGhizzoni/zenium/util"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

var (
	min, max      = big.NewInt(-100), big.NewInt(100)
	rows, columns = big.NewInt(50), big.NewInt(50)
	length        = big.NewInt(10)
	width, amount = big.NewInt(20), big.NewInt(1000)
)

// utility method
func fatalIfErrorNotNil(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// utility method to log file path
func logFullPath(t *testing.T, p string) {
	t.Logf("trying path: %s", p)
	abs, err := filepath.Abs(p)
	t.Logf("abs: %s %v", abs, err)
	dir, file := filepath.Split(abs)
	t.Logf("dir, file: %s %s", dir, file)
}

func createSlice(t *testing.T) []*big.Int {
	generator, err := samples.NewGenerator(min, max)
	fatalIfErrorNotNil(t, err)
	slice, err := generator.Slice(length)
	fatalIfErrorNotNil(t, err)
	return slice
}

func createMatrix(t *testing.T) [][]*big.Int {
	generator, err := samples.NewGenerator(min, max)
	fatalIfErrorNotNil(t, err)
	matrix, err := generator.Matrix(rows, columns)
	fatalIfErrorNotNil(t, err)
	return matrix
}

func createBounds(t *testing.T) []*structures.Bound {
	generator, err := samples.NewGenerator(min, max)
	fatalIfErrorNotNil(t, err)
	bounds, err := generator.Bounds(width, amount)
	fatalIfErrorNotNil(t, err)
	return bounds
}

// this function tests the correct behavior of out.NewPrinter method with
// correct inputs
func TestNewSecurePrinter(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var paths = []string{
		tD + "text.out",
		tD + "some/folder/../.text.out",
		tD + "some/folder/../text.out",
		"", // write to stdout
	}

	for _, path := range paths {
		logFullPath(t, path)

		if _, err := out.NewPrinter(path); err != nil {
			t.Fatalf("out.NewPrinter(%s) must not fail.", path)
		}

		t.Logf("Ok, can open %s", path)
	}
}

// this function tests the correct behavior of out.NewPrinter method with
// wrong inputs
func TestNewSecurePrinter_Arguments(t *testing.T) {
	var paths = []string{
		"/text.out",
		"/../text.out",
		"/.text.out",
		"/../.text.out",
	}

	for _, path := range paths {
		logFullPath(t, path)

		if _, err := out.NewPrinter(path); err == nil {
			t.Fatalf("out.NewPrinter(%s) must fail.", path)
		}

		t.Logf("Ok, can't open %s", path)
	}
}

// this function tests the correct behavior of out.Printer.WriteSlice method
// with correct inputs
func TestSPrinter_WriteSlice(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path  string
		slice []*big.Int
	}{
		{tD + "text.out", createSlice(t)},
		{tD + "text.out", createSlice(t)},
		{tD + "text.out", createSlice(t)},
		{tD + "text.out", createSlice(t)},
		{tD + "text.out", []*big.Int{big.NewInt(0), big.NewInt(0)}},
		{tD + "text.out", []*big.Int{}},

		{tD + ".text.out", createSlice(t)},
		{tD + ".text.out", createSlice(t)},
		{tD + ".text.out", createSlice(t)},
		{tD + ".text.out", createSlice(t)},
		{tD + ".text.out", []*big.Int{big.NewInt(0), big.NewInt(0)}},
		{tD + ".text.out", []*big.Int{}},

		{tD + "dir/.text.out", createSlice(t)},
		{tD + "dir/.text.out", createSlice(t)},
		{tD + "dir/.text.out", createSlice(t)},
		{tD + "dir/.text.out", createSlice(t)},
		{tD + "dir/.text.out", []*big.Int{big.NewInt(0), big.NewInt(0)}},
		{tD + "dir/.text.out", []*big.Int{}},
	}

	for index, test := range tableTest {
		// this is necessary to create a dynamic file name
		test.path += "." + strconv.Itoa(index)
		logFullPath(t, test.path)

		printer, err := out.NewPrinter(test.path)
		fatalIfErrorNotNil(t, err)
		if err := printer.WriteSlice(test.slice); err != nil {
			t.Fatalf(err.Error())
		}

		// get the stat from file already written
		fileStat, err := os.Stat(test.path)
		fatalIfErrorNotNil(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(test.path) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), test.path)
		}

		// checking file size
		if fileStat.Size() == 0 {
			t.Fatal("file already written has size == 0")
		}

		// open new file and check if slice in it is equal to the slice that
		// I have.
		file, err := os.Open(test.path)
		fatalIfErrorNotNil(t, err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		sliceLength, err := util.FromStringToBigInt(scanner.Text(), "Slice Len")
		fatalIfErrorNotNil(t, err)

		sliceFromFile := []*big.Int{}
		if sliceLength.Cmp(big.NewInt(0)) != 0 {
			for scanner.Scan() {
				sliceElementFromFile, err := util.FromStringToBigInt(
					scanner.Text(),
					"Slice Element",
				)
				fatalIfErrorNotNil(t, err)
				sliceFromFile = append(sliceFromFile, sliceElementFromFile)
			}
		}

		// check if while read there war an error
		fatalIfErrorNotNil(t, scanner.Err())

		// check lengths of sliceFromFile
		t.Logf("len(sliceFromFile)=%d", len(sliceFromFile))
		t.Logf("slice read: %v", sliceFromFile)
		if len(sliceFromFile) != len(test.slice) {
			t.Fatalf("len(sliceFromFile) != len(test.slice) : %d, %d",
				len(sliceFromFile),
				len(test.slice))
		}

		// check if every elements from slice's file is equal to input slice
		for index, element := range sliceFromFile {
			if element.Cmp(test.slice[index]) != 0 {
				t.Fatalf("Element from file %d != element from in slice %d",
					element, test.path[index])
			}
		}

		// closing  file. not used defer because I'm in a loop.
		file.Close()
	}
}

// this function tests the correct behavior of out.Printer.WriteSlice method
// with wrong inputs
func TestSPrinter_WriteSlice_Arguments(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var sliceWithNil = []*big.Int{big.NewInt(0), nil}
	var tableTest = []struct {
		path  string
		slice []*big.Int
	}{
		{tD + "text.out", sliceWithNil},
		{tD + ".text.out", sliceWithNil},
		{tD + "dir/.text.out", sliceWithNil},
		{"", sliceWithNil}, // try to print on console
	}

	for _, test := range tableTest {
		logFullPath(t, test.path)

		printer, err := out.NewPrinter(test.path)
		fatalIfErrorNotNil(t, err)

		if err := printer.WriteSlice(test.slice); err == nil {
			t.Fatalf("Printer.WriteSlice(%v) must fail", test.slice)
		}
	}
}

// this function tests the correct behavior of out.Printer.WriteMatrix method
// with correct inputs
func TestSPrinter_WriteMatrix(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path   string
		matrix [][]*big.Int
	}{
		{tD + "text.out", createMatrix(t)},
		{tD + "text.out", createMatrix(t)},
		{tD + "text.out", createMatrix(t)},
		{tD + "text.out", createMatrix(t)},
		{tD + "text.out", [][]*big.Int{{big.NewInt(0), big.NewInt(0)}}},
		{tD + "text.out", [][]*big.Int{{}}},

		{tD + ".text.out", createMatrix(t)},
		{tD + ".text.out", createMatrix(t)},
		{tD + ".text.out", createMatrix(t)},
		{tD + ".text.out", createMatrix(t)},
		{tD + ".text.out", [][]*big.Int{{big.NewInt(0), big.NewInt(0)}}},
		{tD + ".text.out", [][]*big.Int{{}}},

		{tD + "dir/.text.out", createMatrix(t)},
		{tD + "dir/.text.out", createMatrix(t)},
		{tD + "dir/.text.out", createMatrix(t)},
		{tD + "dir/.text.out", createMatrix(t)},
		{tD + "dir/.text.out", [][]*big.Int{{big.NewInt(0), big.NewInt(0)}}},
		{tD + "dir/.text.out", [][]*big.Int{{}}},
	}

	for index, test := range tableTest {
		// this is necessary to create a dynamic file name
		test.path += "." + strconv.Itoa(index)
		logFullPath(t, test.path)

		printer, err := out.NewPrinter(test.path)
		fatalIfErrorNotNil(t, err)
		if err := printer.WriteMatrix(test.matrix); err != nil {
			t.Fatalf(err.Error())
		}

		// get the stat from file already written
		fileStat, err := os.Stat(test.path)
		fatalIfErrorNotNil(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(test.path) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), test.path)
		}

		// checking file size
		if fileStat.Size() == 0 {
			t.Fatal("file already written has size == 0")
		}

		// open new file and check if matrix in it is equal to the matrix that
		// I have.
		file, err := os.Open(test.path)
		fatalIfErrorNotNil(t, err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		matrixRows, err := util.FromStringToBigInt(scanner.Text(), "Matrix Row")
		fatalIfErrorNotNil(t, err)
		t.Logf("Rows read from file: %v", matrixRows)
		scanner.Scan()
		matrixColumns, err := util.FromStringToBigInt(scanner.Text(), "Matrix"+
			" Columns")
		fatalIfErrorNotNil(t, err)
		t.Logf("Columns read from file: %v", matrixColumns)

		matrixFromFile := [][]*big.Int{}
		row, column := big.NewInt(0), big.NewInt(0)
		for ; row.Cmp(matrixRows) == -1; row.Add(row, util.One) {
			rowN := []*big.Int{}
			for ; column.Cmp(matrixColumns) == -1; column.Add(column, util.One) {
				scanner.Scan()
				elementFromFile, err := util.FromStringToBigInt(
					scanner.Text(), "matrix element",
				)
				fatalIfErrorNotNil(t, err)
				rowN = append(rowN, elementFromFile)
			}
			matrixFromFile = append(matrixFromFile, rowN)
			column = big.NewInt(0)
		}

		// check if while read there war an error
		fatalIfErrorNotNil(t, scanner.Err())

		// check lengths of matrixFromFile
		// TODO this not gonna always work because:
		// I have used big.Int as dimension to build the matrix, so len(matrix)
		// return a Int64 and can exceed that dimension.
		t.Logf("# of rows: len(matrixFromFile)=%d", len(matrixFromFile))
		t.Logf("matri read: %v", matrixFromFile)
		if len(matrixFromFile) != len(test.matrix) {
			t.Fatalf("len(matrixFromFile) != len(test.matrix) : %d, %d",
				len(matrixFromFile),
				len(test.matrix))
		}

		// check if every elements from matrix's file is equal to input matrix
		for rowIndex, row := range matrixFromFile {
			for columnIndex, element := range row {
				if element.Cmp(test.matrix[rowIndex][columnIndex]) != 0 {
					t.Fatalf("Element from file %v != element from given "+
						"matrix %v", element,
						test.matrix[rowIndex][columnIndex])
				}
			}
		}

		// closing  file. not used defer because I'm in a loop.
		file.Close()
	}
}

// this function tests the correct behavior of out.Printer.WriteMatrix method
// with wrong inputs
func TestSPrinter_WriteMatrix_Arguments(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var matrixWithNil = [][]*big.Int{{big.NewInt(0), nil}}
	var tableTest = []struct {
		path   string
		matrix [][]*big.Int
	}{
		{tD + "text.out", matrixWithNil},
		{tD + "text.out", [][]*big.Int{}},

		{tD + ".text.out", matrixWithNil},
		{tD + ".text.out", [][]*big.Int{}},

		{tD + "dir/.text.out", matrixWithNil},
		{tD + "dir/text.out", [][]*big.Int{}},

		{"", matrixWithNil}, // try to print on console
		{"", [][]*big.Int{}},
	}

	for _, test := range tableTest {
		logFullPath(t, test.path)

		printer, err := out.NewPrinter(test.path)
		fatalIfErrorNotNil(t, err)

		if err := printer.WriteMatrix(test.matrix); err == nil {
			t.Fatalf("Printer.WriteMatrix(%v) must fail", test.matrix)
		}
	}
}

// this function tests the correct behavior of out.Printer.WriteBounds method
// with correct inputs
func TestSPrinter_WriteBounds(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path   string
		bounds []*structures.Bound
	}{
		{tD + "text.out", createBounds(t)},
		{tD + "text.out", []*structures.Bound{}},

		{tD + ".text.out", createBounds(t)},
		{tD + ".text.out", []*structures.Bound{}},

		{tD + "dir/text.out", createBounds(t)},
		{tD + "dir/text.out", []*structures.Bound{}},
	}

	for index, test := range tableTest {
		// this is necessary to create a dynamic file name
		test.path += "." + strconv.Itoa(index)
		logFullPath(t, test.path)

		printer, err := out.NewPrinter(test.path)
		fatalIfErrorNotNil(t, err)
		if err := printer.WriteBounds(test.bounds); err != nil {
			t.Fatalf(err.Error())
		}

		// get the stat from file already written
		fileStat, err := os.Stat(test.path)
		fatalIfErrorNotNil(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(test.path) {
			t.Fatalf("file name mismatch: %s != %s", fileStat.Name(), test.path)
		}

		// checking file size
		if fileStat.Size() == 0 {
			t.Fatal("file already written has size == 0")
		}

		// open new file and check if bounds in it is equal to the bounds that
		// I have.
		file, err := os.Open(test.path)
		fatalIfErrorNotNil(t, err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		numberOfBounds, err := util.FromStringToBigInt(scanner.Text(),
			"# of bounds")
		fatalIfErrorNotNil(t, err)
		t.Logf("Bounds to read from file: %v", numberOfBounds)

		bounds := []*structures.Bound{}
		var i = big.NewInt(0)
		for ; i.Cmp(numberOfBounds) == -1; i.Add(i, util.One) {
			scanner.Scan()
			bLow, err := util.FromStringToBigInt(scanner.Text(),
				"lower bound")
			fatalIfErrorNotNil(t, err)

			scanner.Scan()
			bUp, err := util.FromStringToBigInt(scanner.Text(),
				"upper bound")
			fatalIfErrorNotNil(t, err)

			boundRead := structures.NewBound(bLow, bUp)
			t.Logf("bound read from file: %s", boundRead.String())
			bounds = append(bounds, boundRead)
		}

		// TODO len(test.bounds) can truncate the actual size of slice,
		// because is build using a big.Int
		lenBounds := big.NewInt(int64(len(test.bounds)))
		if numberOfBounds.Cmp(lenBounds) != 0 {
			t.Fatalf("Number of bounds into file != writen bounds: %v != %v",
				numberOfBounds, lenBounds,
			)
		}

		for index, boundFromFile := range bounds {
			if boundFromFile.Low().Cmp(test.bounds[index].Low()) != 0 {
				t.Fatalf("Lower bound from file != input lower bound: %v != %v",
					boundFromFile.Low(), test.bounds[index].Low())
			}

			if boundFromFile.Up().Cmp(test.bounds[index].Up()) != 0 {
				t.Fatalf("Upper bound from file != input upper bound: %v != %v",
					boundFromFile.Up(), test.bounds[index].Up())
			}
		}
	}
}

// this function tests the correct behavior of out.Printer.WriteBounds method
// with wrong inputs
func TestSPrinter_WriteBounds_Arguments(t *testing.T) {
	tD := "_test/"
	defer os.RemoveAll(tD)

	var tableTest = []struct {
		path   string
		bounds []*structures.Bound
	}{
		{tD + "text.out", nil},
		{tD + "text.out", []*structures.Bound{nil}},
	}

	for index, test := range tableTest {
		// this is necessary to create a dynamic file name
		test.path += "." + strconv.Itoa(index)
		logFullPath(t, test.path)

		printer, err := out.NewPrinter(test.path)
		fatalIfErrorNotNil(t, err)

		// try to write
		if err := printer.WriteBounds(test.bounds); err == nil {
			t.Fatalf("With bound: %v Printer.WriteBounds must fail",
				test.bounds)
		}
		t.Logf("Printer.WriteBounds(%v) failed -> OK.", test.bounds)
	}
}
