package lists

import "github.com/qianxi0410/gogds/containers"

type EnumerableWithIndex[T comparable] interface {
	// Filter returns a new container containing all elements for which the given function returns a true value.
	Filter(func(index int, value T) bool) Lists[T]

	containers.EnumerableWithIndex[T]
}
