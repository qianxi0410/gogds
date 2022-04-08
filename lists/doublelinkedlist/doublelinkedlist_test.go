package doublelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	l := New[int]()

	l.Add(1, 2, 3)

	assert.Equal(t, 3, l.size)
	assert.Equal(t, l.first.val, 1)
	assert.Equal(t, l.last.val, 3)
}

func TestGet(t *testing.T) {
	l := New(1, 2, 3, 4, 6, 9)

	v, ok := l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = l.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = l.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = l.Get(3)
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = l.Get(4)
	assert.True(t, ok)
	assert.Equal(t, 6, v)

	v, ok = l.Get(5)
	assert.True(t, ok)
	assert.Equal(t, 9, v)

	v, ok = l.Get(6)
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestPrepend(t *testing.T) {
	l := New[int]()

	l.Prepend(1, 2, 3)

	assert.Equal(t, 3, l.size)

	assert.Equal(t, l.first.val, 1)
	assert.Equal(t, l.last.val, 3)

}

func TestRemove(t *testing.T) {
	l := New(1)

	v, ok := l.Remove(0)
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	assert.Nil(t, l.first)
	assert.Nil(t, l.last)

	l.Add(1, 2, 3, 4)

	v, ok = l.Remove(0)
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = l.Remove(2)
	assert.True(t, ok)
	assert.Equal(t, 4, v)
}

func TestContains(t *testing.T) {
	l := New(1, 2, 3)

	assert.True(t, l.Contains(1))
	assert.True(t, l.Contains(2))
	assert.True(t, l.Contains(3))
	assert.False(t, l.Contains(4))
	assert.True(t, l.Contains(1, 3))
	assert.False(t, l.Contains(1, 4))
}

func TestIndexOf(t *testing.T) {
	l := New(1, 2, 3)

	assert.Equal(t, 0, l.IndexOf(1))
	assert.Equal(t, 1, l.IndexOf(2))
	assert.Equal(t, 2, l.IndexOf(3))
	assert.Equal(t, -1, l.IndexOf(4))
}

func TestSort(t *testing.T) {
	l := New(3, 2, 1)

	l.Sort(func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, 1, l.first.val)
	assert.Equal(t, 3, l.last.val)

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

	l.Insert(8, 9, 10)
	// 1 7 8 2 3 5 6 4 9 10
	assert.Equal(t, l.Values()[8], 9)
	assert.Equal(t, l.Values()[9], 10)
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
