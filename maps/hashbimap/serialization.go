package hashbimap

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/qianxi0410/gogds/containers"
)

// nolint
func assertSerializationImplementation[K comparable, V comparable]() {
	var _ containers.JSONSerializer = New[K, V]()
	var _ containers.JSONDeserializer = New[K, V]()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ToJSON outputs the JSON representation of list's elements.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	return m.kv.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	eles := make(map[K]V)
	err := json.Unmarshal(data, &eles)

	if err == nil {
		m.Clear()
		for k, v := range eles {
			m.Put(k, v)
		}
	}

	return err
}
