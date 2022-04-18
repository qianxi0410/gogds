package circularqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := New[int](2)

	assert.True(t, q.Empty())
	assert.False(t, q.Full())

	q.Enqueue(1)
	assert.Equal(t, q.Size(), 1)
	assert.False(t, q.Empty())
	assert.False(t, q.Full())

	q.Enqueue(2)
	assert.Equal(t, q.Size(), 2)
	assert.False(t, q.Empty())
	assert.True(t, q.Full())

	q.Enqueue(3)
	// 1 2 => 2 3
	assert.Equal(t, q.Size(), 2)
	assert.False(t, q.Empty())
	assert.True(t, q.Full())
}

func TestDequeue(t *testing.T) {
	q := New[int](2)

	q.Enqueue(1)
	q.Enqueue(2)

	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, v, 1)
	assert.Equal(t, q.Size(), 1)
	assert.False(t, q.Empty())
	assert.False(t, q.Full())

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, v, 2)
	assert.Equal(t, q.Size(), 0)
	assert.True(t, q.Empty())
	assert.False(t, q.Full())
}

func TestPeek(t *testing.T) {
	q := New[int](2)

	q.Enqueue(1)
	q.Enqueue(2)

	v, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, v, 1)
	assert.Equal(t, q.Size(), 2)
	assert.False(t, q.Empty())
	assert.True(t, q.Full())

	q.Dequeue()

	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, v, 2)
	assert.Equal(t, q.Size(), 1)
	assert.False(t, q.Empty())
	assert.False(t, q.Full())
}
