package linkedqueue

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/lists/linkedlist"
)

type Queue[T comparable] struct {
	list *linkedlist.List[T]
}

// New returns a new Queue.
func New[T comparable](values ...T) *Queue[T] {
	return &Queue[T]{
		list: linkedlist.New(values...),
	}
}

// Enqueue adds a new item to the queue.
func (q *Queue[T]) Enqueue(e T) {
	q.list.Add(e)
}

// Dequeue removes the item from the queue.
func (q *Queue[T]) Dequeue() (T, bool) {
	return q.list.Remove(0)
}

// Peek returns the first item in the queue.
func (q *Queue[T]) Peek() (T, bool) {
	return q.list.Get(0)
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return q.list.Empty()
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return q.list.Size()
}

// Values returns all items in the queue.
func (q *Queue[T]) Values() []T {
	return q.list.Values()
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	values := make([]string, 0, q.Size())
	for _, value := range q.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str := fmt.Sprintf("LinkedQueue: [%s]\n", strings.Join(values, ", "))
	return str
}

// checkIdx checks if the index is valid for the queue.
func (q *Queue[T]) checkIdx(idx int) bool {
	return idx >= 0 && idx < q.Size()
}
