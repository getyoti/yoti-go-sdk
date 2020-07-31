package check

// RequestedCheck requests creation of a Check to be performed on a document
type RequestedCheck struct {
	Type   string               `json:"type"`
	Config RequestedCheckConfig `json:"config"`
}

// RequestedCheckConfig is the configuration applied when creating a Check
type RequestedCheckConfig struct {
}
