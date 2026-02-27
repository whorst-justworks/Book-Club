package _0

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Misunderstanding Go contexts (#60)

// Deadlines prevent operations from running indefinitely and consuming resources when external services are slow or unresponsive.
// They ensure your application remains responsive by enforcing maximum execution times for operations like API calls or database queries.
func TestContextDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	//ctx.Done() channel acts like a read-only broadcast mechanism - when it closes, all goroutines
	//listening are notified simultaneously

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped, deadline exceeded:", ctx.Err())
				done <- true
				return
			default:
				fmt.Println("working...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	<-done
}

// Cancellation signals allow you to stop goroutines gracefully when their work is no longer needed, preventing goroutine leaks.
// This is critical for scenarios like user-cancelled requests or when a parent operation fails and all child operations should stop.
func TestContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("operation completed")
	case <-ctx.Done():
		fmt.Println("context cancelled:", ctx.Err())
	}
}

// Context values allow request-scoped data like user IDs or trace IDs to flow through your call chain without modifying function signatures.
// This should only be used for request-scoped data that crosses API boundaries, not for passing optional parameters to functions.
func TestContextValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "userID", 42)

	done := make(chan bool)
	//context.Value() does provide immutable key-value storage that propagates down the call tree

	go func(ctx context.Context) {
		userID := ctx.Value("userID")
		fmt.Println("goroutine accessing userID:", userID)
		done <- true
	}(ctx)

	<-done
}
