package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"math/big"
	"testing"
)

func actual_test_SGenerator_Int(t *testing.T, onlyPositive bool) {
	const max_pow = 64
	var pow int64 = 1
	var zero = big.NewInt(0)

	// reusable pointer to big.Int
	var min, max *big.Int = nil, nil
	if onlyPositive {
		min = big.NewInt(0)
	}else{
		max = big.NewInt(0)
	}

	gen := samples.NewSecureGenerator()
	// pow indicates the power that is used to generate the bounds
	for ; pow < max_pow; pow++ {
		d := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(pow), zero)
		// set max to 2^(pow)
		if onlyPositive {
			max = d
		}else{
			min = d.Neg(d)
		}

		t.Logf("Try to generate secure random number between: min=%v, max=%v",
			min, max)

		r, err := gen.Int(min, max)
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Logf("Secure random number generated: %v", r)

		if r.Cmp(min) == -1 { // r < tt.min
			t.Fatalf("Generated number is less then min: %v < %v", r, min)
		}
		if r.Cmp(max) == 1 { // r > tt.max
			t.Fatalf("Generated number is greater then max: %v > %v", r,
				max)
		}
	}
}

func TestSGenerator_Int(t *testing.T) {
    // test SGenerator.Int(0, 2**x) where x = {1,2, ..., 64}
    actual_test_SGenerator_Int(t, true)
	// test SGenerator.Int(-2**x, 0) where x = {1,2, ..., 64}
	actual_test_SGenerator_Int(t, false)
}
