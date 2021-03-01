package request

import (
	"errors"
	"fmt"
)

// ResponseConfig represents the response config
type ResponseConfig struct {
	TaskResults  *TaskResults  `json:"task_results,omitempty"`
	CheckReports *CheckReports `json:"check_reports"`
}

// ResponseConfigBuilder builds ResponseConfig
type ResponseConfigBuilder struct {
	taskResults  *TaskResults
	checkReports *CheckReports
}

// NewResponseConfigBuilder creates a new ResponseConfigBuilder
func NewResponseConfigBuilder() *ResponseConfigBuilder {
	return &ResponseConfigBuilder{}
}

// WithTaskResults adds task results to the response configuration
func (b *ResponseConfigBuilder) WithTaskResults(taskResults TaskResults) *ResponseConfigBuilder {
	b.taskResults = &taskResults
	fmt.Printf("DEBUG: %v\n", taskResults)
	return b
}

// WithCheckReports adds check reports to the response configuration
func (b *ResponseConfigBuilder) WithCheckReports(checkReports CheckReports) *ResponseConfigBuilder {
	b.checkReports = &checkReports
	return b
}

// Build creates ResponseConfig
func (b *ResponseConfigBuilder) Build() (*ResponseConfig, error) {
	responseConfig := &ResponseConfig{
		CheckReports: b.checkReports,
		TaskResults:  b.taskResults,
	}

	if responseConfig.CheckReports == nil {
		return nil, errors.New("Check Reports must be provided")
	}

	return responseConfig, nil
}
