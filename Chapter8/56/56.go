package _56

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// Thinking Concurrency is always faster

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

// startServer starts an HTTP server that simulates slow API responses
func startServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/work", func(w http.ResponseWriter, r *http.Request) {
		// Simulate random processing time between 50-3000ms
		duration := 50 + rand.Intn(2950)
		time.Sleep(time.Duration(duration) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("work completed"))
	})

	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	go server.ListenAndServe()

	// Wait for server to be ready
	time.Sleep(100 * time.Millisecond)

	return server
}

// stopServer shuts down the HTTP server
func stopServer(server *http.Server) {
	if server != nil {
		server.Close()
		time.Sleep(50 * time.Millisecond) // Give time to shutdown
	}
}

// httpWork makes an HTTP request to simulate I/O-bound work
func httpWork() error {
	resp, err := httpClient.Get("http://localhost:9090/api/work")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)
	return nil
}

// sequential20 calls httpWork 20 times sequentially
func sequential20() {
	for i := 0; i < 20; i++ {
		httpWork()
	}
}

// worker processes jobs from the jobs channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		httpWork() // I/O-bound work - goroutine yields during network I/O
		results <- j
	}
}

// parallelGoroutinesTasks executes tasks using a worker pool pattern
func parallelGoroutinesTasks(numGoRoutines int, numRestCalls int) {
	start := time.Now()

	jobs := make(chan int, numRestCalls)
	results := make(chan int, numRestCalls)

	// Start worker goroutines
	for w := 1; w <= numGoRoutines; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to workers
	for j := 1; j <= numRestCalls; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect all results
	for a := 1; a <= numRestCalls; a++ {
		<-results
	}

	elapsed := time.Since(start)
	fmt.Printf("Parallel with %d go routines (%d number of rest calls): %.2f seconds\n", numGoRoutines, numRestCalls, elapsed.Seconds())
}
