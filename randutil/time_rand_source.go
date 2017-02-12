package randutil

import (
	"math/rand"
	"sync"
	"time"
)

// locked to prevent concurrent use of the underlying source
type TimeSeed struct {
	lock sync.Mutex // protects src
	src  rand.Source
}

func NewTimeSeed() *TimeSeed {
	return &TimeSeed{
		src: rand.NewSource(time.Now().UnixNano()),
	}
}

// to satisfy rand.Source interface
func (r *TimeSeed) Int63() int64 {
	r.lock.Lock()
	ret := r.src.Int63()
	r.lock.Unlock()
	return ret
}

// to satisfy rand.Source interface
func (r *TimeSeed) Seed(seed int64) {
	r.lock.Lock()
	r.src.Seed(seed)
	r.lock.Unlock()
}
