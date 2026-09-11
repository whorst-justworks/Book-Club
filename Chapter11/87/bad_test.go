//go:build bad

package main

import (
	"testing"
	"time"
)

// BAD: The test and the code use different readings of the real clock. The
// test creates an item at the exact boundary, but the function reads the clock
// a moment later, so this test fails.
//
// Run with: go test -v -tags=bad ./...
func TestTrimOlderThanWithRealTime(t *testing.T) {
	created := time.Now().Add(-time.Hour)
	// We expect an item exactly one hour old to be not expired yet. However,
	// isExpiredWithRealTime calls time.Now() again, so the item is already
	// slightly more than one hour old by the time it is checked.
	if isExpiredWithRealTime(created) {
		t.Fatal("expected item not to be expired at the exact boundary")
	}
}

func TestIsExpiredWith59Minutes(t *testing.T) {
	created := time.Now().Add(-time.Hour)
	created = created.Add(59 * time.Minute)

	if isExpiredWithRealTime(created) {
		t.Fatal("item should not be expired after 59 minutes")
	}
}

func isExpiredWithRealTime(created time.Time) bool {
	expiresAt := created.Add(time.Hour)
	now := time.Now()
	return now.After(expiresAt) || now.Equal(expiresAt)
}
