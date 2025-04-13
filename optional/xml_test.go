package optional_test

import (
	"encoding/xml"
	"testing"

	"github.com/KrischanCS/go-tool-box/optional"

	"github.com/stretchr/testify/assert"
)

type xmlTestType struct {
	XMLName      xml.Name                 `xml:"TestTagName"`
	PresentInt   optional.Optional[int]   `xml:"PresentInt"`
	AbsentInt    optional.Optional[int]   `xml:"AbsentInt"`
	PresentState optional.Optional[state] `xml:"PresentState"`
	AbsentState  optional.Optional[state] `xml:"AbsentState"`
}

func TestOptional_UnmarshalXML(t *testing.T) {
	t.Parallel()

	// Arrange
	xmlContent := `<TestTagName>
	<PresentInt>123</PresentInt>
	<PresentState>
		<State>France</State>
		<Capital>Paris</Capital>
	</PresentState>
</TestTagName>`

	var opt xmlTestType

	// Act
	err := xml.Unmarshal([]byte(xmlContent), &opt)

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

func TestOptional_MarshalXML(t *testing.T) {
	t.Parallel()

	// Arrange
	input := xmlTestType{
		PresentInt:   optional.Of(123),
		PresentState: optional.Of(state{"France", "Paris"}),
	}

	// Act
	got, err := xml.MarshalIndent(input, "", "\t")

	// Assert
	want := `<TestTagName>
	<PresentInt>123</PresentInt>
	<PresentState>
		<State>France</State>
		<Capital>Paris</Capital>
	</PresentState>
</TestTagName>`

	assert.NoError(t, err)
	assert.Equal(t, want, string(got))
}

func TestOptional_UnmarshalXML_err(t *testing.T) {
	t.Parallel()

	// Arrange
	xmlContent := `<TestTagName>
	<PresentInt>Not an Integer</PresentInt>
</TestTagName>`

	var dst xmlTestType

	// Act
	err := xml.Unmarshal([]byte(xmlContent), &dst)

	// Assert
	assert.Error(t, err)
}
