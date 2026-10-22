package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if len(in) == 0 {
		return make([]Result[R], 0), nil
	}
	res := make([]Result[R], len(in))
	var wg sync.WaitGroup
	ch := make(chan int)
	for range min(workers, len(in)) {
		wg.Go(func() {
			for i := range ch {
				workerRes, err := fn(ctx, in[i])
				if err != nil {
					res[i] = Result[R]{Err: err}
					continue
				}
				res[i] = Result[R]{workerRes, nil}
			}
		})
	}
dispatch:
	for i := range in {
		select {
		case ch <- i:
		case <-ctx.Done():
			break dispatch
		}
	}
	close(ch)
	wg.Wait()
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
