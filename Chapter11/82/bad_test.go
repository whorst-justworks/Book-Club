//go:build bad

package main

import (
	"testing"
	"time"
)

// BAD: No categorization — unit and integration tests all in one file.
// Running `go test ./...` forces developers to wait for slow tests every time.
//
// Run with: go test -v -tags=bad ./...

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	if result != 12 {
		t.Errorf("expected 12, got %d", result)
	}
}

func TestDatabaseInsert(t *testing.T) {
	time.Sleep(2 * time.Second) // Simulates slow DB connection

	conn, err := ConnectToDatabase()
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	err = InsertRecord(conn, "test-record")
	if err != nil {
		t.Errorf("failed to insert: %v", err)
	}
}

func TestDatabaseMultipleInserts(t *testing.T) {
	time.Sleep(3 * time.Second) // Simulates slow DB setup

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
