package iterator_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/iterator/reducer"
)

func ExampleReduce() {
	i := iterator.Of(1, 2, 3, 4, 5)

	var sum int

	iterator.Reduce(i, &sum, reducer.Sum)

	fmt.Println(sum)

	// Output: 15
}

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
			input:       iterator.Of[int](),
			accumulator: 0,
			fn:          reducer.Sum[int],
			want:        0,
		},
		{
			name:        "Should not modify the accumulator if the fn does nothing",
			input:       iterator.Of(1, 2, 3, 4, 5),
			accumulator: 0,
			fn:          func(_ *int, _ int) {},
			want:        0,
		},
		{
			name:        "Should set acc to value if the input is one value, fn is the sum function and the accumulator is 0",
			input:       iterator.Of(1),
			accumulator: 0,
			fn:          reducer.Sum[int],
			want:        1,
		},
		{
			name:        "Should set the accumulator to the sum of all values in input",
			input:       iterator.Of(1, 2, 3, 4, 5),
			accumulator: 0,
			fn:          reducer.Sum[int],
			want:        15,
		},
		{
			name: "Should set the accumulator to the sum of all values" +
				" plus initial values if accumulator has an initial value",
			input:       iterator.Of(1, 2, 3, 4, 5),
			accumulator: 5,
			fn:          reducer.Sum[int],
			want:        20,
		},
		{
			name:        "Should set the accumulator to the product of all values in input",
			input:       iterator.Of(1, 2, 3, 4, 5),
			accumulator: 1,
			fn:          reducer.Product[int],
			want:        120,
		},
		{
			name:        "Should not modify the accumulator if it's initial value is the neutral element for fn",
			input:       iterator.Of(1, 2, 3, 4, 5),
			accumulator: 0,
			fn:          reducer.Product[int],
			want:        0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			iterator.Reduce(tt.input, &tt.accumulator, tt.fn)

			// Assert
			assert.Equal(t, tt.want, tt.accumulator)
		})
	}
}
