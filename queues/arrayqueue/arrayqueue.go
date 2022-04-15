package arrayqueue

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/lists/arraylist"
	"github.com/qianxi0410/gogds/queues"
)

// nolint
func assertListImplementation[T comparable]() {
	var _ queues.Queues[T] = New[T]()
}

// Queue is the structure holding a arraylist.
type Queue[T comparable] struct {
	list *arraylist.List[T]
}

// New creates a new queue.
func New[T comparable](values ...T) *Queue[T] {
	return &Queue[T]{list: arraylist.New(values...)}
}

// Enqueue inserts a new element at the end of the queue.
func (q *Queue[T]) Enqueue(e T) {
	q.list.Add(e)
}

// Dequeue removes the element at the front of the queue.
func (q *Queue[T]) Dequeue() (T, bool) {
	return q.list.Remove(0)
}

// Peek returns the element at the front of the queue without removing it.
func (q *Queue[T]) Peek() (T, bool) {
	return q.list.Get(0)
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return q.list.Empty()
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return q.list.Size()
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.list.Clear()
}

// Values returns all elements in the queue.
func (q *Queue[T]) Values() []T {
	return q.list.Values()
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	values := make([]string, 0, q.Size())
	for _, value := range q.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str := fmt.Sprintf("ArrayQueue: [%s]\n", strings.Join(values, ", "))
	return str
}

// checkIdx checks if the index is valid.
func (q *Queue[T]) checkIdx(idx int) bool {
	return idx >= 0 && idx < q.Size()
}
