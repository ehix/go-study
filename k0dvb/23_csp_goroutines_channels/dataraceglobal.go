package main

import (
	"fmt"
	"log"
	"net/http"
)

// var nextID int
var nextID = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	// WRONG WAY:
	// fmt.Fprint(w, "<h1>You got %d<h1>", nextID)
	// The webserver built into the standard lib is concurrent.
	// nextID++ //UNSAFE - read, modify, write operation.

	// WITH CHAN:
	fmt.Fprint(w, "<h1>You got %d<h1>", <-nextID)

}

func counter() {
	// Keep generating numbers and put into channel.
	// A channel can't be written to, unless somethings reading from it.
	// If the channel is not ready to be read from, it's blocked.
	for i := 0; ; i++ {
		nextID <- i
		// the iteration gets managed by the channel.
	}
}

func main() {
	go counter()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
