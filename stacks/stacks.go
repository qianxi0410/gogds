package stacks

import "github.com/qianxi0410/gogds/containers"

type Stack[T any] interface {
	Push(T)
	Pop() (T, bool)
	Peek() (T, bool)

	containers.Container[T]
}
