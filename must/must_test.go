package must_test

import (
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/must"
)

func ExampleValue() {
	v := must.Value(func() (string, error) {
		return "test", nil
	}())

	fmt.Println(v)

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		must.Value(func() (string, error) {
			return "", io.EOF
		}())
	}()

	// Output:
	// test
	// Error in call to must.Value: [*errors.errorString] - EOF
}

func ExampleValues() {
	v1, v2 := must.Values(func() (string, float64, error) {
		return "test", 6.28, nil
	}())

	fmt.Println(v1, v2)

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		must.Values(func() (string, float64, error) {
			return "", 0, io.EOF
		}())
	}()

	// Output:
	// test 6.28
	// Error in call to must.Values: [*errors.errorString] - EOF
}

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
