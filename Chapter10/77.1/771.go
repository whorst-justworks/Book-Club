package main

import (
	"encoding/json"
	"fmt"
)

// Using map[string]any: Issues with type safety and error handling

// badExample uses map[string]any which requires type assertions and loses type safety
func badExample() error {
	fmt.Println("=== Bad Example: Using map[string]any ===")

	b := []byte(`{
		"id": 123.456,
		"name": "John Doe",
		"active": true,
		"address": {
			"city": "New York",
			"state": "NY"
		}
	}`)
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	id := m["id"]
	// If this isn't a string we get a panic
	name := m["name"].(string)
	active := m["active"].(bool)

	fmt.Println("Result ", id, name, active)
	fmt.Printf("Type assertion on ID %T\n", id)

	return nil
}

// Problem 1: Need type assertions everywhere
// Problem 2: Type assertions can panic if type is wrong
// Problem 3: No compile-time type checking
// JSON numbers unmarshal to float64

// goodExample uses a proper struct with defined types
func goodExample() error {
	fmt.Println("\n=== Good Example: Using a struct ===")

	type Address struct {
		City  string `json:"city"`
		State string `json:"state"`
	}

	type User struct {
		ID      int     `json:"id"`
		Name    string  `json:"name"`
		Active  bool    `json:"active"`
		Address Address `json:"address"`
	}

	b := []byte(`{
		"id": 123,
		"name": "John Doe",
		"active": true,
		"address": {
			"city": "New York",
			"state": "NY"
		}
	}`)
	var user User
	err := json.Unmarshal(b, &user)
	if err != nil {
		return err
	}

	fmt.Println("Result ", user.ID, user.Name, user.Active)
	fmt.Printf("Type assertion on ID %T\n", user.ID)
	fmt.Printf("City: %s\n", user.Address.City)

	return nil
}

func main() {
	//badExample()
	goodExample()
}
