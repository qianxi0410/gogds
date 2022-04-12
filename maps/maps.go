package maps

import "github.com/qianxi0410/gogds/containers"

// Map interface that all maps implement
type Map[K any, V any] interface {
	Put(K, V)
	Get(K) (V, bool)
	Remove(K)
	Keys() []K
	Vlues() []V

	containers.Container[any]
}
