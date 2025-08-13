package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Define the handler function
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	// Create an HTTP server
	server := &http.Server{Addr: ":8080"}

	// Channel to listen for OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server
	go func() {
		fmt.Println("Starting server on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()

	// Wait for SIGTERM or SIGINT
	<-signalChan
	fmt.Println("\nSIGTERM received. Waiting 10 seconds to drain requests...")

	// Wait 10 seconds to allow ongoing requests to complete
	time.Sleep(10 * time.Second)

	// Gracefully shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error during shutdown: %s\n", err)
	}

	fmt.Println("Server stopped gracefully.")
}
