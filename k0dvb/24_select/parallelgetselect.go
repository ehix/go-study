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
	// Listening to two channels simultaneously, one for results another with a timeout.
	// Adding stopper, signaling channel to send timeout message
	stopper := time.After(3 * time.Second)
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

	for range list {
		// Listen on two channels simultaniously:
		select {
		case r := <-chResults:
			if r.err != nil {
				log.Printf("%-20s %s\n", r.url, r.err)
			} else {
				log.Printf("%-20s %s\n", r.url, r.latency)
			}
		// Assign out of stopper to t
		case t := <-stopper:
			// Kills the program and prints timeout message.
			log.Fatal("timeout  %s", t)
		}
	}
}
