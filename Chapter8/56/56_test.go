package _56

import (
	"testing"
)

func Test_sequential20(t *testing.T) {
	server := startServer()
	defer stopServer(server)
	parallelGoroutinesTasks(1, 20)
}

func Test_parallel16GoroutinesWith200Tasks(t *testing.T) {
	server := startServer()
	defer stopServer(server)
	parallelGoroutinesTasks(16, 32)
}

func Test_parallel200GoroutinesWith200Tasks(t *testing.T) {
	server := startServer()
	defer stopServer(server)
	parallelGoroutinesTasks(200, 200)
}

func Test_parallel500GoroutinesWith500Tasks(t *testing.T) {
	server := startServer()
	defer stopServer(server)
	parallelGoroutinesTasks(500, 500)
}

func Test_parallel8_000_000GoroutinesWith500Tasks(t *testing.T) {
	server := startServer()
	defer stopServer(server)
	parallelGoroutinesTasks(8_000_000, 500)
}
