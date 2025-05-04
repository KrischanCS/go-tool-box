// Package statistics provides reducers for gathering statistics from an
// iterator.
package statistics

import (
	"fmt"

	"github.com/KrischanCS/go-toolbox/constraints"
)

// MinMaxAccumulator is the accumulator type for the [MinMax] reducer.
type MinMaxAccumulator[T any] interface {
	Min() T
	Max() T

	setMin(newMin T)
	setMax(newMax T)
}

// NewMinMaxAccumulator creates a new [MinMaxAccumulator] with the correct
// initial values.
func NewMinMaxAccumulator[T constraints.RealNumber]() MinMaxAccumulator[T] {
	minMaxer, ok := initMinMax[T]()
	if !ok {
		var t T

		panic(fmt.Sprintf("Conversion failed for type %T", t))
	}

	return minMaxer
}

// MeanAccumulator is the accumulator type for the [Mean] reducer.
type MeanAccumulator[T constraints.RealNumber] struct {
	sum   T
	count int
}

// Mean returns the arithmetic mean of the gathered values.
func (m MeanAccumulator[T]) Mean() float64 {
	return float64(m.sum) / float64(m.count)
}
