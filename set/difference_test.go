package set_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/set"
)

func ExampleDifferenceOf() {
	setA := set.Of(1, 2, 3, 4)
	setB := set.Of(3, 6)
	setC := set.Of(4, 7)

	diff := set.DifferenceOf(setA, setB)
	fmt.Println("A - B:", diff)
	fmt.Println("Original is not modified:", setA)

	fmt.Println()

	diff = set.DifferenceOf(setA, setB, setC)
	fmt.Println("A - B - C:", diff)

	// Output:
	// A - B: (Set[int]: [1 2 4])
	// Original is not modified: (Set[int]: [1 2 3 4])
	//
	// A - B - C: (Set[int]: [1 2])
}

func ExampleSet_Difference() {
	setA := set.Of(1, 2, 3, 4)
	setB := set.Of(3, 4, 5, 6)

	setA.Difference(setB)
	fmt.Println("A = A - B:", setA)

	setC := set.Of(3, 5)
	setD := set.Of(6, 7)

	setB.Difference(setC, setD)
	fmt.Println("B = B - C - D:", setB)

	// Output:
	// A = A - B: (Set[int]: [1 2])
	// B = B - C - D: (Set[int]: [4])
}

//nolint:funlen
func TestSet_Difference(t *testing.T) {
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
			name:      "Should return empty set if all values are common",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](1, 2, 3)},
			want:      []any{},
		},
		{
			name:      "Should not modify set if all values different",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](4, 5, 6)},
			want:      []any{1, 2, 3},
		},
		{
			name:      "Should remove values which they are present in other sets",
			set:       set.Of[any](6.283185, 2.718, 9.81),
			otherSets: []set.Set[any]{set.Of[any](1.6605, 1.618), set.Of[any](6.283185, 2.718)},
			want:      []any{9.81},
		},
		{
			name: "Should remove all values, when all are present in other sets",
			set:  set.Of[any](point{1, 2}, point{3, 4}, point{5, 6}),
			otherSets: []set.Set[any]{
				set.Of[any](point{1, 2}, point{5, 6}),
				set.Of[any](point{3, 4}),
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
		sets []set.Set[any]
		want []any
	}

	tests := []test{
		{
			name: "Should create empty set if no sets are given",
			sets: []set.Set[any]{},
			want: []any{},
		},
		{
			name: "Should create copy of first set if no other sets are given",
			sets: []set.Set[any]{set.Of[any]("a", "b", "c")},
			want: []any{"a", "b", "c"},
		},
		{
			name: "Should create empty set if all sets are empty",
			sets: []set.Set[any]{set.Of[any](), set.Of[any]()},
			want: []any{},
		},
		{
			name: "Should return empty set if all values are common",
			sets: []set.Set[any]{set.Of[any](1, 2, 3), set.Of[any](1, 2, 3)},
			want: []any{},
		},
		{
			name: "Should create copy of the first set if all values are different",
			sets: []set.Set[any]{set.Of[any](1, 2, 3), set.Of[any](4, 5, 6)},
			want: []any{1, 2, 3},
		},
		{
			name: "Should remove values which they are present in other sets",
			sets: []set.Set[any]{
				set.Of[any](6.283185, 2.718, 9.81),
				set.Of[any](1.6605, 1.618),
				set.Of[any](6.283185, 2.718),
			},
			want: []any{9.81},
		},
		{
			name: "Should remove all values, when all are present in other sets",
			sets: []set.Set[any]{
				set.Of[any](point{1, 2}, point{3, 4}, point{5, 6}),
				set.Of[any](point{1, 2}, point{5, 6}),
				set.Of[any](point{3, 4}),
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
