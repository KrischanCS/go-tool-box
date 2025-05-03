package object

import (
	"regexp"
	"strconv"
)

var setArrayElemPattern = regexp.MustCompile(`^(.+)\[(\d*)]$`)

// Set sets the value at the given path. If the path does not exist, it will be
// created.
//
// Setting a value in an array or nested in an array, requires the type of the
// array to be []any/[]interface{}.
//
// If an array index is included the index is >= len(array), the value
// will be appended to the array. Using [] without an index will also append the
//
// If the path is empty, the function does nothing.
func Set(object Object, value any, path ...string) {
	if len(path) == 0 {
		return
	}

	set(object, value, path...)
}

func set(object Object, value any, path ...string) {
	nextKey, path := path[0], path[1:]

	nextKeyParts := setArrayElemPattern.FindStringSubmatch(nextKey)
	if nextKeyParts != nil {
		nextKey := nextKeyParts[1]
		index := nextKeyParts[2]

		setArrayElem(object, value, nextKey, index, path...)

		return
	}

	if len(path) == 0 {
		object[nextKey] = value
		return
	}

	nextElem, ok := object[nextKey].(Object)
	if !ok {
		object[nextKey] = buildObject(path, value)
		return
	}

	set(nextElem, value, path...)
}

func setArrayElem(object Object, value any, nextKey, index string, path ...string) {
	next, ok := object[nextKey]
	if !ok {
		array := Array{buildObject(path, value)}
		object[nextKey] = array

		return
	}

	var i int

	if index != "" {
		var err error

		i, err = strconv.Atoi(index)
		if err != nil {
			panic("Parsing index in object-path: " + err.Error())
		}
	}

	nextArray, ok := next.([]any)
	if !ok {
		array := Array{buildObject(path, buildObject(path, value))}
		object[nextKey] = array

		return
	}

	if index == "" || i >= len(nextArray) {
		nextArray = append(nextArray, buildObject(path, value))
		object[nextKey] = nextArray

		return
	}

	nextArray[i] = buildObject(path, value)
}

func buildObject(path []string, value any) any {
	if len(path) == 0 {
		return value
	}

	return Object{
		path[0]: buildObject(path[1:], value),
	}
}
