package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/set"
)

// point is a simple struct as example for testing.
type point struct {
	X int
	Y int
}

func TestOf_empty(t *testing.T) {
	t.Parallel()

	// Act
	s := set.Of[string]()

	// Assert
	assert.True(t, s.IsEmpty())
}

func TestOf_withValues(t *testing.T) {
	t.Parallel()

	// Act
	s := set.Of[string]("a", "b", "c")

	// Assert
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 3, s.Len())
	assert.True(t, s.Contains("a"))
	assert.True(t, s.Contains("b"))
	assert.True(t, s.Contains("c"))
}

func TestOf_withDuplicateValues(t *testing.T) {
	t.Parallel()

	// Act
	s := set.Of[string]("a", "b", "c", "a")

	// Assert
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 3, s.Len())
	assert.True(t, s.Contains("a"))
	assert.True(t, s.Contains("b"))
	assert.True(t, s.Contains("c"))
}

func TestSet_String(t *testing.T) {
	t.Parallel()

	// Arrange
	type testCase struct {
		name  string
		input []string
		want  string
	}

	tests := []testCase{
		{
			name:  "Empty",
			input: []string{},
			want:  "(Set[string]: <empty>)",
		},
		{
			name:  "Single value",
			input: []string{"a"},
			want:  "(Set[string]: [a])",
		},
		{
			name:  "Multiple values",
			input: []string{"1", "2", "3", "4", "5"},
			want:  "(Set[string]: [1 2 3 4 5])",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			s := set.Of(tc.input...)

			// Assert
			assert.Equal(t, tc.want, s.String())
		})
	}
}

func TestSet_Add(t *testing.T) {
	t.Parallel()

	// Arrange
	s := set.Of[string]("a", "b")

	// Act
	s.Add("c")
	s.Add("d", "e")
	s.Add("a", "c", "e", "f")

	// Assert
	assert.ElementsMatch(t, []string{"a", "b", "c", "d", "e", "f"}, s.Values())
}

func TestSet_Remove(t *testing.T) {
	t.Parallel()

	// Arrange
	s := set.Of[string]("a", "b", "c", "d", "e")

	// Act
	s.Remove("a")
	s.Remove("b", "c", "d")
	s.Remove("a", "c")

	// Assert
	assert.ElementsMatch(t, []string{"e"}, s.Values())
}

func TestSet_Clear(t *testing.T) {
	t.Parallel()

	// Arrange
	s := set.Of[string]("a", "b", "c")

	// Act
	s.Clear()

	assert.Equal(t, 0, s.Len())
}

func TestSet_All(t *testing.T) {
	t.Parallel()

	// Arrange
	s := set.Of[string]("a", "b", "c")
	dst := make([]string, 0, s.Len())

	// Act
	for v := range s.All() {
		dst = append(dst, v)
	}

	// Assert
	assert.ElementsMatch(t, []string{"a", "b", "c"}, dst)
}

func TestSet_All_break(t *testing.T) {
	t.Parallel()

	// Arrange
	s := set.Of[string]("a", "b", "c")
	dst := make([]string, 0, 1)

	// Act
	for v := range s.All() {
		dst = append(dst, v)
		break
	}

	// Assert
	assert.Len(t, dst, 1)
	assert.Subset(t, []string{"a", "b", "c"}, dst)
}
