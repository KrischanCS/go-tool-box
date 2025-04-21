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

	testCases := []testCase[any, any]{
		{
			name:   "empty slice",
			values: Of[any](),
			fn:     func(i any) any { return i },
			want:   []any{},
		},
		{
			name:   "one value",
			values: Of[any](1),
			fn:     func(i any) any { return strconv.Itoa(i.(int)) },
			want:   []any{"1"},
		},
		{
			name:   "multiple values",
			values: Of[any](1, 2, 3, 4, 5),
			fn:     func(i any) any { return i.(int) * 7 },
			want:   []any{7, 14, 21, 28, 35},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]any, 0, len(tc.want))

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
	want := []int{2, 4, 6}
	got := make([]int, 0, len(want))

	// Act
	for v := range Map(values, fn) {
		if v > 6 {
			break
		}
		got = append(got, v)
	}

	// Assert
	assert.Equal(t, want, got)
}
