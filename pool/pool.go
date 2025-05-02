// Package pool provides a convenient way to create a worker pools.
package pool

import (
	"runtime"
	"sync"
)

// Options contains the options for the worker pool.
type Options struct {
	// PoolSize is the number of workers in the pool. Defaults to GOMAXPROCS.
	PoolSize int
	// OutBufferSize is buffer size of the output channel. Defaults to 0.
	OutBufferSize int
}

// NewPool creates a worker pool. It will call fn for each value from in and
// send the results to out. There is no guarantee about the order of the
// results.
//
// The pool will run until in is closed. As soon as all values are written out,
// out will be closed.
//
// The behavior of the pool can be configured with opts, if opts is nil,
// defaults will be used (see [Options]).
func NewPool[IN, OUT any](fn func(IN) OUT, inChan <-chan IN, options *Options) <-chan OUT {
	opts := initOptions(options)

	out := make(chan OUT, opts.OutBufferSize)

	var wg sync.WaitGroup
	for i := range opts.PoolSize {
		wg.Add(1)

		go func() {
			for v := range inChan {
				out <- fn(v)
			}

			wg.Done()

			if i == 0 {
				wg.Wait()
				close(out)
			}
		}()
	}

	return out
}

func initOptions(opts *Options) Options {
	if opts == nil {
		return Options{
			PoolSize:      runtime.GOMAXPROCS(0),
			OutBufferSize: 0,
		}
	}

	// When opts are provided but PoolSize is not set, the default is 0, which
	// would mean no worker is started, using the default instead.
	if opts.PoolSize == 0 {
		opts.PoolSize = runtime.GOMAXPROCS(0)
	}

	return *opts
}
