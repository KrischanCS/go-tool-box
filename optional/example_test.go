package optional_test

import (
	"encoding/json"
	"fmt"

	"github.com/KrischanCS/go-toolbox/optional"
)

func ExampleOptional_MarshalJSON() {
	type Test struct {
		Value1 optional.Optional[string] `json:"value1"`
		Value2 optional.Optional[string] `json:"value2"`
		Value3 optional.Optional[string] `json:"value3,omitempty"` // omitempty will be kept, event if optional is empty
		Value4 optional.Optional[string] `json:"value4,omitzero"`  // omitzero will correctly omit an empty optional
	}

	t := Test{
		Value1: optional.Of("value1"),
		Value2: optional.Empty[string](),
		// Not setting Value3 & Value4, zero values are empty optionals,
		// equivalent as setting Value2 to empty explicitly
	}

	jsonBytes, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))

	// Output:
	// {
	//   "value1": "value1",
	//   "value2": null,
	//   "value3": null
	// }
}
