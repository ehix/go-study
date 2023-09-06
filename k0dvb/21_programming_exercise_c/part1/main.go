package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// NOTE: don't do this in real life!
// Floats can't properly represent money.
type dollars float32

// Wrapper around the float value to print currency symbol.
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// Methods doesn't need return type,
// writes to ResponseWriter to go back to original client.
// curl http://localhost:8080/list -i
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// curl http://localhost:8080/create?item=blah\&price=6 -i
func (db database) add(w http.ResponseWriter, req *http.Request) {
	// Can also get via JSON in the body of the request.
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %v", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %v", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "item %s added with price %s\n", item, db[item])
}

// $ curl http://localhost:8080/update?item=socks\&price=10 -i
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %v", item)
		http.Error(w, msg, http.StatusNotFound) // 400
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %v", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "new price %s for item %s\n", db[item], item)
}

// curl localhost:8080/read?item=socks -i
func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	value, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %v", item)
		http.Error(w, msg, http.StatusNotFound) // 400
		return
	}
	fmt.Fprintf(w, "item %s has price %s\n", item, value)
}

// curl localhost:8080/delete?item=socks -i
func (db database) drop(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %v", item)
		http.Error(w, msg, http.StatusNotFound) // 400
		return
	}
	// Just a no-op if doesn't exist.
	delete(db, item)
	fmt.Fprintf(w, "dropped %s\n", item)
}

func main() {
	// A simple map
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	// These are all method values (closing over the object "db").
	// Therefore only need to pass the route and method value.
	// Every method value takes the same parameters:
	// http.ResponseWriter, *http.Request
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.drop)
	http.HandleFunc("/read", db.fetch)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
