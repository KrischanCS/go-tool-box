package iterator

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestReduce(t *testing.T) {
	t.Parallel()

	type testCase[IN any, ACC any] struct {
		name        string
		input       iter.Seq[IN]
		accumulator ACC
		fn          func(*ACC, IN)
		want        ACC
	}

	tests := []testCase[int, int]{
		{
			name:        "Should not modify the accumulator if the input is empty",
			input:       Of[int](),
			accumulator: 0,
			fn: func(acc *int, v int) {
				*acc += v
			},
			want: 0,
		},
		{
			name:        "Should not modify the accumulator if the fn does nothing",
			input:       Of[int](1, 2, 3, 4, 5),
			accumulator: 0,
			fn:          func(_ *int, _ int) {},
			want:        0,
		},
		{
			name:        "Should set acc to value if the input is one value, fn is the sum function and the accumulator is 0",
			input:       Of[int](1),
			accumulator: 0,
			fn: func(acc *int, v int) {
				*acc += v
			},
			want: 1,
		},
		{
			name:        "Should set the accumulator to the sum of all values in input",
			input:       Of[int](1, 2, 3, 4, 5),
			accumulator: 0,
			fn: func(acc *int, v int) {
				*acc += v
			},
			want: 15,
		},
		{
			name: "Should set the accumulator to the sum of all values" +
				" plus initial values if accumulator has an initial value",
			input:       Of[int](1, 2, 3, 4, 5),
			accumulator: 5,
			fn: func(acc *int, v int) {
				*acc += v
			},
			want: 20,
		},
		{
			name:        "Should set the accumulator to the product of all values in input",
			input:       Of[int](1, 2, 3, 4, 5),
			accumulator: 1,
			fn: func(acc *int, v int) {
				*acc *= v
			},
			want: 120,
		},
		{
			name:        "Should not modify the accumulator if it's initial value is the neutral element for fn",
			input:       Of[int](1, 2, 3, 4, 5),
			accumulator: 0,
			fn: func(acc *int, v int) {
				*acc *= v
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			Reduce(tt.input, &tt.accumulator, tt.fn)

			// Assert
			assert.Equal(t, tt.want, tt.accumulator)
		})
	}
}
