package iterator

import "iter"

// Unique yields each unique value of input once.
func Unique[T comparable](input iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		set := make(map[T]struct{})

		for v := range input {
			if _, ok := set[v]; ok {
				continue
			}

			if !yield(v) {
				return
			}

			set[v] = struct{}{}
		}
	}
}
