package set_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-tool-box/set"
)

func FuzzSet(f *testing.F) {
	f.Add(true, 1)
	f.Add(false, 2)
	f.Add(true, -47)
	f.Add(false, 256)
	f.Add(true, 0)
	f.Add(false, 1024)
	f.Add(true, 43217894231)
	f.Add(false, -123456789)

	s := set.Of[int]()

	f.Fuzz(func(t *testing.T, add bool, value int) {
		l := s.Len()

		if add {
			addAndCheck(t, s, value, l)
		} else {
			RemoveAndCheck(t, s, value, l)
		}
	})
}

func addAndCheck(t *testing.T, s set.Set[int], value int, l int) {
	t.Helper()

	notContained := s.Add(value)
	assert.True(t, s.Contains(value))

	if notContained {
		assert.Equal(t, l+1, s.Len())
	} else {
		assert.Equal(t, l, s.Len())
	}
}

func RemoveAndCheck(t *testing.T, s set.Set[int], value int, l int) {
	t.Helper()

	contained := s.Remove(value)
	assert.False(t, s.Contains(value))

	if contained {
		assert.Equal(t, l-1, s.Len())
	} else {
		assert.Equal(t, l, s.Len())
	}
}

func FuzzSet_Operations(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(32))
	f.Add(int64(47))
	f.Add(int64(-13))
	f.Add(int64(2))
	f.Add(int64(102))
	f.Add(int64(12348348193478192))
	f.Add(int64(-32438914312))

	f.Fuzz(func(_ *testing.T, seed int64) {
		//nolint:gosec
		rand := rand.New(rand.NewSource(seed))

		numSets := rand.Intn(7) + 1

		sets := make([]set.Set[int], numSets)

		for i := range numSets {
			numElements := rand.Intn(32)

			s := set.WithCapacity[int](numElements)
			for range numSets {
				s.Add(rand.Intn(16))
			}

			sets[i] = s
		}

		_ = set.UnionOf(sets...)
		_ = set.IntersectionOf(sets...)
		_ = set.DifferenceOf(sets...)
		_ = set.UniqueOf(sets...)
	})
}
