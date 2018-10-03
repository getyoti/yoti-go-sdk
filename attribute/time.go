package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/anchor"
)

//Time is a Yoti attribute which returns a time as its value
type Time struct {
	Name    string
	Value   *time.Time
	Type    AttrType
	Anchors []*anchor.Anchor
	Err     error
}

//NewTime creates a new Time attribute
func NewTime(a *Attribute) *Time {
	parsedTime, err := time.Parse("2006-01-02", string(a.Value))
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", a.Value, err)
	}

	return &Time{
		Name:    a.Name,
		Value:   &parsedTime,
		Type:    AttrTypeTime,
		Anchors: a.Anchors,
		Err:     err,
	}
}
