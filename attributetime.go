package yoti

import (
	"log"
	"time"
)

//TimeAttribute is a Yoti attribute which returns a time as its value
type TimeAttribute struct {
	Name    string
	Value   *time.Time
	Type    AttrType
	Anchors []*Anchor
	Err     error
}

func newTimeAttribute(a *Attribute) *TimeAttribute {
	parsedTime, err := time.Parse("2006-01-02", string(a.Value))
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", a.Value, err)
	}

	return &TimeAttribute{
		Name:    a.Name,
		Value:   &parsedTime,
		Type:    AttrTypeTime,
		Anchors: a.Anchors,
		Err:     err,
	}
}
