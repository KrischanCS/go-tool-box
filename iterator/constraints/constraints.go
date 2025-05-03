// Package constraints provides type constraints used by the iterator package.
package constraints

// RealNumber is a type constraint that matches all real numeric types.
type RealNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Number is a type constraint that matches all numeric types.
type Number interface {
	RealNumber | ~complex64 | ~complex128
}
