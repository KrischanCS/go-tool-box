package set_test

import (
	"math/rand"
	"testing"

	"github.com/KrischanCS/go-tool-box/set"
)

func BenchmarkIntersectionOf(b *testing.B) {

	r := rand.New(rand.NewSource(0))

	sets := []*set.Set[string]{
		set.New("a", "b", "c"),
		set.New("d", "e", "f"),
		set.New("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"),
		set.New("d", "e", "f", "g", "h"),
		set.New("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"),
		set.New("p", "q"),
		set.New("r", "s", "t", "u", "v", "w"),
		set.New("x", "y", "z", "a", "b", "c", "d"),
		set.New("m", "n", "o", "p", "q", "r", "s", "t"),
		set.New("e", "f", "g"),
		set.New("h", "i", "j", "k", "l", "m", "n"),
		set.New("y", "z", "a", "b", "c", "d", "e", "f", "g", "h", "i")}

	chosenSets := make([]*set.Set[string], 0, 5)
	for b.Loop() {
		chosenSets = chosenSets[:0]

		amount := rand.Intn(5) + 1

		for range amount {
			chosenSets = append(chosenSets, sets[r.Intn(len(sets))])
		}

		set.IntersectionOf(chosenSets...)
	}
}
