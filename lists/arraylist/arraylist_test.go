package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNew(t *testing.T) {
	list1 := New[int]()

	assert.True(t, list1.Size() == 0)

	list2 := New("a", "b")

	assert.Equal(t, list2.Size(), 2)

	if v, ok := list2.Get(0); ok {
		assert.Equal(t, v, "a")
	}

	if v, ok := list2.Get(1); ok {
		assert.Equal(t, v, "b")
	}
}

func TestSize(t *testing.T) {
	list := New(1, 2, 3)

	assert.Equal(t, list.Size(), 3)

	list.Add(4, 5, 6)
	assert.Equal(t, list.Size(), 6)

	list.Remove(0)
	assert.Equal(t, list.Size(), 5)
}

func TestAdd(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	for i := range list.Values() {
		assert.Equal(t, list.Values()[i], i+1)
	}

	list.Add(4, 5, 6)
	assert.Equal(t, list.Size(), 6)

	for i := range list.Values() {
		assert.Equal(t, list.Values()[i], i+1)
	}
}

func TestGet(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	for i := range list.Values() {
		v, ok := list.Get(i)
		assert.Equal(t, ok, true)
		assert.Equal(t, v, i+1)
	}
}

func TestRemove(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	v, ok := list.Remove(0)
	assert.True(t, ok)
	assert.Equal(t, v, 1)
	assert.Equal(t, list.Size(), 2)

	for i := range list.Values() {
		v, ok := list.Get(i)
		assert.Equal(t, ok, true)
		assert.Equal(t, v, i+2)
	}
}

func TestContains(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	assert.Equal(t, list.Contains(1), true)
	assert.Equal(t, list.Contains(4), false)
	assert.Equal(t, list.Contains(1, 3), true)

	assert.Equal(t, list.Contains(1, 3, 5), false)
}

func TestIndexOf(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	assert.Equal(t, list.IndexOf(1), 0)
	assert.Equal(t, list.IndexOf(4), -1)
	assert.Equal(t, list.IndexOf(1), 0)

	assert.Equal(t, list.IndexOf(6), -1)
}

func TestSort(t *testing.T) {
	list := New(3, 2, 1, 0)
	assert.Equal(t, list.Size(), 4)

	list.Sort(func(a, b int) bool {
		return a < b
	})
	assert.Equal(t, list.Size(), 4)

	for i := range list.Values() {
		assert.Equal(t, list.Values()[i], i)
	}
}

func TestSwap(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	list.Swap(0, 1)
	assert.Equal(t, list.Size(), 3)

	assert.Equal(t, list.Values()[0], 2)
	assert.Equal(t, list.Values()[1], 1)
	assert.Equal(t, list.Values()[2], 3)
}

func TestInsert(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	list.Insert(0, 0)
	assert.Equal(t, list.Size(), 4)

	assert.Equal(t, list.Values()[0], 0)
	assert.Equal(t, list.Values()[1], 1)
	assert.Equal(t, list.Values()[2], 2)
	assert.Equal(t, list.Values()[3], 3)

	list.Insert(1, 4)
	assert.Equal(t, list.Size(), 5)

	assert.Equal(t, list.Values()[0], 0)
	assert.Equal(t, list.Values()[1], 4)
}

func TestSet(t *testing.T) {
	list := New(1, 2, 3)
	assert.Equal(t, list.Size(), 3)

	list.Set(0, 0)
	assert.Equal(t, list.Size(), 3)

	assert.Equal(t, list.Values()[0], 0)
	list.Set(1, 10)
	assert.Equal(t, list.Values()[1], 10)
}
