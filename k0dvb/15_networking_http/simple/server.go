package main

import (
	"fmt"
	"log"
	"net/http"
)

// This is a concurrent server.
// The http package will handle as many connections as current hardware will support.

func handler(w http.ResponseWriter, r *http.Request) {
	// All handlers have the same signatures, requests and responses.
	// Fprintf first param is just something that can be written to, e.g. Stderr.
	fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])
	// URL path from the client is returned.
	// e.g. curl http://localhost:8080/Alex -i
	// > Hello, world! from Alex
}

func main() {
	http.HandleFunc("/", handler)                // Bind a handler to the top level route
	log.Fatal(http.ListenAndServe(":8080", nil)) // Open TCP socket that can except http requests
}
