package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Going to start 1000 goroutines
// which will try and increment n.
func do() int {
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		// the critical section:
		// r, m, w of a shared variable.
		go func() {
			n++ // DATA RACE
			w.Done()
		}()
	}

	w.Wait()
	// In theory, this should print 1000, but won't.
	return int(n)
}

// Replica of above with protected critical section
func doCS() int {
	// make a counting semaphore of 1
	m := make(chan bool, 1)
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		// r, m, w of a shared variable.
		go func() {
			// push onto channel, so noone else uses it
			m <- true
			n++ // DATA RACE
			<-m
			w.Done()
		}()
	}

	w.Wait()
	// In theory, this should print 1000, but won't.
	return int(n)
}

// replica with mutex
// more efficent and more straightforward
func doMutex() int {
	// make a counting semaphore of 1
	var m sync.Mutex
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		// r, m, w of a shared variable.
		go func() {
			// push onto channel, so noone else uses it
			m.Lock()
			n++ // DATA RACE
			m.Unlock()
			w.Done()
		}()
	}

	w.Wait()
	// In theory, this should print 1000, but won't.
	return int(n)
}

func doAtomic() int {
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		// Go and do an atomic inc at the hardware level.
		go func() {
			atomic.AddInt64(&n, 1)
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

func main() {
	fmt.Println("without anything", do())
	fmt.Println("with counting semaphore", doCS())
	fmt.Println("with mutex", doMutex())
	fmt.Println("with atomic", doAtomic())
}
