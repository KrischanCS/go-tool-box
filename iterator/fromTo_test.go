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
