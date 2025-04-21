package iterator

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name      string
		input     []T
		condition func(T) bool
		want      []T
	}

	testCases := []testCase[int]{
		{
			name:      "empty slice",
			input:     []int{},
			condition: nil,
			want:      []int{},
		},
		{
			name:  "all elements pass",
			input: []int{1, 2, 3, 4, 5},
			condition: func(_ int) bool {
				return true
			},

			want: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "no elements pass",
			input: []int{1, 2, 3, 4, 5},
			condition: func(_ int) bool {
				return false
			},
			want: []int{},
		},
		{
			name:  "some elements pass",
			input: []int{1, 2, 3, 4, 5},
			condition: func(e int) bool {
				return e%2 == 0
			},
			want: []int{2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]int, 0, len(tc.want))

			// Act
			for e := range Filter(PickRight(slices.All(tc.input)), tc.condition) {
				got = append(got, e)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFilter_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	s := []int{1, 2, 3, 4, 5, 6}
	got := make([]int, 0, len(s))

	// Act
	for e := range Filter(PickRight(slices.All(s)), func(_ int) bool {
		return true
	}) {
		if e == 3 {
			break
		}

		got = append(got, e)
	}

	// Assert
	want := []int{1, 2}
	assert.Equal(t, want, got)
}
