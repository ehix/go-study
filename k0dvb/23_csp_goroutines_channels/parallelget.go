package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// Run each get in a goroutine itself, putting resp in a channel.
// When we pass a channel to a function, we can restrict to the read/write end.
// Where, <- write, -> read
func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{}
		ch <- result{url, err, 0} // won't calc latency
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	// Need a channel to send stuff in, using make.
	chResults := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	for _, url := range list {
		go get(url, chResults)
	}

	// Need to read only len(list) results.
	// Bc channels block, and we have one channel.
	// If there's no data to read, we have to wait for data.
	// If we ranged over the channel, it would read until it closed.
	// We're not closing the channel, so we don't want to read more times than there's data in the channel to read.
	// But, we'll always be getting a result due to the struct.
	// Why don't we close the channel, we can only close it once.
	// Who closes the channel on n go routines?
	// We don't know which one will be last.
	// Simplfy: n go routines, n responses.
	for range list {
		// Read the results out, using <- before the var.
		r := <-chResults
		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}

// $ time go run parallelget.go
// 2023/09/06 19:10:25 https://wsj.com      69ms
// 2023/09/06 19:10:25 https://google.com   244ms
// 2023/09/06 19:10:25 https://nytimes.com  266ms
// 2023/09/06 19:10:26 https://amazon.com   575ms

// real    0m0.936s
// user    0m0.418s
// sys     0m0.149s
