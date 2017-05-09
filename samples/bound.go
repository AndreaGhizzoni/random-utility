package samples

import "fmt"

// Bound structure to represents a bound between a lower and maximum.
type Bound struct {
	lower, upper int64
}

// This function creates a new Bound structure with a lower and upper bound
// given.
func NewBound(bLow, bUp int64) *Bound {
	return &Bound{
		lower: bLow,
		upper: bUp,
	}
}

// This function returns the width of the bound.
func (b *Bound) Width() int64 {
	return b.upper - b.lower
}

// This function returns the lower bound of the bound.
func (b *Bound) Low() int64 {
	return b.lower
}

// This function returns the upper bound of the bound.
func (b *Bound) Up() int64 {
	return b.upper
}

// This function returns a string representation of Bound structure.
func (b *Bound) String() string {
	return fmt.Sprintf("(%d %d)", b.lower, b.upper)
}
