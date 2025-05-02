package optional_test

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/must"
	"github.com/KrischanCS/go-toolbox/optional"
)

func TestOptional_UnmarshalJson_string(t *testing.T) {
	t.Parallel()

	// Arrange
	input := `"test"`
	want := "test"

	var opt optional.Optional[string]

	// Act
	err := json.Unmarshal([]byte(input), &opt)

	// Assert
	assert.NoError(t, err)

	got, ok := opt.Get()
	assert.True(t, ok)
	assert.Equal(t, want, got)
}

func TestOptional_UnmarshalJson_int(t *testing.T) {
	t.Parallel()

	// Arrange
	input := `123`
	want := 123

	var opt optional.Optional[int]

	// Act
	err := json.Unmarshal([]byte(input), &opt)

	// Assert
	assert.NoError(t, err)

	got, ok := opt.Get()
	assert.True(t, ok)
	assert.Equal(t, want, got)
}

func TestOptional_UnmarshalJson_float(t *testing.T) {
	t.Parallel()

	// Arrange
	input := `1.23`
	want := 1.23

	var opt optional.Optional[float64]

	// Act
	err := json.Unmarshal([]byte(input), &opt)

	// Assert
	assert.NoError(t, err)

	got, ok := opt.Get()
	assert.True(t, ok)
	assert.InEpsilon(t, want, got, 0.0001)
}

func TestOptional_UnmarshalJson_bool(t *testing.T) {
	t.Parallel()

	for _, b := range []bool{true, false} {
		t.Run(strconv.FormatBool(b), func(t *testing.T) {
			// Arrange
			input := strconv.FormatBool(b)
			want := b

			var opt optional.Optional[bool]

			// Act
			err := json.Unmarshal([]byte(input), &opt)

			// Assert
			assert.NoError(t, err)

			got, ok := opt.Get()
			assert.True(t, ok)
			assert.Equal(t, want, got)
		})
	}
}

func TestOptional_UnmarshalJson_timeTime(t *testing.T) {
	t.Parallel()

	// Arrange
	input := `"2025-04-12T01:02:03Z"`
	want := must.Value(time.Parse(time.DateTime, "2025-04-12 01:02:03"))

	var opt optional.Optional[time.Time]

	// Act
	err := json.Unmarshal([]byte(input), &opt)

	// Assert
	assert.NoError(t, err)

	got, ok := opt.Get()
	assert.True(t, ok)
	assert.Equal(t, want, got)
}

func TestOptional_UnmarshalJson_struct(t *testing.T) {
	t.Parallel()

	// Arrange
	input := `{"state":"France","capital":"Paris"}`
	want := state{"France", "Paris"}

	var opt optional.Optional[state]

	// Act
	err := json.Unmarshal([]byte(input), &opt)

	// Assert
	assert.NoError(t, err)

	got, ok := opt.Get()
	assert.True(t, ok)
	assert.Equal(t, want, got)
}

func TestOptional_UnmarshalJson_err(t *testing.T) {
	t.Parallel()

	// Arrange
	input := `"123"`

	var opt optional.Optional[int]

	// Act
	err := json.Unmarshal([]byte(input), &opt)

	// Assert
	assert.Error(t, err)
	assert.Zero(t, opt)
}

type jsonTestTypeOmitZero struct {
	PresentInt   optional.Optional[int]   `json:"presentInt,omitzero"`
	AbsentInt    optional.Optional[int]   `json:"absentInt,omitzero"`
	PresentState optional.Optional[state] `json:"presentState,omitzero"`
	AbsentState  optional.Optional[state] `json:"absentState,omitzero"`
}

func TestOptional_UnmarshalJson_Object(t *testing.T) {
	t.Parallel()

	// Arrange
	jsonContent := `{
		"presentInt": 123,
		"presentState": {
			"state": "France",
			"capital": "Paris"
		}
	}`

	var opt jsonTestTypeOmitZero

	// Act
	err := json.Unmarshal([]byte(jsonContent), &opt)

	// Assert
	assert.NoError(t, err)

	v1, ok := opt.PresentInt.Get()
	assert.True(t, ok)
	assert.Equal(t, 123, v1)

	v2, ok := opt.AbsentInt.Get()
	assert.False(t, ok)
	assert.Zero(t, v2)

	v3, ok := opt.PresentState.Get()
	assert.True(t, ok)
	assert.Equal(t, state{"France", "Paris"}, v3)

	v4, ok := opt.AbsentState.Get()
	assert.False(t, ok)
	assert.Zero(t, v4)
}

func TestOptional_MarshalJson_Object(t *testing.T) {
	t.Parallel()

	// Arrange
	input := jsonTestTypeOmitZero{
		PresentInt:   optional.Of(123),
		PresentState: optional.Of(state{"France", "Paris"}),
	}

	// Act
	got, err := json.MarshalIndent(input, "", "\t")

	// Assert
	assert.NoError(t, err)

	want := `{
	"presentInt": 123,
	"presentState": {
		"state": "France",
		"capital": "Paris"
	}
}`
	assert.Equal(t, want, string(got))
}

type jsonTestTypeNulls struct {
	PresentInt   optional.Optional[int]   `json:"presentInt"`
	AbsentInt    optional.Optional[int]   `json:"absentInt"`
	PresentState optional.Optional[state] `json:"presentState"`
	AbsentState  optional.Optional[state] `json:"absentState"`
}

func TestOptional_MarshalJson_Nulls(t *testing.T) {
	t.Parallel()

	// Arrange
	input := jsonTestTypeNulls{
		PresentInt:   optional.Of(123),
		PresentState: optional.Of(state{"France", "Paris"}),
	}

	// Act
	got, err := json.MarshalIndent(input, "", "\t")

	// Assert
	assert.NoError(t, err)

	want := `{
	"presentInt": 123,
	"absentInt": null,
	"presentState": {
		"state": "France",
		"capital": "Paris"
	},
	"absentState": null
}`
	assert.Equal(t, want, string(got))
}

func TestOptional_UnmarshalJson_Null(t *testing.T) {
	t.Parallel()

	// Arrange
	input := `{
	"presentInt": 123,
	"absentInt": null,
	"presentState": {
		"state": "France",
		"capital": "Paris"
	},
	"absentState": null
}`

	var dst jsonTestTypeNulls

	// Act
	err := json.Unmarshal([]byte(input), &dst)

	// Assert
	assert.NoError(t, err)

	want := jsonTestTypeNulls{
		PresentInt:   optional.Of(123),
		AbsentInt:    optional.Empty[int](),
		PresentState: optional.Of(state{"France", "Paris"}),
		AbsentState:  optional.Empty[state](),
	}

	assert.Equal(t, want, dst)
}
