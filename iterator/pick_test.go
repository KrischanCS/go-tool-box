package iterator

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPickRight(t *testing.T) {
	t.Parallel()

	// Arrange
	s := []int{1, 2, 3, 4, 5, 6}
	slicesIter := slices.All(s)

	got := make([]int, 0, len(s))

	// Act
	for e := range PickRight(slicesIter) {
		got = append(got, e)
	}

	// Assert
	assert.Equal(t, s, got)
}

func TestPickLeft(t *testing.T) {
	t.Parallel()

	// Arrange
	s := []int{1, 2, 3, 4, 5, 6}
	slicesIter := slices.All(s)

	got := make([]int, 0, len(s))

	// Act
	for e := range PickLeft(slicesIter) {
		got = append(got, e)
	}

	// Assert
	want := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, want, got)
}

func TestCombine(t *testing.T) {
	t.Parallel()

	// Arrange
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	mapIter := maps.All(m)

	got := make([]Pair[int, string], 0, 3)

	// Act
	for e := range Combine(mapIter) {
		got = append(got, e)
	}

	// Assert
	want := []Pair[int, string]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	assert.ElementsMatch(t, want, got)
}

func TestPickRight_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	s := []int{1, 2, 3, 4, 5, 6}
	slicesIter := slices.All(s)

	got := make([]int, 0, len(s))

	// Act
	for e := range PickRight(slicesIter) {
		got = append(got, e)

		if e == 3 {
			break
		}
	}

	// Assert
	assert.Equal(t, s[:3], got)
}

func TestPickLeft_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	s := []int{1, 2, 3, 4, 5, 6}
	slicesIter := slices.All(s)

	got := make([]int, 0, len(s))

	// Act
	for e := range PickLeft(slicesIter) {
		got = append(got, e)

		if e == 2 {
			break
		}
	}

	// Assert
	want := []int{0, 1, 2}
	assert.Equal(t, want, got)
}

func TestCombine_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	mapIter := maps.All(m)

	want := "b"

	// Act
	var got string
	for e := range Combine(mapIter) {
		got = e.Right

		if e.Left == 2 {
			break
		}
	}

	assert.Equal(t, want, got)
}
