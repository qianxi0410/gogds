package binaryheap

import "github.com/qianxi0410/gogds/containers"

// nolint
func assertIteratorImplementation[T comparable]() {
	var _ containers.IteratorWithIndex[T] = &Iterator[T]{}

	var _ containers.ReverseIteratorWithIndex[T] = &Iterator[T]{}
}

// Iterator is an iterator for BinaryHeap.
type Iterator[T comparable] struct {
	tree  *Heap[T]
	index int
}

// Iterator returns a new iterator for BinaryHeap.
func (h *Heap[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{h, -1}
}

// Next returns the true if the iterator has a next item.
// and move the iterator to the next item.
func (it *Iterator[T]) Next() bool {
	if it.index < it.tree.Size() {
		it.index++
	}

	return it.tree.checkIdx(it.index)
}

// Prev returns the true if the iterator has a previous item.
// and move the iterator to the previous item.
func (it *Iterator[T]) Prev() bool {
	if it.index >= 0 {
		it.index--
	}

	return it.tree.checkIdx(it.index)
}

// Value returns the index item of the iterator.
// don't modify the iterator's state.
func (it *Iterator[T]) Value() T {
	start, end := evaluateRange(it.index)
	if end > it.tree.Size() {
		end = it.tree.Size()
	}
	tmpHeap := New(it.tree.Cmp)
	for n := start; n < end; n++ {
		value, _ := it.tree.list.Get(n)
		tmpHeap.Push(value)
	}
	for n := 0; n < it.index-start; n++ {
		tmpHeap.Pop()
	}
	value, _ := tmpHeap.Pop()
	return value
}

// numOfBits counts the number of bits of an int
func numOfBits(n int) uint {
	var count uint
	for n != 0 {
		count++
		n >>= 1
	}
	return count
}

// evaluateRange evaluates the index range [start,end)
// of same level nodes in the heap as the index.
func evaluateRange(index int) (start int, end int) {
	bits := numOfBits(index+1) - 1
	start = 1<<bits - 1
	end = start + 1<<bits
	return
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (it *Iterator[T]) Index() int {
	return it.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (it *Iterator[T]) Begin() {
	it.index = -1
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (it *Iterator[T]) End() {
	it.index = it.tree.Size()
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
