package set_test

import (
	"math/rand"
	"testing"

	"github.com/KrischanCS/go-tool-box/set"
)

//nolint:gochecknoglobals
var operationBenchmarkSets = []*set.Set[string]{
	set.Of("a", "b", "c"),
	set.Of("d", "e", "f"),
	set.Of("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "s", "t", "u", "v", "w", "x", "y", "z"),
	set.Of("d", "e", "f", "g", "h"),
	set.Of("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"),
	set.Of("p", "q"),
	set.Of("s", "t", "u", "v", "w"),
	set.Of("x", "y", "z", "a", "b", "c", "d"),
	set.Of("m", "n", "o", "p", "q", "s", "t"),
	set.Of("e", "f", "g"),
	set.Of("h", "i", "j", "k", "l", "m", "n"),
	set.Of("y", "z", "a", "b", "c", "d", "e", "f", "g", "h", "i")}

func BenchmarkSet_Union(b *testing.B) {
	benchOp(b, func(sets []*set.Set[string]) {
		sets[0].Union(sets[1:]...)
	})
}

func BenchmarkSet_Intersection(b *testing.B) {
	benchOp(b, func(sets []*set.Set[string]) {
		sets[0].Intersection(sets[1:]...)
	})
}

func BenchmarkSet_Difference(b *testing.B) {
	benchOp(b, func(sets []*set.Set[string]) {
		sets[0].Difference(sets[1:]...)
	})
}
func BenchmarkSet_Unique(b *testing.B) {
	benchOp(b, func(sets []*set.Set[string]) {
		sets[0].Unique(sets[1:]...)
	})
}

func BenchmarkUnionOf(b *testing.B) {
	b.ReportAllocs()

	//nolint:gosec
	rand := rand.New(rand.NewSource(0))

	chosenSets := make([]*set.Set[string], 0, 5)
	for b.Loop() {
		chosenSets = chosenSets[:0]

		amount := rand.Intn(5) + 1

		for range amount {
			chosenSets = append(chosenSets, operationBenchmarkSets[rand.Intn(len(operationBenchmarkSets))])
		}

		_ = set.UnionOf(chosenSets...)
	}
}

func BenchmarkDifferenceOf(b *testing.B) {
	b.ReportAllocs()

	//nolint:gosec
	rand := rand.New(rand.NewSource(0))

	chosenSets := make([]*set.Set[string], 0, 5)
	for b.Loop() {
		chosenSets = chosenSets[:0]

		amount := rand.Intn(5) + 1

		for range amount {
			chosenSets = append(chosenSets, operationBenchmarkSets[rand.Intn(len(operationBenchmarkSets))])
		}

		_ = set.DifferenceOf(chosenSets...)
	}
}

func BenchmarkIntersectionOf(b *testing.B) {
	b.ReportAllocs()

	//nolint:gosec
	rand := rand.New(rand.NewSource(0))

	chosenSets := make([]*set.Set[string], 0, 5)
	for b.Loop() {
		chosenSets = chosenSets[:0]

		amount := rand.Intn(5) + 1

		for range amount {
			chosenSets = append(chosenSets, operationBenchmarkSets[rand.Intn(len(operationBenchmarkSets))])
		}

		set.IntersectionOf(chosenSets...)
	}
}

func BenchmarkUniqueOf(b *testing.B) {
	b.ReportAllocs()

	//nolint:gosec
	rand := rand.New(rand.NewSource(0))

	chosenSets := make([]*set.Set[string], 0, 5)
	for b.Loop() {
		chosenSets = chosenSets[:0]

		amount := rand.Intn(5) + 1

		for range amount {
			chosenSets = append(chosenSets, operationBenchmarkSets[rand.Intn(len(operationBenchmarkSets))])
		}

		_ = set.UnionOf(chosenSets...)
	}
}

func benchOp(b *testing.B, fn func(sets []*set.Set[string])) {
	b.Helper()
	b.ReportAllocs()

	//nolint:gosec
	rand := rand.New(rand.NewSource(0))

	chosenSets := make([]*set.Set[string], 0, 5)
	for b.Loop() {
		chosenSets = chosenSets[:0]

		amount := rand.Intn(5) + 1

		for range amount {
			chosenSets = append(chosenSets, operationBenchmarkSets[rand.Intn(len(operationBenchmarkSets))])
		}

		fn(chosenSets)
	}
}
