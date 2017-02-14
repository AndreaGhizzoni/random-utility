package samplegen

import (
	"testing"
)

func TestCreateRandomSlice(t *testing.T) {
	var actualLength int64 = 10
	var min int64 = 10
	var max int64 = 100

	gen := New()
	randomSlice, err := gen.Slice(actualLength, min, max)
	t.Logf("random slice generated: %v", randomSlice)

	if err != nil {
		t.Fatal(err)
	}

	lenRandomSlice := int64(len(randomSlice))
	if lenRandomSlice != actualLength {
		t.Fatalf("length mismatch: %d != %d", lenRandomSlice, actualLength)
	}

	for i, e := range randomSlice {
		if e < min || e >= max {
			t.Fatalf("element randomSlice[%d]=%d is out of bound (%d, %d)",
				i, e, min, max)
		}
	}
}

func TestArgumentSlice(t *testing.T) {
	gen := New()

	_, err1 := gen.Slice(-1, 1, 10)
	if err1 == nil {
		t.Fatalf("With negative length, Slice() needs to return (_, nil)")
	}

	_, err2 := gen.Slice(10, 10, 1)
	if err2 == nil {
		t.Fatalf("With min > max as argument, Slice() needs to return (_, nil)")
	}
}

func BenchmarkGenerateSlice(b *testing.B) {
	gen := New()
	for i := 0; i < b.N; i++ {
		gen.Slice(100, 1, 10)
	}
}
