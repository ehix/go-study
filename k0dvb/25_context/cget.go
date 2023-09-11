package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// Does http request and returns result on a channel.
func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()

	// can make the req and add context later, or ...
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	// instead of .Get(), use DefaultClient.Do() with everything passed in.
	// The http.Get now has a 3 second timeout injected inside.
	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	// context.Background() serves as the parent, as there isn't one.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel() // first thing to do.

	for _, url := range list {
		go get(ctx, url, results)
	}

	for range list {
		r := <-results
		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
