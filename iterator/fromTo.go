package iterator

import (
	"iter"

	"github.com/KrischanCS/go-toolbox/constraints"
)

// FromTo creates an iterator returning the values from start to end exclusive
// (Half open range).
func FromTo(start, endExcluded int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range endExcluded - start {
			if !yield(start + i) {
				return
			}
		}
	}
}

// FromToInclusive creates an iterator returning the values from start to end
// inclusive (Closed range).
func FromToInclusive(start, endIncluded int) iter.Seq[int] {
	return FromTo(start, endIncluded+1)
}

// FromStepTo creates an iterator returning the values from start to end,
// where the value is increased by step every call and end will not be included,
// even if it is met exactly.
//
// If step is 0, it panics.
//
// If step is the wrong sign (never reaching end), the sign is inverted.
func FromStepTo[T constraints.RealNumber](start, step, endExcluded T) iter.Seq[T] {
	switch {
	case step == 0:
		panic("step must not be 0")
	case endExcluded < start:
		return backwardsFromStepTo(start, step, endExcluded)
	case step < 0:
		step = -step
	}

	return func(yield func(T) bool) {
		for v := start; v < endExcluded; v += step {
			if !yield(v) {
				return
			}
		}
	}
}

func backwardsFromStepTo[T constraints.RealNumber](start T, step T, endExcluded T) iter.Seq[T] {
	if step > 0 {
		step = -step
	}

	return func(yield func(T) bool) {
		for v := start; v > endExcluded; v += step {
			if !yield(v) {
				return
			}
		}
	}
}
