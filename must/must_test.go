package must_test

import (
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-tool-box/must"
)

func TestValue(t *testing.T) {
	t.Parallel()

	// Arrange
	testCase := []struct {
		name  string
		value any
	}{
		{name: "string", value: "test"},
		{name: "int", value: 123},
		{name: "float", value: 123.456},
		{name: "bool", value: true},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := must.Value(tc.value, nil)

			// Assert
			assert.Equal(t, tc.value, got)
		})
	}
}

func TestValue_err(t *testing.T) {
	t.Parallel()

	// Arrange
	value := ""
	err := io.EOF

	// Act & Assert
	assert.PanicsWithValue(t,
		fmt.Sprintf("Error in call to must.Value: [%T] - %s", err, err),
		func() { must.Value(value, err) },
	)
}

func TestValues(t *testing.T) {
	t.Parallel()

	// Arrange
	testCase := []struct {
		name   string
		valueT any
		valueR any
	}{
		{name: "string", valueT: "test", valueR: "test2"},
		{name: "int", valueT: 123, valueR: -21},
		{name: "float", valueT: 123.456, valueR: "false"},
		{name: "bool", valueT: true, valueR: net.IPv4(192, 0, 2, 127)},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			gotT, gotR := must.Values(tc.valueT, tc.valueR, nil)

			// Assert
			assert.Equal(t, tc.valueT, gotT)
			assert.Equal(t, tc.valueR, gotR)
		})
	}
}

func TestValues_err(t *testing.T) {
	t.Parallel()

	// Arrange
	value := ""
	err := io.EOF

	// Act & Assert
	assert.PanicsWithValue(t,
		fmt.Sprintf("Error in call to must.Values: [%T] - %s", err, err),
		func() { must.Values(value, value, err) },
	)
}
