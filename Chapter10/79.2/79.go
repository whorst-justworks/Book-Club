package main

import (
	"database/sql"
	"fmt"
	"log"
)

// leakyQuery does NOT close the sql.Rows, leaking the underlying connection.
func leakyQuery(db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}

	// BUG: rows is never closed!
	// The underlying database connection is never returned to the pool.
	// Eventually the pool is exhausted and queries start blocking or failing.
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

// properQuery closes sql.Rows via defer, returning the connection to the pool.
func properQuery(db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

func main() {
	// Not runnable without a real DB — just showing the pattern.
	fmt.Println("See leakyQuery() and properQuery() for the example.")
}
