package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairOf(t *testing.T) {
	t.Parallel()

	// Arrange
	first := 1
	second := "a"

	// Act
	got := PairOf(first, second)

	// Assert
	want := Pair[int, string]{first: first, second: second}
	assert.Equal(t, want, got)
}

func TestPair_First(t *testing.T) {
	t.Parallel()

	// Arrange
	p := PairOf(1, "a")

	// Act
	first := p.First()

	// Assert
	assert.Equal(t, 1, first)
}

func TestPair_Second(t *testing.T) {
	t.Parallel()

	// Arrange
	p := PairOf(1, "a")

	// Act
	second := p.Second()

	// Assert
	assert.Equal(t, "a", second)
}

func TestPair_Unpack(t *testing.T) {
	t.Parallel()

	// Arrange
	p := PairOf(1, "a")

	// Act
	first, second := p.Unpack()

	// Assert
	assert.Equal(t, 1, first)
	assert.Equal(t, "a", second)
}

func TestPair_String(t *testing.T) {
	t.Parallel()

	// Arrange
	p := PairOf(1, "a")

	// Act
	got := p.String()

	// Assert
	want := "(Pair[int, string]: [1; a])"

	assert.Equal(t, want, got)
}
