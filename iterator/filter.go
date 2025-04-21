package iterator

import "iter"

// Filter creates a [iter.Seq] that only yields values where conditions returns
// true.
func Filter[T any](s iter.Seq[T], condition func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s {
			if !condition(e) {
				continue
			}

			if !yield(e) {
				return
			}
		}
	}
}
