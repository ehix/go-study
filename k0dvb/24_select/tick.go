package main

import (
	"log"
	"time"
)

func main() {
	const tickRate = 2 * time.Second
	stopper := time.After(5 * tickRate)
	// returns a ticker obj, .C is a channel that keeps returning time.
	ticker := time.NewTicker(tickRate).C
	log.Println("start")
loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}
	log.Println("finish")

}
