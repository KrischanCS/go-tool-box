package iterator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/tuple"
)

func ExampleEnumerate() {
	it := iterator.Of("a", "b", "c")

	for i, v := range iterator.Enumerate(it) {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Output:
	// 0: a
	// 1: b
	// 2: c
}

func TestEnumerate(t *testing.T) {
	t.Parallel()

	// Arrange

	input := iterator.Of("a", "b", "c", "d", "e")
	got := make([]tuple.Pair[int, string], 0, 5)

	// Act
	for i, v := range iterator.Enumerate(input) {
		got = append(got, tuple.PairOf[int, string](i, v))
	}

	// Assert
	want := []tuple.Pair[int, string]{
		tuple.PairOf(0, "a"),
		tuple.PairOf(1, "b"),
		tuple.PairOf(2, "c"),
		tuple.PairOf(3, "d"),
		tuple.PairOf(4, "e"),
	}

	assert.Equal(t, want, got)
}

func TestEnumerate_shouldStopOnBreak(t *testing.T) {
	t.Parallel()

	// Arrange

	input := iterator.Of("a", "b", "c", "d", "e")
	got := make([]tuple.Pair[int, string], 0, 3)

	// Act
	for i, v := range iterator.Enumerate(input) {
		if i == 3 {
			break
		}

		got = append(got, tuple.PairOf[int, string](i, v))
	}

	// Assert
	want := []tuple.Pair[int, string]{
		tuple.PairOf(0, "a"),
		tuple.PairOf(1, "b"),
		tuple.PairOf(2, "c"),
	}

	assert.Equal(t, want, got)
}
