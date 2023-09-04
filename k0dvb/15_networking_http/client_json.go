package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Server which will return json, accessed via /todos
const url = "https://jsonplaceholder.typicode.com"

// Fields have to be in upper for them to be decoded
// Not all incoming JSON fields need to be included
type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func main() {
	resp, err := http.Get(url + "/todos/1")
	check(err)

	// Must close body otherwise sockets will continue to be opened
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		check(err)

		var item todo // create a space to decode into
		err = json.Unmarshal(body, &item)
		check(err)
		fmt.Printf("%#v\n", item)
	}
}
