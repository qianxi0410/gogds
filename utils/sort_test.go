package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasicSort(t *testing.T) {
	tests := [][2][]int{
		{{1, 2, 3}, {1, 2, 3}},
		{{3, 2, 1}, {1, 2, 3}},
		{{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		Sort(test[0], func(a, b int) bool {
			return a < b
		})

		if !assert.Equal(t, test[0], test[1]) {
			t.Errorf("Sort(%v) = %v, want %v", test[0], test[0], test[1])
		}
	}

	testsf := [][2][]float64{
		{{1.1, 2.2, 3.3}, {1.1, 2.2, 3.3}},
		{{3.3, 2.2, 1.1}, {1.1, 2.2, 3.3}},
		{{10.1, 9.1, 8.1, 7.1, 6.1, 5.1, 4.1, 3.1, 2.1, 1.1}, {1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1}},
	}

	for _, test := range testsf {
		Sort(test[0], func(a, b float64) bool {
			return a < b
		})
		if !assert.Equal(t, test[0], test[1]) {
			t.Errorf("Sort(%v) = %v, want %v", test[0], test[0], test[1])
		}
	}

	testss := [][2][]string{
		{{"a", "b", "c"}, {"a", "b", "c"}},
		{{"c", "b", "a"}, {"a", "b", "c"}},
		{{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}, {"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}},
	}

	for _, test := range testss {
		Sort(test[0], func(a, b string) bool {
			return a < b
		})
		if !assert.Equal(t, test[0], test[1]) {
			t.Errorf("Sort(%v) = %v, want %v", test[0], test[0], test[1])
		}
	}

	testsb := [][2][]bool{
		{{true, false, true}, {false, true, true}},
		{{true, false, true}, {false, true, true}},
	}

	for _, test := range testsb {
		Sort(test[0], func(a, b bool) bool {
			if a && !b {
				return false
			} else {
				return true
			}
		})
		if !assert.Equal(t, test[0], test[1]) {
			t.Errorf("Sort(%v) = %v, want %v", test[0], test[0], test[1])
		}
	}

	now := time.Now()
	before := now.Add(-1 * time.Hour)
	after := now.Add(1 * time.Hour)

	testt := [][2][]time.Time{
		{{now, before, after}, {before, now, after}},
	}

	for _, test := range testt {
		Sort(test[0], func(a, b time.Time) bool {
			return a.Before(b)
		})
		if !assert.Equal(t, test[0], test[1]) {
			t.Errorf("Sort(%v) = %v, want %v", test[0], test[0], test[1])
		}
	}
}

type Coustom struct {
	id   int
	name string
}

func (c Coustom) Compare(another Coustom) int {
	if c.id < another.id {
		return -1
	} else if c.id > another.id {
		return 1
	} else {
		return 0
	}
}

func TestCoustomSort(t *testing.T) {
	tests := [][2][]Coustom{
		{
			{
				{1, "a"},
				{2, "b"},
			},
			{
				{1, "a"},
				{2, "b"},
			},
		},
		{
			{
				{2, "a"},
				{1, "b"},
			},
			{
				{1, "b"},
				{2, "a"},
			},
		},
	}

	for _, test := range tests {
		Sort(test[0], func(a, b Coustom) bool {
			return a.id < b.id
		})
		if !assert.Equal(t, test[0], test[1]) {
			t.Errorf("Sort(%v) = %v, want %v", test[0], test[0], test[1])
		}
	}
}
