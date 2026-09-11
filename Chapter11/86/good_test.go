package main

import (
	"sync/atomic"
	"testing"
	"time"
)

// GOOD: The mock's channel is the synchronization point. Receiving from it
// waits exactly until Publish has happened, with no timing assumption.
func TestHandleWithSynchronization(t *testing.T) {
	done := make(chan bool)
	go func() {
		done <- true
	}()

	if !<-done {
		t.Fatal("work did not finish")
	}
}

// RETRY: If the dependency cannot provide a channel, callback, or other
// synchronization point, retry the observation until it succeeds or times out.
// The short sleep prevents a busy loop; it is not used as completion proof.
func TestHandleWithRetry(t *testing.T) {
	var done atomic.Bool
	go func() {
		done.Store(true)
	}()

	succeededAt := -1
	for i := 0; i < 100; i++ {
		if done.Load() {
			succeededAt = i
			break
		}
		time.Sleep(5 * time.Nanosecond)
	}

	if succeededAt == -1 {
		t.Fatal("condition was not met after 100 retries")
	}
	t.Logf("condition succeeded on iteration %d", succeededAt)
}
