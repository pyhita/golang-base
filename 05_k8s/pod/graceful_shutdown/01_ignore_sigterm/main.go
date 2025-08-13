package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the handler function for the root path
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Respond with a greeting
		fmt.Fprintln(w, "Hello World!")
	})

	// Specify the port to listen on
	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)

	// Start the HTTP server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
