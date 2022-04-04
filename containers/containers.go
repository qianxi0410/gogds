package containers

import "github.com/qianxi0410/gogds/utils"

// Container is a basic generic container interface.
type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
}

// GetSortedValues returns the values of the container in sorted order.
func GetSortedValues[T comparable](c Container[T], comparator utils.Comparator[T]) []T {
	values := c.Values()
	if c.Size() <= 1 {
		return values
	}

	utils.Sort(values, comparator)
	return values
}
