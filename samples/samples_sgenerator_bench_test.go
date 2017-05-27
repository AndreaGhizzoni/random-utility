package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"math/big"
	"testing"
)

var min, max = big.NewInt(-1000), big.NewInt(1000)
var length = big.NewInt(1000)
var width, amount = big.NewInt(50), big.NewInt(10000)
var rows, columns = big.NewInt(50), big.NewInt(50)
var generator, _ = samples.NewSecureGenerator(min, max)

func BenchmarkSGenerator_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generator.Int()
	}
}

func BenchmarkSGenerator_Slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generator.Slice(length)
	}
}

func BenchmarkSGenerator_Matrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generator.Matrix(rows, columns)
	}
}

func BenchmarkSGenerator_Bounds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generator.Bounds(width, amount)
	}
}
