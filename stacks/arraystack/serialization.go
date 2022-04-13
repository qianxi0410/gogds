package arraystack

import (
	"github.com/qianxi0410/gogds/containers"
)

// nolint
func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = New[T]()
	var _ containers.JSONDeserializer = New[T]()
}

// ToJSON outputs the JSON representation of list's elements.
func (s *Stack[T]) ToJSON() ([]byte, error) {
	return s.list.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (s *Stack[T]) FromJSON(data []byte) error {
	return s.list.FromJSON(data)
}
