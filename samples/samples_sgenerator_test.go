package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"math/big"
	"testing"
)

// TODO add description
func Test_SGenerator_Int(t *testing.T) {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	// reusable pointer to big.Int
	var min, max *big.Int = nil, nil
	generate := samples.NewSecureGenerator()
	// pow indicates the power that is used to generate the bounds
	pow := big.NewInt(1)
	max_pow := big.NewInt(64)
	for ; pow.Cmp(max_pow) == -1; pow.Add(pow, one) {
		max = big.NewInt(2).Exp(two, pow, zero)
		min = big.NewInt(2).Exp(two, pow, zero)
		min.Neg(min)

		t.Logf("Try to generate secure random number between: min= %v, max= %v",
			min, max)

		r, err := generate.Int(min, max)
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Logf("Secure random number generated: %v", r)

		if r.Cmp(min) == -1 { // r < min
			t.Fatalf("Generated number is less then min: %v < %v", r, min)
		}
		if r.Cmp(max) == 1 { // r > max
			t.Fatalf("Generated number is greater then max: %v > %v", r, max)
		}
	}
}

// TODO add description
func TestSGenerator_Slice_(t *testing.T) {
	var lengths = []*big.Int{
		big.NewInt(1),
		big.NewInt(10),
		big.NewInt(100),
		big.NewInt(1000),
	}
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	// reusable pointer to big.Int
	var min, max *big.Int = nil, nil
	var pow *big.Int
	generate := samples.NewSecureGenerator()
	// pow indicates the power that is used to generate the bounds
	max_pow := big.NewInt(64)

	for _, l := range lengths {
		pow = big.NewInt(1)
		for ; pow.Cmp(max_pow) == -1; pow.Add(pow, one) {
			max = big.NewInt(2).Exp(two, pow, zero)
			min = big.NewInt(2).Exp(two, pow, zero)
			min.Neg(min)

			t.Logf("Try to generate random secure slice with: l= %v, min= %d,"+
				" max= %v", l, min, max)

			s, err := generate.Slice(l, min, max)
			if err != nil {
				t.Fatal(err.Error())
			}

			lenS := int64(len(s))
			if l.Cmp(big.NewInt(lenS)) != 0 {
				t.Fatalf("Generated slice length mismatch: %v != %v", l, lenS)
			}

			for _, e := range s {
				if e.Cmp(min) == -1 { // e < min
					t.Fatalf("number in slice is less then min: %v < %v", e,
						min)
				}
				if e.Cmp(max) == 1 { // e > max
					t.Fatalf("number in slice is greater then max: %v > %v", e,
						max)
				}
			}
		}
	}
}
