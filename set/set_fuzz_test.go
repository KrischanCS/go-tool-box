package set_test

import (
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

func addAndCheck(t *testing.T, s *set.Set[int], value int, l int) {
	t.Helper()

	contained := s.Contains(value)

	s.Add(value)
	assert.True(t, s.Contains(value))

	if contained {
		assert.Equal(t, l, s.Len())
	} else {
		assert.Equal(t, l+1, s.Len())
	}
}

func RemoveAndCheck(t *testing.T, s *set.Set[int], value int, l int) {
	t.Helper()

	contained := s.Remove(value)
	assert.False(t, s.Contains(value))

	if contained {
		assert.Equal(t, l-1, s.Len())
	} else {
		assert.Equal(t, l, s.Len())
	}
}
