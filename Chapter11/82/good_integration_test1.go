//go:build integration

package main1

import (
	"testing"
	"time"
)

// GOOD: Integration tests are behind a build tag.
// They are EXCLUDED from `go test ./...` by default.
// Only run when explicitly requested.
//
// Run with: go test -v -tags=integration ./...

func TestDatabaseInsertIntegration(t *testing.T) {
	time.Sleep(2 * time.Second)

	conn, err := ConnectToDatabase()
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	err = InsertRecord(conn, "test-record")
	if err != nil {
		t.Errorf("failed to insert: %v", err)
	}
}

func TestDatabaseMultipleInsertsIntegration(t *testing.T) {
	time.Sleep(3 * time.Second)

	conn, err := ConnectToDatabase()
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	for _, name := range []string{"alice", "bob", "charlie"} {
		err = InsertRecord(conn, name)
		if err != nil {
			t.Errorf("failed to insert %s: %v", name, err)
		}
	}
}
