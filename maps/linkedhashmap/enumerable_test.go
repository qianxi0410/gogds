package linkedhashmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEach(t *testing.T) {
	m := New[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)

	m.Each(func(idx, value int) {
		assert.Equal(t, idx, value)
	})
}

func TestFiler(t *testing.T) {
	m := New[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	m.Put(5, 5)
	m.Put(6, 6)
	m.Put(7, 7)
	m.Put(8, 8)
	m.Put(9, 9)
	m.Put(10, 10)

	m.Filter(func(index, value int) bool {
		return value%2 == 0
	}).Each(func(_, value int) {
		assert.True(t, value%2 == 0)
	})
}

func TestAny(t *testing.T) {
	m := New[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	m.Put(5, 5)
	m.Put(6, 6)
	m.Put(7, 7)
	m.Put(8, 8)
	m.Put(9, 9)
	m.Put(10, 10)

	assert.True(t, m.Any(func(index, value int) bool {
		return value%2 == 0
	}))

	assert.False(t, m.Any(func(index, value int) bool {
		return value > 100
	}))
}

func TestAll(t *testing.T) {
	m := New[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	m.Put(5, 5)
	m.Put(6, 6)
	m.Put(7, 7)
	m.Put(8, 8)
	m.Put(9, 9)
	m.Put(10, 10)

	assert.False(t, m.All(func(index, value int) bool {
		return value%2 == 0
	}))

	assert.True(t, m.All(func(index, value int) bool {
		return value <= 10
	}))
}

func TestFind(t *testing.T) {
	m := New[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	m.Put(5, 5)
	m.Put(6, 6)
	m.Put(7, 7)
	m.Put(8, 8)
	m.Put(9, 9)
	m.Put(10, 10)

	idx, v := m.Find(func(_, v int) bool {
		return v%2 == 0
	})

	assert.Equal(t, 2, idx)
	assert.Equal(t, 2, v)

	idx, v = m.Find(func(_, v int) bool {
		return v > 100
	})
	assert.Equal(t, 0, idx)
	assert.Equal(t, 0, v)
}

func TestMap(t *testing.T) {
	m := New[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	m.Put(5, 5)
	m.Put(6, 6)
	m.Put(7, 7)
	m.Put(8, 8)
	m.Put(9, 9)
	m.Put(10, 10)

	m2 := Map_(m, func(index, value int) string {
		return fmt.Sprint(value)
	})

	assert.Equal(t, "1", m2.Values()[0])
	assert.Equal(t, "2", m2.Values()[1])
	assert.Equal(t, "3", m2.Values()[2])
}
