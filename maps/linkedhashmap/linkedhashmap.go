package linkedhashmap

import (
	"bytes"
	"fmt"

	"github.com/qianxi0410/gogds/lists/doublelinkedlist"
	"github.com/qianxi0410/gogds/maps"
)

// nolint
func assertMapImplementation[K comparable, V any]() {
	var _ maps.Map[K, V] = New[K, V]()
}

// Map is a hash table that uses a linked list for collision resolution.
// LinkedHashMap keeps a doubly-linked list of all entries in the map.
type Map[K comparable, V any] struct {
	table map[K]V
	list  *doublelinkedlist.List[K]
}

// New returns a new hash table.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		table: make(map[K]V),
		list:  doublelinkedlist.New[K](),
	}
}

// Put adds a new key-value pair to the map.
func (m *Map[K, V]) Put(key K, value V) {
	if _, ok := m.table[key]; !ok {
		m.list.Add(key)
	}
	m.table[key] = value
}

// Get returns the value for the given key.
func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.table[key]
	return
}

// Remove removes the key-value pair with the given key.
func (m *Map[K, V]) Remove(key K) {
	if _, ok := m.table[key]; ok {
		delete(m.table, key)
		index := m.list.IndexOf(key)
		m.list.Remove(index)
	}
}

// Empty returns true if the map is empty.
func (m *Map[K, V]) Empty() bool {
	return len(m.table) == 0
}

// Keys returns all keys in the map.
func (m *Map[K, V]) Keys() []K {
	return m.list.Values()
}

// Size returns the number of elements in the map.
func (m *Map[K, V]) Size() int {
	return len(m.table)
}

// Values returns all values in the map.
func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, m.Size())

	it := m.Iterator()
	for it.Next() {
		values = append(values, it.Value())
	}

	return values
}

// Clear removes all key-value pairs from the map.
func (m *Map[K, V]) Clear() {
	m.table = make(map[K]V)
	m.list.Clear()
}

// String returns a string representation of the map.
func (m *Map[K, V]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("LinkedHashMap: [")
	it := m.Iterator()
	for it.Next() {
		buffer.WriteString(fmt.Sprintf("(%v: %v) ", it.Key(), it.Value()))
	}
	buffer.WriteString("]\n")
	return buffer.String()
}
