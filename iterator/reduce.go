package iterator

import "iter"

// Reduce takes an [iter.Seq] and applies fn on all yielded values
// consecutively. The result is collected in the given accumulator
func Reduce[IN, ACC any](input iter.Seq[IN], accumulator *ACC, fn func(*ACC, IN)) {
	for v := range input {
		fn(accumulator, v)
	}
}
