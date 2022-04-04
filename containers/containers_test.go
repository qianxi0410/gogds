package containers

import (
	"testing"
)

type TestContainer[T comparable] struct {
	values []T
}

func (container TestContainer[any]) Empty() bool {
	return len(container.values) == 0
}

func (container TestContainer[any]) Size() int {
	return len(container.values)
}

func (container TestContainer[Comparable]) Clear() {
	container.values = []Comparable{}
}

func (container TestContainer[T]) Values() []T {
	return container.values
}

func TestGetSortedValuesInts(t *testing.T) {
	container := TestContainer[int]{
		values: []int{1, 2, 3, 4, 5, -1, -2, -3, -4, -5},
	}
	values := GetSortedValues[int](container, func(a, b int) bool {
		return a < b
	})
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesString(t *testing.T) {
	container := TestContainer[string]{
		values: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
	}
	values := GetSortedValues[string](container, func(a, b string) bool {
		return a < b
	})
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}
