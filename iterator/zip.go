package iterator

import "iter"

// Pair combines two arbitrary types in single struct.
type Pair[L, R any] struct {
	Left  L
	Right R
}

// Zip create a new [iter.Seq] from the left and right iterators, which yields
// pairs of values from both.
//
// The resulting iterator will stop when the shorter of the two iterators stops.
func Zip[L, R any](left iter.Seq[L], right iter.Seq[R]) iter.Seq[Pair[L, R]] {
	return func(yield func(Pair[L, R]) bool) {
		valuesRight, stop := iter.Pull(right)
		defer stop()

		for valueL := range left {
			valueR, ok := valuesRight()
			if !ok {
				return
			}

			if !yield(Pair[L, R]{valueL, valueR}) {
				return
			}
		}
	}
}
