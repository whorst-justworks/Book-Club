package _1

import (
	"net"
	"net/http"
	"time"
)

func main() {

	// The default http.Client has no timeouts set:
	//   http.DefaultClient = &http.Client{}
	//
	// Without these timeouts:
	// - No Timeout:               A slow or unresponsive server can block your goroutine forever,
	//                              leaking connections and eventually exhausting memory.
	// - No DialContext.Timeout:    DNS resolution or TCP handshake hangs indefinitely if the
	//                              host is unreachable (e.g. behind a firewall that drops packets).
	// - No TLSHandshakeTimeout:   A misbehaving server can stall during TLS negotiation,
	//                              holding the connection open without ever completing.
	// - No ResponseHeaderTimeout: The server accepts the connection but never sends headers,
	//                              leaving the client waiting with no way to reclaim the goroutine.
	//
	// Always create a client with explicit timeouts:

	http.DefaultClient = &http.Client{}

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second,
		},
	}
	_ = client
}
