package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	// receive channels
	for i := 0; i < 12; i++ {
		// // Polling would be:
		// m0 := <-chans[0]
		// m1 := <-chans[1]

		// Instead, select:
		// which ever is ready first, read from it.
		select {
		case m0 := <-chans[0]:
			log.Println("received", m0)
		case m1 := <-chans[1]:
			log.Println("received", m1)
		}
	}
}
