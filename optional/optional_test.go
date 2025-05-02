package optional_test

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/KrischanCS/go-toolbox/must"
	"github.com/KrischanCS/go-toolbox/optional"

	"github.com/stretchr/testify/assert"
)

type state struct {
	State   string `json:"state,omitzero"   xml:"State,omitempty"`
	Capital string `json:"capital,omitzero" xml:"Capital,omitempty"`
}

type stateStringer state

func (s stateStringer) String() string {
	return fmt.Sprintf("[%s - Capital: %s]", s.State, s.Capital)
}

func TestNewEmptyOptional(t *testing.T) {
	t.Parallel()

	// Arrange
	opt := optional.Empty[string]()

	// Act
	value, ok := opt.Get()

	// Assert
	assert.False(t, ok)
	assert.Empty(t, value)
}

func TestNewOptional(t *testing.T) {
	t.Parallel()

	// Arrange
	values := []any{
		"test",
		1,
		1.23,
		true,
		state{
			State:   "France",
			Capital: "Paris",
		},
	}

	for _, value := range values {
		t.Run(reflect.TypeOf(value).Kind().String(), func(t *testing.T) {
			// Act
			opt := optional.Of(value)
			got, ok := opt.Get()

			// Assert
			assert.True(t, ok)
			assert.Equal(t, value, got)
		})
	}
}

func TestOptional_Clear(t *testing.T) {
	t.Parallel()

	// Arrange
	values := []any{
		"test",
		1,
		1.23,
		true,
		state{
			State:   "France",
			Capital: "Paris",
		},
	}

	for _, value := range values {
		t.Run(reflect.TypeOf(value).Kind().String(), func(t *testing.T) {
			opt := optional.Of(value)

			// Act
			opt.Clear()

			// Assert
			got, ok := opt.Get()
			assert.False(t, ok)
			assert.Zero(t, got)
		})
	}
}

func TestOptional_Set(t *testing.T) {
	t.Parallel()

	// Arrange
	values := []any{
		"test",
		1,
		1.23,
		true,
		state{
			State:   "France",
			Capital: "Paris",
		},
	}

	opt := optional.Empty[any]()

	for _, value := range values {
		t.Run(reflect.TypeOf(value).Kind().String(), func(t *testing.T) {
			// Act
			opt.Set(value)

			// Assert
			got, ok := opt.Get()
			assert.True(t, ok)
			assert.Equal(t, value, got)
		})
	}
}

func TestOptional_String(t *testing.T) {
	t.Parallel()

	// Arrange
	tests := []struct {
		value  any
		string string
	}{
		{"test", `(Optional[string]: test)`},
		{1, `(Optional[int]: 1)`},
		{1.23, `(Optional[float64]: 1.23)`},
		{true, `(Optional[bool]: true)`},
		{
			value:  must.Value(time.Parse(time.DateTime, "2025-04-12 01:02:03")),
			string: "(Optional[time.Time]: 2025-04-12 01:02:03 +0000 UTC)"},
		{
			value:  state{"France", "Paris"},
			string: "(Optional[optional_test.state]: {France Paris})",
		},
		{
			value:  stateStringer{"France", "Paris"},
			string: "(Optional[optional_test.stateStringer]: [France - Capital: Paris])",
		},
	}

	for _, tc := range tests {
		t.Run(reflect.TypeOf(tc.value).Name(), func(t *testing.T) {
			opt := optional.Of(tc.value)

			// Act
			got := opt.String()

			// Assert
			assert.Equal(t, tc.string, got)
		})
	}
}

func TestOptional_String_empty_string(t *testing.T) {
	t.Parallel()

	// Arrange
	opt := optional.Empty[string]()

	// Act
	got := opt.String()

	// Assert
	assert.Equal(t, "(Optional[string]: <empty>)", got)
}

func TestOptional_String_netIP(t *testing.T) {
	t.Parallel()

	// Arrange
	opt := optional.Empty[net.IP]()

	// Act
	got := opt.String()

	// Assert
	assert.Equal(t, "(Optional[net.IP]: <empty>)", got)
}
