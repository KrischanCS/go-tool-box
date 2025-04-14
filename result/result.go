// Package result provides a simple result, wrapping either a value or an error.
//
// Main use case is to send it over channels, when possible errors should be
// propagated to the receiver.
//
// Apart from that, after calling `Get()`, it can be used with go's idiomatic
// error handling, like:
//
//	value, err := result.Get()
//	if err != nil {
//	  // handle error
//	}
package result

import (
	"fmt"
)

// Result wraps a value or an error.
type Result[T any] struct {
	value T
	err   error
}

// Of creates a new Result with the given value and error.
func Of[T any](value T, err error) Result[T] {
	return Result[T]{value: value, err: err}
}

// OfValue creates a successful Result with the given value.
func OfValue[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// OfError creates a failed Result with the given error.
//
// Panics if given error is nil.
func OfError[T any](err error) Result[T] {
	if err == nil {
		panic("result.OfError called with nil error")
	}

	return Result[T]{err: err}
}

// Get returns the value and error of the Result.
func (r Result[T]) Get() (value T, err error) {
	return r.value, r.err
}

// Must returns the value of the Result and panics if the Result contains an
// error.
//
// Not recommended for production code, but can be useful in tests or small
// scripts/playground projects.
func (r Result[T]) Must() T {
	if r.err != nil {
		panic(fmt.Sprintf("Must called on result with error: [%T]: %v", r.err, r.err))
	}

	return r.value
}

// String returns a string in the format:
//   - If Successful Result: "(Result[{{type}}]: '{{value}}')"
//   - If error Result: "(Result[{{type}}]: <error[{{error type}}]: {{error}}>)"
func (r Result[T]) String() string {
	if r.err != nil {
		return fmt.Sprintf("(Result[%T]: <error[%T]: %v>)", r.value, r.err, r.err)
	}

	return fmt.Sprintf("(Result[%T]: %v)", r.value, r.value)
}
