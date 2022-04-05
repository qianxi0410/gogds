package arraylist

import (
	"time"

	"github.com/qianxi0410/gogds/containers"
)

func assertIteratorImplementation() {
	var _ containers.IteratorWithIndex[int] = &Iterator[int]{}
	var _ containers.IteratorWithIndex[string] = &Iterator[string]{}
	var _ containers.IteratorWithIndex[float32] = &Iterator[float32]{}
	var _ containers.IteratorWithIndex[complex128] = &Iterator[complex128]{}
	var _ containers.IteratorWithIndex[time.Time] = &Iterator[time.Time]{}
	var _ containers.IteratorWithIndex[bool] = &Iterator[bool]{}
	var _ containers.IteratorWithIndex[struct{}] = &Iterator[struct{}]{}
}

// Iterator holds the iterator's state
type Iterator[T comparable] struct {
	list *List[T]
	cur  int
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (l *List[T]) Iterator() Iterator[T] {
	return Iterator[T]{list: l, cur: -1}
}

// Next returns true if there was a next element in the iteration.
func (it *Iterator[T]) Next() bool {
	if it.cur < it.list.size {
		it.cur++
	}
	return it.list.checkIdx(it.cur)
}

// Prev returns true if there was a previous element in the iteration.
func (it *Iterator[T]) Prev() bool {
	if it.cur >= 0 {
		it.cur--
	}

	return it.list.checkIdx(it.cur)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
// in iterator we don't need to check the index,
// because we know that the index is valid
func (it *Iterator[T]) Value() T {
	return it.list.elements[it.cur]
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (it *Iterator[T]) Index() int {
	return it.cur
}

// Begin resets the iterator to its initial state
// so that Next() will return the first element from the list.
func (it *Iterator[T]) Begin() {
	it.cur = -1
}

// End moves the iterator past the last element in the list
func (it *Iterator[T]) End() {
	it.cur = it.list.size
}

// First moves the iterator to the first element and returns true if there was a first element in the list.
// If the list is empty, First() returns false.
// Modifies the state of the iterator.
func (it *Iterator[T]) First() bool {
	it.Begin()
	return it.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the list.
// If the list is empty, Last() returns false.
// Modifies the state of the iterator.
func (it *Iterator[T]) Last() bool {
	it.End()
	return it.Prev()
}
