package iterator

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestSlidingWindow(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name       string
		values     iter.Seq[T]
		windowSize int
		want       [][]int
	}

	testCases := []testCase[int]{
		{
			name:       "empty slice",
			values:     Of[int](),
			windowSize: 0,
			want:       [][]int{},
		},
		{
			name:       "one element, windowSize = 0",
			values:     Of(1),
			windowSize: 0,
			want:       [][]int{},
		},
		{
			name:       "windowSize < 0",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: -1,
			want:       [][]int{},
		},
		{
			name:       "one element, windowSize = 1",
			values:     Of(1),
			windowSize: 1,
			want:       [][]int{{1}},
		},
		{
			name:       "multiple elements, windowSize = 1",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 1,
			want:       [][]int{{1}, {2}, {3}, {4}, {5}, {6}},
		},
		{
			name:       "multiple elements, windowSize = 2",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 2,
			want: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
			},
		},
		{
			name:       "multiple elements, windowSize = 3",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 3,
			want: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
				{4, 5, 6},
			},
		},
		{
			name:       "multiple elements, windowSize = len(values)",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 6,
			want:       [][]int{{1, 2, 3, 4, 5, 6}},
		},
		{
			name:       "multiple elements, windowSize > len(values)",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 7,
			want:       [][]int{{1, 2, 3, 4, 5, 6}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([][]int, 0, len(tc.want))

			// Act
			for group := range SlidingWindow(tc.values, tc.windowSize) {
				got = append(got, slices.Clone(group))
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

//nolint:funlen
func TestFixedWindow(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name       string
		values     iter.Seq[T]
		windowSize int
		want       [][]int
	}

	testCases := []testCase[int]{
		{
			name:       "empty slice",
			values:     Of[int](),
			windowSize: 0,
			want:       [][]int{},
		},
		{
			name:       "one element, windowSize = 0",
			values:     Of(1),
			windowSize: 0,
			want:       [][]int{},
		},
		{
			name:       "windowSize < 0",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: -1,
			want:       [][]int{},
		},
		{
			name:       "one element, windowSize = 1",
			values:     Of(1),
			windowSize: 1,
			want:       [][]int{{1}},
		},
		{
			name:       "multiple elements, windowSize = 1",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 1,
			want:       [][]int{{1}, {2}, {3}, {4}, {5}, {6}},
		},
		{
			name:       "multiple elements, windowSize = 2",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 2,
			want:       [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:       "multiple elements, windowSize = 3",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 3,
			want:       [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name:       "multiple elements, windowSize = len(values)",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 6,
			want:       [][]int{{1, 2, 3, 4, 5, 6}},
		},
		{
			name:       "multiple elements, windowSize > len(values)",
			values:     Of(1, 2, 3, 4, 5, 6),
			windowSize: 7,
			want:       [][]int{{1, 2, 3, 4, 5, 6}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([][]int, 0, len(tc.want))

			// Act
			for ts := range FixedWindow(tc.values, tc.windowSize) {
				got = append(got, slices.Clone(ts))
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFixedWindow_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	values := Of(1, 2, 3, 4, 5, 6)
	windowSize := 2

	got := make([][]int, 0, windowSize)

	// Act
	for group := range FixedWindow(values, windowSize) {
		if slices.Equal(group, []int{5, 6}) {
			break
		}

		got = append(got, slices.Clone(group))
	}

	// Assert
	want := [][]int{{1, 2}, {3, 4}}
	assert.Equal(t, want, got)
}

func TestFixedWindow2Window_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	slice := []int{1, 2, 3, 4, 5, 6}
	values := PickRight(slices.All(slice))

	windowSize := 2

	got := make([][]int, 0, windowSize)

	// Act
	for group := range FixedWindow(values, windowSize) {
		if slices.Equal(group, []int{3, 4}) {
			break
		}

		got = append(got, slices.Clone(group))
	}

	// Assert
	want := [][]int{{1, 2}}
	assert.Equal(t, want, got)
}

func TestSlidingWindow_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	slice := []int{1, 2, 3, 4, 5, 6}
	values := PickRight(slices.All(slice))

	windowSize := 2

	got := make([][]int, 0, windowSize)

	// Act
	for group := range SlidingWindow(values, windowSize) {
		if slices.Equal(group, []int{4, 5}) {
			break
		}

		got = append(got, slices.Clone(group))
	}

	// Assert
	want := [][]int{{1, 2}, {2, 3}, {3, 4}}
	assert.Equal(t, want, got)
}
