package main

import (
	"testing"
	"time"
)

// Mistake #84: Not Using Test Execution Modes
//
// t.Parallel() marks tests to run concurrently instead of sequentially.
// This speeds up suites with slow, independent tests (e.g., HTTP calls, DB queries).
//
// Run sequentially (slow):
//   go test -v -run "Sequential" ./...
//
// Run in parallel (fast):
//   go test -v -run "Parallel" ./...
//
// Control max parallelism:
//   go test -v -run "Parallel" -parallel 2 ./...

func TestSequential_FetchEndpoints(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "FetchUserProfile"},
		{name: "FetchUserOrders"},
		{name: "FetchUserPreferences"},
		{name: "FetchUserNotifications"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(1 * time.Second) // Simulates slow API call
			t.Logf("fetched %s", tt.name)
		})
	}
}

// Same table tests but each subtest calls t.Parallel() — they all run concurrently.

func TestParallel_FetchEndpoints(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "FetchUserProfile"},
		{name: "FetchUserOrders"},
		{name: "FetchUserPreferences"},
		{name: "FetchUserNotifications"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			time.Sleep(1 * time.Second)
			t.Logf("fetched %s", tt.name)
		})
	}
}
