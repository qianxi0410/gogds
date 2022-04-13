package linkedhashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	m := New[int, int]()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)

	it := m.Iterator()

	for it.Next() {
		assert.Equal(t, it.Value(), it.Key())
	}

	assert.True(t, it.Last())

	it.End()
	for it.Prev() {
		assert.Equal(t, it.Value(), it.Key())
	}

	assert.True(t, it.First())
	assert.True(t, it.Last())

	m.Clear()
	it = m.Iterator()

	assert.False(t, it.First())
	assert.False(t, it.Last())
}
