package iterator_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
)

func isEven(i int) bool {
	return i%2 == 0
}

func ExampleFilter() {
	s := iterator.Of(1, 2, 3, 4, 5, 6)

	for e := range iterator.Filter(s, isEven) {
		fmt.Println(e)
	}

	// Output:
	// 2
	// 4
	// 6
}

func TestFilter(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name      string
		input     iter.Seq[T]
		condition func(T) bool
		want      []T
	}

	testCases := []testCase[int]{
		{
			name:      "empty slice",
			input:     iterator.Of[int](),
			condition: nil,
			want:      []int{},
		},
		{
			name:  "all elements pass",
			input: iterator.Of(1, 2, 3, 4, 5),
			condition: func(_ int) bool {
				return true
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "no elements pass",
			input: iterator.Of(1, 2, 3, 4, 5),
			condition: func(_ int) bool {
				return false
			},
			want: []int{},
		},
		{
			name:      "some elements pass",
			input:     iterator.Of(1, 2, 3, 4, 5),
			condition: isEven,
			want:      []int{2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]int, 0, len(tc.want))

			// Act
			for e := range iterator.Filter(tc.input, tc.condition) {
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
	i := iterator.Of(1, 2, 3, 4, 5, 6)
	got := make([]int, 0, 6)

	// Act
	for e := range iterator.Filter(i, func(_ int) bool {
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
