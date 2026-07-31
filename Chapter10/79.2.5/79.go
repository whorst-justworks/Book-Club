package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

//   go1.27rc1 run Chapter10/79.2.5/79.go

// leakyGoroutine sends on a channel that nobody receives from.
// The goroutine blocks forever — a classic goroutine leak.
func leakyGoroutine() {
	ch := make(chan int)

	go func() {
		// This goroutine will block forever because nothing reads from ch.
		ch <- 42
	}()

	// We "forget" about ch — the sender is permanently stuck.
}

// anotherLeak blocks on receiving from a channel that nobody sends to.
func anotherLeak() {
	done := make(chan struct{})

	go func() {
		// Waiting for a signal that never comes.
		<-done
	}()

	// We never close(done) or send on it — goroutine leaked.
}

func main() {
	//fmt.Println("=== Goroutine Leak Demo with Go 1.27 Leak Profile ===\n")

	//baseline := runtime.NumGoroutine()
	//fmt.Printf("Baseline goroutines: %d\n", baseline)

	// Create some leaked goroutines
	for i := 0; i < 10; i++ {
		leakyGoroutine()
	}
	for i := 0; i < 5; i++ {
		anotherLeak()
	}

	// Let goroutines settle into their blocked state, then GC to trigger
	// reachability analysis for the leak profile.
	time.Sleep(100 * time.Millisecond)
	runtime.GC()

	// Go 1.27 introduces the "goroutineleak" profile in runtime/pprof.
	// It uses GC reachability analysis to detect goroutines blocked on
	// channels, mutexes, or condition variables that are unreachable —
	// meaning no other goroutine can ever unblock them.
	leakProfile := pprof.Lookup("goroutineleak")
	if leakProfile == nil {
		os.Exit(1)
	}

	// WriteTo triggers a GC cycle internally to perform reachability analysis.
	leakProfile.WriteTo(os.Stdout, 1)
}

// go1.27rc1 run 79.go
