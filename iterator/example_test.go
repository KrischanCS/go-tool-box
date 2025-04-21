package iterator

import (
	"fmt"
	"slices"
)

func ExampleZip() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	letters := []string{"a", "b", "c", "d", "e", "f"}

	pairs := Zip(
		PickRight(slices.All(numbers)),
		PickRight(slices.All(letters)))

	filtered := Filter(pairs, func(p Pair[int, string]) bool {
		return p.Left%2 == 0
	})

	for s := range SlidingWindow(filtered, 5) {
		for _, p := range s {
			fmt.Println(p.Left, p.Right)
		}
	}

	// Output:
	// 2 b
	// 4 d
	// 6 f
}
