package iterator

import "iter"

// SlidingWindow creates a [iter.Seq] which yields slices with overlapping
// windows of the given values and windowSize.
//
// The yielded slices reuse the same slice, so if you plan to store them, you
// must copy them first.
func SlidingWindow[T any](values iter.Seq[T], windowSize int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if windowSize <= 0 {
			return
		}

		yieldSlidingWindows(yield, values, windowSize)
	}
}

// FixedWindow creates a [iter.Seq] which yields slices of non overlapping
// windows of the given values and windowSize.
//
// The yielded slices reuse the same slice, so if you plan to store them, you
// must copy them first.
func FixedWindow[T any](values iter.Seq[T], windowsSize int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if windowsSize <= 0 {
			return
		}

		yieldFixedWindows(yield, values, windowsSize)
	}
}

func yieldFixedWindows[T any](yield func([]T) bool, values iter.Seq[T], windowsSize int) {
	window := make([]T, 0, windowsSize)

	for v := range values {
		if len(window) == windowsSize {
			if !yield(window) {
				return
			}

			window = window[:0]
		}

		window = append(window, v)
	}

	yield(window)
}

func yieldSlidingWindows[T any](yield func([]T) bool, values iter.Seq[T], windowSize int) {
	window := make([]T, 0, windowSize)

	for v := range values {
		if len(window) < windowSize {
			window = append(window, v)
			continue
		}

		if !yield(window) {
			return
		}

		slideWindow(window, v)
	}

	yield(window)
}

func slideWindow[T any](window []T, v T) {
	for i := range window[:len(window)-1] {
		window[i] = window[i+1]
	}

	window[len(window)-1] = v
}
