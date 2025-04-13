package set

// Unique modifies the set, to only contain values which appear only in one set,
// including the set itself and all given other sets.
func (s *Set[T]) Unique(others ...*Set[T]) {
	if len(others) == 0 {
		return
	}

	m := make(map[T]int, len(s.m))

	for v := range s.m {
		m[v]++
	}

	for _, other := range others {
		for v := range other.m {
			m[v]++
		}
	}

	for v, count := range m {
		if count > 1 {
			s.Remove(v)
		} else {
			s.Add(v)
		}
	}
}

// UniqueOf creates a new set that contains all values which appear only in one
// of the given sets.
func UniqueOf[T comparable](sets ...*Set[T]) *Set[T] {
	switch len(sets) {
	case 0:
		return New[T]()
	case 1:
		return sets[0].Clone()
	}

	m := make(map[T]int, sets[0].Len())

	for _, s := range sets {
		for v := range s.m {
			m[v]++
		}
	}

	s := Set[T]{
		m: make(map[T]struct{}, len(m)),
	}

	for v, count := range m {
		if count > 1 {
			s.Remove(v)
		} else {
			s.Add(v)
		}
	}

	return &s
}
