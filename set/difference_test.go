package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-tool-box/set"
)

//nolint:funlen
func TestSet_Difference(t *testing.T) {
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
			name:      "Should return empty set if all sets are empty",
			set:       set.New[any](),
			otherSets: []*set.Set[any]{set.New[any](), set.New[any]()},
			want:      []any{},
		},
		{
			name:      "Should return empty set if all values are common",
			set:       set.New[any](1, 2, 3),
			otherSets: []*set.Set[any]{set.New[any](1, 2, 3)},
			want:      []any{},
		},
		{
			name:      "Should not modify set if all values different",
			set:       set.New[any](1, 2, 3),
			otherSets: []*set.Set[any]{set.New[any](4, 5, 6)},
			want:      []any{1, 2, 3},
		},
		{
			name:      "Should remove values which they are present in other sets",
			set:       set.New[any](6.283185, 2.718, 9.81),
			otherSets: []*set.Set[any]{set.New[any](1.6605, 1.618), set.New[any](6.283185, 2.718)},
			want:      []any{9.81},
		},
		{
			name: "Should remove all values, when all are presen in other sets",
			set:  set.New[any](point{1, 2}, point{3, 4}, point{5, 6}),
			otherSets: []*set.Set[any]{
				set.New[any](point{1, 2}, point{5, 6}),
				set.New[any](point{3, 4}),
			},
			want: []any{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.set
			others := tc.otherSets

			// Act
			s.Difference(others...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}

//nolint:funlen
func TestDifferenceOf(t *testing.T) {
	t.Parallel()

	// Arrange
	type test struct {
		name string
		sets []*set.Set[any]
		want []any
	}

	tests := []test{
		{
			name: "Should create empty set if no sets are given",
			sets: []*set.Set[any]{},
			want: []any{},
		},
		{
			name: "Should create copy of first set if no other sets are given",
			sets: []*set.Set[any]{set.New[any]("a", "b", "c")},
			want: []any{"a", "b", "c"},
		},
		{
			name: "Should create empty set if all sets are empty",
			sets: []*set.Set[any]{set.New[any](), set.New[any]()},
			want: []any{},
		},
		{
			name: "Should return empty set if all values are common",
			sets: []*set.Set[any]{set.New[any](1, 2, 3), set.New[any](1, 2, 3)},
			want: []any{},
		},
		{
			name: "Should create copy of the first set if all values are different",
			sets: []*set.Set[any]{set.New[any](1, 2, 3), set.New[any](4, 5, 6)},
			want: []any{1, 2, 3},
		},
		{
			name: "Should remove values which they are present in other sets",
			sets: []*set.Set[any]{
				set.New[any](6.283185, 2.718, 9.81),
				set.New[any](1.6605, 1.618),
				set.New[any](6.283185, 2.718),
			},
			want: []any{9.81},
		},
		{
			name: "Should remove all values, when all are present in other sets",
			sets: []*set.Set[any]{
				set.New[any](point{1, 2}, point{3, 4}, point{5, 6}),
				set.New[any](point{1, 2}, point{5, 6}),
				set.New[any](point{3, 4}),
			},
			want: []any{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			s := set.DifferenceOf(tc.sets...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}
