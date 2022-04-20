package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	heap := New(func(a, b int) bool {
		return a <= b
	})

	heap.Push(3, 2, 1)

	it := heap.Iterator()

	it.Next()
	assert.Equal(t, 1, it.Value())
	it.Next()
	assert.Equal(t, 2, it.Value())
	it.Next()
	assert.Equal(t, 3, it.Value())
	it.Next()
	assert.False(t, it.Next())
}
