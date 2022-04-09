package doublelinkedlist

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/lists"
	"github.com/qianxi0410/gogds/utils"
)

func assertListImplementation[T comparable]() {
	var _ lists.Lists[T] = New[T]()
}

// node represents a node in the list
// each node has a value and a pointer to the next and previous nodes
// in the list.
type node[T any] struct {
	next, prev *node[T]
	val        T
}

type List[T comparable] struct {
	first, last *node[T]
	size        int
}

// New returns a new list.
func New[T comparable](values ...T) *List[T] {
	l := &List[T]{}
	if len(values) > 0 {
		l.Add(values...)
	}

	return l
}

// Add adds values to the list.
func (l *List[T]) Add(values ...T) {
	for _, v := range values {
		n := &node[T]{val: v, prev: l.last}
		// if the list is empty, the new node is the first and last node
		if l.last == nil {
			l.first = n
			l.last = n
		} else {
			l.last.next = n
			l.last = n
		}
	}

	l.size += len(values)
}

// Append adds values to the end of the list.
func (l *List[T]) Append(values ...T) {
	l.Add(values...)
}

// Prepend adds values to the beginning of the list.
func (l *List[T]) Prepend(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		n := &node[T]{val: values[i], next: l.first}
		// if the list is empty, the new node is the first and last node
		if l.first == nil {
			l.first = n
			l.last = n
		} else {
			l.first.prev = n
			l.first = n
		}
	}

	l.size += len(values)
}

// Get returns the value at index i.
func (l *List[T]) Get(i int) (t T, _ bool) {

	if !l.checkIdx(i) {
		return t, false
	}

	var cur *node[T]
	if i > l.size/2 {
		// start from the end
		cur = l.last
		for j := l.size - 1; j > i; j-- {
			cur = cur.prev
		}
	} else {
		// start from the beginning
		cur = l.first
		for j := 0; j < i; j++ {
			cur = cur.next
		}
	}

	return cur.val, true
}

// Remove removes the value at index i.
func (l *List[T]) Remove(i int) (t T, _ bool) {
	if !l.checkIdx(i) {
		return t, false
	}

	if l.size == 1 {
		t = l.first.val
		l.Clear()
		return t, true
	}

	var cur *node[T]
	if i > l.size/2 {
		cur = l.last
		for j := l.size - 1; j > i; j-- {
			cur = cur.prev
		}
	} else {
		cur = l.first
		for j := 0; j < i; i++ {
			cur = cur.next
		}
	}

	if cur.prev != nil {
		cur.prev.next = cur.next
		if cur.next != nil {
			cur.next.prev = cur.prev
		}
	} else {
		l.first = cur.next
	}

	cur.next = nil
	cur.prev = nil
	l.size--

	return cur.val, true
}

// Contains checks if the list contains the value.
func (l *List[T]) Contains(values ...T) bool {
	for _, v := range values {
		flag := false
		for cur := l.first; cur != nil; cur = cur.next {
			if cur.val == v {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}

	return true
}

// Values returns a slice of all the values in the list.
func (l *List[T]) Values() []T {
	values := make([]T, l.size)
	for i, cur := 0, l.first; cur != nil; i, cur = i+1, cur.next {
		values[i] = cur.val
	}

	return values
}

// IndexOf returns the index of the first occurrence of the value.
func (l *List[T]) IndexOf(value T) int {
	for i, cur := 0, l.first; cur != nil; i, cur = i+1, cur.next {
		if cur.val == value {
			return i
		}
	}

	return -1
}

// Empty returns true if the list is empty.
func (l *List[T]) Empty() bool {
	return l.size == 0
}

// Size returns the number of values in the list.
func (l *List[T]) Size() int {
	return l.size
}

// Sort sorts the list.
func (l *List[T]) Sort(c utils.Comparator[T]) {
	if l.size < 2 {
		return
	}

	values := l.Values()

	utils.Sort(values, c)
	l.Clear()
	l.Add(values...)
}

// Swap swaps the values at the given indices.
func (l *List[T]) Swap(i, j int) {
	if !l.checkIdx(i) || !l.checkIdx(j) {
		return
	}

	var curi, curj *node[T]

	if i > l.size/2 {
		curi = l.last
		for k := l.size - 1; k > i; k-- {
			curi = curi.prev
		}
	} else {
		curi = l.first
		for k := 0; k < i; k++ {
			curi = curi.next
		}
	}

	if j > l.size/2 {
		curj = l.last
		for k := l.size - 1; k > j; k-- {
			curj = curj.prev
		}
	} else {
		curj = l.first
		for k := 0; k < j; k++ {
			curj = curj.next
		}
	}

	curi.val, curj.val = curj.val, curi.val
}

// Insert inserts the value at the given index.
func (l *List[T]) Insert(i int, values ...T) {
	if !l.checkIdx(i) {
		if i == l.size {
			l.Add(values...)
		}

		return
	}

	if i == 0 {
		l.Prepend(values...)
		return
	}

	// find the insert position
	var cur *node[T]
	if i > l.size/2 {
		cur = l.last
		for j := l.size - 1; j > i-1; j-- {
			cur = cur.prev
		}
	} else {
		cur = l.first
		for j := 0; j < i-1; j++ {
			cur = cur.next
		}
	}

	// insert the values
	next := cur.next

	for _, v := range values {
		n := &node[T]{val: v, prev: cur}
		cur.next = n
		cur = cur.next
	}

	cur.next = next
	next.prev = cur
	l.size += len(values)
}

// Set sets the value at index i.
func (l *List[T]) Set(i int, value T) {
	if !l.checkIdx(i) {
		return
	}

	var cur *node[T]
	if i > l.size/2 {
		cur = l.last
		for j := l.size - 1; j > i; j-- {
			cur = cur.prev
		}
	} else {
		cur = l.first
		for j := 0; j < i; j++ {
			cur = cur.next
		}
	}

	cur.val = value
}

// String returns a string representation of container
func (l *List[T]) String() string {
	values := make([]string, 0, l.size)
	for element := l.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.val))
	}
	str := fmt.Sprintf("DoubleLinkedList: [%s]\n", strings.Join(values, ", "))
	return str
}

// Clear removes all values from the list.
func (l *List[T]) Clear() {
	l.first = nil
	l.last = nil
	l.size = 0
}

// checkIdx checks if the index is within the bounds of the list.
func (l *List[T]) checkIdx(index int) bool {
	return index >= 0 && index < l.size
}
