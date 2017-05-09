package samples

import "fmt"

// TODO add doc
type Bound struct {
	lower, upper int64
}

// TODO add doc
func NewBound(bLow, bUp int64) *Bound {
	return &Bound{
		lower: bLow,
		upper: bUp,
	}
}

// TODO add doc
func (b *Bound) Width() int64 {
	return b.upper - b.lower
}

func (b *Bound) Low() int64 {
	return b.lower
}

func (b *Bound) Up() int64 {
	return b.upper
}

// TODO add doc
func (b *Bound) String() string {
	return fmt.Sprintf("(%d %d)", b.lower, b.upper)
}
