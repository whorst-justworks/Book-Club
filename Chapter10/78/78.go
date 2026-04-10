package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3" // SQLite driver for demo
)

// Mistake #78: Forgetting that sql.Open doesn't establish connections

var dsn = ":memory:" // In-memory SQLite database for demo

// badExample assumes sql.Open establishes a connection to the database
func badExample() error {
	fmt.Println("=== Bad Example: Not verifying connection ===")

	// Problem: sql.Open only validates the DSN format
	// It does NOT actually connect to the database or verify it's reachable
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}
	defer db.Close()

	// At this point, we THINK we're connected, but we might not be!
	// If the DSN was wrong or database is unreachable, we won't know until
	// we try to execute a query later in the application
	fmt.Println("sql.Open succeeded - but are we actually connected?")

	// The error might not surface until much later when we try to use it
	// This could cause issues during application startup vs runtime

	return nil
}

// ⏺ DSN stands for Data Source Name. It's a string that contains all the information needed to connect to a database.
//
//  A DSN typically includes:
//  - Username and password for authentication
//  - Host (server address) and port number
//  - Database name to connect to
//  - Optional connection parameters (timeouts, charset, etc.)

// When you call sql.Open(), it validates that the DSN string is in the correct format for that driver, but it doesn't actually try to connect to the database server. That's why you
//  need Ping() to verify the connection actually works.

// goodExample verifies the connection is actually established
func goodExample() error {
	fmt.Println("\n=== Good Example: Verifying connection with Ping ===")

	// Step 1: Open the database (validates DSN format)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}
	defer db.Close()

	// Step 2: Verify the connection is actually established
	// Ping() will actually connect to the database and verify it's reachable
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	fmt.Println("sql.Open succeeded AND Ping confirmed connection!")

	// Now we know for certain the database is reachable
	// This is especially important during application startup
	// to fail fast if database is misconfigured

	return nil
}

func main() {

	badExample()
	goodExample()

}
