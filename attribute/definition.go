package attribute

// AttributeDefinition contains information about the attribute(s) issued by a third party.
type AttributeDefinition struct {
	name string
}

// Name of the attribute to be issued.
func (a AttributeDefinition) Name() string {
	return a.name
}
