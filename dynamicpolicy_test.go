package yoti

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
	constraint := (&SourceConstraintBuilder{}).New().WithDrivingLicence("").Build()

	policy := (&DynamicPolicyBuilder{}).New().WithAgeOver(18, constraint).Build()

	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"date_of_birth","derivation":"age_over:18","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":false}}]}
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

func ExampleDynamicPolicyBuilder_WithFullName() {
	constraint := (&SourceConstraintBuilder{}).New().WithPassport("").Build()
	policy := (&DynamicPolicyBuilder{}).New().WithFullName(&constraint).Build()

	json, _ := policy.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"wanted":[{"name":"full_name","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASSPORT","sub_type":""}],"soft_preference":false}}]}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder() {
	policy := (&DynamicPolicyBuilder{}).New().WithFullName().
		WithPinAuth().WithWantedRememberMe().Build()
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"full_name"}],"wanted_auth_types":[2],"wanted_remember_me":true}
}
