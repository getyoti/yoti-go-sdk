package attribute

//String is a Yoti attribute which returns a string as its value
type String struct {
	Attribute
	Value string
}

//NewString creates a new String attribute
func NewString(a *Attribute) *String {
	return &String{
		Attribute: Attribute{
			Name:    a.Name,
			Type:    AttrTypeString,
			Anchors: a.Anchors,
		},
		Value: string(a.Value),
	}
}
