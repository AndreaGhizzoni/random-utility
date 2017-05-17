package structures

import (
	"fmt"
	"math/big"
)

// Bound structure to represents a bound between a lower and maximum.
type Bound struct {
	lower, upper *big.Int
}

// This function creates a new Bound structure with a lower and upper bound
// given.
func NewBound(bLow, bUp *big.Int) *Bound {
	// TODO check bounds
	return &Bound{
		lower: bLow,
		upper: bUp,
	}
}

// This function returns the width of the bound.
func (b *Bound) Width() *big.Int{
	return big.NewInt(0).Sub(b.upper, b.lower)
}

// This function returns the lower bound of the bound.
func (b *Bound) Low() *big.Int {
	return b.lower
}

// This function returns the upper bound of the bound.
func (b *Bound) Up()  *big.Int {
	return b.upper
}

// This function returns a string representation of Bound structure.
func (b *Bound) String() string {
	return fmt.Sprintf("(%v %v)", b.lower, b.upper)
}
