package linkedqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	it := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).Iterator()
	for it.Next() {
		assert.Equal(t, it.Value()-1, it.Index())
	}

	assert.True(t, it.First())
}
