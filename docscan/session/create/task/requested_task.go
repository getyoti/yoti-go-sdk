package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedTask requests creation of a Task to be performed on each document
type RequestedTask struct {
	Type   constants.IDDocument `json:"type"`
	Config RequestedTaskConfig  `json:"config"`
}

// RequestedTaskConfig  is the configuration applied when creating a Task
type RequestedTaskConfig struct {
}
