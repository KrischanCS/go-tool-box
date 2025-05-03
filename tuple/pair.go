// Package tuple provides simple types for combining values.
package tuple

import "fmt"

// Pair is a tuple combining two values.
type Pair[T1, T2 any] struct {
	first  T1
	second T2
}

// PairOf creates a new [Pair] from the given values.
func PairOf[T1, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{first, second}
}

// First returns the first value of the pair.
func (p Pair[T1, T2]) First() T1 {
	return p.first
}

// Second returns the second value of the pair.
func (p Pair[T1, T2]) Second() T2 {
	return p.second
}

// Unpack returns the values of the pair.
func (p Pair[T1, T2]) Unpack() (T1, T2) {
	return p.first, p.second
}

// String returns a string representation of the pair in the format:
//
//	(Pair[{{type1}}, {{type2}}]: [{{first}}; {{second}}])
func (p Pair[T1, T2]) String() string {
	return fmt.Sprintf("(Pair[%T, %T]: [%v; %v])",
		p.first, p.second,
		p.first, p.second)
}
