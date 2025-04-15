package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	const (
		numAdds            = 1500
		numDifferentValues = 333
	)

	ints := make([]int, numAdds)
	for i := range numAdds {
		ints[i] = i % numDifferentValues
	}

	var s Set[int]
	for b.Loop() {
		s = Of[int](ints...)
	}

	assert.Equal(b, numDifferentValues, s.Len())
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()

	const (
		numAdds            = 1500
		numDifferentValues = 333
	)

	values := make([]int, numAdds)
	for i := range numAdds {
		values[i] = i % numDifferentValues
	}

	var s Set[int]
	for b.Loop() {
		s = Of[int]()

		for _, v := range values {
			s.Add(v)
		}
	}

	assert.Equal(b, numDifferentValues, s.Len())
}
