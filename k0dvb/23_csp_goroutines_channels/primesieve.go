package main

import "fmt"

// Don't run this forever
// Write only channel
func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i
	}

	close(ch)
}

// scr read only channel, dst write only channel.
func filter(src <-chan int, dst chan<- int, prime int) {
	// Can continually get ints from src,
	// untils src closes and the loop breaks.
	for i := range src {
		// If not divisible by the prime, pass it on.
		// i.e. if there's a remainder from modulo.
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst)
}

func sieve(limit int) {
	// ch back to main from generator
	ch := make(chan int)
	go generator(limit, ch)
	// infinite for
	for {
		// can read ok like with a map, is the channel closed?
		// read the channel, if it's a prime number..
		prime, ok := <-ch
		if !ok {
			break
		}
		// ... make a new channel...
		ch1 := make(chan int)
		// ... start a new filter with scr old channel, dest new channel
		go filter(ch, ch1, prime)

		// ... swap view of channel reading from.
		// If it's closed the loop will break
		ch = ch1

		fmt.Print(prime, " ")
	}
}

func main() {
	// Fixed upper limit, the runtime sucks.
	// Good example of channels and goroutines,
	// but would be more performant as a loop of arithmetic.
	sieve(100) // 2, 3, 5, 7, 11, 13, 17, 19 ...

}
