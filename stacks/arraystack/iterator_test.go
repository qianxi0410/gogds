package arraystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	it := s.Iterator()

	// 3 2 1
	for it.Next() {
		assert.Equal(t, it.Value(), 3-it.Index())
	}

	it.End()

	// 1 2 3
	for it.Prev() {
		assert.Equal(t, it.Value(), 3-it.Index())
	}
}
