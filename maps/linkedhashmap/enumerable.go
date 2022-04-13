package linkedhashmap

import (
	"github.com/qianxi0410/gogds/containers"
)

// nolint
func assertEnumerableImplementation[K comparable, V any]() {
	var _ EnumerableWithKey[K, V] = New[K, V]()
}

type EnumerableWithKey[K comparable, V any] interface {
	Filter(f func(k K, v V) bool) *Map[K, V]

	containers.EnumerableWithKey[K, V]
}

// Each calls the given function for each element.
func (m *Map[K, V]) Each(f func(k K, v V)) {
	it := m.Iterator()
	for it.Next() {
		f(it.Key(), it.Value())
	}
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (m *Map[K, V]) Any(f func(k K, v V) bool) bool {
	it := m.Iterator()
	for it.Next() {
		if f(it.Key(), it.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (m *Map[K, V]) All(f func(k K, v V) bool) bool {
	it := m.Iterator()
	for it.Next() {
		if !f(it.Key(), it.Value()) {
			return false
		}
	}
	return true
}

// Find returns the first value that the given predicate returns true for.
// If none matches, then ok is false.
func (m *Map[K, V]) Find(f func(k K, v V) bool) (_ K, _ V) {
	it := m.Iterator()
	for it.Next() {
		if f(it.Key(), it.Value()) {
			return it.Key(), it.Value()
		}
	}
	return
}

// Filter returns a new container with all elements that the given predicate returns true for.
func (m *Map[K, V]) Filter(f func(k K, v V) bool) *Map[K, V] {
	newMap := New[K, V]()
	it := m.Iterator()
	for it.Next() {
		if f(it.Key(), it.Value()) {
			newMap.Put(it.Key(), it.Value())
		}
	}
	return newMap
}

// Map returns a new container with the results of applying the given function to each element in this container.
func Map_[K comparable, V1 any, V2 any](m *Map[K, V1], f func(k K, v V1) V2) *Map[K, V2] {
	newMap := New[K, V2]()
	it := m.Iterator()
	for it.Next() {
		newMap.Put(it.Key(), f(it.Key(), it.Value()))
	}
	return newMap
}
