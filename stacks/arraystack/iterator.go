package arraystack

import "github.com/qianxi0410/gogds/containers"

func assertIteratorImplementation[T comparable]() {
	var _ containers.IteratorWithIndex[T] = &Iterator[T]{}

	var _ containers.ReverseIteratorWithIndex[T] = &Iterator[T]{}
}

// Iterator represents an iterator of array stack.
type Iterator[T comparable] struct {
	stack *Stack[T]
	index int
}

// Iterator returns an iterator of the array stack.
func (s *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{stack: s, index: -1}
}

// Next returns true if there are more elements in the iterator.
func (it *Iterator[T]) Next() bool {
	it.index++

	if !it.stack.checkIdx(it.index) {
		it.index = it.stack.Size()
		return false
	}

	return true
}

// Prev returns true if there are more elements in the iterator.
func (it *Iterator[T]) Prev() bool {
	it.index--

	if !it.stack.checkIdx(it.index) {
		it.index = -1
		return false
	}

	return true
}

// Value returns the current element of the iterator.
// Does not modify the state of the iterator.
func (it *Iterator[T]) Value() T {
	v, _ := it.stack.list.Get(it.stack.Size() - 1 - it.index)
	return v
}

// Index returns the current index of the iterator.
// Does not modify the state of the iterator.
func (it *Iterator[T]) Index() int {
	return it.index
}

// Begin resets the iterator to its initial state.
func (it *Iterator[T]) Begin() {
	it.index = -1
}

// End resets the iterator to its final state.
func (it *Iterator[T]) End() {
	it.index = it.stack.Size()
}

// First moves the iterator to its first element.
func (it *Iterator[T]) First() bool {
	it.Begin()
	return it.Next()
}

// Last moves the iterator to its last element.
func (it *Iterator[T]) Last() bool {
	it.End()
	return it.Prev()
}
