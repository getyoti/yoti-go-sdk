package dynamicsharingservice

import (
	"encoding/json"
	"github.com/getyoti/yoti-go-sdk/v2/dynamicsharingservice/policy"
)

// DynamicScenarioBuilder builds a Dynamic Scenario
type DynamicScenarioBuilder struct {
	scenario DynamicScenario
}

// DynamicScenario represents a dynamic scenario
type DynamicScenario struct {
	policy           policy.DynamicPolicy
	extensions       []interface{}
	callbackEndpoint string
}

// New initializes the state of a DynamicScenarioBuilder before its use
func (builder *DynamicScenarioBuilder) New() *DynamicScenarioBuilder {
	builder.scenario.policy = (&policy.DynamicPolicyBuilder{}).New().Build()
	builder.scenario.callbackEndpoint = ""
	return builder
}

// WithPolicy attaches a DynamicPolicy to the DynamicScenario
func (builder *DynamicScenarioBuilder) WithPolicy(policy policy.DynamicPolicy) *DynamicScenarioBuilder {
	builder.scenario.policy = policy
	return builder
}

// WithCallbackEndpoint sets the callback URL
func (builder *DynamicScenarioBuilder) WithCallbackEndpoint(endpoint string) *DynamicScenarioBuilder {
	builder.scenario.callbackEndpoint = endpoint
	return builder
}

// Build constructs the DynamicScenario
func (builder *DynamicScenarioBuilder) Build() DynamicScenario {
	return builder.scenario
}

// MarshalJSON ...
func (scenario *DynamicScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Policy           policy.DynamicPolicy `json:"policy"`
		Extensions       []interface{}        `json:"extensions"`
		CallbackEndpoint string               `json:"callback_endpoint"`
	}{
		Policy:           scenario.policy,
		Extensions:       make([]interface{}, 0),
		CallbackEndpoint: scenario.callbackEndpoint,
	})
}
