// Package object provides utility methods to work with values where the
// underlying type is 'map[string]any'.
//
// This can be handy when working with unknown or dynamic json/yaml/toml/… data.
package object

// Object is an alias for map[string]any, for compliance with naming in
// JSON/YAML/TOML/….
type Object = map[string]any

// Array is an alias for []any, for compliance with naming in JSON/YAML/TOML/….
type Array = []any
