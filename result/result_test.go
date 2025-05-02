package result_test

import (
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/result"
)

func TestOf_success(t *testing.T) {
	t.Parallel()

	// Act
	strResult := result.Of("test", nil)
	str, err := strResult.Get()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "test", str)
}

func TestOf_error(t *testing.T) {
	t.Parallel()

	// Act
	errResult := result.Of("", io.EOF)
	str, err := errResult.Get()

	// Assert
	assert.Error(t, err)
	assert.Empty(t, str)
}

func TestOfValue(t *testing.T) {
	t.Parallel()

	// Act
	strResult := result.OfValue("test")
	str, err := strResult.Get()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "test", str)
}

func TestOfError(t *testing.T) {
	t.Parallel()

	// Act
	errResult := result.OfError[string](io.EOF)
	str, err := errResult.Get()

	// Assert
	assert.Error(t, err)
	assert.Empty(t, str)
}

func TestOfError_panicsOnNilError(t *testing.T) {
	t.Parallel()

	// Act & Assert
	assert.PanicsWithValue(t,
		"result.OfError called with nil error",
		func() { result.OfError[string](nil) },
	)
}

func TestResult_Must(t *testing.T) {
	t.Parallel()

	// Arrange
	strResult := result.Of("test", nil)

	// Act
	str := strResult.Must()

	// Assert
	assert.Equal(t, "test", str)
}

func TestResult_Must_panicsOnError(t *testing.T) {
	t.Parallel()

	// Arrange
	errResult := result.Of("", io.EOF)

	// Act & Assert
	assert.PanicsWithValue(t,
		"Must called on result with error: [*errors.errorString]: EOF",
		func() { errResult.Must() },
	)
}

func TestResult_String(t *testing.T) {
	t.Parallel()

	// Arrange

	type testCase struct {
		name  string
		value any
		err   error
		want  string
	}

	tests := []testCase{
		{
			name:  "successful string result",
			value: "test",
			err:   nil,
			want:  "(Result[string]: test)",
		},
		{
			name:  "successful int result",
			value: 123,
			err:   nil,
			want:  "(Result[int]: 123)",
		},
		{
			name:  "successful float result",
			value: 123.456,
			err:   nil,
			want:  "(Result[float64]: 123.456)",
		},
		{
			name:  "successful net.IP result",
			value: net.IPv4(192, 0, 2, 127),
			err:   nil,
			want:  "(Result[net.IP]: 192.0.2.127)",
		},
		{
			name:  "error result",
			value: "",
			err:   io.EOF,
			want:  "(Result[string]: <error[*errors.errorString]: EOF>)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := result.Of(tc.value, tc.err)

			// Act
			got := r.String()

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}
