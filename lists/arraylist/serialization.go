package arraylist

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
func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List[T]) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}
