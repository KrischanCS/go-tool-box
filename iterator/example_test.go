package iterator

import (
	"fmt"
	"slices"

	"github.com/KrischanCS/go-toolbox/tuple"
)

func ExampleZip() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	letters := []string{"a", "b", "c", "d", "e", "f"}

	pairs := Zip(
		PickRight(slices.All(numbers)),
		PickRight(slices.All(letters)))

	filtered := Filter(pairs, func(p tuple.Pair[int, string]) bool {
		return p.First()%2 == 0
	})

	for s := range SlidingWindow(filtered, 5) {
		for _, p := range s {
			fmt.Println(p.First(), p.Second())
		}
	}

	// Output:
	// 2 b
	// 4 d
	// 6 f
}
