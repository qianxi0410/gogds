package circularqueue

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/qianxi0410/gogds/containers"
)

// nolint
func assertSerializationImplementation[T comparable]() {
	var _ containers.JSONSerializer = New[T](0)
	var _ containers.JSONDeserializer = New[T](0)
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ToJSON outputs the JSON representation of list's elements.
func (q *Queue[T]) ToJSON() ([]byte, error) {
	return json.Marshal(q.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (q *Queue[T]) FromJSON(data []byte) error {
	values := make([]T, q.maxSize)
	err := json.Unmarshal(data, &values)
	if err == nil {
		q.Clear()
		for _, v := range values {
			q.Enqueue(v)
		}
	}
	return err
}
