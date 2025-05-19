package set_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/set"
)

func ExampleUniqueOf() {
	setA := set.Of(1, 2, 3, 4)
	setB := set.Of(3, 6)
	setC := set.Of(4, 7)

	fmt.Println("A ∆ B:", set.UniqueOf(setA, setB))
	fmt.Println("B ∆ C:", set.UniqueOf(setB, setC))
	fmt.Println("C ∆ A:", set.UniqueOf(setC, setA))
	fmt.Println("A ∆ B ∆ C:", set.UniqueOf(setA, setB, setC))

	fmt.Println()

	fmt.Println("Originals are not modified:")

	fmt.Println("A:", setA)
	fmt.Println("B:", setB)
	fmt.Println("C:", setC)

	// Output:
	// A ∆ B: (Set[int]: [1 2 4 6])
	// B ∆ C: (Set[int]: [3 4 6 7])
	// C ∆ A: (Set[int]: [1 2 3 7])
	// A ∆ B ∆ C: (Set[int]: [1 2 6 7])
	//
	// Originals are not modified:
	// A: (Set[int]: [1 2 3 4])
	// B: (Set[int]: [3 6])
	// C: (Set[int]: [4 7])
}

func ExampleSet_Unique() {
	setA := set.Of(1, 2, 3, 4)
	setB := set.Of(3, 6)

	setA.Unique(setB)
	fmt.Println("A = A ∆ B:", setA)

	setC := set.Of(3, 1, 5)
	setD := set.Of(3, 4)
	setC.Unique(setC, setD)
	fmt.Println("B = B ∆ C ∆ D:", setC)

	// Output:
	// A = A ∆ B: (Set[int]: [1 2 4 6])
	// B = B ∆ C ∆ D: (Set[int]: [4])
}

//nolint:funlen
func TestSet_Unique(t *testing.T) {
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
			name:      "Should be empty set if all sets are empty",
			set:       set.Of[any](),
			otherSets: []set.Set[any]{},
			want:      []any{},
		},
		{
			name:      "Should be empty if all values are in common",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](1, 2, 3)},
			want:      []any{},
		},
		{
			name:      "Should contain all values if they are all different",
			set:       set.Of[any](1, 2, 3),
			otherSets: []set.Set[any]{set.Of[any](4, 5, 6)},
			want:      []any{1, 2, 3, 4, 5, 6},
		},
		{
			name:      "Should contain all values if they are all different with multiple sets",
			set:       set.Of[any](6.283185, 2.718, 9.81),
			otherSets: []set.Set[any]{set.Of[any](1.6605, 1.618), set.Of[any](1.38, 1.602)},
			want:      []any{6.283185, 2.718, 9.81, 1.6605, 1.618, 1.38, 1.602},
		},
		{
			name: "Should contain all values which are unique over all sets",
			set:  set.Of[any](point{1, 2}, point{3, 4}),
			otherSets: []set.Set[any]{
				set.Of[any](point{9, 10}, point{3, 4}, point{5, 6}),
				set.Of[any](point{3, 4}, point{7, 8}, point{9, 10}),
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
		sets []set.Set[any]
		want []any
	}

	tests := []test{
		{
			name: "Should create a new, empty set if no sets are given",
			sets: []set.Set[any]{},
			want: []any{},
		},
		{
			name: "Should create a copy of given set if only one is given",
			sets: []set.Set[any]{set.Of[any]("a", "b", "c")},
			want: []any{"a", "b", "c"},
		},
		{
			name: "Should create empty set if all sets are empty",
			sets: []set.Set[any]{set.Of[any](), set.Of[any](), set.Of[any]()},
			want: []any{},
		},
		{
			name: "Should create an empty set if all values are common",
			sets: []set.Set[any]{
				set.Of[any](1, 2, 3),
				set.Of[any](1, 2, 3),
			},
			want: []any{},
		},
		{
			name: "Should create a set with all values if they are all different",
			sets: []set.Set[any]{
				set.Of[any](1, 2, 3),
				set.Of[any](4, 5, 6),
			},
			want: []any{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Should create a set with all values if they are all different with multiple sets",
			sets: []set.Set[any]{
				set.Of[any](6.283185, 2.718, 9.81),
				set.Of[any](1.6605, 1.618),
				set.Of[any](1.38, 1.602),
			},
			want: []any{6.283185, 2.718, 9.81, 1.6605, 1.618, 1.38, 1.602},
		},
		{
			name: "Should create a set with all values which are unique over all sets",
			sets: []set.Set[any]{
				set.Of[any](point{1, 2}, point{3, 4}),
				set.Of[any](point{9, 10}, point{3, 4}, point{5, 6}),
				set.Of[any](point{3, 4}, point{7, 8}, point{9, 10}),
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

func BenchmarkSet_Unique(b *testing.B) {
	b.ReportAllocs()

	setA, setB, setC := createThreeBigSetsWithDifferentOverlaps()

	for b.Loop() {
		b.StopTimer()

		setA := setA.Clone()

		b.StartTimer()

		setA.Unique(setB, setC)
	}
}

func BenchmarkUniqueOf(b *testing.B) {
	b.ReportAllocs()

	setA, setB, setC := createThreeBigSetsWithDifferentOverlaps()

	for b.Loop() {
		_ = set.UniqueOf(setA, setB, setC)
	}
}

func BenchmarkSet_UniqueHuge(b *testing.B) {
	b.ReportAllocs()

	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	for b.Loop() {
		b.StopTimer()

		setA := setA.Clone()

		b.StartTimer()

		setA.Unique(setB, setC)
	}
}

func BenchmarkSet_UniqueHuge_2Sets(b *testing.B) {
	b.ReportAllocs()

	setA, setB, _ := createThreeHugeSetsWithDifferentOverlaps()

	for b.Loop() {
		b.StopTimer()

		setA := setA.Clone()

		b.StartTimer()

		setA.Unique(setB)
	}
}

func BenchmarkSet_UniqueHuge_6Sets(b *testing.B) {
	b.ReportAllocs()

	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	for b.Loop() {
		b.StopTimer()

		setAClone := setA.Clone()

		b.StartTimer()

		setAClone.Unique(setB, setC, setA, setB, setC)
	}
}

func BenchmarkUniqueOfHuge(b *testing.B) {
	b.ReportAllocs()

	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	for b.Loop() {
		_ = set.UniqueOf(setA, setB, setC)
	}
}

func BenchmarkUniqueOfHuge_2Sets(b *testing.B) {
	b.ReportAllocs()

	setA, setB, _ := createThreeHugeSetsWithDifferentOverlaps()

	for b.Loop() {
		_ = set.UniqueOf(setA, setB)
	}
}

func BenchmarkUniqueOfHuge_6Sets(b *testing.B) {
	b.ReportAllocs()

	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	for b.Loop() {
		_ = set.UniqueOf(setA, setB, setA, setB, setC)
	}
}

// createThreeBigSetsWithDifferentOverlaps creates three sets with the followin Overlaps:
//
//	|     -100----0----100--------400--425--500--475--------800---900----999
//	| A:          AAAAAAAAAAAAAAAAAAAAAAAAAA
//	| B:                          BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
//	| C:  CCCCCCCCCCCCCCCC             CCCCCCCCCC           CCCCCCCCCCCCCCCC
func createThreeBigSetsWithDifferentOverlaps() (set.Set[int], set.Set[int], set.Set[int]) {
	startA := 0
	endA := 500
	startB := 400
	endB := 900
	startC1 := -100
	endC1 := 100
	startC2 := 425
	endC2 := 475
	startC3 := 800
	endC3 := 1000

	return createTestSets(endA, startA, endB, startB, endC1, startC1, endC2, startC2, endC3, startC3)
}

// createThreeBigSetsWithDifferentOverlaps creates three sets with the followin Overlaps:
//
//	|     -10000----0----10000--------40000--42500--50000--47500--------80000---90000----99999
//	| A:            AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
//	| B:                              BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
//	| C:  CCCCCCCCCCCCCCC             CCCCCCCCCCCCCCCCCCCCC             CCCCCCCCCCCCCCCCCCCCCC
func createThreeHugeSetsWithDifferentOverlaps() (set.Set[int], set.Set[int], set.Set[int]) {
	startA := 0
	endA := 50_000
	startB := 40_000
	endB := 90_000
	startC1 := -1_0000
	endC1 := 10_000
	startC2 := 42_500
	endC2 := 47_500
	startC3 := 80_000
	endC3 := 10_0000

	return createTestSets(endA, startA, endB, startB, endC1, startC1, endC2, startC2, endC3, startC3)
}

func createTestSets(
	endA int, startA int,
	endB int, startB int,
	endC1 int, startC1 int,
	endC2 int, startC2 int,
	endC3 int, startC3 int) (set.Set[int], set.Set[int], set.Set[int]) {
	setA := set.WithCapacity[int](endA - startA)
	iterator.Reduce(iterator.FromTo(startA, endA), &setA, func(acc *set.Set[int], i int) {
		acc.Add(i)
	})

	setB := set.WithCapacity[int](endB - startB)
	iterator.Reduce(iterator.FromTo(startB, endB), &setB, func(acc *set.Set[int], i int) {
		acc.Add(i)
	})

	setC := set.WithCapacity[int](endC1 - startC1 + endC2 - startC2 + endC3 - startC3)
	iterator.Reduce(
		iterator.Concat(
			iterator.FromTo(startC1, endC1),
			iterator.FromTo(startC2, endC2),
			iterator.FromTo(startC3, endC3),
		),
		&setC,
		func(acc *set.Set[int], i int) {
			acc.Add(i)
		},
	)

	return setA, setB, setC
}
