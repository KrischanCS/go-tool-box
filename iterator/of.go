// Package iterator implements a some handy iterator functions
//
// The package is at the moment mostly build for trying out iterators and their
// characteristics.
//
// The file benchmark_test.go compares the usage of iterators with the
// equivalent implementation using for loops.
package iterator

import "iter"

// Of creates a [iter.Seq] of the given values.
func Of[T any](values ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}
