package iterator

import (
	"iter"

	"github.com/KrischanCS/go-toolbox/tuple"
)

// PickLeft takes a [iter.Seq2] and returns a [iter.Seq] of the left/first
// elements of the given sequence.
func PickLeft[L, UNUSED any](seq iter.Seq2[L, UNUSED]) iter.Seq[L] {
	return func(yield func(L) bool) {
		for l := range seq {
			if !yield(l) {
				return
			}
		}
	}
}

// PickRight takes a [iter.Seq2] and returns a [iter.Seq] of the right/second
// elements of the given sequence.
func PickRight[R, UNUSED any](seq iter.Seq2[UNUSED, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for _, r := range seq {
			if !yield(r) {
				return
			}
		}
	}
}

// Combine takes a [iter.Seq2] and returns a [iter.Seq] combining both return values
// of the given sequence into a [Pair].
func Combine[L, R any](seq iter.Seq2[L, R]) iter.Seq[tuple.Pair[L, R]] {
	return func(yield func(tuple.Pair[L, R]) bool) {
		for l, r := range seq {
			if !yield(tuple.PairOf[L, R](l, r)) {
				return
			}
		}
	}
}
