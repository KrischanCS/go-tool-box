package statistics

import (
	"cmp"

	"github.com/KrischanCS/go-toolbox/constraints"
	"github.com/KrischanCS/go-toolbox/iterator"
)

// Min is a [iterator.Reducer], which collects the minimum value of a stream.
// acc must be initialized to the maximum value of the type.
func Min[T cmp.Ordered](acc *T, in T) {
	*acc = min(*acc, in)
}

// Max is a [iterator.Reducer], which collects the maximum value of a stream.
// acc must be initialized to the minimum value of the type.
func Max[T cmp.Ordered](acc *T, in T) {
	*acc = max(*acc, in)
}

// Mean is a [iterator.Reducer], which collects the mean value of a stream.
func Mean[T constraints.RealNumber](acc *MeanAccumulator[T], in T) {
	acc.sum += in
	acc.count++
}

var _ iterator.Reducer[MinMaxAccumulator[int], int] = MinMax[int]

// MinMax is a [iterator.Reducer], which collects the minimum and maximum value
// of a stream.
//
// acc min must be initialized to the maximum value of the type,
// acc max must be initialized to the minimum value of the type.
//
// For real numbers, a correctly initialized MinMaxAccumulator can be created with
// [NewMinMaxAccumulator].
func MinMax[T cmp.Ordered](acc *MinMaxAccumulator[T], in T) {
	(*acc).setMin(min((*acc).Min(), in))
	(*acc).setMax(max((*acc).Max(), in))
}
