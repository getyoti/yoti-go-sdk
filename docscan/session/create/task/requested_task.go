package task

// RequestedTask requests creation of a Task to be performed on each document
type RequestedTask interface {
	Type() string
	Config() RequestedTaskConfig
	MarshalJSON() ([]byte, error)
}

// RequestedTaskConfig  is the configuration applied when creating a Task
type RequestedTaskConfig interface {
}
