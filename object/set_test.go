package object_test

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/object"
)

//nolint:errcheck,forcetypeassert
func ExampleSet() {
	obj := object.Object{
		"int":           23,
		"stringToFloat": "text",
		"object": object.Object{
			"array": []any{1, 2, 3},
		},
	}

	object.Set(obj, 42, "int")
	fmt.Println(obj["int"])

	object.Set(obj, 6.28, "stringToFloat")
	fmt.Println(obj["stringToFloat"])

	object.Set(obj, 23, "object", "array[0]")
	fmt.Println(obj["object"].(map[string]any)["array"])

	object.Set(obj, 2.718, "object", "array[]")
	fmt.Println(obj["object"].(map[string]any)["array"])

	object.Set(obj, true, "object", "array[47]")
	fmt.Println(obj["object"].(map[string]any)["array"])

	object.Set(obj, true, "object", "newBool")
	fmt.Println(obj["object"])

	// Output:
	// 42
	// 6.28
	// [23 2 3]
	// [23 2 3 2.718]
	// [23 2 3 2.718 true]
	// map[array:[23 2 3 2.718 true] newBool:true]
}

//nolint:funlen
func TestSet_any(t *testing.T) {
	t.Parallel()

	type test struct {
		name  string
		input object.Object
		path  []string
		value any
		want  object.Object
	}

	tests := []test{
		{
			name:  "Should overwrite value in root",
			input: object.Object{"int": 23},
			path:  []string{"int"},
			value: 42,
			want:  object.Object{"int": 42},
		},
		{
			name:  "Should create new value in root",
			input: object.Object{"int": 23},
			path:  []string{"bool"},
			value: true,
			want:  object.Object{"int": 23, "bool": true},
		},
		{
			name:  "Should update specific value when multiple exists",
			input: object.Object{"int": 23, "bool": true, "float": 3.14, "anotherFloat": 2.718},
			path:  []string{"float"},
			value: 6.283185,
			want:  object.Object{"int": 23, "bool": true, "float": 6.283185, "anotherFloat": 2.718},
		},
		{
			name:  "Should create new value in nested object",
			input: object.Object{"object": object.Object{"int": 23}},
			path:  []string{"object", "string"},
			value: "text",
			want:  object.Object{"object": object.Object{"int": 23, "string": "text"}},
		},
		{
			name:  "Should update a value in nested object",
			input: object.Object{"object": object.Object{"int": 23, "string": "text"}},
			path:  []string{"object", "int"},
			value: 42,
			want:  object.Object{"object": object.Object{"int": 42, "string": "text"}},
		},
		{
			name:  "Should create a deeper nested value according to given path",
			input: object.Object{"object": object.Object{"int": 23}},
			path:  []string{"object", "nested", "nestedDeeper", "int"},
			value: 42,
			want: object.Object{"object": object.Object{"int": 23,
				"nested": object.Object{"nestedDeeper": object.Object{"int": 42}}}},
		},
		{
			name:  "Should not modify object if path is empty",
			input: object.Object{"int": 23, "bool": true, "float": 3.14},
			path:  []string{},
			value: 42,
			want:  object.Object{"int": 23, "bool": true, "float": 3.14},
		},
		{
			name:  "Should modify a value at a specified array index",
			input: object.Object{"array": []any{1, 2, 3}},
			path:  []string{"array[1]"},
			value: 42,
			want:  object.Object{"array": []any{1, 42, 3}},
		},
		{
			name:  "Should append a value to an array when index is length of the array (cap == len, must grow)",
			input: object.Object{"array": []any{1, 2, 3}},
			path:  []string{"array[3]"},
			value: 4,
			want:  object.Object{"array": []any{1, 2, 3, 4}},
		},
		{
			name:  "Should append a value to an array when index is not specified (cap == len, must grow)",
			input: object.Object{"array": []any{1, 2, 3}},
			path:  []string{"array[]"},
			value: 4,
			want:  object.Object{"array": []any{1, 2, 3, 4}},
		},
		{
			name:  "Should append a value to an array when index is greater than length of the array (cap > len, can reuse)",
			input: object.Object{"array": slices.Grow([]any{1, 2, 3}, 6)},
			path:  []string{"array[47]"},
			value: 4,
			want:  object.Object{"array": []any{1, 2, 3, 4}},
		},
		{
			name:  "Should create an array with the one given element if no array exists at path",
			input: object.Object{},
			path:  []string{"object[]"},
			value: 6.28,
			want:  object.Object{"object": []any{6.28}},
		},
		{
			name:  "Should replace a non array value with an array if an index is privided",
			input: object.Object{"object": object.Object{"field": 23}},
			path:  []string{"object", "field[5]"},
			value: 42,
			want:  object.Object{"object": object.Object{"field": []any{42}}},
		},
		{
			name: "Should modify a value in an object nested in an array",
			input: object.Object{
				"array": []any{
					object.Object{"int": 0},
					object.Object{"int": 23},
					object.Object{"int": 5},
				},
			},
			path:  []string{"array[1]", "int"},
			value: 42,
			want: object.Object{
				"array": []any{
					object.Object{"int": 0},
					object.Object{"int": 42},
					object.Object{"int": 5},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			object.Set(tc.input, tc.value, tc.path...)

			// Assert
			assert.Equal(t, tc.want, tc.input)
		})
	}
}

func TestSet_ShouldPanicWhenArrayIndexIsGreaterThanMaxInr(t *testing.T) {
	t.Parallel()

	// Arrange
	input := object.Object{"array": []any{1, 2, 3}}
	pathElemnt := fmt.Sprintf("array[%d]", uint64(math.MaxInt)+1)

	// Act & Assert
	assert.PanicsWithValue(t,
		`Parsing index in object-path: strconv.Atoi: parsing "9223372036854775808": value out of range`,
		func() {
			object.Set(input, 42, pathElemnt)
		},
	)
}
