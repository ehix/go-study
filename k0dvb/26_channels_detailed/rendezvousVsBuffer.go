package main

import (
	"fmt"
	"time"
)

// <!> Will create a race condition

type T struct {
	i byte
	b bool
}

// Take channel and create T, put pointer in channel.
// After it's sent, it's modified.
func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t

	// If rendezvous, it's too late, the send value will have been read,
	// and t will have been garbage collected.
	// If buffered, there will be time for the pointer to be used to
	// modify the value before it's read.
	t.b = true // UNSAFE AT ANY SPEED, don't ever do this.

}

func main() {
	vs := make([]T, 5)    // slice with capacity, means no append, no waiting for memory allocation
	ch := make(chan *T)   // unbuffered channel, rendezvous behaviour
	ch = make(chan *T, 5) // buffer

	for i := range vs {
		go send(i, ch) // start goroutines
	}

	time.Sleep(1 * time.Second) // all goroutines will have started

	// copy quickly!
	for i := range vs {
		// read channel, deference pointer, and copy it into vs
		vs[i] = *<-ch
	}

	// print later, after all i/o is finished above
	for _, v := range vs {
		fmt.Println(v)
	}

	// With unbuffered:
	// {4 false}
	// {0 false}
	// {1 false}
	// {2 false}
	// {3 false}
	// With buffer:
	// {4 true}
	// {0 true}
	// {1 true}
	// {2 true}
	// {3 true}

}
