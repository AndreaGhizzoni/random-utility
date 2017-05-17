package samples_test

import (
	"github.com/AndreaGhizzoni/zenium/samples"
	"github.com/AndreaGhizzoni/zenium/structures"
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
func TestSGenerator_Slice(t *testing.T) {
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

			t.Logf("Try to generate random secure slice with: l= %v, min= %v,"+
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

// TODO add description
func TestSGenerator_Matrix(t *testing.T) {
	var lengths = []struct {
		r, c *big.Int
	}{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(10), big.NewInt(10)},
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

			t.Logf("Try to generate random secure matrix with: r= %v, c= %v, "+
				"min= %v, max= %v", l.r, l.c, min, max)

			m, err := generate.Matrix(l.r, l.c, min, max)
			if err != nil {
				t.Fatal(err.Error())
			}

			rows := int64(len(m))
			if l.r.Cmp(big.NewInt(rows)) != 0 {
				t.Fatalf("Generated matrix rows mismatch: %v != %v", l.r, rows)
			}

			cols := int64(len(m[0]))
			if l.c.Cmp(big.NewInt(cols)) != 0 {
				t.Fatalf("Generated matrix cols mismatch: %v != %v", l.c, cols)
			}

			for _, r := range m {
				for _, c := range r {
					if c.Cmp(min) == -1 { // e < min
						t.Fatalf("number in matrix is less then min: %v < "+
							"%v", c, min)
					}
					if c.Cmp(max) == 1 { // e > max
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
	var tableTest = []struct {
		amount, width *big.Int
	}{
		{big.NewInt(10), big.NewInt(1000)},
		{big.NewInt(100), big.NewInt(1000)},
		{big.NewInt(1000), big.NewInt(1000)},
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
	for _, l := range tableTest {
		pow = big.NewInt(10) // power start to 10 because of fixed bound width
		for ; pow.Cmp(max_pow) == -1; pow.Add(pow, one) {
			max = big.NewInt(2).Exp(two, pow, zero)
			min = big.NewInt(2).Exp(two, pow, zero)
			min.Neg(min)

			t.Logf("Try to generate random secure bound with: a= %v, w= %v, "+
				"min= %v, max= %v", l.amount, l.width, min, max)

			bounds := []*structures.Bound{}
			i := big.NewInt(0)
			for ; i.Cmp(l.amount) == -1; i.Add(i, one) {
				b, err := generate.Bound(min, max, l.width)
				if err != nil {
					t.Fatal(err.Error())
				}
				bounds = append(bounds, b)
			}

			lenBounds := int64(len(bounds))
			if l.amount.Cmp(big.NewInt(lenBounds)) != 0 {
				t.Fatalf("Generated amount of bounds mismatch: %v != %v",
					l.amount, lenBounds)
			}

			for _, b := range bounds {
				if b.Low().Cmp(min) == -1 { // lower bound < min
					t.Fatalf("lower bound is less then min: %v < %v",
						b.Low(), min)
				}
				if b.Up().Cmp(max) == 1 { // upper bound > max
					t.Fatalf("upper bound is greater then max: %v > %v",
						b.Up(), max)
				}
				if b.Width().Cmp(l.width) != 0 { // bound width != width
					t.Fatalf("bound width mismatch: %v != %v",
						b.Width(), l.width)
				}
			}
		}
	}
}
