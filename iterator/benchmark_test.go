package iterator

import (
	"encoding/json"
	"maps"
	"slices"
	"strconv"
	"testing"
)

// This benchmarks compares the performance between the iterators and the
// equivalent implementation using a for loop.

// FromTo

const (
	from    = 137
	to      = 2341
	breakAt = 1742
	step    = 0.879
)

func BenchmarkFromTo(b *testing.B) {
	iterator := FromTo(from, to)

	for b.Loop() {
		res := 0
		for v := range iterator {
			res += v*3 - 1

			if v == breakAt {
				break
			}
		}
	}
}

func BenchmarkFromTo_For(b *testing.B) {
	for b.Loop() {
		res := 0

		for i := from; i < to; i++ {
			res += i*3 - 1

			if i == breakAt {
				break
			}
		}
	}
}

func BenchmarkFromToInclusive(b *testing.B) {
	iterator := FromToInclusive(from, to)

	for b.Loop() {
		res := 0
		for v := range iterator {
			res += v*3 - 1

			if v == breakAt {
				break
			}
		}
	}
}

func BenchmarkFromToInclusive_For(b *testing.B) {
	for b.Loop() {
		res := 0

		for i := from; i <= to; i++ {
			res += i*3 - 1

			if i == breakAt {
				break
			}
		}
	}
}

func BenchmarkFromStepTo(b *testing.B) {
	iterator := FromStepTo(float64(from), step, float64(to))

	for b.Loop() {
		res := 0
		for v := range iterator {
			res += int(v*3) - 1

			if v >= breakAt {
				break
			}
		}
	}
}

func BenchmarkFromStepTo_For(b *testing.B) {
	for b.Loop() {
		res := float64(0)

		for i := float64(from); i < float64(to); i += step {
			res += i*3 - 1

			if i >= breakAt {
				break
			}
		}
	}
}

// Filter

func BenchmarkFilter(b *testing.B) {
	iterator := FromTo(from, to)

	divisibleByThree := func(i int) bool {
		return i%3 == 0
	}

	for b.Loop() {
		res := 0

		for v := range Filter(iterator, divisibleByThree) {
			res += v*3 - 1

			if v >= breakAt {
				break
			}
		}
	}
}

//nolint:gocognit
func BenchmarkFilter_For(b *testing.B) {
	divisibleByThree := func(i int) bool {
		return i%3 == 0
	}

	for b.Loop() {
		res := 0

		for i := from; i < to; i++ {
			if !divisibleByThree(i) {
				continue
			}

			res += i*3 - 1

			if i >= breakAt {
				break
			}
		}
	}
}

// Of

func BenchmarkOf(b *testing.B) {
	iterator := Of(slices.Collect(FromTo(from, to))...)

	for b.Loop() {
		res := 0
		for v := range iterator {
			res += v*3 - 1

			if v == breakAt {
				break
			}
		}
	}
}

func BenchmarkOf_For(b *testing.B) {
	slice := slices.Collect(FromTo(from, to))

	for b.Loop() {
		res := 0

		for _, v := range slice {
			res += v*3 - 1

			if v == breakAt {
				break
			}
		}
	}
}

// Pick

func BenchmarkPickLeft(b *testing.B) {
	m := make(map[string]int)

	for v := range FromTo(from, to) {
		str := strconv.Itoa(v)
		m[str] = v
	}

	values := maps.All(m)
	breakAt := strconv.Itoa(breakAt)

	for b.Loop() {
		res := ""
		for value := range PickLeft(values) {
			res += value

			if value == breakAt {
				break
			}
		}
	}
}

func BenchmarkPickLeft_For(b *testing.B) {
	m := make(map[string]int)

	for v := range FromTo(from, to) {
		str := strconv.Itoa(v)
		m[str] = v
	}

	breakAt := strconv.Itoa(breakAt)

	for b.Loop() {
		res := ""
		for key := range m {
			res += key

			if key == breakAt {
				break
			}
		}
	}
}

func BenchmarkPickRight(b *testing.B) {
	m := make(map[string]int)

	for v := range FromTo(from, to) {
		str := strconv.Itoa(v)
		m[str] = v
	}

	values := maps.All(m)

	for b.Loop() {
		res := 0
		for value := range PickRight(values) {
			res += value*3 - 1

			if value == breakAt {
				break
			}
		}
	}
}

func BenchmarkPickRight_For(b *testing.B) {
	m := make(map[string]int)

	for v := range FromTo(from, to) {
		str := strconv.Itoa(v)
		m[str] = v
	}

	for b.Loop() {
		res := 0
		for _, value := range m {
			res += value*3 - 1

			if value == breakAt {
				break
			}
		}
	}
}

func BenchmarkCombine(b *testing.B) {
	m := make(map[string]int)

	for v := range FromTo(from, to) {
		str := strconv.Itoa(v)
		m[str] = v
	}

	values := maps.All(m)

	for b.Loop() {
		strRes := ""
		intRes := 0

		for pair := range Combine(values) {
			strRes += pair.Left
			intRes += pair.Right*3 - 1

			if pair.Right == breakAt {
				break
			}
		}
	}
}

func BenchmarkCombine_For(b *testing.B) {
	m := make(map[string]int)

	for v := range FromTo(from, to) {
		str := strconv.Itoa(v)
		m[str] = v
	}

	for b.Loop() {
		strRes := ""
		intRes := 0

		for key, value := range m {
			strRes += key
			intRes += value*3 - 1

			if value == breakAt {
				break
			}
		}
	}
}

// Window

//nolint:gocognit
func BenchmarkFixedWindow(b *testing.B) {
	iterator := FromTo(from, to)

	for b.Loop() {
		res := 0

	WINDOW:
		for window := range FixedWindow(iterator, 4) {
			for _, v := range window {
				res += v*3 - 1
				if v == breakAt {
					break WINDOW
				}
			}
		}
	}
}

//nolint:gocognit
func BenchmarkFixedWindow_For(b *testing.B) {
	slice := slices.Collect(FromTo(from, to))
	windowLen := 4

	for b.Loop() {
		res := 0

	WINDOW:
		for i := 0; i < len(slice); i += windowLen {
			for j := range windowLen {
				if j >= len(slice) {
					break WINDOW
				}

				v := slice[j]
				res += v*3 - 1

				if v == breakAt {
					break WINDOW
				}
			}
		}
	}
}

//nolint:gocognit
func BenchmarkSlidingWindow(b *testing.B) {
	iterator := FromTo(from, to)

	for b.Loop() {
		res := 0

	WINDOW:
		for window := range SlidingWindow(iterator, 4) {
			for _, v := range window {
				res += v*3 - 1
				if v == breakAt {
					break WINDOW
				}
			}
		}
	}
}

//nolint:gocognit
func BenchmarkSlidingWindow_For(b *testing.B) {
	slice := slices.Collect(FromTo(from, to))
	windowLen := 4

	for b.Loop() {
		res := 0

	WINDOW:
		for range len(slice) - windowLen {
			for j := range windowLen {
				v := slice[j]
				res += v*3 - 1

				if v == breakAt {
					break WINDOW
				}
			}
		}
	}
}

// Zip

func BenchmarkZip(b *testing.B) {
	iterator1 := FromTo(from, to)

	slice2 := make([]string, 0, breakAt-from)
	for i := range iterator1 {
		slice2 = append(slice2, strconv.Itoa(i))
	}

	iterator2 := Of(slice2...)

	for b.Loop() {
		intRes := 0
		strRes := ""

		for pair := range Zip[int, string](iterator1, iterator2) {
			intRes += pair.Left*3 - 1
			strRes += pair.Right

			if pair.Left == breakAt {
				break
			}
		}
	}
}

//nolint:gocognit
func BenchmarkZip_For(b *testing.B) {
	slice1 := slices.Collect(FromTo(from, to))

	slice2 := make([]string, 0, breakAt-from)
	for i := range slice1 {
		slice2 = append(slice2, strconv.Itoa(i))
	}

	for b.Loop() {
		intRes := 0
		strRes := ""

		for i, v := range slice1 {
			if i > len(slice2)-1 {
				break
			}

			intRes += v*3 - 1
			strRes += slice2[i]

			if v == breakAt {
				break
			}
		}
	}
}

// Map

func BenchmarkMap(b *testing.B) {
	iterator := FromTo(from, to)

	for b.Loop() {
		res := 0
		for v := range Map(iterator, func(i int) int { return i*3 - 1 }) {
			res += v

			if v == breakAt {
				break
			}
		}
	}
}

func BenchmarkMap_For(b *testing.B) {
	slice := slices.Collect(FromTo(from, to))

	for b.Loop() {
		res := 0

		for _, v := range slice {
			v = v*3 - 1
			res += v

			if v == breakAt {
				break
			}
		}
	}
}

// Reduce

func BenchmarkReduce(b *testing.B) {
	iterator := FromTo(from, to)

	for b.Loop() {
		acc := 0
		Reduce(iterator, &acc, func(acc *int, v int) {
			*acc += v*3 - 1

			if v == breakAt {
				return
			}
		})
	}
}

func BenchmarkReduce_For(b *testing.B) {
	slice := slices.Collect(FromTo(from, to))

	for b.Loop() {
		res := 0
		acc := &res

		for _, v := range slice {
			*acc = v*3 - 1

			if v == breakAt {
				break
			}
		}
	}
}

// Unique

//nolint:gocognit
func BenchmarkUnique(b *testing.B) {
	tmp := FromTo(from, to)

	slice := make([]int, 0, (breakAt-from)*5)

	for v := range tmp {
		for range 5 {
			slice = append(slice, v)
		}
	}

	iterator := Of(slice...)

	for b.Loop() {
		res := 0
		for v := range Unique(iterator) {
			res += v*3 - 1

			if v == breakAt {
				break
			}
		}
	}
}

//nolint:gocognit
func BenchmarkUnique_For(b *testing.B) {
	tmp := FromTo(from, to)

	slice := make([]int, 0, (breakAt-from)*5)

	for v := range tmp {
		for range 5 {
			slice = append(slice, v)
		}
	}

	for b.Loop() {
		set := make(map[int]struct{})

		res := 0

		for v := range slice {
			if _, ok := set[v]; ok {
				continue
			}

			res += v*3 - 1

			if v == breakAt {
				break
			}

			set[v] = struct{}{}
		}
	}
}

// Complex Iterators

func BenchmarkComplexIterators(b *testing.B) {
	numbers := make([]int, 0, 1000)
	letters := make([]string, 0, 1000)

	for i := range 1000 {
		numbers = append(numbers, i)
		letters = append(letters, string(byte(i%26+'a')))
	}

	var res []Pair[int, string]

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		pairs := Zip(
			PickRight(slices.All(numbers)),
			PickRight(slices.All(letters)),
		)

		filtered := Filter(pairs, func(p Pair[int, string]) bool {
			return p.Left%2 == 0
		})

		for s := range SlidingWindow(filtered, 5) {
			res = slices.Clone(s)
		}
	}

	_ = res
}

//nolint:gocognit
func BenchmarkComplexIteration_For(b *testing.B) {
	numbers := make([]int, 0, 1000)
	letters := make([]string, 0, 1000)

	for i := range 1000 {
		numbers = append(numbers, i)
		letters = append(letters, string(byte(i%26+'a')))
	}

	var res []Pair[int, string]

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		n := 5

		state := make([]Pair[int, string], 0, n)

		for i, number := range numbers {
			if number%2 == 0 {
				continue
			}

			current := Pair[int, string]{number, letters[i]}

			switch {
			case len(state) < n-1:
				state = append(state, current)
				continue
			case len(state) == n-1:
				state = append(state, current)
			default:
				for i := range state[:len(state)-1] {
					state[i] = state[i+1]
				}

				state[len(state)-1] = current
			}

			res = slices.Clone(state)
		}
	}

	_ = res
}

// Complex iterators and higher workload

//nolint:gocognit
func BenchmarkComplexIteratorsAndWorkload(b *testing.B) {
	numbers := make([]int, 0, 1000)
	letters := make([]string, 0, 1000)

	for i := range 1000 {
		numbers = append(numbers, i)
		letters = append(letters, string(byte(i%26+'a')))
	}

	var res []any

	type intString struct {
		Number int    `json:"number"`
		Text   string `json:"text"`
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		res = res[:0]

		pairs := Zip(
			PickRight(slices.All(numbers)),
			PickRight(slices.All(letters)),
		)

		filtered := Filter(pairs, func(p Pair[int, string]) bool {
			return p.Left%2 == 0
		})

		for s := range SlidingWindow(filtered, 5) {
			t := []intString{
				{s[0].Left, s[0].Right},
				{s[1].Left, s[1].Right},
				{s[2].Left, s[2].Right},
				{s[3].Left, s[3].Right},
				{s[4].Left, s[4].Right},
			}

			v, err := json.Marshal(t)
			if err != nil {
				b.Fatal(err)
			}

			var dst any

			err = json.Unmarshal(v, &dst)
			if err != nil {
				panic(err)
			}

			res = append(res, dst)
		}
	}

	_ = res
}

//nolint:gocognit,funlen
func BenchmarkComplexIteratorsAndWorkload_For(b *testing.B) {
	numbers := make([]int, 0, 1000)
	letters := make([]string, 0, 1000)

	for i := range 1000 {
		numbers = append(numbers, i)
		letters = append(letters, string(byte(i%26+'a')))
	}

	var res []any

	type intString struct {
		Number int    `json:"number"`
		Text   string `json:"text"`
	}

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		res = res[:0]

		n := 5
		state := make([]Pair[int, string], 0, n)

		for i, number := range numbers {
			if number%2 != 0 {
				continue
			}

			current := Pair[int, string]{number, letters[i]}

			switch {
			case len(state) < n-1:
				state = append(state, current)
				continue
			case len(state) == n-1:
				state = append(state, current)
			default:
				for i := range state[:len(state)-1] {
					state[i] = state[i+1]
				}

				state[len(state)-1] = current
			}

			t := []intString{
				{state[0].Left, state[0].Right},
				{state[1].Left, state[1].Right},
				{state[2].Left, state[2].Right},
				{state[3].Left, state[3].Right},
				{state[4].Left, state[4].Right},
			}

			v, err := json.Marshal(t)
			if err != nil {
				b.Fatal(err)
			}

			var dst any

			err = json.Unmarshal(v, &dst)
			if err != nil {
				panic(err)
			}

			res = append(res, dst)
		}
	}
}
