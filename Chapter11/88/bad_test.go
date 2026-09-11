//go:build bad

package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// BAD: Manually creating HTTP test doubles is verbose and easy to get wrong.
//
// This test has to build a request and maintain a custom ResponseWriter just
// to call a handler. The custom writer only implements the small part of the
// HTTP contract used by this handler. It does not behave like the writer used
// by a real HTTP server and could hide bugs involving headers, body writes, or
// optional ResponseWriter interfaces.
//
// httptest.NewRequest and httptest.NewRecorder provide standard HTTP-aware
// test objects, so a good test can focus on the handler's result instead of
// maintaining infrastructure that imitates net/http.
//
// Run with: go test -v -tags=bad ./...
func TestHandlerWithoutHttptest(t *testing.T) {
	request := &http.Request{
		Method: http.MethodGet,
		Body:   io.NopCloser(strings.NewReader("")),
	}
	response := &responseWriter{header: make(http.Header)}

	Handler(response, request)

	if response.status != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, response.status)
	}
}

// This is a partial and fragile implementation of http.ResponseWriter. If
// the handler later depends on another ResponseWriter capability, this test
// double will not behave like a real server.
type responseWriter struct {
	header http.Header
	status int
	body   []byte
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
}

func (w *responseWriter) Write(body []byte) (int, error) {
	w.body = append(w.body, body...)
	return len(body), nil
}

func TestReadAllWithoutIotest(t *testing.T) {
	// This is another hand-written test double. The test must define its state,
	// error type, and exact sequence of reads correctly. io.ErrNoProgress is
	// also not a real timeout error, so this may not exercise the failure
	// behavior the test is intended to cover.
	reader := &timeoutReader{reader: strings.NewReader("hello")}

	got, err := readAll(reader, 3)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "hello" {
		t.Fatalf("expected hello, got %q", got)
	}
}

// iotest.TimeoutReader already provides a reader that injects timeout-like
// failures. Reusing it avoids duplicating reader behavior and makes the test
// intention clear.
type timeoutReader struct {
	reader io.Reader
	failed bool
}

func (r *timeoutReader) Read(p []byte) (int, error) {
	if !r.failed {
		r.failed = true
		return 0, io.ErrNoProgress
	}
	return r.reader.Read(p)
}

// Other bad reasons to use real HTTP servers in unit tests:
//   - They are slower because every test pays for server startup and shutdown.
//   - They can be flaky due to ports, sockets, timing, and local resource limits.
//   - Tests may leak goroutines, connections, or listeners when cleanup is missed.
//   - Network scheduling makes failures less deterministic and harder to reproduce.
//   - Transport behavior can obscure the handler or client behavior under test.
//   - Tests become sensitive to machine, operating system, and network configuration.
//   - Parallel tests can contend for ports and other shared network resources.
//   - Injecting precise responses, delays, and failures is more difficult than with
//     httptest helpers or direct test doubles.
