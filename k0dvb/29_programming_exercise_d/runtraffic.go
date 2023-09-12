package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Generates traffic to the server.

type sku struct {
	item, price string
}

var items = []sku{
	{"shoes", "46"},
	{"socks", "6"},
	{"sandals", "27"},
	{"clogs", "36"},
	{"pants", "30"},
	{"shorts", "20"},
}

func doQuery(cmd, params string) error {
	url := "http://localhost:8080/" + cmd + "?" + params
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err %s = %v\n", params, err)
		return err
	}

	defer resp.Body.Close()

	fmt.Fprintf(os.Stderr, "got %s = %d (no err)\n", params, resp.StatusCode)
	return nil
}

func runAdds() {
	for {
		for _, s := range items {
			params := "item=" + s.item + "&price=" + s.price
			if err := doQuery("create", params); err != nil {
				return
			}
		}
	}
}

func runUpdates() {
	for {
		for _, s := range items {
			params := "item=" + s.item + "&price=" + s.price
			if err := doQuery("update", params); err != nil {
				return
			}
		}
	}
}

func runDeletes() {
	for {
		for _, s := range items {
			params := "item=" + s.item
			if err := doQuery("delete", params); err != nil {
				return
			}
		}
	}
}

func main() {
	go runAdds()
	go runUpdates()
	go runDeletes()

	time.Sleep(10 * time.Second)
}

// When run with the server from the previous task:
// go run server.go -race
// fatal error: concurrent map writes
