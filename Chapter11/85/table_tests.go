package main

// Mistake #85: Not Using Table-Driven Tests
//
// Table-driven tests reduce duplication, make it easy to add new cases,
// and give each case a descriptive name that shows up in test output.

import (
	"errors"
	"strings"
)

var ErrInvalidAge = errors.New("age must be between 0 and 150")
var ErrEmptyName = errors.New("name cannot be empty")

type User struct {
	Name  string
	Email string
	Age   int
}

func ValidateUser(u User) error {
	if strings.TrimSpace(u.Name) == "" {
		return ErrEmptyName
	}
	if u.Age < 0 || u.Age > 150 {
		return ErrInvalidAge
	}
	return nil
}
