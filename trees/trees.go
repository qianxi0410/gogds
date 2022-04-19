package trees

import "github.com/qianxi0410/gogds/containers"

// Tree is a interface that defines a tree structure.
type Tree[T any] interface {
	containers.Container[T]
}
