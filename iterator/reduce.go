package iterator

import "iter"

// Reduce takes an [iter.Seq] and applies fn on all yielded values
// consecutively. The result is collected in the given accumulator.
func Reduce[IN, ACC any](input iter.Seq[IN], accumulator *ACC, fn Reducer[ACC, IN]) {
	for v := range input {
		fn(accumulator, v)
	}
}

// Reducer is the function signature for the reducer function.
// Each call must take in and apply its operation to the accumulator.
type Reducer[ACC, IN any] func(accumulator *ACC, value IN)
