package statistics

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinMax_int(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[int]()

	// Assert
	assert.Equal(t, math.MaxInt, mm.Min())
	assert.Equal(t, math.MinInt, mm.Max())
}

func TestNewMinMax_int8(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[int8]()

	// Assert
	assert.Equal(t, int8(math.MaxInt8), mm.Min())
	assert.Equal(t, int8(math.MinInt8), mm.Max())
}

func TestNewMinMax_int16(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[int16]()

	// Assert
	assert.Equal(t, int16(math.MaxInt16), mm.Min())
	assert.Equal(t, int16(math.MinInt16), mm.Max())
}

func TestNewMinMax_int32(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[int32]()

	// Assert
	assert.Equal(t, int32(math.MaxInt32), mm.Min())
	assert.Equal(t, int32(math.MinInt32), mm.Max())
}

func TestNewMinMax_int64(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[int64]()

	// Assert
	assert.Equal(t, int64(math.MaxInt64), mm.Min())
	assert.Equal(t, int64(math.MinInt64), mm.Max())
}

func TestNewMinMax_uint(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[uint]()

	// Assert
	assert.Equal(t, uint(math.MaxUint), mm.Min())
	assert.Equal(t, uint(0), mm.Max())
}

func TestNewMinMax_uint8(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[uint8]()

	// Assert
	assert.Equal(t, uint8(math.MaxUint8), mm.Min())
	assert.Equal(t, uint8(0), mm.Max())
}

func TestNewMinMax_uint16(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[uint16]()

	// Assert
	assert.Equal(t, uint16(math.MaxUint16), mm.Min())
	assert.Equal(t, uint16(0), mm.Max())
}

func TestNewMinMax_uint32(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[uint32]()

	// Assert
	assert.Equal(t, uint32(math.MaxUint32), mm.Min())
	assert.Equal(t, uint32(0), mm.Max())
}

func TestNewMinMax_uint64(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[uint64]()

	// Assert
	assert.Equal(t, uint64(math.MaxUint64), mm.Min())
	assert.Equal(t, uint64(0), mm.Max())
}

func TestNewMinMax_uintptr(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[uintptr]()

	// Assert
	assert.Equal(t, uintptr(math.MaxUint), mm.Min())
	assert.Equal(t, uintptr(0), mm.Max())
}

func TestNewMinMax_float32(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[float32]()

	// Assert
	assert.InEpsilon(t, float32(math.MaxFloat32), mm.Min(), 0.0001)
	assert.InEpsilon(t, float32(-math.MaxFloat32), mm.Max(), 0.0001)
}

func TestNewMinMax_float64(t *testing.T) {
	t.Parallel()

	// Act
	mm := NewMinMaxAccumulator[float64]()

	// Assert
	assert.InEpsilon(t, math.MaxFloat64, mm.Min(), 0.0001)
	assert.InEpsilon(t, -math.MaxFloat64, mm.Max(), 0.0001)
}
