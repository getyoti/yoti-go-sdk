package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// DateAttribute is a Yoti attribute which returns a date as *time.Time for its value
type DateAttribute struct {
	attributeDetails
	value *time.Time
}

// NewDate creates a new Date attribute
func NewDate(a *yotiprotoattr.Attribute) (*DateAttribute, error) {
	parsedTime, err := time.Parse("2006-01-02", string(a.Value))
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", a.Value, err)
		return nil, err
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &DateAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
		},
		value: &parsedTime,
	}, nil
}

// Value returns the value of the TimeAttribute as *time.Time
func (a *DateAttribute) Value() *time.Time {
	return a.value
}
