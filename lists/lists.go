package lists

import (
	"github.com/qianxi0410/gogds/containers"
	"github.com/qianxi0410/gogds/utils"
)

type Lists[T any] interface {
	Get(int) (T, bool)
	Remove(int) (T, bool)
	Add(values ...T)
	Contains(values ...T) bool
	Sort(utils.Comparator[T])
	Swap(int, int)
	Insert(int, ...T)
	Set(int, T)
	IndexOf(T) int
	String() string

	Prepend(values ...T)
	Append(values ...T)

	containers.Container[T]
}
