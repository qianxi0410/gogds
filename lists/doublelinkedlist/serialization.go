package doublelinkedlist

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
