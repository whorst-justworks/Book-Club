package main

import "testing"

// -shuffle randomizes test execution order to expose hidden dependencies.
//
// Tests that pass only in a specific order have implicit coupling — one test
// sets up state that another depends on. Shuffling catches this.
//
// Run in default order (may hide coupling bugs):
//   go test -v -run "Shuffle" ./...
//
// Run shuffled (exposes order-dependent tests):
//   go test -v -run "Shuffle" -shuffle=on ./...
//
// Run with a specific seed (reproducible):
//   go test -v -run "Shuffle" -shuffle=12345 ./...

// Shared package-level state — this is the problem.
var sharedDB = map[string]string{}

// These are separate top-level tests with a hidden dependency:
// CreateUser must run before GetUser or DeleteUser.
// Without -shuffle, Go runs them in source order so the bug stays hidden.

func TestShuffle_CreateUser(t *testing.T) {
	sharedDB["user1"] = "Alice"
	t.Log("created user1=Alice")
}

func TestShuffle_GetUser(t *testing.T) {
	name, ok := sharedDB["user1"]
	if !ok {
		t.Fatal("user1 not found — depends on TestShuffle_CreateUser running first!")
	}
	if name != "Alice" {
		t.Errorf("expected Alice, got %s", name)
	}
	t.Log("got user1=Alice")
}

func TestShuffle_UpdateUser(t *testing.T) {
	sharedDB["user1"] = "Alice Updated"
	t.Log("updated user1=Alice Updated")
}

func TestShuffle_DeleteUser(t *testing.T) {
	_, ok := sharedDB["user1"]
	if !ok {
		t.Fatal("user1 not found — depends on TestShuffle_CreateUser running first!")
	}
	delete(sharedDB, "user1")
	t.Log("deleted user1")
}
