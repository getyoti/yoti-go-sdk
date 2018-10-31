package attribute

//Generic is a Yoti attribute which returns a generic value
type Generic struct {
	Attribute
	Value string
	Data  []byte
}

//NewGeneric creates a new generic attribute
func NewGeneric(a *Attribute) *Generic {

	return &Generic{
		Attribute: Attribute{
			Name:    a.Name,
			Type:    AttrTypeInterface,
			Anchors: a.Anchors,
		},
		Value: string(a.Value),
		Data:  a.Value,
	}
}
