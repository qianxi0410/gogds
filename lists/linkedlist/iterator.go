package linkedlist

import "github.com/qianxi0410/gogds/containers"

// nolint
func assertIteratorImplementation[T comparable]() {
	var _ containers.IteratorWithIndex[T] = &Iterator[T]{}
}

// Iterator holds the iterator's state
type Iterator[T comparable] struct {
	list  *List[T]
	index int
	// need a pointer to the element
	// point to the cur index element
	cur *node[T]
}

func (l *List[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{list: l, index: -1, cur: nil}
}

// Next returns true if there was a next element in the iteration.
func (it *Iterator[T]) Next() bool {
	it.index++

	if !it.list.checkIdx(it.index) {
		it.cur = nil
		return false
	}

	if it.cur == nil {
		it.cur = it.list.first
	} else {
		it.cur = it.cur.next
	}

	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
// in iterator we don't need to check the index,
// because we know that the index is valid
func (it *Iterator[T]) Value() T {
	return it.cur.val
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (it *Iterator[T]) Index() int {
	return it.index
}

// Begin resets the iterator to its initial state.
func (it *Iterator[T]) Begin() {
	it.cur = nil
	it.index = -1
}

// First moves the iterator to the first element and returns true if there was a first element in the iteration.
func (it *Iterator[T]) First() bool {
	it.Begin()
	return it.Next()
}
