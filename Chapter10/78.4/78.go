package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// When a column can be NULL, you can't scan into a plain Go type —
	// you need the sql.Null* wrapper types:
	//
	//   sql.NullString  — for nullable text/varchar columns
	//   sql.NullBool    — for nullable boolean columns
	//   sql.NullInt64   — for nullable integer columns
	//   sql.NullFloat64 — for nullable float/numeric columns
	//   sql.NullTime    — for nullable timestamp/date columns

	var name string
	var nickname sql.NullString // nickname column may be NULL

	err = db.QueryRow("SELECT name, nickname FROM users WHERE id = $1", 1).
		Scan(&name, &nickname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", name)

	// Check .Valid to see if the value is non-NULL
	if nickname.Valid {
		fmt.Println("Nickname:", nickname.String)
	} else {
		fmt.Println("Nickname: (none)")
	}
}
