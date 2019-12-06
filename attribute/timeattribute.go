package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

// TimeAttribute is a Yoti attribute which returns a time as its value
type TimeAttribute struct {
	attribute *yotiprotoattr.Attribute
	value     *time.Time
	anchors   []*anchor.Anchor
}

// NewTime creates a new Time attribute
func NewTime(a *yotiprotoattr.Attribute) (*TimeAttribute, error) {
	parsedTime, err := time.Parse("2006-01-02", string(a.Value))
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", a.Value, err)
		parsedTime = time.Time{}
		return nil, err
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &TimeAttribute{
		attribute: &yotiprotoattr.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		value:   &parsedTime,
		anchors: parsedAnchors,
	}, nil
}

// Value returns the value of the TimeAttribute as *time.Time
func (a *TimeAttribute) Value() *time.Time {
	return a.value
}

// Name returns the name as a string
func (a *TimeAttribute) Name() string {
	return a.attribute.Name
}

// ContentType returns the content type as a string
func (a *TimeAttribute) ContentType() string {
	return a.attribute.ContentType.String()
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *TimeAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *TimeAttribute) Sources() []*anchor.Anchor {
	return anchor.GetSources(a.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *TimeAttribute) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(a.anchors)
}
