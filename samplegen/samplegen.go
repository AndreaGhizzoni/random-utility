package samplegen

import (
	"math/rand"
	"time"
)

type SampleGenerator struct {
	rand *rand.Rand // or seed ?
	seed int
}

func New() *SampleGenerator {
	s := rand.NewSource(time.Now().UnixNano())
	return &SampleGenerator{
		rand: rand.New(s),
	}
}

func (s *SampleGenerator) Slice(len int64, min, max int64) ([]int64, error) {
	return nil, nil
}

func (s *SampleGenerator) Matrix(rows, cols int64, min, max int64) ([][]int64, error) {
	return nil, nil
}

func (s *SampleGenerator) Bound(min, max int64, width int64) (int64, int64) {
	return -1, -1
}

func (s *SampleGenerator) OrderedSlice(len int64, min, max int64) ([]int64, error) {
	return nil, nil
}
