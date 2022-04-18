package circularqueue

import "github.com/qianxi0410/gogds/containers"

type EnumerableWithIndex[T comparable] interface {
	// Filter returns a new container containing all elements for which the given function returns a true value.
	Filter(func(index int, value T) bool) *Queue[T]

	containers.EnumerableWithIndex[T]
}

// nolint
func assertEnumerableImplementation[T comparable]() {
	var _ EnumerableWithIndex[T] = New[T](0)
}

// Each calls the given function once for each element, passing that element's index and value.
func (q *Queue[T]) Each(f func(index int, value T)) {
	it := q.Iterator()
	for it.Next() {
		f(it.Index(), it.Value())
	}
}

// Filter returns a new container containing all elements for which the given function returns a true value.
func (q *Queue[T]) Filter(f func(index int, value T) bool) *Queue[T] {
	it := q.Iterator()

	newQueue := &Queue[T]{
		start:   q.start,
		end:     q.start,
		maxSize: q.maxSize,
		values:  make([]T, q.maxSize),
	}

	for it.Next() {
		if f(it.Index(), it.Value()) {
			newQueue.Enqueue(it.Value())
		}
	}
	return newQueue
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (q *Queue[T]) Any(f func(index int, value T) bool) bool {
	it := q.Iterator()

	for it.Next() {
		if f(it.Index(), it.Value()) {
			return true
		}
	}

	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (q *Queue[T]) All(f func(index int, value T) bool) bool {
	it := q.Iterator()

	for it.Next() {
		if !f(it.Index(), it.Value()) {
			return false
		}
	}

	return true
}

// Find returns the first element of the container for which the given function returns a true value.
func (q *Queue[T]) Find(f func(index int, value T) bool) (index int, value T) {
	it := q.Iterator()

	for it.Next() {
		if f(it.Index(), it.Value()) {
			return it.Index(), it.Value()
		}
	}

	return -1, value
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func Map[K any, V any](q *Queue[K], f func(index int, value K) V) *Queue[V] {
	newQueue := &Queue[V]{
		start:   q.start,
		end:     q.start,
		maxSize: q.maxSize,
		values:  make([]V, q.maxSize),
	}

	it := q.Iterator()

	for it.Next() {
		newQueue.Enqueue(f(it.Index(), it.Value()))
	}

	return newQueue
}
