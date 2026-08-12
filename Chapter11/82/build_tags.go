package main

// Mistake #82: Not Categorizing Tests
//
// Use build tags to separate unit tests from integration/e2e tests.
// This way `go test ./...` only runs fast unit tests by default.
//
// Usage:
//   go test ./...                          → runs only unit tests (fast)
//   go test -tags=integration ./...        → runs unit + integration tests

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func ConnectToDatabase() (string, error) {
	// Simulate a slow database connection
	return "connected", nil
}

func InsertRecord(conn string, name string) error {
	// Simulate a slow database insert
	return nil
}
