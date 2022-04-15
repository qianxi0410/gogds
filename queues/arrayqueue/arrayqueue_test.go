package arrayqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 1, q.Values()[0])

	q.Enqueue(2)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 2, q.Values()[1])

	q.Enqueue(3)
	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 3, q.Values()[2])
}

func TestDequeue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
}
