package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"github.com/AndreaGhizzoni/zenium/structures"
	"math/big"
	"testing"
)

var (
	modZero = big.NewInt(0)
	one     = big.NewInt(1)
	two     = big.NewInt(2)
)

// TODO add description
func Test_SGenerator_Int(t *testing.T) {
	// reusable pointer to big.Int
	var min, max *big.Int = nil, nil
	// power indicates the power of 2 that is used to generate the upper and
	// lower bounds
	var power *big.Int = nil
	var max_power = big.NewInt(64)

	generate := samples.NewSecureGenerator()
	for ; power.Cmp(max_power) == -1; power.Add(power, one) {
		max = big.NewInt(2).Exp(two, power, modZero)
		min = big.NewInt(2).Exp(two, power, modZero)
		min.Neg(min)

		t.Logf("Try to generate secure random number between: min= %v, max= %v",
			min, max)

		random, err := generate.Int(min, max)
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Logf("Secure random number generated: %v", random)

		if random.Cmp(min) == -1 { // random < min
			t.Fatalf("Generated number is less then min: %v < %v", random, min)
		}
		if random.Cmp(max) == 1 { // random > max
			t.Fatalf("Generated number is greater then max: %v > %v", random,
				max)
		}
	}
}

// TODO add description
func TestSGenerator_Slice(t *testing.T) {
	var lengths = []*big.Int{
		big.NewInt(1),
		big.NewInt(10),
		big.NewInt(100),
		big.NewInt(1000),
	}

	// reusable pointer to big.Int
	var min, max *big.Int = nil, nil
	// power indicates the power of 2 that is used to generate the upper and
	// lower bounds
	var power *big.Int = nil
	var max_power = big.NewInt(64)

	generate := samples.NewSecureGenerator()
	for _, length := range lengths {
		power = big.NewInt(1)
		for ; power.Cmp(max_power) == -1; power.Add(power, one) {
			max = big.NewInt(2).Exp(two, power, modZero)
			min = big.NewInt(2).Exp(two, power, modZero)
			min.Neg(min)

			t.Logf("Try to generate random secure slice with: length= %v, "+
				"min= %v, max= %v", length, min, max)

			slice, err := generate.Slice(length, min, max)
			if err != nil {
				t.Fatal(err.Error())
			}

			actualSliceLength := int64(len(slice))
			if length.Cmp(big.NewInt(actualSliceLength)) != 0 {
				t.Fatalf("Generated slice length mismatch: %v != %v", length,
					actualSliceLength)
			}

			for _, element := range slice {
				if element.Cmp(min) == -1 { // element < min
					t.Fatalf("number in slice is less then min: %v < %v",
						element, min)
				}
				if element.Cmp(max) == 1 { // element > max
					t.Fatalf("number in slice is greater then max: %v > %v",
						element, max)
				}
			}
		}
	}
}

// TODO add description
func TestSGenerator_Matrix(t *testing.T) {
	var matrixDimensions = []struct {
		r, c *big.Int
	}{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(10), big.NewInt(10)},
	}

	// reusable pointer to big.Int
	var min, max *big.Int = nil, nil
	// power indicates the power of 2 that is used to generate the upper and
	// lower bounds
	var power *big.Int = nil
	var max_power = big.NewInt(64)

	generate := samples.NewSecureGenerator()
	for _, dimension := range matrixDimensions {
		power = big.NewInt(1)
		for ; power.Cmp(max_power) == -1; power.Add(power, one) {
			max = big.NewInt(2).Exp(two, power, modZero)
			min = big.NewInt(2).Exp(two, power, modZero)
			min.Neg(min)

			t.Logf("Try to generate random secure matrix with: r= %v, c= %v, "+
				"min= %v, max= %v", dimension.r, dimension.c, min, max)

			matrix, err := generate.Matrix(dimension.r, dimension.c, min, max)
			if err != nil {
				t.Fatal(err.Error())
			}

			rows := int64(len(matrix))
			if dimension.r.Cmp(big.NewInt(rows)) != 0 {
				t.Fatalf("Generated matrix rows mismatch: %v != %v",
					dimension.r, rows)
			}

			cols := int64(len(matrix[0]))
			if dimension.c.Cmp(big.NewInt(cols)) != 0 {
				t.Fatalf("Generated matrix cols mismatch: %v != %v",
					dimension.c, cols)
			}

			for _, r := range matrix {
				for _, c := range r {
					if c.Cmp(min) == -1 { // c < min
						t.Fatalf("number in matrix is less then min: %v < "+
							"%v", c, min)
					}
					if c.Cmp(max) == 1 { // c > max
						t.Fatalf("number in matrix is greater then max: %v > "+
							"%v", c, max)
					}
				}
			}
		}
	}
}

// TODO add description
func TestSGenerator_Bound(t *testing.T) {
	boundWidth := big.NewInt(1000)
	var numberOfBounds = []*big.Int{
		big.NewInt(10),
		big.NewInt(100),
		big.NewInt(1000),
	}

	// reusable pointer to big.Int
	var min, max *big.Int = nil, nil
	// power indicates the power of 2 that is used to generate the upper and
	// lower bounds
	var power *big.Int = nil
	var max_power = big.NewInt(64)

	generate := samples.NewSecureGenerator()
	for _, amount := range numberOfBounds {
		// It starts at 10 because of fixed boundWidth
		power = big.NewInt(10)
		for ; power.Cmp(max_power) == -1; power.Add(power, one) {
			max = big.NewInt(2).Exp(two, power, modZero)
			min = big.NewInt(2).Exp(two, power, modZero)
			min.Neg(min)

			t.Logf("Try to generate random secure bound with: amount= %v, "+
				"w= %v, min= %v, max= %v", amount, boundWidth, min, max)

			bounds := []*structures.Bound{}
			i := big.NewInt(0)
			for ; i.Cmp(amount) == -1; i.Add(i, one) {
				b, err := generate.Bound(min, max, boundWidth)
				if err != nil {
					t.Fatal(err.Error())
				}
				bounds = append(bounds, b)
			}

			lenBounds := int64(len(bounds))
			if amount.Cmp(big.NewInt(lenBounds)) != 0 {
				t.Fatalf("Generated amount of bounds mismatch: %v != %v",
					amount, lenBounds)
			}

			for _, bound := range bounds {
				if bound.Low().Cmp(min) == -1 { // lower bound < min
					t.Fatalf("lower bound is less then min: %v < %v",
						bound.Low(), min)
				}
				if bound.Up().Cmp(max) == 1 { // upper bound > max
					t.Fatalf("upper bound is greater then max: %v > %v",
						bound.Up(), max)
				}
				if bound.Width().Cmp(boundWidth) != 0 { // width != boundWidth
					t.Fatalf("bound boundWidth mismatch: %v != %v",
						bound.Width(), boundWidth)
				}
			}
		}
	}
}
