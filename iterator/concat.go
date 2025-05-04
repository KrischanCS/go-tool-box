package iterator

import "iter"

// Concat creates an iterator which yields the values of all given iterators in order.
func Concat[T any](inputs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, input := range inputs {
			if !yieldAll(input, yield) {
				return
			}
		}
	}
}

func yieldAll[T any](input iter.Seq[T], yield func(T) bool) bool {
	for v := range input {
		if !yield(v) {
			return false
		}
	}

	return true
}
