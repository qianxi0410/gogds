package arraylist

import (
	"time"

	"github.com/qianxi0410/gogds/containers"
)

type EnumerableWithIndex[T comparable] interface {
	// Filter returns a new container containing all elements for which the given function returns a true value.
	Filter(func(index int, value T) bool) *List[T]

	containers.EnumerableWithIndex[T]
}

func assertEnumerableImplementation() {
	var _ containers.EnumerableWithIndex[int] = New[int]()
	var _ containers.EnumerableWithIndex[string] = New[string]()
	var _ containers.EnumerableWithIndex[float32] = New[float32]()
	var _ containers.EnumerableWithIndex[complex128] = New[complex128]()
	var _ containers.EnumerableWithIndex[time.Time] = New[time.Time]()
	var _ containers.EnumerableWithIndex[bool] = New[bool]()
	var _ containers.EnumerableWithIndex[struct{}] = New[struct{}]()

	var _ EnumerableWithIndex[int] = New[int]()
	var _ EnumerableWithIndex[string] = New[string]()
	var _ EnumerableWithIndex[float32] = New[float32]()
	var _ EnumerableWithIndex[complex128] = New[complex128]()
	var _ EnumerableWithIndex[time.Time] = New[time.Time]()
	var _ EnumerableWithIndex[bool] = New[bool]()
	var _ EnumerableWithIndex[struct{}] = New[struct{}]()
}

// Each calls the given function once for each element, passing that element's index and value.
func (l *List[T]) Each(f func(index int, value T)) {
	it := l.Iterator()
	for it.Next() {
		f(it.Index(), it.Value())
	}
}

// Filter returns a new container containing all elements for which the given function returns a true value.
func (l *List[T]) Filter(f func(index int, value T) bool) *List[T] {
	// pre-allocate some space for the result
	newList := &List[T]{}

	it := l.Iterator()

	for it.Next() {
		if f(it.Index(), it.Value()) {
			newList.Add(it.Value())
		}
	}

	return newList
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (l *List[T]) Any(f func(index int, value T) bool) bool {
	it := l.Iterator()

	for it.Next() {
		if f(it.Index(), it.Value()) {
			return true
		}
	}

	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (l *List[T]) All(f func(index int, value T) bool) bool {
	it := l.Iterator()

	for it.Next() {
		if !f(it.Index(), it.Value()) {
			return false
		}
	}

	return true
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (l *List[T]) Find(f func(index int, value T) bool) (_ int, t T) {
	it := l.Iterator()

	for it.Next() {
		if f(it.Index(), it.Value()) {
			return it.Index(), it.Value()
		}
	}

	return -1, t
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func Map[K comparable, V comparable](l *List[K], f func(index int, value K) V) *List[V] {
	newList := &List[V]{}

	it := l.Iterator()
	for it.Next() {
		newList.Add(f(it.Index(), it.Value()))
	}

	return newList
}
