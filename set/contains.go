package set

// Contains checks if the set contains all the given values.
func (s Set[T]) Contains(values ...T) bool {
	for _, v := range values {
		_, ok := s.keySetMap[v]
		if !ok {
			return false
		}
	}

	return true
}

// ContainsExactly checks if the set contains all the given values and no more.
func (s Set[T]) ContainsExactly(values ...T) bool {
	if len(s.keySetMap) != len(values) {
		return false
	}

	return s.Contains(values...)
}
