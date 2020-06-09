package sample

import (
	"math"
	"math/rand"
	"sort"
)

type Reservoir struct {
	seen, size int
	samples []float64
	min, max float64

	// welford's
	m, s float64
}

func NewReservoir(size int) Reservoir {
	return Reservoir{
		size: size,
		samples: make([]float64, 0, size),
	}
}

func (r *Reservoir) Reset() {
	r.seen = 0
	r.size = 0
	r.samples = r.samples[:0]
	r.m = 0.0
	r.s = 0.0
}

func (r *Reservoir) Sample(v float64) {
	if r.seen == 0 {
		r.min = v
		r.max = v
	} else {
		if v < r.min {
			r.min =v
		}
		if v > r.max {
			r.max = v
		}
	}
	r.seen++

	oldM := r.m
	kf := float64(r.seen)
	r.m = r.m + (v-r.m)/kf
	r.s = r.s + (v-r.m)*(v-oldM)

	if len(r.samples) < cap(r.samples) {
		r.samples = append(r.samples, v)
	} else {
		idx := rand.Intn(len(r.samples))
		if rand.Intn(r.seen) == 1 {
			r.samples[idx] = v;
		}
	}
}

func (r Reservoir) Samples() int {
	return len(r.samples)
}

func (r Reservoir) Seen() int {
	return r.seen
}

func (r Reservoir) Minimum() float64 {
	if r.seen == 0 {
		return math.NaN()
	}
	return r.min
}

func (r Reservoir) Maximum() float64 {
	if r.seen == 0 {
		return math.NaN()
	}
	return r.max
}

func (r Reservoir) Median() float64 {
	n := len(r.samples)
	if n == 0 {
		return math.NaN()
	}

	sort.Slice(r.samples, func (i, j int) bool {
		return r.samples[i] < r.samples[j]
	})
	if n % 2 == 0 {
		return (r.samples[(n/2)-1] + r.samples[(n/2)]) / 2
	}
	return r.samples[(n/2)]
}

func (r Reservoir) Stdev() float64 {
	if len(r.samples) < 2 {
		return math.NaN()
	}

	return math.Sqrt(r.s/(float64(r.seen)-1))
}
