package iterator_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/tuple"
)

func ExampleZip() {
	numbers := iterator.Of(1, 2, 3, 4)
	letters := iterator.Of("a", "b", "c")

	for pair := range iterator.Zip[int, string](numbers, letters) {
		fmt.Println(pair.First(), pair.Second())
	}

	// Output:
	// 1 a
	// 2 b
	// 3 c
}

//nolint:funlen
func TestZip(t *testing.T) {
	t.Parallel()

	type testCase[L, R any] struct {
		name       string
		leftInput  iter.Seq[L]
		rightInput iter.Seq[R]
		want       []tuple.Pair[L, R]
	}

	testCases := []testCase[int, string]{
		{
			name:       "empty slices",
			leftInput:  iterator.Of[int](),
			rightInput: iterator.Of[string](),
			want:       []tuple.Pair[int, string]{},
		},
		{
			name:       "one element",
			leftInput:  iterator.Of(1),
			rightInput: iterator.Of("a"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
			},
		},
		{
			name:       "multiple elements",
			leftInput:  iterator.Of(1, 2, 3),
			rightInput: iterator.Of("a", "b", "c"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
				tuple.PairOf(2, "b"),
				tuple.PairOf(3, "c"),
			},
		},
		{
			name:       "len(leftInput) < len(rightInput)",
			leftInput:  iterator.Of(1, 2, 3),
			rightInput: iterator.Of("a", "b"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
				tuple.PairOf(2, "b"),
			},
		},
		{
			name:       "len(leftInput) > len(rightInput)",
			leftInput:  iterator.Of(1, 2),
			rightInput: iterator.Of("a", "b", "c"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
				tuple.PairOf(2, "b"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]tuple.Pair[int, string], 0, len(tc.want))

			// Act
			for p := range iterator.Zip(tc.leftInput, tc.rightInput) {
				got = append(got, p)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestZip_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	l := iterator.Of(1, 2, 3, 4, 5, 6)

	r := iterator.Of("a", "b", "c", "d", "e", "f")

	got := make([]tuple.Pair[int, string], 0, 2)

	stop := tuple.PairOf[int, string](3, "c")

	// Act
	for e := range iterator.Zip(l, r) {
		if e == stop {
			break
		}

		got = append(got, e)
	}

	// Assert
	want := []tuple.Pair[int, string]{
		tuple.PairOf(1, "a"),
		tuple.PairOf(2, "b"),
	}
	assert.Equal(t, want, got)
}
