package out

import "testing"

func TestConvert(t *testing.T) {
	slices := [][]int64{
		{1, 2, 3, 4},
		{10, 9, 8, 7},
	}

	for _, e := range slices {
		byteSlice, err := convert(e)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("% x", byteSlice)
	}
}
