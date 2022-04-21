package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := New(func(a, b int) bool {
		return a > b
	})

	q.Enqueue(1)
	assert.Equal(t, 1, q.Size())

	q.Enqueue(2)
	assert.Equal(t, 2, q.Size())
}

func TestDequeue(t *testing.T) {
	q := New(func(a, b int) bool {
		return a > b
	})

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, ok := q.Dequeue()
	assert.Equal(t, v, 3)
	assert.True(t, ok)

	v, ok = q.Dequeue()
	assert.Equal(t, v, 2)
	assert.True(t, ok)

	v, ok = q.Dequeue()
	assert.Equal(t, v, 1)
	assert.True(t, ok)

	v, ok = q.Dequeue()
	assert.Equal(t, v, 0)
	assert.False(t, ok)
}

func TestPeek(t *testing.T) {
	q := New(func(a, b int) bool {
		return a > b
	})

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, ok := q.Peek()
	assert.Equal(t, v, 3)
	assert.True(t, ok)

	v, ok = q.Peek()
	assert.Equal(t, v, 3)
	assert.True(t, ok)

	v, ok = q.Dequeue()
	assert.Equal(t, v, 3)
	assert.True(t, ok)
}
