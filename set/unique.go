package set

// Unique modifies the set, to only contain values which appear only in one set,
// including the set itself and all given other sets.
func (s Set[T]) Unique(others ...Set[T]) {
	if len(others) == 0 {
		return
	}

	m := createElementCounts(append(others, s)...)

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

	elementCounts := createElementCounts(sets...)

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

// createElementCounts creates a map with each element appearing in any of the
// given sets, and the number of times it appears.
func createElementCounts[T comparable](sets ...Set[T]) map[T]int {
	if len(sets) == 0 {
		return nil
	}

	m := make(map[T]int, sets[0].Len())

	for _, s := range sets {
		for v := range s.keySetMap {
			m[v]++
		}
	}

	return m
}
