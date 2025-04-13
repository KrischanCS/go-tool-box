// Package must provides functions which panics on error and returns the
// value(s) passed to them.
//
// This is usually not recommended, but can be helpful in tests or
// initialization of an application.
//
// It can be used like [regex.MustCompile] or [template.MustParse] but for
// arbitrary functions.
package must

import "fmt"

// Value returns the value passed to it err is nil and panics otherwise.
func Value[T any](t T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("Error in call to must.Value: [%T] - %s", err, err))
	}

	return t
}

// Values returns the values passed to it if err is nil and panics otherwise.
func Values[T, R any](t T, r R, err error) (T, R) {
	if err != nil {
		panic(fmt.Sprintf("Error in call to must.Values: [%T] - %s", err, err))
	}

	return t, r
}
