package object

import (
	"reflect"
	"regexp"
	"strconv"
)

var getArrayElemPattern = regexp.MustCompile(`^(.+)\[(\d+)]$`)

// Get returns the value at the given path, when it exists and is of type T.
func Get[T any](object Object, path ...string) (value T, ok bool) {
	return get[T](object, path...)
}

func get[T any](object any, path ...string) (value T, ok bool) {
	if len(path) == 0 {
		return objectAsT[T](object)
	}

	o, ok := object.(Object)
	if !ok {
		return value, false
	}

	next, path := path[0], path[1:]

	nextParts := getArrayElemPattern.FindStringSubmatch(next)
	if nextParts != nil {
		next := nextParts[1]
		index := nextParts[2]

		return getArrayElem[T](o, next, index, path...)
	}

	nextObject, ok := o[next]
	if !ok {
		return value, false
	}

	return get[T](nextObject, path...)
}

func getArrayElem[T any](object Object, next, index string, path ...string) (value T, ok bool) {
	i, err := strconv.Atoi(index)
	if err != nil {
		panic("Parsing index in object-path: " + err.Error())
	}

	nextObj, ok := object[next]
	if !ok {
		return value, false
	}

	reflectValue := reflect.ValueOf(nextObj)
	if reflectValue.Kind() != reflect.Slice && reflectValue.Kind() != reflect.Array {
		return value, false
	}

	if i >= reflectValue.Len() {
		return value, false
	}

	nextObject := reflectValue.Index(i).Interface()

	if len(path) == 0 {
		return objectAsT[T](nextObject)
	}

	return get[T](nextObject, path...)
}

// objectAsT checks if the object is of type T and returns it.
func objectAsT[T any](object any) (value T, ok bool) {
	o, ok := object.(T)
	if !ok {
		return value, false
	}

	return o, true
}
