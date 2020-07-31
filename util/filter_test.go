package util

import (
	"testing"

	"gotest.tools/v3/assert"
)

type myStruct struct {
	s string
	i int8
}

func TestFilter_Ints(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	filteredArr := Filter(arr, func(val interface{}) bool {
		return val == 5 || val == 1
	})

	result := filteredArr.([]int)
	assert.Equal(t, result[0], 1)
	assert.Equal(t, result[1], 5)
}

func TestFilter_Strings(t *testing.T) {
	arr := []string{"a", "b", "c", "b"}
	filteredArr := Filter(arr, func(val interface{}) bool {
		return val != "b"
	})

	result := filteredArr.([]string)

	assert.Equal(t, result[0], "a")
	assert.Equal(t, result[1], "c")
}

func TestFilter_CustomStruct(t *testing.T) {
	arr := make([]myStruct, 2)

	item1 := myStruct{
		s: "a",
		i: 1,
	}

	item2 := myStruct{
		s: "b",
		i: 2,
	}

	arr = append(arr, item1, item2)

	filteredArr := Filter(arr, func(val interface{}) bool {
		return val.(myStruct).i > 1
	})

	result := filteredArr.([]myStruct)

	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0], item2)
}
