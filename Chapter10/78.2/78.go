package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("MariaDB", "MariaDB://localhost/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SetMaxOpenConns: limits total open connections (in-use + idle).
	// Prevents overwhelming the database with too many connections.
	// Default is 0 (unlimited). Set this to match your DB's capacity.
	db.SetMaxOpenConns(25)

	// SetMaxIdleConns: how many connections to keep ready in the pool.
	// Higher = faster response (no dial overhead), but uses more memory.
	// Lower = less memory, but more latency creating new connections.
	// Default is 2. Should be <= MaxOpenConns.
	db.SetMaxIdleConns(10)

	// SetConnMaxIdleTime: how long an idle connection sits unused before closing.
	// Helps shrink the pool during low-traffic periods.
	// Prevents holding connections open that aren't needed.
	// Default is 0 (no limit).
	db.SetConnMaxIdleTime(5 * time.Minute)

	// SetConnMaxLifetime: max total age of a connection, regardless of use.
	// Forces recycling so connections don't go stale (firewalls, DNS changes,
	// DB failovers). Also spreads reconnection load over time instead of
	// all connections dying at once.
	// Default is 0 (no limit). Should be shorter than any DB/infra timeout.
	db.SetConnMaxLifetime(1 * time.Hour)

	// Use the pool — sql.DB handles checkout/checkin automatically
	var now time.Time
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB time:", now)

	// Check pool stats
	stats := db.Stats()
	fmt.Printf("Open: %d, InUse: %d, Idle: %d\n",
		stats.OpenConnections, stats.InUse, stats.Idle)
}
