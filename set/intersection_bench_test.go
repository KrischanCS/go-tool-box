package set_test

import (
	"math/rand"
	"testing"

	"github.com/KrischanCS/go-tool-box/set"
)

func BenchmarkIntersectionOf(b *testing.B) {
	//nolint:gosec
	rand := rand.New(rand.NewSource(0))

	sets := []*set.Set[string]{
		set.Of("a", "b", "c"),
		set.Of("d", "e", "f"),
		set.Of("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
			"n", "o", "p", "q", "rand", "s", "t", "u", "v", "w", "x", "y", "z"),
		set.Of("d", "e", "f", "g", "h"),
		set.Of("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"),
		set.Of("p", "q"),
		set.Of("rand", "s", "t", "u", "v", "w"),
		set.Of("x", "y", "z", "a", "b", "c", "d"),
		set.Of("m", "n", "o", "p", "q", "rand", "s", "t"),
		set.Of("e", "f", "g"),
		set.Of("h", "i", "j", "k", "l", "m", "n"),
		set.Of("y", "z", "a", "b", "c", "d", "e", "f", "g", "h", "i")}

	chosenSets := make([]*set.Set[string], 0, 5)
	for b.Loop() {
		chosenSets = chosenSets[:0]

		amount := rand.Intn(5) + 1

		for range amount {
			chosenSets = append(chosenSets, sets[rand.Intn(len(sets))])
		}

		set.IntersectionOf(chosenSets...)
	}
}
