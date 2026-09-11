package main

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-API-VERSION", "1.0")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("hello"))
}

func readAll(r io.Reader, retries int) ([]byte, error) {
	var result []byte
	buffer := make([]byte, 1)

	for {
		n, err := r.Read(buffer)
		result = append(result, buffer[:n]...)
		if err == nil {
			continue
		}
		if err == io.EOF {
			return result, nil
		}
		if retries == 0 {
			return result, err
		}
		retries--
	}
}
