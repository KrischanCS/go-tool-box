package set

// Difference removes all values from the set that are contained in the other
// sets.
func (s Set[T]) Difference(others ...Set[T]) {
	for _, other := range others {
		for v := range other.keySetMap {
			s.Remove(v)
		}
	}
}

// DifferenceOf creates a new set that contains all values from the first set
// which are not contained in any of the other sets.
func DifferenceOf[T comparable](sets ...Set[T]) Set[T] {
	switch len(sets) {
	case 0:
		return Of[T]()
	case 1:
		return sets[0].Clone()
	}

	s := sets[0].Clone()

	s.Difference(sets[1:]...)

	return s
}
