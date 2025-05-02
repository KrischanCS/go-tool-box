package iterator_test

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
)

func TestConcat(t *testing.T) {
	t.Parallel()

	type test struct {
		name   string
		inputs []iter.Seq[int]
		want   []int
	}

	tests := []test{
		{
			name: "Should yield values of all given iterators in order",
			inputs: []iter.Seq[int]{
				iterator.Of(1, 2, 3),
				iterator.Of(4, 5, 6),
				iterator.Of(7, 8, 9),
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Should yield same values as given iterator if only one is given",
			inputs: []iter.Seq[int]{
				iterator.Of(1, 2, 3),
			},
			want: []int{1, 2, 3},
		},
		{
			name:   "Should yield nothing if no iterators are given",
			inputs: []iter.Seq[int]{},
			want:   []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]int, 0, 16)

			// Act
			for v := range iterator.Concat(tc.inputs...) {
				got = append(got, v)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestConcat_mustStopOnBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]int, 0, 16)
	inputs := []iter.Seq[int]{
		iterator.Of(1, 2, 3),
		iterator.Of(4, 5, 6),
		iterator.Of(7, 8, 9),
	}

	// Act
	for v := range iterator.Concat(inputs...) {
		if v == 5 {
			break
		}

		got = append(got, v)
	}

	// Assert
	want := []int{1, 2, 3, 4}
	assert.Equal(t, want, got)
}
