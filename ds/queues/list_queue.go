package queues

import (
	"container/list"
	"fmt"
)

// ListQueue is a basic FIFO queue based on a linked list.
// Simple wrapper for container/list, which is not exported.
type ListQueue struct {
	list *list.List
}

// Initalises and returns a new ListQueue.
func NewListQueue() *ListQueue {
	return &ListQueue{
		list: list.New(),
	}
}

func (q ListQueue) Enqueue(e interface{}) {
	q.list.PushBack(e)
}

func (q ListQueue) Dequeue() (interface{}, error) {
	if q.list.Len() == 0 {
		return nil, fmt.Errorf("dequeue from an empty queue")
	}

	e := q.list.Front() // First element
	q.list.Remove(e)
	return e.Value, nil
}
