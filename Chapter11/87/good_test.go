package main

import (
	"testing"
	"time"
)

func TestIsExpired(t *testing.T) {
	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	created := now.Add(-time.Hour)

	if !isExpired(created, now) {
		t.Fatal("expected item to be expired at the exact boundary")
	}
}

func TestIsExpiredWithInjectedTime(t *testing.T) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	now := start
	events := 0
	getNow := func() time.Time {
		now = now.Add(10 * time.Millisecond)
		events++ // An event can be triggered here.
		return now
	}
	created := start.Add(-time.Hour)

	if !isExpiredWithNow(created, getNow) {
		t.Fatal("expected item to be expired")
	}

	if events != 1 {
		t.Fatalf("expected one event, got %d", events)
	}
}

// Other ways to simulate time:
//   - Use a manually controlled fake clock with Advance(duration), so tests
//     move time explicitly instead of advancing on every Now call.
//   - Use a scripted clock that returns a predefined sequence of timestamps.
