package main

import (
	"fmt"
	"time"
)

func wrongDuration() {

	// Mistake 1: Using plain integer (this is nanoseconds, not seconds!)
	seconds := 5
	duration := time.Duration(seconds)
	fmt.Printf("Intended 5 seconds, got %v (%d nanoseconds)\n", duration, duration)

	// Mistake 2: Thinking time.Second * 1000 gives 1 second
	// This actually gives 1000 seconds (16.6 minutes)!
	wrongTimeout := time.Second * 1000
	fmt.Printf("Intended 1 second, got %v\n", wrongTimeout)

	// Mistake 3: Config value without time unit
	configTimeout := 100 // meant to be 100ms
	timeout := time.Duration(configTimeout)
	fmt.Printf("Intended 100ms, got %v\n\n", timeout)
}

// rightDuration demonstrates correct ways to specify time durations
func rightDuration() {

	// Correct 1: Multiply integer by time unit
	seconds := 5
	duration := time.Duration(seconds) * time.Second
	fmt.Printf("5 seconds: %v\n", duration)

	// Correct 2: Use time constants directly
	timeout := 30 * time.Second
	fmt.Printf("30 seconds: %v\n", timeout)

	// Correct 3: Config value with proper time unit
	configTimeout := 100 // milliseconds
	properTimeout := time.Duration(configTimeout) * time.Millisecond
	fmt.Printf("100ms: %v\n", properTimeout)

	// Correct 4: Combine multiple time units
	combined := 2*time.Hour + 30*time.Minute + 15*time.Second
	fmt.Printf("Combined: %v\n", combined)

}

func main() {

	wrongDuration()
	rightDuration()

}
