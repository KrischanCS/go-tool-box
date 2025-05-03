package result_test

import (
	"bufio"
	"fmt"
	"io"

	"github.com/KrischanCS/go-toolbox/result"
)

func ExampleOf() {
	resultChan := make(chan result.Result[int])

	go func() {
		resultChan <- result.Of(getValue(1))
		resultChan <- result.Of(getError(io.EOF))
		resultChan <- result.OfValue(2)
		resultChan <- result.OfError[int](bufio.ErrFinalToken)

		close(resultChan)
	}()

	for res := range resultChan {
		v, err := res.Get()
		if err != nil {
			fmt.Println("Got error:", err)
			continue
		}

		fmt.Println("Got value:", v)
	}

	// Output:
	// Got value: 1
	// Got error: EOF
	// Got value: 2
	// Got error: final token
}

func ExampleResult_Must() {
	resultSuccess := result.Of(getValue(1))

	fmt.Println(resultSuccess.Must())

	func() {
		defer func() {
			v := recover()
			if v != nil {
				fmt.Println("Panicked:", v)
			}
		}()

		resultError := result.Of(getError(io.EOF))
		fmt.Println(resultError.Must())
	}()

	// Output:
	// 1
	// Panicked: Must called on result with error: [*errors.errorString]: EOF
}

func ExampleResult_String() {
	resultSuccess := result.Of(getValue(1))
	resultError := result.Of(getError(io.EOF))

	fmt.Println(resultSuccess.String())
	fmt.Println(resultError.String())

	// Output:
	// (Result[int]: 1)
	// (Result[int]: <error[*errors.errorString]: EOF>)
}

func getValue(i int) (int, error) {
	return i, nil
}

func getError(err error) (int, error) {
	return 0, err
}
