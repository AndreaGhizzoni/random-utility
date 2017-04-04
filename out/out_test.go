package out_test

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/AndreaGhizzoni/zenium/out"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func failIf(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestWrite(t *testing.T) {
	var tableTest = []struct {
		path  string
		slice []int64
	}{
		{"text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{"text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"text.out", []int64{0000000, 00000000}},

		{".text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{".text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{".text.out", []int64{0000000, 00000000}},

		{"dir/.text.out", []int64{1, 1, 2, 3, 5, 8, 13, 21}},
		{"dir/.text.out", []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"dir/.text.out", []int64{0000000, 00000000}},
	}

	for _, tt := range tableTest {
		t.Logf("trying path: %s", tt.path)

		abs, err := filepath.Abs(tt.path)
		t.Logf("abs: %s %v", abs, err)
		dir, file1 := filepath.Split(abs)
		t.Logf("dir, file: %s %s", dir, file1)

		if err := out.Write(tt.slice, tt.path); err != nil {
			t.Fatal(err)
		}

		fileStat, err := os.Stat(tt.path)
		failIf(t, err)

		// checking file name
		if fileStat.Name() != filepath.Base(tt.path) {
			t.Fatal(fmt.Errorf("file name mismatch: %s != %s",
				fileStat.Name(), tt.path))
		}

		// checking file size
		if fileStat.Size() == 0 {
			t.Fatal(errors.New("file size == 0"))
		}

		// Now check the actual content
		file, err := os.Open(tt.path)
		failIf(t, err)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		sliceLen, err := strconv.ParseInt(scanner.Text(), 10, 64)
		failIf(t, err)

		result := make([]int64, sliceLen)
		i := 0
		for scanner.Scan() {
			x, err := strconv.ParseInt(scanner.Text(), 10, 64)
			failIf(t, err)
			result[i] = x
			i += 1
		}

		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}

		t.Logf("len(result)=%d", len(result))
		for _, v := range result {
			t.Logf("%d ", v)
		}

		file.Close()
		if dir != "" {
			os.RemoveAll(dir)
		} else {
			os.Remove(file1)
		}
	}
}
