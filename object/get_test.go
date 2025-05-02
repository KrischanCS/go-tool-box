package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/object"
)

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
