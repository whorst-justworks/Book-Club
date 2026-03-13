package main

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

func handler1(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for i, circle := range circles {
		// Creating new i variable and circle (see mistake 63)
		i := i
		circle := circle
		go func() {
			defer wg.Done()

			result, err := foo(ctx, circle)
			if err != nil {
				// What should we do with multiple errors here?
				// Slice of errors alligned with the result slice
				// Mutex error slice
				// Channel with a parent go routine handling the errors
			}

			// writing to different elements in the slice avoids a data reace here.
			results[i] = result
		}()
	}

	wg.Wait()
	// ...
	return results, nil
}

// Solution: Use errgroup!
func handler2(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))

	// errgroup has a single function withContext that returns a *Group struct
	// This struct allows us to synchronize errors
	g, ctx := errgroup.WithContext(ctx)
	// errgroup.Group has two public methods
	// g.Go() triggers a new go routine
	// g.Wait() blocks until the go routines have completed

	for i, circle := range circles {
		i := i
		circle := circle
		g.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				// Just return the error here
				return err
			}
			results[i] = result
			return nil
		})
	}

	// if we get an error in one of the goroutines, the shared context will cancel and we will return that first error
	// avoids long waits on multiple goroutines if one routine errors quickly
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func foo(context.Context, Circle) (Result, error) {
	return Result{}, nil
}

type (
	Circle struct{}
	Result struct{}
)
