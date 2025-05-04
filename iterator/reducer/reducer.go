// Package reducer provides a few often used Reducers for [iterator.Reduce]
package reducer

import (
	"strings"

	"github.com/KrischanCS/go-toolbox/constraints"
	"github.com/KrischanCS/go-toolbox/iterator"
)

// Count increases acc on each call, effectively counting the number of inputs.
func Count[T any](acc *int, _ T) {
	*acc++
}

// Sum adds the given value to acc.
func Sum[T constraints.RealNumber](acc *T, in T) {
	*acc += in
}

// Product multiplies acc by input.
func Product[T constraints.RealNumber](acc *T, in T) {
	*acc *= in
}

// Join creates a [iterator.Reducer] which joins all given strings with the given separator.
func Join(separator string) iterator.Reducer[strings.Builder, string] {
	return func(acc *strings.Builder, in string) {
		if acc.Len() > 0 {
			acc.WriteString(separator)
		}

		acc.WriteString(in)
	}
}

// GroupBy creates a [iterator.Reducer] which groups the given inputs by the result of the given keyFunc.
func GroupBy[KEY comparable, IN any](keyFunc func(IN) KEY) iterator.Reducer[map[KEY][]IN, IN] {
	return func(m *map[KEY][]IN, in IN) {
		key := keyFunc(in)

		s, ok := (*m)[key]
		if !ok {
			//nolint:mnd
			s = make([]IN, 0, 8)
		}

		s = append(s, in)

		(*m)[key] = s
	}
}
