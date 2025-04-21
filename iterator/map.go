package iterator

import "iter"

// Map applies the given fn to each value from the given [iter.Seq]
// and yields the result.
func Map[IN, OUT any](in iter.Seq[IN], fn func(IN) OUT) iter.Seq[OUT] {
	return func(yield func(OUT) bool) {
		for v := range in {
			if !yield(fn(v)) {
				return
			}
		}
	}
}
