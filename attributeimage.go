package yoti

//ImageAttribute is a Yoti attribute which returns an image as its value
type ImageAttribute struct {
	Name    string
	Value   *Image
	Type    AttrType
	Anchors []*Anchor
}

func newImageAttribute(a *Attribute) *ImageAttribute {
	return &ImageAttribute{
		Name: a.Name,
		Value: &Image{
			Data: a.Value,
		},
		Type:    a.Type,
		Anchors: a.Anchors,
	}
}
