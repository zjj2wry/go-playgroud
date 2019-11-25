package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var srv http.Server

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, %v\n", time.Now())
	})

	srv.Addr = ":7070"
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
