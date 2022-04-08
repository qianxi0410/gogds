package linkedlist

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/lists"
	"github.com/qianxi0410/gogds/utils"
)

func assertListImplementation[T comparable]() {
	var _ lists.Lists[T] = New[T]()
}

// node represents a node in a linked list
// each node has a value and a pointer to the next node.
type node[T any] struct {
	val  T
	next *node[T]
}

// List represents a linked list
// the first node is the head of the list and the last node is the tail of the list.
type List[T comparable] struct {
	first *node[T]
	last  *node[T]
	size  int
}

// New returns a new list.
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}

	if len(values) > 0 {
		list.Add(values...)
	}

	return list
}

// Add adds a value to the list.
func (l *List[T]) Add(values ...T) {
	for _, v := range values {
		n := &node[T]{val: v}
		if l.size == 0 {
			l.first = n
			l.last = n
		} else {
			l.last.next = n
			l.last = n
		}
		l.size++
	}
}

// Append adds a value to the end of the list.
func (l *List[T]) Append(values ...T) {
	l.Add(values...)
}

// Prepend adds a value to the beginning of the list.
func (l *List[T]) Prepend(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		n := &node[T]{val: values[i], next: l.first}

		l.first = n
		if l.size == 0 {
			l.last = n
		}

		l.size++
	}
}

// Get returns the value at the given index.
func (l *List[T]) Get(index int) (t T, _ bool) {

	if !l.checkIdx(index) {
		return t, false
	}

	cur := l.first
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val, true
}

// Remove removes the value at the given index.
func (l *List[T]) Remove(index int) (t T, _ bool) {
	if !l.checkIdx(index) {
		return t, false
	}

	if index == 0 {
		t = l.first.val

		if l.first == l.last {
			l.last = nil
		}
		l.first = l.first.next
		l.size--
		return t, true
	}

	if l.size == 1 {
		l.Clear()
		return
	}

	cur := l.first
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}

	t = cur.next.val
	cur.next = cur.next.next
	l.size--

	return t, true
}

// Contains checks if the given value is in the list.
func (l *List[T]) Contains(values ...T) bool {
	if len(values) == 0 {
		return true
	}

	if l.size == 0 {
		return false
	}

	for _, v := range values {
		found := false
		for ele := l.first; ele != nil; ele = ele.next {
			if ele.val == v {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values return a copy of the list's elements.
func (l *List[T]) Values() []T {
	values := make([]T, l.size)

	for e, ele := 0, l.first; ele != nil; e, ele = e+1, ele.next {
		values[e] = ele.val
	}
	return values
}

// IndexOf return the index of the first element matching the given value.
func (l *List[T]) IndexOf(value T) int {
	if l.size == 0 {
		return -1
	}

	for idx, ele := 0, l.first; ele != nil; idx, ele = idx+1, ele.next {
		if ele.val == value {
			return idx
		}
	}

	return -1
}

// Empty returns true if the list is empty.
func (l *List[T]) Empty() bool {
	return l.size == 0
}

// Size returns the number of elements in the container.
func (l *List[T]) Size() int {
	return l.size
}

// Sort sort values in this container.
func (l *List[T]) Sort(c utils.Comparator[T]) {
	// FIXME: this way is too slow, maybe use some other sort algorithm.
	if l.size < 2 {
		return
	}

	values := l.Values()
	utils.Sort(values, c)

	l.Clear()
	l.Add(values...)
}

// Swap swaps values in this container.
func (l *List[T]) Swap(i, j int) {
	if l.checkIdx(i) && l.checkIdx(j) && i != j {
		var ele1, ele2 *node[T]

		for e, cur := 0, l.first; cur != nil; e, cur = e+1, cur.next {
			if i == e {
				ele1 = cur
			} else if j == e {
				ele2 = cur
			}
		}

		ele1.val, ele2.val = ele2.val, ele1.val
	}
}

// Insert inserts values in this container.
// The values are inserted before the given index.
func (l *List[T]) Insert(index int, values ...T) {
	if !l.checkIdx(index) {
		if index == l.size {
			l.Append(values...)
		}
		return
	}

	l.size += len(values)

	dummy := &node[T]{next: l.first}
	cur := dummy

	for i := 0; i < index; i++ {
		cur = cur.next
	}

	next := cur.next
	for _, v := range values {
		node := &node[T]{val: v, next: cur.next}
		cur.next = node
		cur = node
	}

	cur.next = next

	l.first = dummy.next
}

// Set sets the value at the given index.
// If the index is equal to the sie, the value is appended to the list.
func (l *List[T]) Set(index int, value T) {
	if !l.checkIdx(index) {
		if index == l.size {
			l.Add(value)
		}
		return
	}

	cur := l.first
	for index > 0 {
		cur = cur.next
		index--
	}

	cur.val = value
}

// String returns a string representation of the list.
func (l *List[T]) String() string {
	str := "LinkedList: "
	values := []string{}
	for element := l.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.val))
	}
	str += strings.Join(values, ", ")
	str = "[" + str + "]"
	return str
}

// Clear removes all elements in the container.
func (l *List[T]) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}

// checkIdx checks if the given index is valid.
func (l *List[T]) checkIdx(index int) bool {
	return index >= 0 && index < l.size
}
