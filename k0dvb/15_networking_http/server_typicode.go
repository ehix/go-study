package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func httpError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"

	// e.g. todo/1
	resp, err := http.Get(base + r.URL.Path[1:])
	// If typicode can't be resolved
	httpError(err, w)

	defer resp.Body.Close()

	var item todo
	// resp.Body is a ReadCloser, can be passed into decoder to read body of resp directly.
	err = json.NewDecoder(resp.Body).Decode(&item)
	httpError(err, w)

	tmpl := template.New("mine")
	tmpl.Parse(form)
	tmpl.Execute(w, item)
}

// Can do a client request against another server, and formats into a web browser
func main() {
	http.HandleFunc("/", handler)                // Bind a handler to the top level route
	log.Fatal(http.ListenAndServe(":8080", nil)) // Open TCP socket that can except http requests
}
