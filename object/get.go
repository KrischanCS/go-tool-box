package object

// Get returns the value at the given path, when it exists and is of type T.
func Get[T any](object Object, path ...string) (value T, ok bool) {
	return get[T](object, path...)
}

func get[T any](object any, path ...string) (value T, ok bool) {
	if len(path) == 0 {
		o, ok := object.(T)
		if !ok {
			return value, false
		}

		return o, true
	}

	next, path := path[0], path[1:]

	if len(path) == 0 {
		get[T](object)
	}

	o, ok := object.(Object)
	if !ok {
		return value, false
	}

	nextObject, ok := o[next]
	if !ok {
		return value, false
	}

	return get[T](nextObject, path...)
}
