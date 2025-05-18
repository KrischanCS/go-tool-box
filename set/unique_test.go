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

//nolint:funlen
func TestSet_UniqueAlternative(t *testing.T) {
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
			s.UniqueAlternative(others...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}

//nolint:funlen
func TestUniqueOfAlternative(t *testing.T) {
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
			s := set.UniqueOfAlternative(tc.sets...)

			// Assert
			assert.ElementsMatch(t, tc.want, s.Values())
		})
	}
}

func BenchmarkSet_Unique(b *testing.B) {
	setA, setB, setC := createThreeBigSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		b.StopTimer()
		setA := setA.Clone()
		b.StartTimer()

		setA.Unique(setB, setC)
	}
}

func BenchmarkSet_UniqueAlternative(b *testing.B) {
	setA, setB, setC := createThreeBigSetsWithDifferentOverlaps()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		setA := setA.Clone()
		b.StartTimer()

		setA.UniqueAlternative(setB, setC)
	}
}

func BenchmarkUniqueOf(b *testing.B) {
	setA, setB, setC := createThreeBigSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOf(setA, setB, setC)
	}
}

func BenchmarkUniqueOfAlternative(b *testing.B) {
	setA, setB, setC := createThreeBigSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOfAlternative(setA, setB, setC)
	}
}

func BenchmarkSet_UniqueHuge(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		b.StopTimer()
		setA := setA.Clone()
		b.StartTimer()

		setA.Unique(setB, setC)
	}
}

func BenchmarkSet_UniqueAlternativeHuge(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		setA := setA.Clone()
		b.StartTimer()

		setA.UniqueAlternative(setB, setC)
	}
}

func BenchmarkSet_UniqueHuge_2Sets(b *testing.B) {
	setA, setB, _ := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		b.StopTimer()
		setA := setA.Clone()
		b.StartTimer()

		setA.Unique(setB)
	}
}

func BenchmarkSet_UniqueAlternativeHuge_2Sets(b *testing.B) {
	setA, setB, _ := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		setA := setA.Clone()
		b.StartTimer()

		setA.UniqueAlternative(setB)
	}
}

func BenchmarkSet_UniqueHuge_6Sets(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		b.StopTimer()
		setAClone := setA.Clone()
		b.StartTimer()

		setAClone.Unique(setB, setC, setA, setB, setC)
	}
}

func BenchmarkSet_UniqueAlternativeHuge_6Sets(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		setAClone := setA.Clone()
		b.StartTimer()

		setAClone.UniqueAlternative(setB, setC, setA, setB, setC)
	}
}

func BenchmarkUniqueOfHuge(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOf(setA, setB, setC)
	}
}

func BenchmarkUniqueOfAlternativeHuge(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOfAlternative(setA, setB, setC)
	}
}

func BenchmarkUniqueOfHuge_2Sets(b *testing.B) {
	setA, setB, _ := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOf(setA, setB)
	}
}

func BenchmarkUniqueOfAlternativeHuge_2Sets(b *testing.B) {
	setA, setB, _ := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOfAlternative(setA, setB)
	}
}

func BenchmarkUniqueOfHuge_6Sets(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOf(setA, setB, setA, setB, setC)
	}
}

func BenchmarkUniqueOfAlternativeHuge_6Sets(b *testing.B) {
	setA, setB, setC := createThreeHugeSetsWithDifferentOverlaps()

	b.ReportAllocs()
	for b.Loop() {
		_ = set.UniqueOfAlternative(setA, setB, setA, setB, setC)
	}
}

// createThreeBigSetsWithDifferentOverlaps creates three sets with the followin Overlaps:
//
//	|     -100----0----100--------400--425--500--475--------800---900----999
//	| A:          AAAAAAAAAAAAAAAAAAAAAAAAAA
//	| B:                          BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
//	| C:  CCCCCCCCCCCCCCCC             CCCCCCCCCC           CCCCCCCCCCCCCCCC
func createThreeBigSetsWithDifferentOverlaps() (set.Set[int], set.Set[int], set.Set[int]) {
	setA := set.WithCapacity[int](500)

	iterator.Reduce(iterator.FromTo(0, 500), &setA, func(acc *set.Set[int], i int) {
		acc.Add(i)
	})

	setB := set.WithCapacity[int](500)
	iterator.Reduce(iterator.FromTo(400, 900), &setB, func(acc *set.Set[int], i int) {
		acc.Add(i)
	})

	setC := set.WithCapacity[int](150)
	iterator.Reduce(
		iterator.Concat(
			iterator.FromTo(-100, 100),
			iterator.FromTo(425, 475),
			iterator.FromTo(800, 1000),
		),
		&setC,
		func(acc *set.Set[int], i int) {
			acc.Add(i)
		},
	)
	return setA, setB, setC
}

// createThreeBigSetsWithDifferentOverlaps creates three sets with the followin Overlaps:
//
//	|     -10000----0----10000--------40000--42500--50000--47500--------80000---90000----99999
//	| A:            AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
//	| B:                              BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
//	| C:  CCCCCCCCCCCCCCC             CCCCCCCCCCCCCCCCCCCCC             CCCCCCCCCCCCCCCCCCCCCC
func createThreeHugeSetsWithDifferentOverlaps() (set.Set[int], set.Set[int], set.Set[int]) {
	setA := set.WithCapacity[int](50_000)

	iterator.Reduce(iterator.FromTo(0, 50_000), &setA, func(acc *set.Set[int], i int) {
		acc.Add(i)
	})

	setB := set.WithCapacity[int](50_000)
	iterator.Reduce(iterator.FromTo(40_000, 90_000), &setB, func(acc *set.Set[int], i int) {
		acc.Add(i)
	})

	setC := set.WithCapacity[int](15_000)
	iterator.Reduce(
		iterator.Concat(
			iterator.FromTo(-10_000, 10_000),
			iterator.FromTo(42_500, 47_500),
			iterator.FromTo(80_000, 10_000),
		),
		&setC,
		func(acc *set.Set[int], i int) {
			acc.Add(i)
		},
	)
	return setA, setB, setC
}
