package iterator

import "iter"

// Enumerate takes an [iter.Seq] and creates an [iter.Seq2] which yields an index
// starting at 0 together with the values yielded by the given [iter.Seq].
func Enumerate[T any](input iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for t := range input {
			if !yield(i, t) {
				return
			}

			i++
		}
	}
}
