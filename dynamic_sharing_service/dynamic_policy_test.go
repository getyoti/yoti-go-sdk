package dynamic_sharing_service

import (
	"fmt"
)

func ExampleDynamicPolicyBuilder_WithFamilyName() {
	policy, err := (&DynamicPolicyBuilder{}).New().WithFamilyName().Build()
	if err != nil {
		return
	}
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"family_name"}
}

func ExampleDynamicPolicyBuilder_WithSelfie() {
	policy, err := (&DynamicPolicyBuilder{}).New().WithSelfie().Build()
	if err != nil {
		return
	}
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"selfie"}
}

func ExampleDynamicPolicyBuilder_WithAgeOver() {
	constraint, err := (&SourceConstraintBuilder{}).New().WithDrivingLicence("").Build()
	if err != nil {
		return
	}

	policy, err := (&DynamicPolicyBuilder{}).New().WithAgeOver(18, constraint).Build()
	if err != nil {
		return
	}

	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"date_of_birth","derivation":"age_over:18","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":false}}]}
}

func ExampleDynamicPolicyBuilder_WithSelfieAuth() {
	policy, err := (&DynamicPolicyBuilder{}).New().WithSelfieAuth().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[1],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithWantedRememberMe() {
	policy, err := (&DynamicPolicyBuilder{}).New().WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[],"wanted_remember_me":true}
}

func ExampleDynamicPolicyBuilder_WithFullName() {
	constraint, err := (&SourceConstraintBuilder{}).New().WithPassport("").Build()
	if err != nil {
		return
	}

	policy, err := (&DynamicPolicyBuilder{}).New().WithFullName(&constraint).Build()
	if err != nil {
		return
	}

	json, _ := policy.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"wanted":[{"name":"full_name","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASSPORT","sub_type":""}],"soft_preference":false}}]}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder() {
	policy, err := (&DynamicPolicyBuilder{}).New().WithFullName().
		WithPinAuth().WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"full_name"}],"wanted_auth_types":[2],"wanted_remember_me":true}
}
