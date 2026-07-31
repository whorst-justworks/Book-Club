package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"time"
)

// go run 79.go leak
// go run 79.go proper

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run 79.go [leak|proper]")
		os.Exit(1)
	}

	// Server returns a large body (64KB) so partial reads leave data buffered.
	payload := make([]byte, 1024*64)
	for i := range payload {
		payload[i] = 'A'
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(payload)
	}))
	defer server.Close()

	baseGoroutines := runtime.NumGoroutine()
	fmt.Printf("Baseline goroutines: %d\n", baseGoroutines)

	switch os.Args[1] {
	case "leak":
		fmt.Println("\n=== Leaky requests (body never closed, partially read) ===")
		for i := 0; i < 50; i++ {
			leakyRequest(server.URL)
		}
	case "proper":
		fmt.Println("\n=== Proper requests (body closed with defer) ===")
		for i := 0; i < 50; i++ {
			properRequest(server.URL)
		}
	default:
		fmt.Println("Usage: go run 79.go [leak|proper]")
		os.Exit(1)
	}

	time.Sleep(500 * time.Millisecond)

	//runtime.GC() triggers a manual garbage collection cycle. In this example, it forces Go to clean up any
	// unreferenced objects so we get an accurate goroutine count — ensuring the
	//leaked goroutines we see are genuinely stuck (pinned by unclosed bodies) and not just waiting to be collected.
	runtime.GC()

	// runtime.NumGoroutine() returns the number of currently active goroutines
	finalGoroutines := runtime.NumGoroutine()
	fmt.Printf("Goroutines now: %d (leaked: %d)\n", finalGoroutines, finalGoroutines-baseGoroutines)
}

// leakyRequest does NOT close the response body.
// It only reads a small portion, leaving the connection pinned open.
func leakyRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	// BUG: We only read part of the body and never close it.
	// The underlying TCP connection cannot be returned to the pool or released.
	// Each request pins a goroutine that waits to finish reading.
	buf := make([]byte, 128)
	n, _ := resp.Body.Read(buf)
	return string(buf[:n])
}

// properRequest closes the response body via defer, even if not fully read.
func properRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 128)
	n, _ := resp.Body.Read(buf)
	return string(buf[:n])
}
