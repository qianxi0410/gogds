package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	l := New[int]()

	assert.Equal(t, l.first, l.last)
	assert.Nil(t, l.first)
	assert.Nil(t, l.last)

	l.Add(1)
	assert.Equal(t, 1, l.first.val)
	assert.Equal(t, 1, l.last.val)
	assert.Equal(t, l.last, l.first)

	l.Add(2)
	assert.Equal(t, 2, l.last.val)
}

func TestPrepend(t *testing.T) {
	l := New(1)

	l.Prepend(2)
	assert.Equal(t, 2, l.first.val)

	l.Prepend(3, 4)
	assert.Equal(t, 3, l.first.val)
}

func TestGet(t *testing.T) {
	l := New(1, 2, 3)

	v, ok := l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = l.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = l.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = l.Get(10)
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestRemove(t *testing.T) {
	l := New(1, 2, 3, 4, 5)

	v, ok := l.Remove(0)
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = l.Remove(2)
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	l = New(1)
	l.Remove(0)
	assert.Nil(t, l.first)
	assert.Nil(t, l.last)
}

func TestContains(t *testing.T) {
	l := New(1, 2, 3)

	ok := l.Contains(1)
	assert.True(t, ok)

	ok = l.Contains(1, 3)
	assert.True(t, ok)

	ok = l.Contains()
	assert.True(t, ok)

	ok = l.Contains(-1)
	assert.False(t, ok)
}

func TestIndexOf(t *testing.T) {
	l := New("1", "2", "3")

	idx := l.IndexOf("1")
	assert.Equal(t, 0, idx)

	idx = l.IndexOf("3")
	assert.Equal(t, 2, idx)

	idx = l.IndexOf("-1")

	assert.Equal(t, -1, idx)
}

func TestSort(t *testing.T) {
	l := New(-1, 10, 1, 2, 3)

	l.Sort(func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, l.Values()[0], -1)
	assert.Equal(t, l.Values()[l.Size()-1], 10)
}

func TestSwap(t *testing.T) {
	l := New(1, 2, 3, 4)

	l.Swap(0, 3)

	assert.Equal(t, l.Values()[0], 4)
	assert.Equal(t, l.Values()[3], 1)
}

func TestInsert(t *testing.T) {
	l := New[int]()

	l.Insert(0, 1, 2, 3, 4)

	// 1 2 3 4
	assert.Equal(t, l.Values()[0], 1)
	assert.Equal(t, l.Values()[1], 2)
	assert.Equal(t, l.Values()[2], 3)
	assert.Equal(t, l.Values()[3], 4)

	l.Insert(3, 5, 6)
	// 1 2 3 5 6 4
	assert.Equal(t, l.Values()[3], 5)
	assert.Equal(t, l.Values()[4], 6)
	assert.Equal(t, l.Values()[5], 4)

	l.Insert(1, 7, 8)
	// 1 7 8 2 3 5 6 4
	assert.Equal(t, l.Values()[1], 7)
	assert.Equal(t, l.Values()[2], 8)
	assert.Equal(t, l.Values()[3], 2)
}

func TestSet(t *testing.T) {
	l := New(1, 2, 3, 4)

	l.Set(0, 5)
	assert.Equal(t, l.Values()[0], 5)

	l.Set(1, 6)
	assert.Equal(t, l.Values()[1], 6)

	l.Set(2, 7)
	assert.Equal(t, l.Values()[2], 7)

	l.Set(3, 8)
	assert.Equal(t, l.Values()[3], 8)
}
