package chapter1

import (
	"errors"
	"fmt"
)

// Anti-pattern: Overly Nested Code (Chapter 2)
// Problem: Deep nesting makes code hard to read and understand

// BadNestedCode demonstrates overly nested code
func BadNestedCode(items []string) error {
	if len(items) > 0 {
		for _, item := range items {
			if item != "" {
				if len(item) > 3 {
					if item[0] == 'a' {
						fmt.Println("Processing:", item)
						return nil
					} else {
						return errors.New("item must start with 'a'")
					}
				} else {
					return errors.New("item too short")
				}
			} else {
				return errors.New("empty item")
			}
		}
	} else {
		return errors.New("no items provided")
	}
	return nil
}

// GoodNestedCode demonstrates the improved approach using early returns
// and guard clauses to reduce nesting
func GoodNestedCode(items []string) error {
	if len(items) == 0 {
		return errors.New("no items provided")
	}

	for _, item := range items {
		if item == "" {
			return errors.New("empty item")
		}
		if len(item) <= 3 {
			return errors.New("item too short")
		}
		if item[0] != 'a' {
			return errors.New("item must start with 'a'")
		}
		fmt.Println("Processing:", item)
		return nil
	}
	return nil
}
