package task

// RequestedTask requests creation of a Task to be performed on each document
type RequestedTask struct {
	Type   string              `json:"type"`
	Config RequestedTaskConfig `json:"config"`
}

// RequestedTaskConfig  is the configuration applied when creating a Task
type RequestedTaskConfig struct {
}
