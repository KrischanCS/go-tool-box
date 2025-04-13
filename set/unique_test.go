package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-tool-box/set"
)

//nolint:funlen
func TestSet_Unique(t *testing.T) {
	t.Parallel()

	// Arrange
	type test struct {
		name      string
		set       *set.Set[any]
		otherSets []*set.Set[any]
		want      []any
	}

	tests := []test{
		{
			name:      "Should not modify set if no other sets are given",
			set:       set.New[any]("a", "b", "c"),
			otherSets: []*set.Set[any]{},
			want:      []any{"a", "b", "c"},
		},
		{
			name:      "Should be empty set if all sets are empty",
			set:       set.New[any](),
			otherSets: []*set.Set[any]{},
			want:      []any{},
		},
		{
			name:      "Should be empty if all values are in common",
			set:       set.New[any](1, 2, 3),
			otherSets: []*set.Set[any]{set.New[any](1, 2, 3)},
			want:      []any{},
		},
		{
			name:      "Should contain all values if they are all different",
			set:       set.New[any](1, 2, 3),
			otherSets: []*set.Set[any]{set.New[any](4, 5, 6)},
			want:      []any{1, 2, 3, 4, 5, 6},
		},
		{
			name:      "Should contain all values if they are all different with multiple sets",
			set:       set.New[any](6.283185, 2.718, 9.81),
			otherSets: []*set.Set[any]{set.New[any](1.6605, 1.618), set.New[any](1.38, 1.602)},
			want:      []any{6.283185, 2.718, 9.81, 1.6605, 1.618, 1.38, 1.602},
		},
		{
			name: "Should contain all values which are unique over all sets",
			set:  set.New[any](point{1, 2}, point{3, 4}),
			otherSets: []*set.Set[any]{
				set.New[any](point{9, 10}, point{3, 4}, point{5, 6}),
				set.New[any](point{3, 4}, point{7, 8}, point{9, 10}),
			},
			want: []any{point{1, 2}, point{5, 6}, point{7, 8}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.set
			others := tc.otherSets

			// Act
			s.Unique(others...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}

//nolint:funlen
func TestUniqueOf(t *testing.T) {
	t.Parallel()

	// Arrange
	type test struct {
		name string
		sets []*set.Set[any]
		want []any
	}

	tests := []test{
		{
			name: "Should create a new, empty set if no sets are given",
			sets: []*set.Set[any]{},
			want: []any{},
		},
		{
			name: "Should create a copy of given set if only one is given",
			sets: []*set.Set[any]{set.New[any]("a", "b", "c")},
			want: []any{"a", "b", "c"},
		},
		{
			name: "Should create empty set if all sets are empty",
			sets: []*set.Set[any]{set.New[any](), set.New[any](), set.New[any]()},
			want: []any{},
		},
		{
			name: "Should create an empty set if all values are common",
			sets: []*set.Set[any]{
				set.New[any](1, 2, 3),
				set.New[any](1, 2, 3),
			},
			want: []any{},
		},
		{
			name: "Should create a set with all values if they are all different",
			sets: []*set.Set[any]{
				set.New[any](1, 2, 3),
				set.New[any](4, 5, 6),
			},
			want: []any{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Should create a set with all values if they are all different with multiple sets",
			sets: []*set.Set[any]{
				set.New[any](6.283185, 2.718, 9.81),
				set.New[any](1.6605, 1.618),
				set.New[any](1.38, 1.602),
			},
			want: []any{6.283185, 2.718, 9.81, 1.6605, 1.618, 1.38, 1.602},
		},
		{
			name: "Should create a set with all values which are unique over all sets",
			sets: []*set.Set[any]{
				set.New[any](point{1, 2}, point{3, 4}),
				set.New[any](point{9, 10}, point{3, 4}, point{5, 6}),
				set.New[any](point{3, 4}, point{7, 8}, point{9, 10}),
			},
			want: []any{point{1, 2}, point{5, 6}, point{7, 8}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			s := set.UniqueOf(tc.sets...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}
