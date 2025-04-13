package set

// Intersection removes all values from the set that are not contained in all
// other given sets.
func (s *Set[T]) Intersection(others ...*Set[T]) {
	for v := range s.m {
		for _, other := range others {
			if !other.Contains(v) {
				s.Remove(v)
				break
			}
		}
	}
}

// IntersectionOf creates a new set that contains the values which are present
// in all given sets.
func IntersectionOf[T comparable](sets ...*Set[T]) *Set[T] {
	if len(sets) == 0 {
		return New[T]()
	}

	s := sets[0].Clone()

	if len(sets) == 1 {
		return s
	}

	s.Intersection(sets[1:]...)

	return s
}
