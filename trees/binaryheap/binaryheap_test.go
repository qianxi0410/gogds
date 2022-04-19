package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	heap := New(func(a, b int) bool {
		return a < b
	})

	assert.True(t, heap.Empty())

	heap.Push(3)
	assert.Equal(t, 1, heap.Size())
	v, ok := heap.Peek()
	assert.Equal(t, 3, v)
	assert.True(t, ok)

	heap.Push(2)
	assert.Equal(t, 2, heap.Size())
	v, ok = heap.Peek()
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	heap.Push(1)
	assert.Equal(t, 3, heap.Size())
	v, ok = heap.Peek()
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	heap.Push(0)
	assert.Equal(t, 4, heap.Size())
	v, ok = heap.Peek()
	assert.Equal(t, 0, v)
	assert.True(t, ok)
}

func TestPop(t *testing.T) {
	heap := New(func(a, b int) bool {
		return a < b
	})

	heap.Push(4, 3, 2, 1)

	v, ok := heap.Pop()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	assert.Equal(t, 3, heap.Size())

	v, ok = heap.Pop()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	assert.Equal(t, 2, heap.Size())

	v, ok = heap.Pop()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	assert.Equal(t, 1, heap.Size())

	v, ok = heap.Pop()
	assert.Equal(t, 4, v)
	assert.True(t, ok)
	assert.Equal(t, 0, heap.Size())

	v, ok = heap.Pop()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
	assert.Equal(t, 0, heap.Size())
}
