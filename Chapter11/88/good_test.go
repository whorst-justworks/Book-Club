package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/iotest"
)

// GOOD: httptest supplies request and response objects that follow the
// net/http testing contract. The test only checks the handler's observable
// behavior: status, header, and body. It does not need to maintain a custom
// ResponseWriter or start a real network server.
func TestHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Handler(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, response.Code)
	}
	if response.Header().Get("X-API-VERSION") != "1.0" {
		t.Fatal("expected API version header")
	}
	if response.Body.String() != "hello" {
		t.Fatalf("expected hello, got %q", response.Body.String())
	}
}

// GOOD: iotest.TimeoutReader simulates a reader that temporarily fails. This
// lets the test verify retry behavior using a standard, reusable test reader
// instead of duplicating the reader protocol and error handling itself.
func TestReadAllWithTimeout(t *testing.T) {
	reader := iotest.TimeoutReader(strings.NewReader("hello"))

	got, err := readAll(reader, 3)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "hello" {
		t.Fatalf("expected hello, got %q", got)
	}
}

// Other useful httptest functionality:
//   - NewServer starts an HTTP test server with a generated URL.
//   - NewTLSServer starts an HTTPS test server and provides its configured client.
//   - NewUnstartedServer creates a server that can be configured before Start.
//   - ResponseRecorder captures status, headers, and body from a handler.
//   - Server.Client returns an HTTP client configured to trust the test server.
//
// Other useful iotest functionality:
//   - DataReader turns a byte slice into a Reader that reports read data.
//   - ErrReader returns a reader that always fails with a supplied error.
//   - HalfReader limits each read to half of the requested buffer.
//   - OneByteReader limits each read to one byte.
//   - OneLineReader limits each read to one line.
//   - TestReader checks Reader behavior against expected data and errors.
//   - TruncateWriter writes only a fixed number of bytes before returning an error.
