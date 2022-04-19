package binaryheap

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/lists/arraylist"
	"github.com/qianxi0410/gogds/trees"
	"github.com/qianxi0410/gogds/utils"
)

// nolint
func assertListImplementation[T comparable]() {
	var _ trees.Tree[T] = New[T](nil)
}

type Heap[T comparable] struct {
	list *arraylist.List[T]
	Cmp  utils.Comparator[T]
}

// Push pushes the element into the heap.
func (h *Heap[T]) Push(elem ...T) {
	if len(elem) == 1 {
		h.list.Add(elem[0])
		h.bubbleUp()
	} else {
		h.list.Add(elem...)
		size := h.list.Size()/2 + 1
		for i := size; i >= 0; i-- {
			h.bubbleDownIndex(i)
		}
	}
}

// Pop pops the element from the heap.
func (h *Heap[T]) Pop() (v T, ok bool) {
	v, ok = h.list.Get(0)
	if !ok {
		return
	}

	lastIdx := h.list.Size() - 1
	h.list.Swap(0, lastIdx)
	h.list.Remove(lastIdx)
	h.bubbleDown()
	return
}

// Values returns the values of the heap.
func (h *Heap[T]) Values() []T {
	return h.list.Values()
}

// Peek returns the element at the top of the heap.
func (h *Heap[T]) Peek() (v T, ok bool) {
	return h.list.Get(0)
}

// Empty returns true if the heap is empty.
func (h *Heap[T]) Empty() bool {
	return h.list.Empty()
}

// Clear clears the heap.
func (h *Heap[T]) Clear() {
	h.list.Clear()
}

// Size returns the size of the heap.
func (h *Heap[T]) Size() int {
	return h.list.Size()
}

// String returns a string representation of the heap.
func (h *Heap[T]) String() string {
	values := make([]string, h.list.Size())

	for i, v := range h.list.Values() {
		values[i] = fmt.Sprintf("%v", v)
	}

	return fmt.Sprintf("BinaryHeap: [%s]\n", strings.Join(values, ", "))
}

// New creates a new binary heap.
func New[T comparable](c utils.Comparator[T]) *Heap[T] {
	return &Heap[T]{
		list: arraylist.New[T](),
		Cmp:  c,
	}
}

// bubbleDownIndex bubbles down the element at the given index.
func (h *Heap[T]) bubbleDownIndex(index int) {
	size := h.list.Size()

	for leftIdx := index*2 + 1; leftIdx < size; leftIdx = index*2 + 1 {
		rightIdx := leftIdx + 1
		smallerIdx := leftIdx

		leftValue, _ := h.list.Get(leftIdx)
		rightValue, _ := h.list.Get(rightIdx)

		if rightIdx < size && h.Cmp(rightValue, leftValue) {
			smallerIdx = rightIdx
		}
		indexValue, _ := h.list.Get(index)
		smallerValue, _ := h.list.Get(smallerIdx)
		if !h.Cmp(indexValue, smallerValue) {
			h.list.Swap(index, smallerIdx)
		} else {
			break
		}
		index = smallerIdx
	}
}

// bubbleDown bubbles down the element at the 0 index.
func (h *Heap[T]) bubbleDown() {
	h.bubbleDownIndex(0)
}

// bubbleUp bubbles up the element.
func (h *Heap[T]) bubbleUp() {
	index := h.list.Size() - 1

	for parentIdx := (index - 1) / 2; index > 0; parentIdx = (index - 1) / 2 {
		indexValue, _ := h.list.Get(index)
		parentValue, _ := h.list.Get(parentIdx)

		if h.Cmp(parentValue, indexValue) {
			break
		}

		h.list.Swap(index, parentIdx)
		index = parentIdx
	}
}

// checkIdx checks the index.
func (h *Heap[T]) checkIdx(idx int) bool {
	return idx >= 0 && idx < h.list.Size()
}
