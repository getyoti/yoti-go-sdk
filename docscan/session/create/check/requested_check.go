package check

// RequestedCheck requests creation of a Check to be performed on a document
type RequestedCheck interface {
	Type() string
	Config() RequestedCheckConfig
	MarshalJSON() ([]byte, error)
}

// RequestedCheckConfig is the configuration applied when creating a Check
type RequestedCheckConfig interface {
}
