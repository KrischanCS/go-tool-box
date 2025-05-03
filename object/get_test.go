package object_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/object"
)

//nolint:funlen
func ExampleGet() {
	obj := object.Object{
		"int":  42,
		"bool": true,
		"object": map[string]any{
			"nestedSlice": []float64{1.618, 2.718, 9.81},
			"nestedObject": map[string]any{
				"deepNestedObject": map[string]any{
					"deepNestedFloat": 6.28,
				},
			},
		},
		"sliceWithObject": []object.Object{
			{
				"string": "blue gopher",
			},
		},
	}

	i, ok := object.Get[int](obj, "int")
	fmt.Println("Gets array from property at root:", ok, i)

	i, ok = object.Get[int](obj, "anotherInt")
	fmt.Println(
		"If property doesn't exist, ok will be false and the value is the zero value of the type:",
		ok,
		i,
	)

	b, ok := object.Get[bool](obj, "int")
	fmt.Println("If type mismatches, ok is also false:", ok, b)

	float, ok := object.Get[float64](
		obj,
		"object",
		"nestedObject",
		"deepNestedObject",
		"deepNestedFloat",
	)
	fmt.Println("Nested objects can be obtained with complete path:", ok, float)

	float, ok = object.Get[float64](obj, "object", "nestedSlice[2]")
	fmt.Println(
		"Values from arrays can be obtained by adding [{{index}}] to a path element:",
		ok,
		float,
	)

	float, ok = object.Get[float64](obj, "object", "nestedSlice[23]")
	fmt.Println(
		"If index is out of bounds, ok is false and value the zero value of the type:",
		ok,
		float,
	)

	s, ok := object.Get[string](obj, "sliceWithObject[0]", "string")
	fmt.Println(
		"Values from nested objects in arrays can be obtained:",
		ok,
		s,
	)

	// Output:
	// Gets array from property at root: true 42
	// If property doesn't exist, ok will be false and the value is the zero value of the type: false 0
	// If type mismatches, ok is also false: false false
	// Nested objects can be obtained with complete path: true 6.28
	// Values from arrays can be obtained by adding [{{index}}] to a path element: true 9.81
	// If index is out of bounds, ok is false and value the zero value of the type: false 0
	// Values from nested objects in arrays can be obtained: true blue gopher
}

//nolint:gochecknoglobals
var testObject = map[string]any{
	"int":    42,
	"string": "string",
	"float":  6.283185,
	"bool":   true,
	"slice":  []int{1, 2, 3},
	"object": map[string]any{
		"nestedInt":   23,
		"nestedSlice": []float64{1.618, 2.718, 9.81},
		"nestedObject": map[string]any{
			"deepNestedObject": map[string]any{
				"deepNestedInt": 5,
			},
		},
	},
	"sliceWithObjects": []object.Object{
		{
			"int":    1,
			"string": "text1",
		},
		{
			"int":    2,
			"string": "text2",
		},
	},
}

//nolint:funlen
func TestGet_any(t *testing.T) {
	t.Parallel()

	type test struct {
		name   string
		path   []string
		exists bool
		want   any
	}

	tests := []test{
		{
			name:   "Should return int from root with simple path",
			path:   []string{"int"},
			exists: true,
			want:   42,
		},
		{
			name:   "Should return string from root with simple path",
			path:   []string{"string"},
			exists: true,
			want:   "string",
		},
		{
			name:   "Should return float from root with simple path",
			path:   []string{"float"},
			exists: true,
			want:   6.283185,
		},
		{
			name:   "Should return bool from root with simple path",
			path:   []string{"bool"},
			exists: true,
			want:   true,
		},
		{
			name:   "Should return slice from root with simple path",
			path:   []string{"slice"},
			exists: true,
			want:   []int{1, 2, 3},
		},
		{
			name:   "Should return ok=false if path not exists",
			path:   []string{"wrongPath"},
			want:   nil,
			exists: false,
		},
		{
			name:   "Should get value from nested object",
			path:   []string{"object", "nestedInt"},
			exists: true,
			want:   23,
		},
		{
			name:   "Should get non primitive value from nested object",
			path:   []string{"object", "nestedObject"},
			exists: true,
			want:   map[string]any{"deepNestedObject": map[string]any{"deepNestedInt": 5}},
		},
		{
			name:   "Should get value from deep nested object",
			path:   []string{"object", "nestedObject", "deepNestedObject", "deepNestedInt"},
			exists: true,
			want:   5,
		},
		{
			name:   "Should return not-ok if path requires nested object, but value is primitive",
			path:   []string{"int", "digIntoInt?"},
			want:   nil,
			exists: false,
		},
		{
			name:   "Should return object itself if path has length 0",
			path:   []string{},
			want:   testObject,
			exists: true,
		},
		{
			name:   "Should return specified element of array if specified in brackets in path",
			path:   []string{"slice[1]"},
			want:   2,
			exists: true,
		},
		{
			name:   "Should return specified element of a nested array",
			path:   []string{"object", "nestedSlice[2]"},
			want:   9.81,
			exists: true,
		},
		{
			name:   "Should return not-ok if specified element of array is out of bounds",
			path:   []string{"object", "nestedSlice[3]"},
			want:   nil,
			exists: false,
		},
		{
			name:   "Should return not-ok if index used on a non-indexable type",
			path:   []string{"int[0]"},
			want:   nil,
			exists: false,
		},
		{
			name:   "Should return not-ok element not exist with an index path",
			path:   []string{"notExisting[0]"},
			want:   nil,
			exists: false,
		},
		{
			name:   "Should return value from an object nested in an array",
			path:   []string{"sliceWithObjects[1]", "int"},
			want:   2,
			exists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, ok := object.Get[any](testObject, tt.path...)

			// Assert
			assert.Equal(t, tt.exists, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGet_shouldReturnConcreteTypeWhenSpecified(t *testing.T) {
	t.Parallel()

	// Act
	got, ok := object.Get[[]float64](testObject, "object", "nestedSlice")

	// Assert
	assert.True(t, ok)

	want := []float64{1.618, 2.718, 9.81}
	assert.Equal(t, want, got)
}

func TestGet_shouldReturnNotOkIfPathExistButTypeMismatches(t *testing.T) {
	t.Parallel()

	// Act
	got, ok := object.Get[[]int](testObject, "object", "nestedSlice")

	// Assert
	want := []int(nil)

	assert.False(t, ok)
	assert.Equal(t, want, got)
}

func TestGet_shouldPanicWhenIndexGreaterThanMaxInt(t *testing.T) {
	t.Parallel()

	// Arrange
	index := uint64(math.MaxInt) + 1
	path := []string{"slice[" + strconv.FormatUint(index, 10) + "]"}

	// Act & Assert
	assert.PanicsWithValue(t,
		`Parsing index in object-path: strconv.Atoi: parsing "9223372036854775808": value out of range`,
		func() {
			object.Get[int](testObject, path...)
		},
	)
}
