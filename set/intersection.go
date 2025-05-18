package set

// Intersection removes all values from the set that are not contained in all
// other given sets.
func (s Set[T]) Intersection(others ...Set[T]) {
	for v := range s.keySetMap {
		if !allContains(others, v) {
			s.Remove(v)
		}
	}
}

// IntersectionOf creates a new set that contains the values which are present
// in all given sets.
func IntersectionOf[T comparable](sets ...Set[T]) Set[T] {
	switch len(sets) {
	case 0:
		return Of[T]()
	case 1:
		return sets[0].Clone()
	}

	// If length of the given sets varies a lot, it is faster to check only it's
	// elements in the other sets. Also less memory is allocated when cloning
	// the shortest set.
	swapShortestFirst(sets)

	s := sets[0].Clone()

	s.Intersection(sets[1:]...)

	return s
}

func allContains[T comparable](others []Set[T], v T) bool {
	for _, other := range others {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

func swapShortestFirst[T comparable](sets []Set[T]) {
	indexShortest := 0
	for i, set := range sets {
		if len(set.keySetMap) < len(sets[indexShortest].keySetMap) {
			indexShortest = i
		}
	}

	if indexShortest != 0 {
		sets[0], sets[indexShortest] = sets[indexShortest], sets[0]
	}
}
