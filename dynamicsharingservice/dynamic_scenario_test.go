package dynamicsharingservice

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v2/dynamicsharingservice/policy"
)

func ExampleDynamicScenarioBuilder() {
	scenario := (&DynamicScenarioBuilder{}).New().Build()
	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[],"wanted_auth_types":[],"wanted_remember_me":false},"extensions":[],"callback_endpoint":""}
}

func ExampleDynamicScenarioBuilder_WithPolicy() {
	policy := (&policy.DynamicPolicyBuilder{}).New().WithEmail().WithPinAuth().Build()
	scenario := (&DynamicScenarioBuilder{}).New().WithPolicy(policy).WithCallbackEndpoint("/foo").Build()

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[{"name":"email_address"}],"wanted_auth_types":[2],"wanted_remember_me":false},"extensions":[],"callback_endpoint":"/foo"}
}
