package linkedstack

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/qianxi0410/gogds/containers"
)

func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = New[T]()
	var _ containers.JSONDeserializer = New[T]()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ToJSON outputs the JSON representation of list's elements.
func (s *Stack[T]) ToJSON() ([]byte, error) {
	return s.list.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (s *Stack[T]) FromJSON(data []byte) error {
	return s.list.FromJSON(data)
}
