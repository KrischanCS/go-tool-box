package reducer_test

import (
	"fmt"
	"strings"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/iterator/reducer"
)

func ExampleSum() {
	i := iterator.Of(1, 2, 3, 4, 5)

	sum := 0
	iterator.Reduce(i, &sum, reducer.Sum)

	fmt.Println(sum)

	// Output: 15
}

func ExampleProduct() {
	i := iterator.Of(3, 4)

	product := 1
	iterator.Reduce(i, &product, reducer.Product)

	fmt.Println(product)

	// Output: 12
}

func ExampleJoin() {
	i := iterator.Of("a", "b", "c")

	var sb strings.Builder

	iterator.Reduce(i, &sb, reducer.Join(", "))

	fmt.Println(sb.String())

	// Output: a, b, c
}

func ExampleGroupBy() {
	type person struct {
		name string
		age  int
	}

	persons := iterator.Of(
		person{"Alice", 30},
		person{"Bob", 2},
		person{"Charlie", 13},
		person{"Dory", 25},
		person{"Edward", 8},
		person{"Fiona", 2},
		person{"Greg", 5},
	)

	keyFunc := func(p person) string {
		switch {
		case p.age < 3:
			return "toddler"
		case p.age < 13:
			return "child"
		case p.age < 18:
			return "teen"
		default:
			return "adult"
		}
	}

	groups := make(map[string][]person)
	iterator.Reduce(persons, &groups, reducer.GroupBy(keyFunc))

	// Print results in a deterministic order
	for _, category := range []string{"toddler", "child", "teen", "adult"} {
		if people, exists := groups[category]; exists {
			fmt.Printf("%s: %+v\n", category, people)
		}
	}

	// Output:
	// toddler: [{name:Bob age:2} {name:Fiona age:2}]
	// child: [{name:Edward age:8} {name:Greg age:5}]
	// teen: [{name:Charlie age:13}]
	// adult: [{name:Alice age:30} {name:Dory age:25}]
}
