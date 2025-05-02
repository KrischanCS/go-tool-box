package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/set"
)

//nolint:funlen
func TestIntersection(t *testing.T) {
	t.Parallel()

	// Arrange

	type test struct {
		name      string
		set       set.Set[any]
		otherSets []set.Set[any]
		want      []any
	}

	tests := []test{
		{
			name:      "Should not modify set if no other sets are given",
			set:       set.Of[any]("a", "b", "c"),
			otherSets: []set.Set[any]{},
			want:      []any{"a", "b", "c"},
		},
		{
			name:      "Should return empty set if all sets are empty",
			set:       set.Of[any](),
			otherSets: []set.Set[any]{set.Of[any](), set.Of[any]()},
			want:      []any{},
		},
		{
			name:      "Should return empty set if no values are in common",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](4, 5, 6)},
			want:      []any{},
		},
		{
			name:      "Should return empty set if no values are in common with multiple sets",
			set:       set.Of[any](6.283185, 2.718, 9.81),
			otherSets: []set.Set[any]{set.Of[any](1.6605, 1.618), set.Of[any](1.38, 1.602)},
			want:      []any{},
		},
		{
			name:      "Should return empty set if one of the other sets is empty",
			set:       set.Of[any](point{1, 2}, point{3, 4}),
			otherSets: []set.Set[any]{set.Of[any](), set.Of[any](point{3, 4}, point{1, 2})},
			want:      []any{},
		},
		{
			name:      "Should not modify set if all values are in common",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](1, 2, 3)},
			want:      []any{1, 2, 3},
		},
		{
			name:      "Should return common values from multiple sets",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](1, 2, 3, 4), set.Of[any](1, 2, 5)},
			want:      []any{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.set
			others := tc.otherSets

			// Act
			s.Intersection(others...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}

//nolint:funlen
func TestIntersectionOf(t *testing.T) {
	t.Parallel()

	// Arrange

	type test struct {
		name string
		sets []set.Set[any]
		want []any
	}

	tests := []test{
		{
			name: "Should create an empty set if no sets are given",
			sets: []set.Set[any]{},
			want: []any{},
		},
		{
			name: "Should create a copy of a set if only one set is given",
			sets: []set.Set[any]{set.Of[any]("a", "b", "c")},
			want: []any{"a", "b", "c"},
		},
		{
			name: "Should create an empty set if all sets are empty",
			sets: []set.Set[any]{set.Of[any](), set.Of[any]()},
			want: []any{},
		},
		{
			name: "Should create an empty set if no values are in common",
			sets: []set.Set[any]{set.Of[any](1, 2, 3), set.Of[any](4, 5, 6)},
			want: []any{},
		},
		{
			name: "Should create an empty set if no values are in common with multiple sets",
			sets: []set.Set[any]{
				set.Of[any](6.283185, 2.718, 9.81),
				set.Of[any](1.6605, 1.618),
				set.Of[any](1.38, 1.602),
			},
			want: []any{},
		},
		{
			name: "Should create an empty set if one set is empty",
			sets: []set.Set[any]{
				set.Of[any](point{1, 2}, point{3, 4}),
				set.Of[any](),
				set.Of[any](point{3, 4}, point{1, 2}),
			},
			want: []any{},
		},
		{
			name: "Should create a set with all elements if all values are in common",
			sets: []set.Set[any]{set.Of[any](1, 2, 3), set.Of[any](1, 2, 3)},
			want: []any{1, 2, 3},
		},
		{
			name: "Should create a set with values which are common in all given sets",
			sets: []set.Set[any]{
				set.Of[any](1, 2, 3),
				set.Of[any](1, 2, 3, 4),
				set.Of[any](1, 2, 5),
			},
			want: []any{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := set.IntersectionOf(tc.sets...)

			// Assert
			assert.ElementsMatch(t, tc.want, got.Values())
		})
	}
}
