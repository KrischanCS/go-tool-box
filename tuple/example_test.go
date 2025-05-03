package tuple_test

import (
	"fmt"

	"github.com/KrischanCS/go-toolbox/tuple"
)

func ExamplePair_Unpack() {
	p := tuple.PairOf(1, "b")

	first, second := p.Unpack()

	fmt.Println(first)
	fmt.Println(second)

	// Output:
	// 1
	// b
}

func ExamplePair_First() {
	p := tuple.PairOf(1, "b")

	first := p.First()
	fmt.Println(first)

	// Output:
	// 1
}

func ExamplePair_Second() {
	p := tuple.PairOf(1, "b")

	second := p.Second()

	fmt.Println(second)

	// Output:
	// b
}

func ExamplePair_String() {
	p := tuple.PairOf(1, "b")

	fmt.Println(p.String())

	// Output:
	// (Pair[int, string]: [1; b])
}

func ExampleTriple_Unpack() {
	t := tuple.TripleOf(1, "b", 6.28)

	first, second, third := t.Unpack()

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	// Output:
	// 1
	// b
	// 6.28
}

func ExampleTriple_First() {
	t := tuple.TripleOf(1, "b", 6.28)

	first := t.First()
	fmt.Println(first)

	// Output:
	// 1
}

func ExampleTriple_Second() {
	t := tuple.TripleOf(1, "b", 6.28)

	second := t.Second()
	fmt.Println(second)

	// Output:
	// b
}

func ExampleTriple_Third() {
	t := tuple.TripleOf(1, "b", 6.28)

	third := t.Third()
	fmt.Println(third)

	// Output:
	// 6.28
}

func ExampleTriple_String() {
	t := tuple.TripleOf(1, "b", 6.28)

	fmt.Println(t.String())

	// Output:
	// (Triple[int, string, float64]: [1; b; 6.28])
}
