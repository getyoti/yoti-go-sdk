package sandbox

type ResponseConfig struct {
	TaskResults  *TaskResults  `json:"task_results,omitempty"`
	CheckReports *CheckReports `json:"check_reports"`
}

type responseConfigBuilder struct {
	taskResults  *TaskResults
	checkReports *CheckReports
	err          error
}

func NewResponseConfigBuilder() *responseConfigBuilder {
	return &responseConfigBuilder{}
}

func (b *responseConfigBuilder) WithTaskResults(taskResults TaskResults) *responseConfigBuilder {
	b.taskResults = &taskResults
	return b
}

func (b *responseConfigBuilder) WithCheckReports(checkReports CheckReports) *responseConfigBuilder {
	b.checkReports = &checkReports
	return b
}

func (b *responseConfigBuilder) Build() (ResponseConfig, error) {
	responseConfig := ResponseConfig{}

	responseConfig.CheckReports = b.checkReports

	if b.taskResults != nil {
		responseConfig.TaskResults = b.taskResults
	}

	return responseConfig, b.err
}
