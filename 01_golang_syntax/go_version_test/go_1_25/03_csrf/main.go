package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Normal handlers
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK - GET request\n")
	})
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK - POST request\n")
	})

	// Create CSRF protection middleware
	csrp := http.NewCrossOriginProtection()

	// Allow requests from these trusted origins
	csrp.AddTrustedOrigin("https://example.com")
	csrp.AddTrustedOrigin("https://*.example.com")

	// Wrap the handlers with CSRF protection
	server := &http.Server{
		Addr:    ":8080",
		Handler: csrp.Handler(mux),
	}

	log.Println("Server running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
