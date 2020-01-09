package dynamic_sharing_service

import (
	"encoding/json"
)

// DynamicScenarioBuilder builds a dynamic scenario
type DynamicScenarioBuilder struct {
	scenario DynamicScenario
	err      error
}

// DynamicScenario represents a dynamic scenario
type DynamicScenario struct {
	policy           *DynamicPolicy
	extensions       []interface{}
	callbackEndpoint string
}

// WithPolicy attaches a DynamicPolicy to the DynamicScenario
func (builder *DynamicScenarioBuilder) WithPolicy(policy DynamicPolicy) *DynamicScenarioBuilder {
	builder.scenario.policy = &policy
	return builder
}

// WithExtension adds an extension to the scenario
func (builder *DynamicScenarioBuilder) WithExtension(extension interface{}) *DynamicScenarioBuilder {
	builder.scenario.extensions = append(builder.scenario.extensions, extension)
	return builder
}

// WithCallbackEndpoint sets the callback URL
func (builder *DynamicScenarioBuilder) WithCallbackEndpoint(endpoint string) *DynamicScenarioBuilder {
	builder.scenario.callbackEndpoint = endpoint
	return builder
}

// Build constructs the DynamicScenario
func (builder *DynamicScenarioBuilder) Build() (DynamicScenario, error) {
	if builder.scenario.extensions == nil {
		builder.scenario.extensions = make([]interface{}, 0)
	}
	if builder.scenario.policy == nil {
		policy, err := (&DynamicPolicyBuilder{}).Build()
		if err != nil {
			return builder.scenario, err
		}
		builder.scenario.policy = &policy
	}
	return builder.scenario, builder.err
}

// MarshalJSON returns the JSON encoding
func (scenario DynamicScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Policy           DynamicPolicy `json:"policy"`
		Extensions       []interface{} `json:"extensions"`
		CallbackEndpoint string        `json:"callback_endpoint"`
	}{
		Policy:           *scenario.policy,
		Extensions:       scenario.extensions,
		CallbackEndpoint: scenario.callbackEndpoint,
	})
}
