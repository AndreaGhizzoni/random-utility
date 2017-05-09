package samples

// TODO add doc
type Bound struct {
	lower, upper int64
}

// TODO add doc
func NewBound(bLow, bUp int64) (*Bound, error) {
	if err := checkArgs(bLow, bUp); err != nil {
		return nil, err
	}

	return &Bound{
		lower: bLow,
		upper: bUp,
	}, nil
}

func checkArgs(bLow, bUp int64) error {

}

// TODO add doc
func (b *Bound) Width() int64 {
	return b.upper - b.lower
}
