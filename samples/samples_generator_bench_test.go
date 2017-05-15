package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"testing"
)

// benchmark random int64 generator
func BenchmarkGenerate_Int64(b *testing.B) {
	gen := samples.NewGenerator()
	for i := 0; i < b.N; i++ {
		gen.Int64(10, 100)
	}
}

// benchmark random slice generator
func BenchmarkGenerate_Slice(b *testing.B) {
	gen := samples.NewGenerator()
	for i := 0; i < b.N; i++ {
		gen.Slice(100, 1, 10)
	}
}

// benchmark random matrix generator
func BenchmarkGenerate_Matrix(b *testing.B) {
	gen := samples.NewGenerator()
	for i := 0; i < b.N; i++ {
		gen.Matrix(10, 10, -10, 10)
	}
}

// benchmark random matrix generator
func BenchmarkGenerate_Bounds(b *testing.B) {
	gen := samples.NewGenerator()
	for i := 0; i < b.N; i++ {
        gen.Bound(-10, 10, 3)
	}
}
