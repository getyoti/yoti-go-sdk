package attribute

import (
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// Item is a structure which contains information about an attribute value
type Item struct {
	// ContentType is the content of the item.
	ContentType yotiprotoattr.ContentType

	// Value is the underlying data of the item.
	Value interface{}
}
