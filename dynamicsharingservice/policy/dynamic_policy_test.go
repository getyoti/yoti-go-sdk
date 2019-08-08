package policy

import (
	"fmt"
)

func ExampleDynamicPolicyBuilder_WithFamilyName() {
	policy := (&DynamicPolicyBuilder{}).New().WithFamilyName().Build()
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"family_name"}
}

func ExampleDynamicPolicyBuilder_WithSelfie() {
	policy := (&DynamicPolicyBuilder{}).New().WithSelfie().Build()
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"selfie"}
}

func ExampleDynamicPolicyBuilder_WithAgeOver() {
	policy := (&DynamicPolicyBuilder{}).New().WithAgeOver(18).Build()
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"date_of_birth","derivation":"age_over:18"}
}

func ExampleDynamicPolicyBuilder_WithSelfieAuth() {
	policy := (&DynamicPolicyBuilder{}).New().WithSelfieAuth().Build()
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[1],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithWantedRememberMe() {
	policy := (&DynamicPolicyBuilder{}).New().WithWantedRememberMe().Build()
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[],"wanted_remember_me":true}
}

func ExampleDynamicPolicyBuilder() {
	policy := (&DynamicPolicyBuilder{}).New().WithFullName().
		WithPinAuth().WithWantedRememberMe().Build()
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"full_name"}],"wanted_auth_types":[2],"wanted_remember_me":true}
}
