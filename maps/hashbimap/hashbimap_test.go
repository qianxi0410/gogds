package hashbimap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	m := New[int, int]()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)

	v, ok := m.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = m.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = m.Get(3)
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = m.Get(4)
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = m.Get(5)
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestGetKey(t *testing.T) {
	m := New[int, int]()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)

	v, ok := m.GetKey(1)
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = m.GetKey(2)
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = m.GetKey(3)
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = m.GetKey(4)
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = m.GetKey(5)
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestRemove(t *testing.T) {
	m := New[int, int]()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)

	m.Remove(1)
	v, ok := m.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 0, v)

	m.Remove(2)
	v, ok = m.Get(2)
	assert.False(t, ok)
	assert.Equal(t, 0, v)

	m.Remove(3)
	v, ok = m.Get(3)
	assert.False(t, ok)
	assert.Equal(t, 0, v)

	m.Remove(4)
	v, ok = m.Get(4)
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}
