package iterator_test

import (
	"fmt"
	"iter"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
)

func ExampleMap() {
	i := iterator.Of(2, 4, 6)

	for squared := range iterator.Map(i, func(i int) int { return i * i }) {
		fmt.Println(squared)
	}

	// Output:
	// 4
	// 16
	// 36
}

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
			values: iterator.Of[int](),
			fn:     func(i int) string { return string(rune(i + 64)) },
			want:   []string{},
		},
		{
			name:   "one value",
			values: iterator.Of(1),
			fn:     strconv.Itoa,
			want:   []string{"1"},
		},
		{
			name:   "multiple values",
			values: iterator.Of(10, 20, 30, 40, 50),
			fn:     strconv.Itoa,
			want:   []string{"10", "20", "30", "40", "50"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]string, 0, len(tc.want))

			// Act
			for v := range iterator.Map(tc.values, tc.fn) {
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
	values := iterator.Of(1, 2, 3, 4, 5)
	fn := func(i int) int { return i * 2 }
	got := make([]int, 0, 3)

	// Act
	for v := range iterator.Map(values, fn) {
		if v > 6 {
			break
		}

		got = append(got, v)
	}

	// Assert
	want := []int{2, 4, 6}
	assert.Equal(t, want, got)
}
