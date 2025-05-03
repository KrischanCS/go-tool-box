package tuple

import "fmt"

// Triple is a tuple combining three values.
type Triple[T1, T2, T3 any] struct {
	first  T1
	second T2
	third  T3
}

// TripleOf creates a new [Triple] from the given values.
func TripleOf[T1, T2, T3 any](first T1, second T2, third T3) Triple[T1, T2, T3] {
	return Triple[T1, T2, T3]{first, second, third}
}

// First returns the first value of the triple.
func (t Triple[T1, T2, T3]) First() T1 {
	return t.first
}

// Second returns the second value of the triple.
func (t Triple[T1, T2, T3]) Second() T2 {
	return t.second
}

// Third returns the third value of the triple.
func (t Triple[T1, T2, T3]) Third() T3 {
	return t.third
}

// Unpack returns the values of the triple.
func (t Triple[T1, T2, T3]) Unpack() (T1, T2, T3) {
	return t.first, t.second, t.third
}

// String returns a string representation of the triple in the format:
//
//	(Triple[{{type1}}, {{type2}}, {{type3}}]: [{{first}}; {{second}}; {{third}}])
func (t Triple[T1, T2, T3]) String() string {
	return fmt.Sprintf("(Triple[%T, %T, %T]: [%v; %v; %v])",
		t.first, t.second, t.third,
		t.first, t.second, t.third,
	)
}
