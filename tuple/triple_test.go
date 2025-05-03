package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripleOf(t *testing.T) {
	t.Parallel()

	// Arrange
	first := 1
	second := "a"
	third := 6.28

	// Act
	got := TripleOf(first, second, third)

	// Assert
	want := Triple[int, string, float64]{first: first, second: second, third: third}
	assert.Equal(t, want, got)
}

func TestTripleOf_First(t *testing.T) {
	t.Parallel()

	// Arrange
	triple := TripleOf(1, "a", 6.28)

	// Act
	first := triple.First()

	// Assert
	assert.Equal(t, 1, first)
}

func TestTripleOf_Second(t *testing.T) {
	t.Parallel()

	// Arrange
	triple := TripleOf(1, "a", 6.28)

	// Act
	second := triple.Second()

	// Assert
	assert.Equal(t, "a", second)
}

func TestTripleOf_Third(t *testing.T) {
	t.Parallel()

	// Arrange
	triple := TripleOf(1, "a", 6.28)

	// Act
	third := triple.Third()

	// Assert
	assert.InEpsilon(t, 6.28, third, 0.0001)
}

func TestTripleOf_Unpack(t *testing.T) {
	t.Parallel()

	// Arrange
	triple := TripleOf(1, "a", 6.28)

	// Act
	first, second, third := triple.Unpack()

	// Assert
	assert.Equal(t, 1, first)
	assert.Equal(t, "a", second)
	assert.InEpsilon(t, 6.28, third, 0.0001)
}

func TestTripleOf_String(t *testing.T) {
	t.Parallel()

	// Arrange
	triple := TripleOf(1, "a", 6.28)

	// Act
	got := triple.String()

	// Assert
	want := "(Triple[int, string, float64]: [1; a; 6.28])"
	assert.Equal(t, want, got)
}
