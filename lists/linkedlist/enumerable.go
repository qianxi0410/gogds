package linkedlist

import "github.com/qianxi0410/gogds/containers"

func assertEnumerableImplementation[T comparable]() {
	var _ containers.EnumerableWithIndex[T] = &List[T]{}

	var _ EnumerableWithIndex[T] = &List[T]{}
}

// Enumerable interface
type EnumerableWithIndex[T comparable] interface {
	// Filter returns a new container containing all elements for which the given function returns a true value.
	Filter(func(index int, value T) bool) *List[T]

	containers.EnumerableWithIndex[T]
}

// Each calls the given function once for each element, passing that element's index and value.
func (l *List[T]) Each(f func(index int, value T)) {
	for i, e := 0, l.first; e != nil; i, e = i+1, e.next {
		f(i, e.val)
	}
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (l *List[T]) Any(f func(index int, value T) bool) bool {
	for i, e := 0, l.first; e != nil; i, e = i+1, e.next {
		if f(i, e.val) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (l *List[T]) All(f func(index int, value T) bool) bool {
	for i, e := 0, l.first; e != nil; i, e = i+1, e.next {
		if !f(i, e.val) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (l *List[T]) Find(f func(index int, value T) bool) (_ int, t T) {
	for i, e := 0, l.first; e != nil; i, e = i+1, e.next {
		if f(i, e.val) {
			return i, e.val
		}
	}
	return -1, t
}

// Filter returns a new container by retaining only the elements for which the given function returns a true value.
func (l *List[T]) Filter(f func(index int, value T) bool) *List[T] {
	newList := &List[T]{}

	for i, e := 0, l.first; e != nil; i, e = i+1, e.next {
		if f(i, e.val) {
			newList.Append(e.val)
		}
	}

	return newList
}

// Map returns a new container by transforming every element with a function f.
func Map[K comparable, V comparable](l *List[K], f func(index int, value K) V) *List[V] {
	newList := &List[V]{}

	for i, e := 0, l.first; e != nil; i, e = i+1, e.next {
		newList.Add(f(i, e.val))
	}

	return newList
}
