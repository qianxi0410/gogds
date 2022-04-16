package linkedqueue

import "github.com/qianxi0410/gogds/containers"

// nolint
func assertIteratorImplementation[T comparable]() {
	var _ containers.IteratorWithIndex[T] = &Iterator[T]{}

	var _ containers.ReverseIteratorWithIndex[T] = &Iterator[T]{}
}

// Iterator is the structure used to iterate over a queue.
type Iterator[T comparable] struct {
	q     *Queue[T]
	index int
}

// Iterator returns a new iterator for the given queue.
func (q *Queue[T]) Iterator() Iterator[T] {
	return Iterator[T]{
		q:     q,
		index: -1,
	}
}

// Next returns true if there are more elements to iterate over.
func (it *Iterator[T]) Next() bool {
	if it.index < it.q.Size() {
		it.index++
	}

	return it.q.checkIdx(it.index)
}

// Prev returns true if there are more elements to iterate over.
func (it *Iterator[T]) Prev() bool {
	if it.index >= 0 {
		it.index--
	}
	return it.q.checkIdx(it.index)
}

// Value returns the current element.
func (it *Iterator[T]) Value() T {
	v, _ := it.q.list.Get(it.index)
	return v
}

// Index returns the current index.
func (it *Iterator[T]) Index() int {
	return it.index
}

// Begin resets the iterator to its initial state.
func (it *Iterator[T]) Begin() {
	it.index = -1
}

// End moves the iterator past the last element of the queue.
func (it *Iterator[T]) End() {
	it.index = it.q.Size()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (it *Iterator[T]) First() bool {
	it.Begin()
	return it.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (it *Iterator[T]) Last() bool {
	it.End()
	return it.Prev()
}
