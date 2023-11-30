package queues

import (
	"reflect"
	"testing"
)

func TestNewListQueue(t *testing.T) {
	q := NewListQueue()
	if q.list.Len() != 0 {
		t.Errorf("New queue length expected to be 0, got %d", q.list.Len())
	}

	expectedType := "*queues.ListQueue"
	if reflect.TypeOf(q).String() != expectedType {
		t.Errorf("Expected type %s, got %s", expectedType, reflect.TypeOf(q).String())
	}
}

func TestListEnqueue(t *testing.T) {
	q := NewListQueue()
	q.Enqueue(1)
	if q.list.Len() != 1 {
		t.Errorf("Queue length expected to be 1, got %d", q.list.Len())
	}
}

func TestListDequeue(t *testing.T) {
	q := NewListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if q.list.Len() != 1 {
		t.Errorf("Queue length expected to be 1, got %d", q.list.Len())
	}
}

func TestListDequeueEmpty(t *testing.T) {
	q := NewListQueue()
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
