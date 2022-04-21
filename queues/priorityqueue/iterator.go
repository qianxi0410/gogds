package priorityqueue

import (
	"github.com/qianxi0410/gogds/containers"
	"github.com/qianxi0410/gogds/trees/binaryheap"
)

// nolint
func assertIteratorImplementation[T comparable]() {
	var _ containers.IteratorWithIndex[T] = &Iterator[T]{}
}

type Iterator[T comparable] struct {
	heap binaryheap.Iterator[T]
}

func (q *Queue[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{
		heap: *q.heap.Iterator(),
	}
}

// Next returns true if there are more elements to iterate over.
func (it *Iterator[T]) Next() bool {
	return it.heap.Next()
}

// Value returns the current element.
func (it *Iterator[T]) Value() T {
	return it.heap.Value()
}

// Index returns the current index.
func (it *Iterator[T]) Index() int {
	return it.heap.Index()
}

// Begin resets the iterator to its initial state.
func (it *Iterator[T]) Begin() {
	it.heap.Begin()
}

// End moves the iterator past the last element of the queue.
func (it *Iterator[T]) End() {
	it.heap.End()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (it *Iterator[T]) First() bool {
	it.Begin()
	return it.Next()
}
