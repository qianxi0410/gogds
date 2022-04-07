package linkedlist

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/qianxi0410/gogds/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = New[int]()
	var _ containers.JSONSerializer = New[string]()
	var _ containers.JSONSerializer = New[float32]()
	var _ containers.JSONSerializer = New[complex128]()
	var _ containers.JSONSerializer = New[time.Time]()
	var _ containers.JSONSerializer = New[bool]()
	var _ containers.JSONSerializer = New[struct{}]()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ToJSON outputs the JSON representation of list's elements.
func (l *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(l.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (l *List[T]) FromJSON(data []byte) error {
	var elements []T
	err := json.Unmarshal(data, &elements)
	if err == nil {
		l.Clear()
		l.Add(elements...)
	}
	return err
}
