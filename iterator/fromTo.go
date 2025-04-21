package iterator

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// Number is a type constraint that matches all numeric types.
type Number interface {
	constraints.Integer | constraints.Float
}

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
// Panics in the following cases:
//   - step == 0
//   - start < end && step < 0
//   - start > end && step > 0
func FromStepTo[T Number](start, step, endExcluded T) iter.Seq[T] {
	return func(yield func(T) bool) {
		v := start
		for i := 0; v < endExcluded; i, v = i+1, v+step {
			if !yield(v) {
				return
			}
		}
	}
}
