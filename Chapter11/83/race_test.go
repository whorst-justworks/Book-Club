package main

import (
	"sync"
	"testing"
)

// This test PASSES without -race, but FAILS with -race.
//
// Run WITHOUT race detector (passes, hides the bug):
//   go test -v ./...
//

func TestCounterConcurrent(t *testing.T) {
	counter := NewCounter()

	var wg sync.WaitGroup
	goroutines := 100

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	// Without -race, this may even produce the "correct" result sometimes,
	// making the bug even harder to find.
	expected := goroutines * 1000
	actual := counter.Value()
	t.Logf("expected %d, got %d (lost %d increments)", expected, actual, expected-actual)
}

// Run WITH race detector (catches the data race):
//   go test -race -v ./...
