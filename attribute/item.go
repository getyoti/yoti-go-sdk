package attribute

import (
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr"
)

// Item is a structure which contains information about an attribute value
type Item struct {
	contentType yotiprotoattr.ContentType
	value       interface{}
}

// GetContentType returns the content type of the item.
func (item *Item) GetContentType() yotiprotoattr.ContentType {
	return item.contentType
}

// GetValue returns the underlying data of the item
func (item *Item) GetValue() interface{} {
	return item.value
}
