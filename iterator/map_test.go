package iterator

import (
	"iter"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()

	type testCase[IN, OUT any] struct {
		name   string
		values iter.Seq[IN]
		fn     func(IN) OUT
		want   []OUT
	}

	testCases := []testCase[int, string]{
		{
			name:   "empty slice",
			values: Of[int](),
			fn:     func(i int) string { return string(rune(i + 64)) },
			want:   []string{},
		},
		{
			name:   "one value",
			values: Of(1),
			fn:     func(i int) string { return strconv.Itoa(i) },
			want:   []string{"1"},
		},
		{
			name:   "multiple values",
			values: Of(10, 20, 30, 40, 50),
			fn:     func(i int) string { return strconv.Itoa(i) },
			want:   []string{"10", "20", "30", "40", "50"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]string, 0, len(tc.want))

			// Act
			for v := range Map(tc.values, tc.fn) {
				got = append(got, v)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMap_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	values := Of(1, 2, 3, 4, 5)
	fn := func(i int) int { return i * 2 }
	got := make([]int, 0, 3)

	// Act
	for v := range Map(values, fn) {
		if v > 6 {
			break
		}

		got = append(got, v)
	}

	// Assert
	want := []int{2, 4, 6}
	assert.Equal(t, want, got)
}
