package result_test

import (
	"bufio"
	"fmt"
	"io"

	"github.com/KrischanCS/go-tool-box/result"
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

func getValue(i int) (int, error) {
	return i, nil
}

func getError(err error) (int, error) {
	return 0, err
}
