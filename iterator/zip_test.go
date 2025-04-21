package iterator

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestZip(t *testing.T) {
	t.Parallel()

	type testCase[L, R any] struct {
		name       string
		leftInput  iter.Seq[L]
		rightInput iter.Seq[R]
		want       []Pair[L, R]
	}

	testCases := []testCase[int, string]{
		{
			name:       "empty slices",
			leftInput:  Of[int](),
			rightInput: Of[string](),
			want:       []Pair[int, string]{},
		},
		{
			name:       "one element",
			leftInput:  Of(1),
			rightInput: Of("a"),
			want: []Pair[int, string]{
				{1, "a"},
			},
		},
		{
			name:       "multiple elements",
			leftInput:  Of(1, 2, 3),
			rightInput: Of("a", "b", "c"),
			want: []Pair[int, string]{
				{1, "a"},
				{2, "b"},
				{3, "c"},
			},
		},
		{
			name:       "len(leftInput) < len(rightInput)",
			leftInput:  Of(1, 2, 3),
			rightInput: Of("a", "b"),
			want: []Pair[int, string]{
				{1, "a"},
				{2, "b"},
			},
		},
		{
			name:       "len(leftInput) > len(rightInput)",
			leftInput:  Of(1, 2),
			rightInput: Of("a", "b", "c"),
			want: []Pair[int, string]{
				{1, "a"},
				{2, "b"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]Pair[int, string], 0, len(tc.want))

			// Act
			for p := range Zip(tc.leftInput, tc.rightInput) {
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
	l := Of(1, 2, 3, 4, 5, 6)

	r := Of("a", "b", "c", "d", "e", "f")

	got := make([]Pair[int, string], 0, 2)

	stop := Pair[int, string]{3, "c"}

	// Act
	for e := range Zip(l, r) {
		if e == stop {
			break
		}

		got = append(got, e)
	}

	// Assert
	want := []Pair[int, string]{
		{1, "a"},
		{2, "b"},
	}
	assert.Equal(t, want, got)
}
