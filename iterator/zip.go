package iterator

import (
	"iter"

	"github.com/KrischanCS/go-toolbox/tuple"
)

// Zip create a new [iter.Seq] from the left and right iterators, which yields
// pairs of values from both.
//
// The resulting iterator will stop when the shorter of the two iterators stops.
func Zip[L, R any](left iter.Seq[L], right iter.Seq[R]) iter.Seq[tuple.Pair[L, R]] {
	return func(yield func(tuple.Pair[L, R]) bool) {
		valuesRight, stop := iter.Pull(right)
		defer stop()

		for valueL := range left {
			valueR, ok := valuesRight()
			if !ok {
				return
			}

			if !yield(tuple.PairOf[L, R](valueL, valueR)) {
				return
			}
		}
	}
}
