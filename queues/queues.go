package queues

import "github.com/qianxi0410/gogds/containers"

// Queue interface that all queues implement
type Queues[V any] interface {
	Enqueue(v V)
	Dequeue() V
	Peek() V

	containers.Container[V]
}
