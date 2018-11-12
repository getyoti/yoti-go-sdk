package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//TimeAttribute is a Yoti attribute which returns a time as its value
type TimeAttribute struct {
	*yotiprotoattr_v3.Attribute
	Value     *time.Time
	anchors   []*anchor.Anchor
	sources   []*anchor.Anchor
	verifiers []*anchor.Anchor
}

//NewTime creates a new Time attribute
func NewTime(a *yotiprotoattr_v3.Attribute) (*TimeAttribute, error) {
	parsedTime, err := time.Parse("2006-01-02", string(a.Value))
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", a.Value, err)
		parsedTime = time.Time{}
		return nil, err
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &TimeAttribute{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value:     &parsedTime,
		anchors:   parsedAnchors,
		sources:   anchor.GetSources(parsedAnchors),
		verifiers: anchor.GetVerifiers(parsedAnchors),
	}, nil
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *TimeAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *TimeAttribute) Sources() []*anchor.Anchor {
	return a.sources
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *TimeAttribute) Verifiers() []*anchor.Anchor {
	return a.verifiers
}
