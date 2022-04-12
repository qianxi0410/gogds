package linkedstack

import (
	"fmt"
	"strings"

	"github.com/qianxi0410/gogds/lists/linkedlist"
	"github.com/qianxi0410/gogds/stacks"
)

func assertStackImplementation[T comparable]() {
	var _ stacks.Stack[T] = New[T]()
}

// Stack is a last-in-first-out (LIFO) stack of elements.
// hold a linkedlist.
type Stack[T comparable] struct {
	list *linkedlist.List[T]
}

// New creates a new stack.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{list: linkedlist.New[T]()}
}

// Push adds a new element to the top of the stack.
func (s *Stack[T]) Push(elem T) {
	s.list.Add(elem)
}

// Pop removes the top element from the stack and returns it.
// It returns (zero-value, false) if the stack is empty.
func (s *Stack[T]) Pop() (t T, _ bool) {
	return s.list.Remove(s.list.Size() - 1)
}

// Peek return the peek value of the stack
// if stack is empty return (zero-value, false)
func (s *Stack[T]) Peek() (t T, _ bool) {
	return s.list.Get(s.list.Size() - 1)
}

// Empty return true if the stakc is empty
func (s *Stack[T]) Empty() bool {
	return s.list.Empty()
}

// Size return size of this stack.
func (s *Stack[T]) Size() int {
	return s.list.Size()
}

// Clear clear this stack.
func (s *Stack[T]) Clear() {
	s.list.Clear()
}

// Values return values the stack hole.
func (s *Stack[T]) Values() []T {
	size := s.list.Size()
	values := make([]T, 0, size)

	for i := 1; i <= size; i++ {
		values[size-i], _ = s.list.Get(i - 1)
	}

	return values
}

// String returns a string represent this stack.
func (s *Stack[T]) String() string {
	values := make([]string, 0, s.list.Size())
	for _, v := range s.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}

	str := fmt.Sprintf("LinkedStack: [%s]\n", strings.Join(values, ", "))
	return str
}

func (s *Stack[T]) checkIdx(index int) bool {
	return index >= 0 && index < s.list.Size()
}
