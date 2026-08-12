package main

// Mistake #83: Not Enabling the Race Flag
//
// Go's race detector (`go test -race ./...`) finds data races at runtime.
// Without it, race conditions silently corrupt data and are hard to reproduce.

// Counter has a race condition: multiple goroutines read/write `count`
// without synchronization.
type Counter struct {
	count int
}

func NewCounter() *Counter {
	return &Counter{}
}

// Increment is NOT safe for concurrent use.
func (c *Counter) Increment() {
	c.count++ // DATA RACE: unsynchronized read-modify-write
}

// Value is NOT safe for concurrent use.
func (c *Counter) Value() int {
	return c.count // DATA RACE: unsynchronized read
}
