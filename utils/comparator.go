package utils

// Comparator is a function that compares two values.
// if a <= b => true
// if a > b => false
type Comparator[T any] func(a, b T) bool
