package dynamic

import (
	"encoding/json"
)

// ScenarioBuilder builds a dynamic scenario
type ScenarioBuilder struct {
	scenario Scenario
	err      error
}

// Scenario represents a dynamic scenario
type Scenario struct {
	policy           *Policy
	extensions       []interface{}
	callbackEndpoint string
	subject          *json.RawMessage
}

// WithPolicy attaches a DynamicPolicy to the DynamicScenario
func (builder *ScenarioBuilder) WithPolicy(policy Policy) *ScenarioBuilder {
	builder.scenario.policy = &policy
	return builder
}

// WithExtension adds an extension to the scenario
func (builder *ScenarioBuilder) WithExtension(extension interface{}) *ScenarioBuilder {
	builder.scenario.extensions = append(builder.scenario.extensions, extension)
	return builder
}

// WithCallbackEndpoint sets the callback URL
func (builder *ScenarioBuilder) WithCallbackEndpoint(endpoint string) *ScenarioBuilder {
	builder.scenario.callbackEndpoint = endpoint
	return builder
}

// WithSubject adds a subject to the scenario. Must be valid JSON.
func (builder *ScenarioBuilder) WithSubject(subject json.RawMessage) *ScenarioBuilder {
	builder.scenario.subject = &subject
	return builder
}

// Build constructs the DynamicScenario
func (builder *ScenarioBuilder) Build() (Scenario, error) {
	if builder.scenario.extensions == nil {
		builder.scenario.extensions = make([]interface{}, 0)
	}
	if builder.scenario.policy == nil {
		policy, err := (&PolicyBuilder{}).Build()
		if err != nil {
			return builder.scenario, err
		}
		builder.scenario.policy = &policy
	}
	return builder.scenario, builder.err
}

// MarshalJSON returns the JSON encoding
func (scenario Scenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Policy           Policy           `json:"policy"`
		Extensions       []interface{}    `json:"extensions"`
		CallbackEndpoint string           `json:"callback_endpoint"`
		Subject          *json.RawMessage `json:"subject,omitempty"`
	}{
		Policy:           *scenario.policy,
		Extensions:       scenario.extensions,
		CallbackEndpoint: scenario.callbackEndpoint,
		Subject:          scenario.subject,
	})
}
