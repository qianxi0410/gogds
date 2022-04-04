package utils

import "sort"

type wrapper[T comparable] struct {
	values     []T
	comparator Comparator[T]
}

func (s wrapper[any]) Len() int {
	return len(s.values)
}

func (s wrapper[any]) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s wrapper[T]) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j])
}

// SortWithComparator sorts the given values using the given comparator.
func Sort[T comparable](values []T, c Comparator[T]) {
	sort.Sort(wrapper[T]{values, c})
}
