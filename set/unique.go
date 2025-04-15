package set

// Unique modifies the set, to only contain values which appear only in one set,
// including the set itself and all given other sets.
func (s Set[T]) Unique(others ...Set[T]) {
	if len(others) == 0 {
		return
	}

	m := countAppearances(s, others)

	for v, count := range m {
		if count == 1 {
			s.Add(v)
		} else {
			s.Remove(v)
		}
	}
}

// UniqueOf creates a new set that contains all values which appear only in one
// of the given sets.
func UniqueOf[T comparable](sets ...Set[T]) Set[T] {
	switch len(sets) {
	case 0:
		return Of[T]()
	case 1:
		return sets[0].Clone()
	}

	elementCounts := countAppearances(sets[0], sets[1:])

	s := setOfUniqueElements(elementCounts)

	return s
}

func setOfUniqueElements[T comparable](elementCounts map[T]int) Set[T] {
	s := WithCapacity[T](len(elementCounts))

	for v, count := range elementCounts {
		if count == 1 {
			s.Add(v)
		}
	}

	return s
}

// countAppearances counts int how many of the given sets each value appears.
func countAppearances[T comparable](set Set[T], sets []Set[T]) map[T]int {
	m := make(map[T]int, set.Len())

	for _, v := range set.Values() {
		m[v]++
	}

	for _, s := range sets {
		for v := range s.m {
			m[v]++
		}
	}

	return m
}
