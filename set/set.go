// Package set provides an implementation of a set, meaning a collection of
// unique, unordered values.
//
// It implements basic collection options as add, remove, contains, len and
// clear as well as common set operations like union, intersection and
// difference.
//
// The implementation is based on normal go map, thus is not thread-safe.
package set

import (
	"fmt"
	"iter"
	"sort"
)

//nolint:gochecknoglobals
var placeholder = struct{}{}

// Set implements a collection of unique, unordered values.
//
// It is not thread-safe.
type Set[T comparable] struct {
	m map[T]struct{}
}

// Of creates a new set with the given values.
func Of[T comparable](values ...T) Set[T] {
	s := Set[T]{m: make(map[T]struct{}, len(values))}

	for _, v := range values {
		s.m[v] = placeholder
	}

	return s
}

// WithCapacity creates a Set with the given capacity.
func WithCapacity[T comparable](capacity int) Set[T] {
	return Set[T]{m: make(map[T]struct{}, capacity)}
}

// Add adds the given value to the set if it is not already present.
//
// It returns true if the value was newly added, and false if it was already
// present.
func (s Set[T]) Add(value T) bool {
	if _, ok := s.m[value]; ok {
		return false
	}

	s.m[value] = placeholder

	return true
}

// Remove removes the given value from the set.
//
// It returns true if the value existed and false if not.
func (s Set[T]) Remove(value T) bool {
	if _, ok := s.m[value]; !ok {
		return false
	}

	delete(s.m, value)

	return true
}

// Clear removes all values from the set.
func (s Set[T]) Clear() {
	clear(s.m)
}

// Values returns a slice of all values in the set without any particular
// order.
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s.m))

	for v := range s.m {
		values = append(values, v)
	}

	return values
}

// All creates an iterator over all values in the set without any particular
// order.
func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

// Len returns the number of values in the set.
func (s Set[T]) Len() int {
	return len(s.m)
}

// IsEmpty returns true if the set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s.m) == 0
}

// Clone creates a shallow copy of the set.
func (s Set[T]) Clone() Set[T] {
	return Of[T](s.Values()...)
}

// String returns a string representation in the format:
//   - If Present: "(Set[{{type}}]: [{{value 1}} {{value 2}} ...])"
//   - If Empty: "(Set[{{type}}]: <empty>)"
//
// The values are sorted by their string representation for easier overview,
// the actual set is not sorted.
func (s Set[T]) String() string {
	if s.IsEmpty() {
		return fmt.Sprintf("(Set[%T]: <empty>)", *new(T))
	}

	values := make([]string, 0, len(s.m))
	for v := range s.m {
		values = append(values, fmt.Sprintf("%v", v))
	}

	sort.Strings(values)

	return fmt.Sprintf("(Set[%T]: %s)", *new(T), values)
}
