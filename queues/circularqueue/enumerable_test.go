package circularqueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEach(t *testing.T) {
	l := New[int](5)
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.Enqueue(5)

	l.Each(func(idx, value int) {
		assert.Equal(t, idx+1, value)
	})
}

func TestFiler(t *testing.T) {
	l := New[int](5)
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.Enqueue(5)

	l.Filter(func(index, value int) bool {
		return value%2 == 0
	}).Each(func(_, value int) {
		assert.True(t, value%2 == 0)
	})
}

func TestAny(t *testing.T) {
	l := New[int](5)
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.Enqueue(5)

	assert.True(t, l.Any(func(index, value int) bool {
		return value%2 == 0
	}))

	assert.False(t, l.Any(func(index, value int) bool {
		return value > 100
	}))
}

func TestAll(t *testing.T) {
	l := New[int](5)
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.Enqueue(5)

	assert.False(t, l.All(func(index, value int) bool {
		return value%2 == 0
	}))

	assert.True(t, l.All(func(index, value int) bool {
		return value <= 10
	}))
}

func TestFind(t *testing.T) {
	l := New[int](5)
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.Enqueue(5)

	idx, v := l.Find(func(_, v int) bool {
		return v%2 == 0
	})

	assert.Equal(t, 1, idx)
	assert.Equal(t, 2, v)

	idx, v = l.Find(func(_, v int) bool {
		return v > 100
	})
	assert.Equal(t, -1, idx)
	assert.Equal(t, 0, v)
}

func TestMap(t *testing.T) {
	l := New[int](3)
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)

	l2 := Map(l, func(index, value int) string {
		return fmt.Sprint(value)
	})

	assert.Equal(t, "1", l2.Values()[0])
	assert.Equal(t, "2", l2.Values()[1])
	assert.Equal(t, "3", l2.Values()[2])
}
