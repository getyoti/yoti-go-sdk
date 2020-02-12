package dynamic_sharing_service

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v2/extension"
)

func ExampleDynamicScenarioBuilder() {
	scenario, err := (&DynamicScenarioBuilder{}).Build()
	if err != nil {
		return
	}
	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[],"wanted_auth_types":[],"wanted_remember_me":false},"extensions":[],"callback_endpoint":""}
}

func ExampleDynamicScenarioBuilder_WithPolicy() {
	policy, err := (&DynamicPolicyBuilder{}).WithEmail().WithPinAuth().Build()
	if err != nil {
		return
	}
	scenario, err := (&DynamicScenarioBuilder{}).WithPolicy(policy).WithCallbackEndpoint("/foo").Build()
	if err != nil {
		return
	}

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[{"name":"email_address"}],"wanted_auth_types":[2],"wanted_remember_me":false},"extensions":[],"callback_endpoint":"/foo"}
}

func ExampleDynamicScenarioBuilder_WithExtension() {
	policy, err := (&DynamicPolicyBuilder{}).WithFullName().Build()
	if err != nil {
		return
	}
	extension, err := (&extension.TransactionalFlowExtensionBuilder{}).
		WithContent("Transactional Flow Extension").
		Build()
	if err != nil {
		return
	}

	scenario, err := (&DynamicScenarioBuilder{}).WithExtension(extension).WithPolicy(policy).Build()
	if err != nil {
		return
	}

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[{"name":"full_name"}],"wanted_auth_types":[],"wanted_remember_me":false},"extensions":[{"type":"TRANSACTIONAL_FLOW","content":"Transactional Flow Extension"}],"callback_endpoint":""}

}