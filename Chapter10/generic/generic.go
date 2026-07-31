package main

import "fmt"

// Before Go 1.27: generic types could have methods, but methods
// themselves couldn't introduce NEW type parameters.
// You had to use package-level functions as a workaround.

type Cache[K comparable, V any] struct {
	data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{data: make(map[K]V)}
}

func (c *Cache[K, V]) Set(key K, val V) {
	c.data[key] = val
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	v, ok := c.data[key]
	return v, ok
}

// Go 1.27: Methods can now declare their OWN type parameters.
// This lets you write a method that operates on a type not known
// when the struct was defined.

// Transform applies a function to a cached value and returns the result
// as a different type — the method introduces its own type parameter R.
func (c *Cache[K, V]) Transform[R any](key K, fn func(V) R) (R, bool) {

	//   The [R any] after the method name is the new part — it declares a type parameter
	//  R that belongs to the method itself, separate from the struct's K and V.
	v, ok := c.data[key]
	if !ok {
		var zero R
		return zero, false
	}
	return fn(v), true
}

// Before 1.27, you'd need a package-level function:
//   func Transform[K comparable, V any, R any](c *Cache[K, V], key K, fn func(V) R) (R, bool)

func main() {
	c := NewCache[string, int]()
	c.Set("age", 30)
	c.Set("score", 95)

	// R = string — transform an int into a formatted string
	result, ok := c.Transform("age", func(v int) string {
		return fmt.Sprintf("%d years old", v)
	})
	fmt.Println(result, ok) // "30 years old" true

	// R = bool — transform an int into a pass/fail check
	passed, ok := c.Transform("score", func(v int) bool {
		return v >= 70
	})
	fmt.Println("passed:", passed, ok) // passed: true true

	// R = float64 — transform into a percentage
	pct, ok := c.Transform("score", func(v int) float64 {
		return float64(v) / 100.0
	})
	fmt.Println("percentage:", pct, ok) // percentage: 0.95 true

	// Key doesn't exist — returns zero value of R
	missing, ok := c.Transform("missing", func(v int) string {
		return "found"
	})
	fmt.Println("missing:", missing, ok) // missing:  false
}

//A way to think of this:
// The method Transform accepts a function which can return any type, and Transform's
// return type becomes is the function's return type

// go1.27rc1 run generic.go
