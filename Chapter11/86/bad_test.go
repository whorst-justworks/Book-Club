//go:build bad

package main

import (
	"testing"
	"time"
)

// BAD: Sleeping guesses how long the goroutine needs to finish. The test is
// either unnecessarily slow or flaky when the machine is under load.
//
// Run with: go test -v -tags=bad ./...
func TestHandleWithSleep(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var done bool
		go func() {
			done = true
		}()

		time.Sleep(time.Nanosecond * 5000)
		if !done {
			t.Fatalf("work did not finish on iteration %d", i)
		}
	}
}
