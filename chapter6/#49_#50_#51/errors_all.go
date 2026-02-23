package _49__50__51

import (
	"errors"
	"fmt"
	"io/fs"
)

// Section #49: Ignoring when to wrap an error
//
// 1. Use fmt.Errorf with %w to wrap errors when adding context
// 2. Wrapping preserves the error chain for errors.Is/As checks
// 3. Return errors as-is when no additional context is needed
// 4. Wrapping creates parent-child relationship in error chain
//
// When to wrap:
// - Adding context about what operation failed
// - Transforming errors between layers (e.g., DB error -> domain error)
// - Want callers to check for underlying error types/values
//
// When NOT to wrap:
// - Passing errors through without adding information
// - Want to hide implementation details from callers
func demonstrateErrorWrapping() error {
	err := doSomething()
	if err != nil {
		return fmt.Errorf("operation failed")
	}

	var SpecificError = errors.New("This is an error")
	// GOOD: Wrapping preserves error chain
	err := SpecificError
	if err != nil {
		newError := fmt.Errorf("operation failed: %w", err) // Use %w to wrap
		// NewError -> SpecificError
		return newError
	}

	// GOOD: Return as-is when no context needed
	func() error {
		return doSomething() // No wrapping needed
	}
}

// Section #50: Comparing an error type inaccurately
//
// NEVER use type assertion or type switch on wrapped errors directly, use errors.As()
// to check if error is or wraps a specific type and to unwrap the error chain automatically

func demonstrateErrorTypeComparison() error {
	err := fmt.Errorf("wrapped: %w", &fs.PathError{Op: "open", Path: "/tmp/file"})

	// BAD: Type assertion doesn't work on wrapped errors
	if _, ok := err.(*fs.PathError); ok {
		// This will NEVER execute because err is fmt.wrapError, not *fs.PathError
		fmt.Println("This won't print")
	}

	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		// This WILL execute - errors.As finds *fs.PathError in chain
		fmt.Printf("Path error: op=%s, path=%s\n", pathErr.Op, pathErr.Path)
		return nil
	}

	return err
}

// Section #51: Comparing an error value inaccurately
// NEVER use == to compare sentinel errors in wrapped errors,
// use errors.Is() to check if error is or wraps a specific value
func demonstrateErrorValueComparison() error {
	var ErrNotFound = errors.New("not found")

	// Simulate wrapped sentinel error
	err := fmt.Errorf("failed to get user: %w", ErrNotFound)

	// BAD: Direct comparison doesn't work on wrapped errors
	if err == ErrNotFound {
		// This will NEVER execute because err is fmt.wrapError, not ErrNotFound
		fmt.Println("This won't print")
	}

	// GOOD: Use errors.Is() to check value in error chain
	if errors.Is(err, ErrNotFound) {
		// This WILL execute - errors.Is finds ErrNotFound in chain
		fmt.Println("Not found error detected")
		return nil
	}

	return err
}

// Why errors.Is():
// - Handles wrapped errors correctly
// - Traverses entire error chain
// - Works with custom Is() methods on error types

func doSomething() error {
	return errors.New("something went wrong")
}
