package arraystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(1)

	assert.Equal(t, 1, s.Size())

	s.Push(2)
	assert.Equal(t, 2, s.Size())

}

func TestPop(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, ok := s.Pop()
	assert.Equal(t, 3, v)
	assert.True(t, ok)

	v, ok = s.Pop()
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	v, ok = s.Pop()
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = s.Pop()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestPeek(t *testing.T) {
	s := New[int]()
	s.Push(1)

	v, ok := s.Peek()
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	s.Pop()
	v, ok = s.Peek()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}
