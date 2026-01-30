package chapter1

import "errors"

// Anti-pattern: Overly Using Getters and Setters (Chapter 2)
type BadUser struct {
	name  string
	email string
	age   int
}

func (u *BadUser) GetName() string {
	return u.name
}

func (u *BadUser) GetEmail() string {
	return u.email
}

func (u *BadUser) GetAge() int {
	return u.age
}

func (u *BadUser) SetName(name string) {
	u.name = name
}

func (u *BadUser) SetEmail(email string) {
	u.email = email
}

func (u *BadUser) SetAge(age int) {
	u.age = age
}

// GoodUser demonstrates the idiomatic Go approach
// Export fields that should be accessible, keep private ones unexported
type GoodUser struct {
	Name  string // Exported - can be accessed directly
	Email string // Exported - can be accessed directly
	age   int    // Unexported - internal only
}

// Only add getters/setters when they provide actual value
// For example, when you need validation or transformation
func (u *GoodUser) SetAge(age int) error {
	if age < 0 || age > 150 {
		return ErrInvalidAge
	}
	u.age = age
	return nil
}

func (u *GoodUser) Age() int {
	return u.age
}

var ErrInvalidAge = errors.New("age must be between 0 and 150")
