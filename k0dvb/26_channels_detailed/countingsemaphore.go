package main

import (
	"fmt"
	"sync"
)

func main() {
	const maxConcurrency = 3 // Adjust the maximum concurrency as needed
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	tasks := []string{"Task 1", "Task 2", "Task 3", "Task 4", "Task 5"}

	for _, task := range tasks {
		wg.Add(1)
		go func(taskName string) {
			// Acquire semaphore (try to send to the channel)
			sem <- struct{}{}
			defer func() {
				// Release semaphore (read from the channel)
				<-sem
				wg.Done()
			}()

			// Perform some work
			fmt.Printf("%s: working...\n", taskName)
			// Simulate work by sleeping for a while
			// In a real-world scenario, replace this with actual work.
			// For demonstration purposes, we'll sleep for a short time.
			// You can replace this with your actual task.
			// time.Sleep(time.Second)

			fmt.Printf("%s: done!\n", taskName)
		}(task)
	}

	// Wait for all tasks to complete
	wg.Wait()
	fmt.Println("All tasks completed.")
}

// In this example:

// We create a buffered channel sem with a capacity of maxConcurrency,
// which limits the number of concurrent tasks that can access the
// shared resource or perform work at any given time.

// We use a sync.WaitGroup to ensure that the main function waits for all tasks
// to complete before exiting.

// Each task tries to acquire the semaphore
// by sending an empty struct {} to the sem channel.
// If the buffer is full (i.e., the maximum concurrency is reached),
// the task will block until a slot in the buffer becomes available.

// After acquiring the semaphore, the task performs its work
// (in this case, simulated by printing messages) and then
// releases the semaphore by reading from the sem channel using <-sem.

// The defer statement ensures that the semaphore is always released,
// even if the task encounters an error or panics.

// Finally, we wait for all tasks to complete using wg.Wait().

// This code demonstrates how to use a counting semaphore to limit t
// he concurrency of tasks in Go.
// You can adjust the maxConcurrency constant to control the maximum number
// of tasks running concurrently.
