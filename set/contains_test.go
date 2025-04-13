package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-tool-box/set"
)

func TestSet_Contains_Add(t *testing.T) {
	t.Parallel()

	// Arrange
	s := set.New(1.0, 2.1, 3.2)

	// Act
	s.Add(4.31)
	s.Add(3.1415)

	// Act  & Assert
	assert.True(t, s.Contains(1.0))
	assert.True(t, s.Contains(2.1))
	assert.True(t, s.Contains(3.2))
	assert.True(t, s.Contains(4.31))
	assert.True(t, s.Contains(3.1415))
}

func TestSet_Contains_Remove(t *testing.T) {
	t.Parallel()

	// Arrange
	s := set.New(1, 2, 3)

	// Act
	contained2 := s.Remove(2)
	contained3 := s.Remove(3)
	contained4 := s.Remove(4)

	// Assert
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	assert.Equal(t, 1, s.Len())

	assert.True(t, contained2)
	assert.True(t, contained3)
	assert.False(t, contained4)
}

//nolint:funlen
func TestSet_Contains_multipleValues(t *testing.T) {
	t.Parallel()

	// Arrange
	type testCase struct {
		name  string
		input []any
		check []any
		want  bool
	}

	tests := []testCase{
		{
			name:  "Should return true if set contains exactly the same values as checked",
			input: []any{"a", "b", "c"},
			check: []any{"a", "b", "c"},
			want:  true,
		},
		{
			name:  "Should return true if set contains the same values in different order",
			input: []any{1, 2, 3},
			check: []any{2, 3, 1},
			want:  true,
		},
		{
			name:  "Should return true if set is empty and no value is checked",
			input: []any{},
			check: []any{},
			want:  true,
		},
		{
			name:  "Should return true if set contains more values than checked, but all checked values are in the set",
			input: []any{6.283185, 2.718, 9.81},
			check: []any{6.283185, 2.718},
			want:  true,
		},
		{
			name:  "Should return false if set contains less values than checked",
			input: []any{point{1, 2}, point{3, 4}, point{5, 6}},
			check: []any{point{1, 2}, point{3, 4}, point{5, 6}, point{7, 8}},
		},
		{
			name:  "Should return false if set is empty and any value is checked",
			input: []any{},
			check: []any{"a"},
			want:  false,
		},
		{
			name:  "Should return false if set contains different values than checked",
			input: []any{"a", "b", "c"},
			check: []any{"d", "e", "f"},
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := set.New(tc.input...)

			// Act
			got := s.Contains(tc.check...)

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Set_Contains_Exactly(t *testing.T) {
	t.Parallel()

	// Arrange
	type test struct {
		name  string
		input []any
		check []any
		want  bool
	}

	tests := []test{
		{
			name:  "Should return true if set contains exactly the same values as checked",
			input: []any{"a", "b", "c"},
			check: []any{"a", "b", "c"},
			want:  true,
		},
		{
			name:  "Should return true if set contains the same values in different order",
			input: []any{1.0, 2.1, 3.2},
			check: []any{2.1, 3.2, 1.0},
			want:  true,
		},
		{
			name:  "Should return false if set contains less values than checked",
			input: []any{point{1, 2}, point{3, 4}, point{5, 6}},
			check: []any{point{1, 2}, point{3, 4}, point{5, 6}, point{7, 8}},
			want:  false,
		},
		{
			name:  "Should return false if set contains more values than checked",
			input: []any{6.283185, 2.718, 9.81},
			check: []any{6.283185, 2.718},
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			s := set.New(tc.input...)
			got := s.ContainsExactly(tc.check...)

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}
