package out

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func failif(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestWrite(t *testing.T) {
	path := "text.out"
	slice := []int64{1, 1, 2, 3, 5, 8, 13, 21}

	if err := Write(slice, path); err != nil {
		t.Fatal(err)
	}

	fileStat, err := os.Stat(path)
	failif(t, err)

	// checking file name
	if fileStat.Name() != path {
		failif(t,
			fmt.Errorf("file name mismatch: %s != %s", fileStat.Name(), path),
		)
	}

	// checking file size
	if fileStat.Size() == 0 {
		failif(t, fmt.Errorf("file size == 0"))
	}

	// Now check the actual content
	file, err1 := os.Open(path)
	failif(t, err1)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	sliceLen, err := strconv.ParseInt(scanner.Text(), 10, 64)
	failif(t, err)

	result := make([]int64, sliceLen)
	i := 0
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 64)
		failif(t, err)
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
}
