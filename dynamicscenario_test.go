package yoti

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v2/extension"
)

func ExampleDynamicScenarioBuilder() {
	scenario := (&DynamicScenarioBuilder{}).New().Build()
	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[],"wanted_auth_types":[],"wanted_remember_me":false},"extensions":[],"callback_endpoint":""}
}

func ExampleDynamicScenarioBuilder_WithPolicy() {
	policy := (&DynamicPolicyBuilder{}).New().WithEmail().WithPinAuth().Build()
	scenario := (&DynamicScenarioBuilder{}).New().WithPolicy(policy).WithCallbackEndpoint("/foo").Build()

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[{"name":"email_address"}],"wanted_auth_types":[2],"wanted_remember_me":false},"extensions":[],"callback_endpoint":"/foo"}
}

func ExampleDynamicScenarioBuilder_WithExtension() {
	policy := (&DynamicPolicyBuilder{}).New().WithFullName().Build()
	extension := (&extension.TransactionalFlowExtensionBuilder{}).New().
		WithContent("Transactional Flow Extension").
		Build()

	scenario := (&DynamicScenarioBuilder{}).New().WithExtension(extension).WithPolicy(policy).Build()

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[{"name":"full_name"}],"wanted_auth_types":[],"wanted_remember_me":false},"extensions":[{"type":"TRANSACTIONAL_FLOW","content":"Transactional Flow Extension"}],"callback_endpoint":""}

}
