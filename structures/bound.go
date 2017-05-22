package structures

import (
	"fmt"
	"math/big"
)

// Bound structure to represents a bound between a lower and maximum.
type Bound struct {
	lower, upper *big.Int
}

// NewBound creates a new Bound structure with a lower and upper bound given.
func NewBound(bLow, bUp *big.Int) *Bound {
	return &Bound{
		lower: bLow,
		upper: bUp,
	}
}

// Width returns the width of the bound.
func (b *Bound) Width() *big.Int {
	return big.NewInt(0).Sub(b.upper, b.lower)
}

// Low returns the lower bound of the bound.
func (b *Bound) Low() *big.Int {
	return b.lower
}

// Up returns the upper bound of the bound.
func (b *Bound) Up() *big.Int {
	return b.upper
}

// String returns a string representation of Bound structure.
func (b *Bound) String() string {
	return fmt.Sprintf("(%v %v)", b.lower, b.upper)
}
