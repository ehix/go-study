package queues

import (
	"reflect"
	"testing"
)

func TestSliceEnqueue(t *testing.T) {
	var q SliceQueue
	q.Enqueue(1)
	if len(q) != 1 {
		t.Errorf("Queue length expected to be 1, got %d", len(q))
	}
}

func TestSliceEnqueueInterfaceTypes(t *testing.T) {
	var q SliceQueue
	var intVal int = 1
	var floatVal float64 = 3.14
	var stringVal string = "Hello"
	var sliceVal []int = []int{1, 2, 3}

	q.Enqueue(intVal)
	q.Enqueue(floatVal)
	q.Enqueue(stringVal)
	q.Enqueue(sliceVal)

	if len(q) != 4 {
		t.Errorf("Queue length expected to be 4, got %d", len(q))
	}

	if val, ok := q[0].(int); !ok || val != intVal {
		t.Errorf("Expected first element to be int 1, got %v", q[0])
	}

	if val, ok := q[1].(float64); !ok || val != floatVal {
		t.Errorf("Expected second element to be float64 1.1, got %v", q[1])
	}

	if val, ok := q[2].(string); !ok || val != stringVal {
		t.Errorf("Expected third element to be string 'a', got %v", q[2])
	}

	if val, ok := q[3].([]int); !ok || !reflect.DeepEqual(val, sliceVal) {
		t.Errorf("Expected fourth element to be []int{1, 2, 3}, got %v", q[3])
	}
}

func TestSliceDequeue(t *testing.T) {
	var q SliceQueue
	q.Enqueue(1)
	q.Enqueue(2)
	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	if len(q) != 1 {
		t.Errorf("Queue length expected to be 1, got %d", len(q))
	}
}

func TestSliceDequeueEmpty(t *testing.T) {
	var q SliceQueue
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
