package priorityqueue

import (
	"github.com/qianxi0410/gogds/containers"
)

// nolint
func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = New[T](nil)
	var _ containers.JSONDeserializer = New[T](nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (q *Queue[T]) ToJSON() ([]byte, error) {
	return q.heap.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (q *Queue[T]) FromJSON(data []byte) error {
	return q.heap.FromJSON(data)
}
