package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	var r result

	start := time.Now()
	ticker := time.NewTicker(1 * time.Second).C // just for show
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}

	for {
		select {
		case ch <- r:
			return
		// no goroutines will tick, bc they will all write to a buffered channel
		case <-ticker:
			log.Panicln("tick")
		}
	}
}

func first(ctx context.Context, urls []string) (*result, error) {
	// as long as there's space in the buffer, a sender with send
	results := make(chan result, len(urls)) // buffer to avoid leaking

	// add cancel to whatever context given, ctx is a decendant!
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	for _, url := range urls { // other go routines will get stuck here
		go get(ctx, url, results)
	}

	select {
	case r := <-results: // read one item out of buffered channel
		return &r, nil // return will cause defered cancel, cancelling all requests.
	case <-ctx.Done():
		// Context is passed into first, therefore its our responsibility
		// to handle case that the context is cancelled from above
		return nil, ctx.Err()
	}

}

func main() {
	// results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	// Return the first response, e.g. imagine calling several microservices,
	// whichever responds first, wins.
	r, _ := first(context.Background(), list)

	if r.err != nil {
		log.Printf("%-20s %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Second)
	log.Println("quit anyway...", runtime.NumGoroutine(), "still running")
}
