package statistics

import (
	"fmt"
	"math"
)

type minMax[T any] struct {
	min T
	max T
}

func (m minMax[T]) Min() T {
	return m.min
}

func (m minMax[T]) Max() T {
	return m.max
}

//nolint:unused // false positive
func (m *minMax[T]) setMin(newMin T) {
	m.min = newMin
}

//nolint:unused // false positive
func (m *minMax[T]) setMax(newMax T) {
	m.max = newMax
}

func initMinMax[T comparable]() (MinMaxAccumulator[T], bool) {
	var minMaxer any

	var t T
	switch any(t).(type) {
	default:
		panic(fmt.Sprintf("unsupported type %T", t))
	case int:
		minMaxer = intMinMax()
	case int8:
		minMaxer = int8MinMax()
	case int16:
		minMaxer = int16MinMax()
	case int32:
		minMaxer = int32MinMax()
	case int64:
		minMaxer = int64MinMax()
	case uint:
		minMaxer = uintMinMax()
	case uint8:
		minMaxer = uint8MinMax()
	case uint16:
		minMaxer = uint16MinMax()
	case uint32:
		minMaxer = uint32MinMax()
	case uint64:
		minMaxer = uint64MinMax()
	case uintptr:
		minMaxer = uintptrMinMax()
	case float32:
		minMaxer = float32MinMax()
	case float64:
		minMaxer = float64MinMax()
	}

	mm, ok := minMaxer.(MinMaxAccumulator[T])

	return mm, ok
}

func intMinMax() MinMaxAccumulator[int] {
	return &minMax[int]{min: math.MaxInt, max: math.MinInt}
}

func int8MinMax() MinMaxAccumulator[int8] {
	return &minMax[int8]{min: math.MaxInt8, max: math.MinInt8}
}

func int16MinMax() MinMaxAccumulator[int16] {
	return &minMax[int16]{min: math.MaxInt16, max: math.MinInt16}
}

func int32MinMax() MinMaxAccumulator[int32] {
	return &minMax[int32]{min: math.MaxInt32, max: math.MinInt32}
}

func int64MinMax() MinMaxAccumulator[int64] {
	return &minMax[int64]{min: math.MaxInt64, max: math.MinInt64}
}

func uintMinMax() MinMaxAccumulator[uint] {
	return &minMax[uint]{min: math.MaxUint, max: 0}
}

func uint8MinMax() MinMaxAccumulator[uint8] {
	return &minMax[uint8]{min: math.MaxUint8, max: 0}
}

func uint16MinMax() MinMaxAccumulator[uint16] {
	return &minMax[uint16]{min: math.MaxUint16, max: 0}
}

func uint32MinMax() MinMaxAccumulator[uint32] {
	return &minMax[uint32]{min: math.MaxUint32, max: 0}
}

func uint64MinMax() MinMaxAccumulator[uint64] {
	return &minMax[uint64]{min: math.MaxUint64, max: 0}
}

func uintptrMinMax() MinMaxAccumulator[uintptr] {
	return &minMax[uintptr]{min: math.MaxUint, max: 0}
}

func float32MinMax() MinMaxAccumulator[float32] {
	return &minMax[float32]{min: math.MaxFloat32, max: -math.MaxFloat32}
}

func float64MinMax() MinMaxAccumulator[float64] {
	return &minMax[float64]{min: math.MaxFloat64, max: -math.MaxFloat64}
}
