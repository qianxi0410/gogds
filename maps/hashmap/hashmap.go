package hashmap

import (
	"bytes"
	"fmt"

	"github.com/qianxi0410/gogds/maps"
)

// nolint
func assertMapImplementation[K comparable, V any]() {
	var _ maps.Map[K, V] = New[K, V]()
}

// Map holds a internal map.
type Map[K comparable, V any] struct {
	m map[K]V
}

// New returns a new hashmap.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: make(map[K]V),
	}
}

// Put inserts a new key-value pair into the map.
func (m *Map[K, V]) Put(key K, value V) {
	m.m[key] = value
}

// Get returns the value for the given key.
func (m *Map[K, V]) Get(key K) (V, bool) {
	value, ok := m.m[key]
	return value, ok
}

// Remove removes the key-value pair from the map.
func (m *Map[K, V]) Remove(key K) {
	delete(m.m, key)
}

// Empty returns true if the map is empty.
func (m *Map[K, V]) Empty() bool {
	return len(m.m) == 0
}

// Size returns the number of key-value pairs in the map.
func (m *Map[K, V]) Size() int {
	return len(m.m)
}

// Keys returns all keys in the map.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, len(m.m))
	i := 0
	for key := range m.m {
		keys[i] = key
		i++
	}
	return keys
}

// Values returns all values in the map.
func (m *Map[K, V]) Values() []V {
	values := make([]V, len(m.m))
	i := 0
	for _, value := range m.m {
		values[i] = value
		i++
	}
	return values
}

// Clear removes all key-value pairs from the map.
func (m *Map[K, V]) Clear() {
	m.m = make(map[K]V)
}

// String returns a string representation of the map.
func (m *Map[K, V]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("HashMap: [")
	for key, value := range m.m {
		buffer.WriteString(fmt.Sprintf("(%v: %v) ", key, value))
	}
	buffer.WriteString("]\n")
	return buffer.String()
}
