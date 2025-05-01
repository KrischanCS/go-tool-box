package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFromTo() {
	for n := range FromTo(3, 7) {
		fmt.Println(n)
	}

	// Output:
	// 3
	// 4
	// 5
	// 6
}

func ExampleFromToInclusive() {
	for n := range FromToInclusive(3, 7) {
		fmt.Println(n)
	}

	// Output:
	// 3
	// 4
	// 5
	// 6
	// 7
}

func ExampleFromStepTo() {
	for n := range FromStepTo(1.0, 1.5, 7.0) {
		fmt.Println(n)
	}

	// Output:
	// 1
	// 2.5
	// 4
	// 5.5
}

func ExampleFromStepTo_backwards() {
	for n := range FromStepTo(7.0, -1.5, 1.0) {
		fmt.Println(n)
	}

	// Output:
	// 7
	// 5.5
	// 4
	// 2.5
}

func TestFromTo_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]int, 0)

	stop := 5

	// Act
	for v := range FromTo(1, 10) {
		if v == stop {
			break
		}

		got = append(got, v)
	}

	// Assert
	want := []int{1, 2, 3, 4}
	assert.Equal(t, want, got)
}

func TestFromToInclusive_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]int, 0)

	stop := 5

	// Act
	for v := range FromToInclusive(1, 10) {
		if v == stop {
			break
		}

		got = append(got, v)
	}

	// Assert
	want := []int{1, 2, 3, 4}
	assert.Equal(t, want, got)
}

//nolint:funlen
func TestFromStepTo(t *testing.T) {
	t.Parallel()

	type test struct {
		Name  string
		Start float64
		Step  float64
		End   float64
		Want  []float64
	}

	tests := []test{
		{
			Name:  "Should yield values including start, excluding end in distance of step",
			Start: 1.0,
			Step:  0.5,
			End:   3.0,
			Want:  []float64{1.0, 1.5, 2.0, 2.5},
		},
		{
			Name:  "Should yield values decreasingly if start > end",
			Start: 3.0,
			Step:  -0.5,
			End:   1.0,
			Want:  []float64{3.0, 2.5, 2.0, 1.5},
		},
		{
			Name:  "Should only yield start if step > (end - start)",
			Start: 1.0,
			Step:  2.0,
			End:   2.0,
			Want:  []float64{1.0},
		},
		{
			Name:  "Should only yield start if step == (end - start)",
			Start: 1.0,
			Step:  1.0,
			End:   2.0,
			Want:  []float64{1.0},
		},
		{
			Name:  "Should invert step if end > start and step < 0",
			Start: 1.0,
			Step:  -1.5,
			End:   10.0,
			Want:  []float64{1.0, 2.5, 4.0, 5.5, 7.0, 8.5},
		},
		{
			Name:  "Should invert step if end < start and step > 0",
			Start: 10.0,
			Step:  1.5,
			End:   1.0,
			Want:  []float64{10.0, 8.5, 7.0, 5.5, 4.0, 2.5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// Arrange
			got := make([]float64, 0)

			// Act
			for v := range FromStepTo(tc.Start, tc.Step, tc.End) {
				got = append(got, v)
			}

			// Assert
			assert.Equal(t, tc.Want, got)
		})
	}
}

func TestFromStepTo_shouldPanicIfStepIsZero(t *testing.T) {
	t.Parallel()

	// Arrange
	start := 1.0
	step := 0.0
	end := 10.0

	// Act
	assert.PanicsWithValue(t, "step must not be 0", func() {
		_ = FromStepTo(start, step, end)
	})
}

func TestFromStepTo_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]float64, 0)

	stop := 5.0

	// Act
	for v := range FromStepTo(1.0, 1.5, 10.0) {
		if v > stop {
			break
		}

		got = append(got, v)
	}

	// Assert
	want := []float64{1.0, 2.5, 4.0}
	assert.Equal(t, want, got)
}

func TestFromStepTo_backwardWithBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]float64, 0)
	stop := 5.0

	// Act
	for v := range FromStepTo(10.0, 2, 1.0) {
		if v < stop {
			break
		}

		got = append(got, v)
	}

	// Assert
	want := []float64{10.0, 8, 6}
	assert.Equal(t, want, got)
}
