package set

// Union adds all values from the given sets to the current set.
func (s Set[T]) Union(sets ...Set[T]) {
	for _, set := range sets {
		for v := range set.m {
			s.Add(v)
		}
	}
}

// UnionOf creates a new set that contains all values from the given sets.
func UnionOf[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return Of[T]()
	}

	s := sets[0].Clone()

	s.Union(sets[1:]...)

	return s
}
