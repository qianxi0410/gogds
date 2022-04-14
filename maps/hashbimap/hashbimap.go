package hashbimap

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/maps"
	"github.com/qianxi0410/gogds/maps/hashmap"
)

// nolint
func assertMapImplementation[K comparable, V comparable]() {
	var _ maps.Map[K, V] = New[K, V]()
}

// Map is the structure that implements the map.
// key map to value and value map to key.
type Map[K comparable, V comparable] struct {
	kv *hashmap.Map[K, V]
	vk *hashmap.Map[V, K]
}

// New returns a new hashbimap.Map.
func New[K comparable, V comparable]() *Map[K, V] {
	return &Map[K, V]{
		kv: hashmap.New[K, V](),
		vk: hashmap.New[V, K](),
	}
}

// Put puts the key-value pair into the map.
func (m *Map[K, V]) Put(key K, value V) {
	if v, ok := m.kv.Get(key); ok {
		m.vk.Remove(v)
	}
	if v, ok := m.vk.Get(value); ok {
		m.kv.Remove(v)
	}

	m.kv.Put(key, value)
	m.vk.Put(value, key)
}

// Get returns the value of the key.
func (m *Map[K, V]) Get(k K) (V, bool) {
	return m.kv.Get(k)
}

// GetKey returns the key of the value.
func (m *Map[K, V]) GetKey(v V) (K, bool) {
	return m.vk.Get(v)
}

// Remove removes the key-value pair from the map.
func (m *Map[K, V]) Remove(k K) {
	if v, ok := m.kv.Get(k); ok {
		m.kv.Remove(k)
		m.vk.Remove(v)
	}
}

// Empty returns true if the map is empty.
func (m *Map[K, V]) Empty() bool {
	return m.kv.Empty()
}

// Size returns the size of the map.
func (m *Map[K, V]) Size() int {
	return m.kv.Size()
}

// Keys returns all keys.
func (m *Map[K, V]) Keys() []K {
	return m.kv.Keys()
}

// Values returns all values.
func (m *Map[K, V]) Values() []V {
	return m.vk.Keys()
}

// Clear removes all key-value pairs from the map.
func (m *Map[K, V]) Clear() {
	m.kv.Clear()
	m.vk.Clear()
}

// String returns a string representation of the map.
func (m *Map[K, V]) String() string {
	strbuf := new(strings.Builder)
	strbuf.WriteString("HashBiMap [")

	for _, key := range m.kv.Keys() {
		v, _ := m.kv.Get(key)
		strbuf.WriteString(fmt.Sprintf("(%v: %v) ", key, v))
	}

	strbuf.WriteString("]\n")
	return strbuf.String()
}
