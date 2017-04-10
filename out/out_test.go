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

// this function tests the correct behavior of out.WriteSlice method with
// not corrected inputs
func TestWriteArgs(t *testing.T) {
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

		if err := out.WriteSlice([]int64{}, p); err == nil {
			t.Fatalf("WriteSlice method with %s must fail.", dir+file1)
		}

		t.Logf("Ok, can't open %s", dir+file1)
	}
}

// this function tests the correct behavior of out.WriteSlice method with correct
// inputs
func TestWrite(t *testing.T) {
	testsDir := "_test"
	defer os.RemoveAll(testsDir)

	var tableTest = []struct {
		path  string
		slice []int64
	}{
		{testsDir + "/text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{testsDir + "/text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{testsDir + "/text.out", []int64{0000000, 00000000}},
		{testsDir + "/text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{testsDir + "/text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},

		{testsDir + "/.text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{testsDir + "/.text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{testsDir + "/.text.out", []int64{0000000, 00000000}},
		{testsDir + "/.text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{testsDir + "/.text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},

		{testsDir + "/dir/.text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{testsDir + "/dir/.text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{testsDir + "/dir/.text.out", []int64{0000000, 00000000}},
		{testsDir + "/dir/.text.out", []int64{-1, -1, -2, -3, -5, -8, -13, -21}},
		{testsDir + "/dir/.text.out", []int64{1, 1, -2, 3, -5, 8, -13, 21}},
	}

	for i, tt := range tableTest {
		// this is necessary to create a dynamic file name
		tt.path += "." + strconv.Itoa(i)

		t.Logf("trying path: %s", tt.path)
		abs, err := filepath.Abs(tt.path)
		t.Logf("abs: %s %v", abs, err)
		dir, f := filepath.Split(abs)
		t.Logf("dir, file: %s %s", dir, f)

		// try to write
		if err := out.WriteSlice(tt.slice, tt.path); err != nil {
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
