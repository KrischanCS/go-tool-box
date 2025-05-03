package set_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/set"
)

func ExampleUnionOf() {
	setA := set.Of(1, 2, 3, 4)
	setB := set.Of(3, 6)
	setC := set.Of(4, 7)

	fmt.Println("A ∪ B:", set.UnionOf(setA, setB))
	fmt.Println("B ∪ C:", set.UniqueOf(setB, setC))
	fmt.Println("C ∪ A:", set.UnionOf(setC, setA))
	fmt.Println("A ∪ B ∪ C:", set.UnionOf(setA, setB, setC))

	fmt.Println()

	fmt.Println("Originals are not modified:")
	fmt.Println("A:", setA)
	fmt.Println("B:", setB)
	fmt.Println("C:", setC)

	// Output:
	// A ∪ B: (Set[int]: [1 2 3 4 6])
	// B ∪ C: (Set[int]: [3 4 6 7])
	// C ∪ A: (Set[int]: [1 2 3 4 7])
	// A ∪ B ∪ C: (Set[int]: [1 2 3 4 6 7])
	//
	// Originals are not modified:
	// A: (Set[int]: [1 2 3 4])
	// B: (Set[int]: [3 6])
	// C: (Set[int]: [4 7])
}

func ExampleSet_Union() {
	setA := set.Of(1, 2, 3, 4)
	setB := set.Of(3, 6)

	setA.Union(setB)
	fmt.Println("A = A ∪ B:", setA)

	setC := set.Of(3, 1, 5)
	setD := set.Of(3, 4)
	setB.Union(setC, setD)

	fmt.Println("B = B ∪ C ∪ D:", setB)

	// Output:
	// A = A ∪ B: (Set[int]: [1 2 3 4 6])
	// B = B ∪ C ∪ D: (Set[int]: [1 3 4 5 6])
}

func TestUnion(t *testing.T) {
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
			name:      "Should return empty set if all sets are empty",
			set:       set.Of[any](),
			otherSets: []set.Set[any]{set.Of[any](), set.Of[any]()},
			want:      []any{},
		},
		{
			name:      "Should not modify set wif no sets are added",
			set:       set.Of[any]("a", "b", "c"),
			otherSets: []set.Set[any]{},
			want:      []any{"a", "b", "c"},
		},
		{
			name:      "Should add values from other sets to set",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](4, 5, 6)},
			want:      []any{1, 2, 3, 4, 5, 6},
		},
		{
			name:      "Should add values from multiple other sets to set",
			set:       set.Of[any](6.283185, 2.718, 9.81),
			otherSets: []set.Set[any]{set.Of[any](1.6605, 1.618), set.Of[any](1.38, 1.602)},
			want:      []any{6.283185, 2.718, 9.81, 1.6605, 1.618, 1.38, 1.602},
		},
		{
			name: "Should add values from multiple other sets to set without adding duplicates",
			set:  set.Of[any](point{1, 2}, point{3, 4}),
			otherSets: []set.Set[any]{
				set.Of[any](point{5, 6}, point{1, 2}),
				set.Of[any](point{3, 4}, point{7, 8}),
			},
			want: []any{point{1, 2}, point{3, 4}, point{7, 8}, point{5, 6}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.set
			others := tc.otherSets

			// Act
			s.Union(others...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}

func TestUnionOf(t *testing.T) {
	t.Parallel()

	// Arrange

	type test struct {
		name string
		sets []set.Set[any]
		want []any
	}

	tests := []test{
		{
			name: "Should create an empty set if all given sets are empty",
			sets: []set.Set[any]{set.Of[any](), set.Of[any]()},
			want: []any{},
		},
		{
			name: "Should create an empty set if no sets are given",
			sets: []set.Set[any]{},
			want: []any{},
		},
		{
			name: "Should create a copy of a set if only one set is given",
			sets: []set.Set[any]{set.Of[any](4, 5, 6)},
			want: []any{4, 5, 6},
		},
		{
			name: "Should create a set with values of given sets",
			sets: []set.Set[any]{set.Of[any](1, 2, 3), set.Of[any](4, 5, 6)},
			want: []any{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Should create a set with values from multiple other sets",
			sets: []set.Set[any]{
				set.Of[any](6.283185, 2.718, 9.81),
				set.Of[any](1.6605, 1.618),
				set.Of[any](1.38, 1.602),
			},
			want: []any{6.283185, 2.718, 9.81, 1.6605, 1.618, 1.38, 1.602},
		},
		{
			name: "Should create a set with values from multiple other sets without adding duplicates",
			sets: []set.Set[any]{
				set.Of[any](point{1, 2}, point{3, 4}),
				set.Of[any](point{5, 6}, point{1, 2}),
				set.Of[any](point{3, 4}, point{7, 8}),
			},
			want: []any{point{1, 2}, point{3, 4}, point{7, 8}, point{5, 6}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := set.UnionOf(tc.sets...)

			// Assert
			assert.ElementsMatch(t, tc.want, got.Values())
		})
	}
}

func TestUnionOf_modify(t *testing.T) {
	t.Parallel()

	// Arrange
	set1 := set.Of[any](1, 2, 3)
	set2 := set.Of[any](4, 5, 6)

	// Act
	union := set.UnionOf(set1, set2)

	// Modify the original sets
	union.Remove(2)
	union.Remove(6)

	// Assert
	assert.ElementsMatch(t, []any{1, 2, 3}, set1.Values())
	assert.ElementsMatch(t, []any{4, 5, 6}, set2.Values())
}
