package _8_5

import (
	"database/sql"
	"fmt"
	"log"
)

func threeErrors(db *sql.DB) error {
	// Error 1: executing the query
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}
	// Error 2: closing the rows (releases connection back to pool)
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("close error:", err)
		}
	}()

	for rows.Next() {
		var name, email string
		// Error 3: scanning a row
		if err := rows.Scan(&name, &email); err != nil {
			return fmt.Errorf("scan error: %w", err)
		}
		fmt.Println(name, email)
	}
	return nil
}

func threeErrorsPlusRowsErr(db *sql.DB) error {
	// Error 1: executing the query
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}
	// Error 2: closing the rows
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("close error:", err)
		}
	}()

	for rows.Next() {
		var name, email string
		// Error 3: scanning a row
		if err := rows.Scan(&name, &email); err != nil {
			return fmt.Errorf("scan error: %w", err)
		}
		fmt.Println(name, email)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows iteration error: %w", err)
	}

	return nil
}

// threeErrorsPlusRowsErr handles the same three errors as above, plus checks
// rows.Err() to determine if the rows.Next() loop ended due to an error
// (e.g. network failure mid-iteration) rather than simply running out of rows.

// Error 4: rows.Err() reports any error encountered during iteration.
// rows.Next() returns false on BOTH end-of-results AND errors.
// Without this check, we'd silently miss iteration failures.

// Situations in which rows.Err() can happen
//1. Network failure — the connection to the database drops mid-iteration (timeout, TCP reset, server crash)
//2. Context cancellation — the context.Context passed to QueryContext is cancelled or hits its deadline while you're still iterating
//3. Driver-level errors — the database driver encounters a protocol error while fetching the next batch of rows (e.g., corrupted response)
//4. Server-side timeout — the database server kills the query mid-stream (e.g., statement_timeout in Postgres fires after some rows were already sent)
//5. Out of memory on the server — the DB runs out of resources while streaming a large result set
