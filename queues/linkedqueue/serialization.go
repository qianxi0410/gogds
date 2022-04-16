package linkedqueue

import (
	"github.com/qianxi0410/gogds/containers"
)

// nolint
func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = New[T]()
	var _ containers.JSONDeserializer = New[T]()
}

// ToJSON outputs the JSON representation of list's elements.
func (q *Queue[T]) ToJSON() ([]byte, error) {
	return q.list.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (q *Queue[T]) FromJSON(data []byte) error {
	return q.list.FromJSON(data)
}
