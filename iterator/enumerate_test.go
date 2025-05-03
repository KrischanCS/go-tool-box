package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumerate(t *testing.T) {
	t.Parallel()

	// Arrange

	input := Of("a", "b", "c", "d", "e")
	got := make([]Pair[int, string], 0, 5)

	// Act
	for i, v := range Enumerate(input) {
		got = append(got, Pair[int, string]{i, v})
	}

	// Assert
	want := []Pair[int, string]{
		{0, "a"},
		{1, "b"},
		{2, "c"},
		{3, "d"},
		{4, "e"},
	}

	assert.Equal(t, want, got)
}

func TestEnumerate_shouldStopOnBreak(t *testing.T) {
	t.Parallel()

	// Arrange

	input := Of("a", "b", "c", "d", "e")
	got := make([]Pair[int, string], 0, 3)

	// Act
	for i, v := range Enumerate(input) {
		if i == 3 {
			break
		}

		got = append(got, Pair[int, string]{i, v})
	}

	// Assert
	want := []Pair[int, string]{
		{0, "a"},
		{1, "b"},
		{2, "c"},
	}

	assert.Equal(t, want, got)
}
