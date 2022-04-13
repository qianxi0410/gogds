package doublelinkedlist

import "github.com/qianxi0410/gogds/containers"

// nolint
func assertIteratorImplementation[T comparable]() {
	var _ containers.IteratorWithIndex[T] = &Iterator[T]{}

	var _ containers.ReverseIteratorWithIndex[T] = &Iterator[T]{}
}

// Iterator represents a forward iterator of a DoubleLinkedList.
type Iterator[T comparable] struct {
	list  *List[T]
	index int
	cur   *node[T]
}

func (l *List[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{list: l, index: -1}
}

// Next returns the next element of the iterator.
func (it *Iterator[T]) Next() bool {
	it.index++

	if !it.list.checkIdx(it.index) {
		it.cur = nil
		it.index = it.list.size
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
func (it *Iterator[T]) Value() T {
	return it.cur.val
}

// Index returns the current element's index.
func (it *Iterator[T]) Index() int {
	return it.index
}

// Begin resets the iterator to its initial state.
func (it *Iterator[T]) Begin() {
	it.index = -1
	it.cur = nil
}

// First moves the iterator to the first element and returns true if there was a first element in the iteration.
func (it *Iterator[T]) First() bool {
	it.Begin()
	return it.Next()
}

// Prev returns the previous element of the iterator.
func (it *Iterator[T]) Prev() bool {
	it.index--

	if !it.list.checkIdx(it.index) {
		it.index = -1
		it.cur = nil
		return false
	}

	if it.cur == nil {
		it.cur = it.list.last
	} else {
		it.cur = it.cur.prev
	}

	return true
}

// End moves the iterator past the last element in the list.
func (it *Iterator[T]) End() {
	it.cur = nil
	it.index = it.list.size
}

// Last moves the iterator to the last element and returns true if there was a last element in the iteration.
func (it *Iterator[T]) Last() bool {
	it.End()
	return it.Prev()
}
