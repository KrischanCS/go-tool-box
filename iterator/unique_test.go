package iterator_test

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-tool-box/iterator"
)

func TestUnique(t *testing.T) {
	t.Parallel()

	type test struct {
		name  string
		input iter.Seq[int]
		want  []int
	}

	tests := []test{
		{
			name:  "Should yield all values if all are different",
			input: iterator.Of(1, 2, 3, 4, 5),
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Should yield only first value if all are the same",
			input: iterator.Of(1, 1, 1, 1, 1),
			want:  []int{1},
		},
		{
			name:  "Should yield first value of each different value",
			input: iterator.Of(1, 1, 1, 2, 2, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5),
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "Should yield all in the order of their first appearance",
			input: iterator.Of(1, 2, 1, 2, 2, 3, 1, 2, 3, 2, 1, 4, 4, 4, 1, 2, 1, 3, 4, 5, 3, 2, 1, 2, 3, 4, 1, 1, 1),
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Should yield nothing if input is empty",
			input: iterator.Of[int](),
			want:  []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]int, 0, 16)

			// Act
			for v := range iterator.Unique(tc.input) {
				got = append(got, v)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUnique_MustStopOnBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	input := iterator.Of(1, 1, 2, 2, 3, 4, 5)
	breakAt := 3
	got := make([]int, 0, 3)

	// Act

	for v := range iterator.Unique(input) {
		got = append(got, v)

		if v == breakAt {
			break
		}
	}

	// Assert
	want := []int{1, 2, 3}
	assert.Equal(t, want, got)
}
