package linkedhashmap

import (
	"github.com/qianxi0410/gogds/containers"
	"github.com/qianxi0410/gogds/lists/doublelinkedlist"
)

// nolint
func assertIteratorImplementation[K comparable, V any]() {
	var _ containers.IteratorWithKey[K, V] = &Iterator[K, V]{}

	var _ containers.ReverseIteratorWithKey[K, V] = &Iterator[K, V]{}
}

// Iterator is an iterator for the linked hash map.
type Iterator[K comparable, V any] struct {
	table    map[K]V
	iterator *doublelinkedlist.Iterator[K]
}

// Iterator returns a new iterator for the map.
func (m *Map[K, V]) Iterator() *Iterator[K, V] {
	return &Iterator[K, V]{
		table:    m.table,
		iterator: m.list.Iterator(),
	}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
func (it *Iterator[K, V]) Next() bool {
	return it.iterator.Next()
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
func (it *Iterator[K, V]) Prev() bool {
	return it.iterator.Prev()
}

// Value returns the current element of the iterator.
func (it *Iterator[K, V]) Value() V {
	return it.table[it.iterator.Value()]
}

// Key returns the current key of the iterator.
func (it *Iterator[K, V]) Key() K {
	return it.iterator.Value()
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (it *Iterator[K, V]) Begin() {
	it.iterator.Begin()
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (it *Iterator[K, V]) End() {
	it.iterator.End()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If the iterator was positioned before the first element, it will be moved to the first element.
// If the first element has not been loaded, loading will be performed.
func (it *Iterator[K, V]) First() bool {
	return it.iterator.First()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If the iterator was positioned before the last element, it will be moved to the last element.
// If the last element has not been loaded, loading will be performed.
func (it *Iterator[K, V]) Last() bool {
	return it.iterator.Last()
}
