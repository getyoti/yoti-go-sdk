package yoti

//StringAttribute is a Yoti attribute which returns a string as its value
type StringAttribute struct {
	Name    string
	Value   string
	Type    AttrType
	Anchors []*Anchor
}

func newStringAttribute(a *Attribute) *StringAttribute {
	return &StringAttribute{
		Name:    a.Name,
		Value:   string(a.Value),
		Type:    AttrTypeString,
		Anchors: a.Anchors,
	}
}
