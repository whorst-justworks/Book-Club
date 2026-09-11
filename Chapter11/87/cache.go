package main

import "time"

// isExpired reports whether an item created at created is at least one hour
// old relative to now. Passing now in keeps the function easy to test.
func isExpired(created, now time.Time) bool {
	expiresAt := created.Add(time.Hour)
	return now.After(expiresAt) || now.Equal(expiresAt)
}
