package util

import "reflect"

// Filter filters either: array, chan, map, slice, or string for a condition.
// As it's a generic filter, it's almost 8 times slower than a specific one.
// Use a specific one when possible.
func Filter(arr interface{}, cond func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(arr)
	contentValue := reflect.ValueOf(arr)

	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); cond(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}

	return newContent.Interface()
}
