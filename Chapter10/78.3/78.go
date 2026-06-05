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

	// Prepared statements have two main benefits:
	//
	// EFFICIENCY: The SQL is parsed and planned once by the database, then
	// executed many times with different parameters. This avoids repeated
	// parsing/planning overhead when running the same query in a loop.
	//
	// SECURITY: Parameters are sent separately from the SQL text, so user
	// input can never be interpreted as SQL. This eliminates SQL injection
	// attacks — the DB knows exactly what is a command vs what is data.

	stmt, err := db.Prepare("SELECT name, email FROM users WHERE age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Now we can execute efficiently with different values
	ages := []int{21, 30, 40}
	for _, age := range ages {
		rows, err := stmt.Query(age)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var name, email string
			if err := rows.Scan(&name, &email); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("age > %d: %s (%s)\n", age, name, email)
		}
		rows.Close()
	}
	
	// CON: You must remember to Close() the prepared statement when done.
	// Each prepared statement holds a connection from the pool until closed.
	// Forgetting to close leaks connections and can exhaust the pool.

}
