//go:build bad

package main

import "testing"

// BAD: Repetitive tests with copy-pasted structure.
// Adding a new case means duplicating an entire function.
// Hard to see what's actually different between tests.
//
// Run with: go test -v -tags=bad ./...

func TestValidateUser_EmptyName(t *testing.T) {
	u := User{Name: "", Email: "alice@example.com", Age: 25}
	err := ValidateUser(u)
	if err == nil {
		t.Fatal("expected error for empty name")
	}
	if err != ErrEmptyName {
		t.Errorf("expected ErrEmptyName, got %v", err)
	}
}

func TestValidateUser_WhitespaceName(t *testing.T) {
	u := User{Name: "   ", Email: "alice@example.com", Age: 25}
	err := ValidateUser(u)
	if err == nil {
		t.Fatal("expected error for whitespace name")
	}
	if err != ErrEmptyName {
		t.Errorf("expected ErrEmptyName, got %v", err)
	}
}

func TestValidateUser_NegativeAge(t *testing.T) {
	u := User{Name: "Alice", Email: "alice@example.com", Age: -1}
	err := ValidateUser(u)
	if err == nil {
		t.Fatal("expected error for negative age")
	}
	if err != ErrInvalidAge {
		t.Errorf("expected ErrInvalidAge, got %v", err)
	}
}

func TestValidateUser_AgeTooHigh(t *testing.T) {
	u := User{Name: "Alice", Email: "alice@example.com", Age: 200}
	err := ValidateUser(u)
	if err == nil {
		t.Fatal("expected error for age too high")
	}
	if err != ErrInvalidAge {
		t.Errorf("expected ErrInvalidAge, got %v", err)
	}
}

func TestValidateUser_ValidUser(t *testing.T) {
	u := User{Name: "Alice", Email: "alice@example.com", Age: 25}
	err := ValidateUser(u)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestValidateUser_ValidEdgeAge0(t *testing.T) {
	u := User{Name: "Baby", Email: "baby@example.com", Age: 0}
	err := ValidateUser(u)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestValidateUser_ValidEdgeAge150(t *testing.T) {
	u := User{Name: "Elder", Email: "elder@example.com", Age: 150}
	err := ValidateUser(u)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
