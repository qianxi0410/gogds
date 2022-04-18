package circularqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	q := New[int](5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	it := q.Iterator()
	it.Next()
	assert.Equal(t, 1, it.Value())
	assert.Equal(t, 0, it.Index())
	it.Next()
	assert.Equal(t, 2, it.Value())
	assert.Equal(t, 1, it.Index())
	it.Next()
	assert.Equal(t, 3, it.Value())
	assert.Equal(t, 2, it.Index())
	it.Next()
	assert.Equal(t, 4, it.Value())
	assert.Equal(t, 3, it.Index())
	it.Next()
	assert.Equal(t, 5, it.Value())
	assert.Equal(t, 4, it.Index())

	assert.True(t, q.full)

	assert.False(t, it.Next())
	assert.False(t, it.Next())

	q.Enqueue(6)
	assert.Equal(t, 5, q.size)
	assert.Equal(t, 1, q.end)
	assert.Equal(t, 6, q.values[0])
}
