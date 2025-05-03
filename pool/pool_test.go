package pool_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/pool"
)

func ExampleNew() {
	doubleFn := func(i int) int {
		return i * 2
	}

	inC := make(chan int)

	outC := pool.New(doubleFn, inC, &pool.Options{
		PoolSize:      3,
		OutBufferSize: 1,
	})

	go func() {
		for _, i := range []int{1, 2, 3, 4, 5} {
			inC <- i
		}

		close(inC)
	}()

	var sum int
	for v := range outC {
		sum += v
	}

	fmt.Println(sum)

	// Output: 30
}

func TestNewPool(t *testing.T) {
	t.Parallel()

	// Arrange
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	double := func(in int) int {
		return in * 2
	}
	inC := make(chan int)

	// Act
	outC := pool.New(double, inC, nil)

	go func() {
		for _, v := range values {
			inC <- v
		}

		close(inC)
	}()

	got := make([]int, 0, len(values))
	for v := range outC {
		got = append(got, v)
	}

	// Assert
	want := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	assert.ElementsMatch(t, want, got, "Result must contain the same values in any order.")
}

func TestNewPool_ChainedPools(t *testing.T) {
	t.Parallel()

	// Arrange
	values := iterator.FromTo(0, 15)

	double := func(in int) int {
		return in * 2
	}
	toFloatPlus50Percent := func(in int) float64 {
		return float64(in) * 1.25
	}
	toString := func(in float64) string {
		return strconv.FormatFloat(in, 'f', 2, 64)
	}

	inC := make(chan int)

	// Act
	intermediate1 := pool.New(double, inC, nil)
	intermediate2 := pool.New(toFloatPlus50Percent, intermediate1, nil)
	outC := pool.New(toString, intermediate2, nil)

	go func() {
		for v := range values {
			inC <- v
		}

		close(inC)
	}()

	got := make([]string, 0, 100)
	for v := range outC {
		got = append(got, v)
	}

	// Assert
	want := []string{
		"0.00",
		"2.50",
		"5.00",
		"7.50",
		"10.00",
		"12.50",
		"15.00",
		"17.50",
		"20.00",
		"22.50",
		"25.00",
		"27.50",
		"30.00",
		"32.50",
		"35.00",
	}

	assert.ElementsMatch(t, want, got, "Result must contain the same values in any order.")
}

func TestNewPool_BufferSizeIsApplied(t *testing.T) {
	t.Parallel()

	// Arrange
	double := func(in int) int {
		return in * 2
	}
	inC := make(chan int)

	// Act
	outC := pool.New(double, inC, &pool.Options{OutBufferSize: 10})

	// Assert
	bufferSize := reflect.ValueOf(outC).Cap()

	assert.Equal(t, 10, bufferSize)
}
