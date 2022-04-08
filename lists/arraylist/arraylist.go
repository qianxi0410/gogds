package arraylist

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/lists"
	"github.com/qianxi0410/gogds/utils"
)

func assertListImplementation[T comparable]() {
	var _ lists.Lists[T] = New[T]()
}

type List[T comparable] struct {
	elements []T
	size     int
}

const (
	growthFactor = float32(2)
	shrinkFactor = float32(0.25)
)

func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}

	return list
}

// Size returns the number of elements within the list.
func (l *List[T]) Size() int {
	return l.size
}

// Add add elements to list.
func (l *List[T]) Add(values ...T) {
	l.grow(len(values))

	for _, value := range values {
		l.elements[l.size] = value
		l.size++
	}
}

// Get get element by index.
func (l *List[T]) Get(index int) (t T, ok bool) {
	if !l.checkIdx(index) {
		return
	}

	return l.elements[index], true
}

// Remove remove element by index.
// and return the removed element.
func (l *List[T]) Remove(index int) (t T, ok bool) {
	if !l.checkIdx(index) {
		return
	}

	t = l.elements[index]
	copy(l.elements[index:], l.elements[index+1:l.size])
	l.size--

	l.shrink()
	return t, true
}

// Contains check whether the list contains all the values.
func (l *List[comparable]) Contains(values ...comparable) bool {
	for _, searchValue := range values {
		found := false
		for _, element := range l.elements {
			if element == searchValue {
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

// Values return all the elements in the list.
func (l *List[T]) Values() []T {
	newEle := make([]T, l.size, l.size)
	copy(newEle, l.elements[:l.size])
	return newEle
}

// to check wheather idx is leegal.
func (l *List[T]) checkIdx(idx int) bool {
	return idx >= 0 && idx < l.size
}

// IndexOf return the index of the first occurrence of the value.
// if not found return -1
func (l *List[comparable]) IndexOf(value comparable) int {
	if l.size == 0 {
		return -1
	}

	for idx, ele := range l.elements {
		if ele == value {
			return idx
		}
	}

	return -1
}

// Empty return whether the list is empty.
func (l *List[T]) Empty() bool {
	return l.size == 0
}

// Clear clean the list.
func (l *List[T]) Clear() {
	l.size = 0
	l.elements = []T{}
}

// Sort sort the list.
func (l *List[comparable]) Sort(c utils.Comparator[comparable]) {
	if l.size < 2 {
		return
	}

	if c != nil {
		utils.Sort(l.elements[:l.size], c)
		return
	}
}

// Swap swap the element at index i and j.
func (l *List[T]) Swap(i, j int) {
	if l.checkIdx(i) && l.checkIdx(j) {
		l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
	}
}

// Insert insert the values at index.
func (l *List[T]) Insert(index int, values ...T) {
	if !l.checkIdx(index) {
		// Append
		if index == l.size {
			l.Add(values...)
		}
		return
	}

	size := len(values)
	l.grow(size)
	l.size += size
	copy(l.elements[index+size:], l.elements[index:l.size-size])
	copy(l.elements[index:], values)
}

// Set set the element at index to value.
func (l *List[T]) Set(index int, value T) {
	if !l.checkIdx(index) {
		if index == l.size {
			l.Add(value)
		}
		return
	}

	l.elements[index] = value
}

// String returns a string representation of container
func (l *List[T]) String() string {
	str := "ArrayList: "
	values := []string{}
	for _, value := range l.elements[:l.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	str = "[" + str + "]"
	return str
}

// grow expand the array if necessary
// i.e. capacity will be reached if we add n elements
func (l *List[T]) grow(n int) {
	curCap := cap(l.elements)
	if l.size+n >= curCap {
		newCap := int(growthFactor * float32(curCap+n))
		l.resize(newCap)
	}
}

// shrink the array if necessary
// i.e. when size is shrinkFactor percent of current capacity
func (l *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// shrink when size is at shrinkFactor * capacity
	curCap := cap(l.elements)
	if l.size <= int(float32(curCap)*shrinkFactor) {
		l.resize(l.size)
	}
}

func (l *List[T]) resize(cap int) {
	newEle := make([]T, cap, cap)
	copy(newEle, l.elements)
	l.elements = newEle
}
