package policy

import (
	"fmt"
)

func ExampleWithFamilyName() {
	policy := (&DynamicPolicyBuilder{}).New().WithFamilyName().Build()
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"family_name"}
}

func ExampleWithSelfie() {
	policy := (&DynamicPolicyBuilder{}).New().WithSelfie().Build()
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"selfie"}
}

func ExampleWithOverAge() {
	policy := (&DynamicPolicyBuilder{}).New().WithAgeOver(18).Build()
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"date_of_birth","derivation":"age_over:18"}
}

func ExampleWithSelfieAuth() {
	policy := (&DynamicPolicyBuilder{}).New().WithSelfieAuth().Build()
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[1],"wanted_remember_me":false}
}

func ExampleWithRememberMeId() {
	policy := (&DynamicPolicyBuilder{}).New().WithWantedRememberMe().Build()
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[],"wanted_remember_me":true}
}

func ExampleFullPolicy() {
	policy := (&DynamicPolicyBuilder{}).New().WithSelfie().WithFullName().WithEmail().
		WithPinAuth().WithWantedRememberMe().Build()
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"selfie"},{"name":"full_name"},{"name":"email_address"}],"wanted_auth_types":[2],"wanted_remember_me":true}
}
