package out

import (
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	slices := [][]int64{
		{1, 2, 3, 4},
		{10, 9, 8, 7},
		{-1, -2, -3, -4},
		{-10, -9, -8, -7},
		{1, -2, 3, -4},
		{-10, -9, 8, 7},
	}

	for _, e := range slices {
		byteSlice, err := convert(e)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("% x", byteSlice)
	}
}

func TestArgumentsCheckRW(t *testing.T) {
	noRead := "no-read.txt"
	nr, err := os.Create(noRead)
	if err != nil {
		t.Fatal(err)
	}
	nr.Chmod(0222) // this sets permissions to --w--w--w-
	defer os.Remove(noRead)

	if err := checkRW(noRead); err == nil {
		t.Fatalf("check no read %v", err)
	}

	noWrite := "no-write.txt"
	nw, err := os.Create(noWrite)
	if err != nil {
		t.Fatal(err)
	}
	nw.Chmod(0444) // this sets permissions to -r--r--r--
	defer os.Remove(noWrite)

	if err := checkRW(noWrite); err == nil {
		t.Fatalf("check no write %v", err)
	}
}
