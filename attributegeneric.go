package yoti

//GenericAttribute is a Yoti attribute which returns a generic value
type GenericAttribute struct {
	Name    string
	Value   string
	Data    []byte
	Type    AttrType
	Anchors []*Anchor
}

func newGenericAttribute(a *Attribute) *GenericAttribute {
	return &GenericAttribute{
		Name:    a.Name,
		Value:   string(a.Value),
		Data:    a.Value,
		Type:    AttrTypeInterface,
		Anchors: a.Anchors,
	}
}
