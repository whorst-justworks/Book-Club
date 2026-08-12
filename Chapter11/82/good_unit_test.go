package main

import "testing"

// GOOD: Unit tests have NO build tag — they always run by default.
// Fast, no external dependencies.
//
// Run with: go test -v ./...

func TestAddUnit(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestMultiplyUnit(t *testing.T) {
	result := Multiply(3, 4)
	if result != 12 {
		t.Errorf("expected 12, got %d", result)
	}
}
