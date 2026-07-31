package _1

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// The default http.Server has no timeouts set:
	//   &http.Server{}
	//
	// Without these timeouts:
	// - No ReadHeaderTimeout: Slowloris attacks — a client opens a connection and sends
	//                         headers byte-by-byte, tying up a goroutine and file descriptor
	//                         indefinitely. Enough of these exhaust your server.
	// - No ReadTimeout:       A client that sends a request body very slowly (or never finishes)
	//                         holds the connection and goroutine open forever.
	// - No Handler timeout:   A handler that blocks (e.g. waiting on a downstream service)
	//                         never frees the goroutine. Under load, goroutines pile up
	//                         until the server OOMs.
	//
	// Always set explicit timeouts:
	defaultServer := &http.Server{}

	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		Handler:           http.TimeoutHandler(handler{}, time.Second, "foo"),
	}
	_ = s
	fmt.Println(defaultServer)
}

type handler struct{}

func (h handler) ServeHTTP(http.ResponseWriter, *http.Request) {}
