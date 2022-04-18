package circularqueue

import (
	"fmt"

	"github.com/qianxi0410/gogds/queues"
)

// nolint
func assertListImplementation[T comparable]() {
	var _ queues.Queues[T] = New[T](0)
}

// Queue holds the value of sice.
// Circular queue:
// The queue is implemented as a ring buffer.
// The size of the queue is fixed.
// The queue is full when the end is equal to the start.
// The queue is empty when the end is equal to the start and the queue is not full.
// see: https://en.wikipedia.org/wiki/Circular_buffer
type Queue[T any] struct {
	values []T
	// start points to the first element in the queue.
	start int
	// end points to the last avaliable insert position in the queue.
	end  int
	full bool
	// maxSize is the cap of the queue.
	maxSize int
	// size is the current size of the queue.
	size int
}

// New creates a new queue.
func New[T any](maxSize int) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid max size, must be greater than 0")
	}

	q := &Queue[T]{maxSize: maxSize}
	q.Clear()
	return q
}

// Enqueue adds a new value to the queue.
// if the queue is full, it will remove the first element in the front of the queue.
func (q *Queue[T]) Enqueue(value T) {
	if q.full {
		q.Dequeue()
	}

	q.values[q.end] = value
	q.end = (q.end + 1) % q.maxSize
	q.size++
	q.full = q.size == q.maxSize
}

// Dequeue removes the first value from the queue.
func (q *Queue[T]) Dequeue() (_ T, _ bool) {
	if q.size == 0 {
		return
	}

	value := q.values[q.start]
	q.start = (q.start + 1) % q.maxSize
	q.size--
	q.full = false

	return value, true
}

// Peek returns the first value in the queue without removing it.
func (q *Queue[T]) Peek() (_ T, _ bool) {
	if q.size == 0 {
		return
	}

	return q.values[q.start], true
}

// Size returns the size of the queue.
func (q *Queue[T]) Size() int {
	return q.size
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return q.size == 0
}

// Full returns true if the queue is full.
func (q *Queue[T]) Full() bool {
	return q.full
}

// Clear removes all values from the queue.
func (q *Queue[T]) Clear() {
	q.values = make([]T, q.maxSize, q.maxSize)
	q.start = 0
	q.end = 0
	q.full = false
	q.size = 0
}

// Values returns all values in the queue.
func (q *Queue[T]) Values() []T {
	values := make([]T, 0, q.size)

	for i := range q.values {
		values = append(values, q.values[i])
	}

	return values
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	str := "CircularQueue: ["
	for i := range q.values {
		if i == q.start {
			str += "|"
		}
		str += fmt.Sprintf("%v", q.values[i])
		if i == q.end-1 {
			str += "|"
		}
		str += ", "
	}
	str += "]\n"
	return str
}

// checkIdx checks if the index is valid.
func (q *Queue[T]) checkIdx(idx int) bool {
	if q.size == 0 {
		return false
	}

	if q.end > q.start {
		return idx >= q.start && idx < q.end
	} else {
		return idx >= q.start || idx < q.end
	}
}
