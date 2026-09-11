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

// Other useful httptest functionality:
//
//   - httptest.NewRequest creates an HTTP request for a handler test.
//   - httptest.NewRecorder captures a handler's status, headers, and body.
//   - httptest.NewServer starts an in-memory HTTP server for client tests.
//   - httptest.NewTLSServer starts an HTTPS test server.
//   - httptest.NewUnstartedServer creates a server that can be configured
//     before it starts.
//   - server.Close stops a test server and releases its resources.
