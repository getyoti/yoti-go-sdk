package dynamic

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/extension"
)

func ExampleScenarioBuilder() {
	scenario, err := (&ScenarioBuilder{}).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[],"wanted_auth_types":[],"wanted_remember_me":false},"extensions":[],"callback_endpoint":""}
}

func ExampleScenarioBuilder_WithPolicy() {
	policy, err := (&PolicyBuilder{}).WithEmail().WithPinAuth().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	scenario, err := (&ScenarioBuilder{}).WithPolicy(policy).WithCallbackEndpoint("/foo").Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[{"name":"email_address"}],"wanted_auth_types":[2],"wanted_remember_me":false},"extensions":[],"callback_endpoint":"/foo"}
}

func ExampleScenarioBuilder_WithExtension() {
	policy, err := (&PolicyBuilder{}).WithFullName().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	builtExtension, err := (&extension.TransactionalFlowExtensionBuilder{}).
		WithContent("Transactional Flow Extension").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	scenario, err := (&ScenarioBuilder{}).WithExtension(builtExtension).WithPolicy(policy).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := scenario.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"policy":{"wanted":[{"name":"full_name"}],"wanted_auth_types":[],"wanted_remember_me":false},"extensions":[{"type":"TRANSACTIONAL_FLOW","content":"Transactional Flow Extension"}],"callback_endpoint":""}

}
