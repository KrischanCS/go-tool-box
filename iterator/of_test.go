package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]string, 0, 3)

	// Act
	iter := Of[string]("a", "b", "c")
	for v := range iter {
		got = append(got, v)
	}

	// Assert
	want := []string{"a", "b", "c"}
	assert.Equal(t, want, got)
}

func TestOf_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]string, 0, 1)

	// Act
	iter := Of[string]("a", "b", "c")
	for v := range iter {
		got = append(got, v)
		break
	}

	// Assert
	want := []string{"a"}
	assert.Equal(t, want, got)
}
