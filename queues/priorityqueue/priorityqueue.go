package priorityqueue

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/queues"
	"github.com/qianxi0410/gogds/trees/binaryheap"
	"github.com/qianxi0410/gogds/utils"
)

// nolint
func assertListImplementation[T comparable]() {
	var _ queues.Queues[T] = New[T](nil)
}

// Queue is the priority queue
type Queue[T comparable] struct {
	heap *binaryheap.Heap[T]
	Cmp  utils.Comparator[T]
}

// New returns a new priority queue
func New[T comparable](cmp utils.Comparator[T]) *Queue[T] {
	return &Queue[T]{
		heap: binaryheap.New(cmp),
		Cmp:  cmp,
	}
}

// Enqueue add an element to the queue.
func (q *Queue[T]) Enqueue(elem T) {
	q.heap.Push(elem)
}

// Dequeue removes the element with the highest priority from the queue.
func (q *Queue[T]) Dequeue() (T, bool) {
	return q.heap.Pop()
}

// Peek returns the element with the highest priority without removing it from the queue.
func (q *Queue[T]) Peek() (T, bool) {
	return q.heap.Peek()
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return q.heap.Empty()
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.heap.Clear()
}

// Values returns all elements in the queue.
func (q *Queue[T]) Values() []T {
	return q.heap.Values()
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return q.heap.Size()
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	values := make([]string, q.Size())

	for i, v := range q.Values() {
		values[i] = fmt.Sprintf("%v", v)
	}

	return fmt.Sprintf("PriorityQueue: [%s]\n", strings.Join(values, ", "))
}
