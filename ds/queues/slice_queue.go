package queues

// Basic FIFO queue using slices.

// Allegedly, using slices is considered less memory efficient,
// as the memory allocated to the slice is not freed.
// https://go.dev/ref/spec#Slice_types

// Furthermore:
// It is recommended that whenever an element is popped or removed from a queue,
// always zero it (the element in the slice) so the value will not remain in memory needlessly.
// This becomes even more critical if your slice contains pointers to big data structures.
// https://stackoverflow.com/questions/28432658/does-go-garbage-collect-parts-of-slices

import (
	"fmt"
)

// SliceQueue is a basic FIFO queue based on a slice that resizes as needed.
// It will accept any type as an element.
type SliceQueue []interface{}

// Enqueue adds an element to the end of the queue.
func (q *SliceQueue) Enqueue(e interface{}) {
	*q = append(*q, e)
}

// Dequeue removes the first element from the queue and returns it.
func (q *SliceQueue) Dequeue() (interface{}, error) {
	if len(*q) == 0 {
		return nil, fmt.Errorf("Dequeue from an empty queue")
	}

	e := (*q)[0]    // Dequeue the first element.
	(*q)[0] = nil   // Overwrite with the zero value.
	(*q) = (*q)[1:] // Slice off the dequeued element from the queue.
	return e, nil
}
