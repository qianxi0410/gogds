package binaryheap

import (
	"github.com/qianxi0410/gogds/containers"
)

// nolint
func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = New[T](nil)
	var _ containers.JSONDeserializer = New[T](nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (h *Heap[T]) ToJSON() ([]byte, error) {
	return h.list.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (h *Heap[T]) FromJSON(data []byte) error {
	return h.list.FromJSON(data)
}
