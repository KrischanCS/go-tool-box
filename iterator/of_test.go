package iterator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
)

func ExampleOf() {
	s := iterator.Of(1, 2, 3, 4, 5)

	for v := range s {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func TestOf(t *testing.T) {
	t.Parallel()

	// Arrange
	got := make([]string, 0, 3)

	// Act
	iter := iterator.Of[string]("a", "b", "c")
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
	iter := iterator.Of[string]("a", "b", "c")
	for v := range iter {
		got = append(got, v)
		break
	}

	// Assert
	want := []string{"a"}
	assert.Equal(t, want, got)
}
