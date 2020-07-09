package dynamic_sharing_service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/consts"
	"github.com/getyoti/yoti-go-sdk/v2/yotierror"
	"gotest.tools/v3/assert"
)

func ExampleDynamicPolicyBuilder_WithFamilyName() {
	policy, err := (&DynamicPolicyBuilder{}).WithFamilyName().Build()
	if err != nil {
		return
	}
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"family_name"}
}

func ExampleDynamicPolicyBuilder_WithSelfie() {
	policy, err := (&DynamicPolicyBuilder{}).WithSelfie().Build()
	if err != nil {
		return
	}
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"selfie"}
}

func ExampleDynamicPolicyBuilder_WithAgeOver() {
	constraint, err := (&SourceConstraintBuilder{}).WithDrivingLicence("").Build()
	if err != nil {
		return
	}

	policy, err := (&DynamicPolicyBuilder{}).WithAgeOver(18, constraint).Build()
	if err != nil {
		return
	}

	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"date_of_birth","derivation":"age_over:18","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":false}}]}
}

func ExampleDynamicPolicyBuilder_WithSelfieAuth() {
	policy, err := (&DynamicPolicyBuilder{}).WithSelfieAuth().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[1],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithWantedRememberMe() {
	policy, err := (&DynamicPolicyBuilder{}).WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[],"wanted_remember_me":true}
}

func ExampleDynamicPolicyBuilder_WithFullName() {
	constraint, err := (&SourceConstraintBuilder{}).WithPassport("").Build()
	if err != nil {
		return
	}

	policy, err := (&DynamicPolicyBuilder{}).WithFullName(&constraint).Build()
	if err != nil {
		return
	}

	json, _ := policy.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"wanted":[{"name":"full_name","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASSPORT","sub_type":""}],"soft_preference":false}}]}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder() {
	policy, err := (&DynamicPolicyBuilder{}).WithFullName().
		WithPinAuth().WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"full_name"}],"wanted_auth_types":[2],"wanted_remember_me":true}
}

func ExampleDynamicPolicyBuilder_WithAgeUnder() {
	policy, err := (&DynamicPolicyBuilder{}).WithAgeUnder(18).Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"date_of_birth","derivation":"age_under:18"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithGivenNames() {
	policy, err := (&DynamicPolicyBuilder{}).WithGivenNames().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"given_names"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithDateOfBirth() {
	policy, err := (&DynamicPolicyBuilder{}).WithDateOfBirth().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"date_of_birth"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithGender() {
	policy, err := (&DynamicPolicyBuilder{}).WithGender().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"gender"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithPostalAddress() {
	policy, err := (&DynamicPolicyBuilder{}).WithPostalAddress().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"postal_address"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithStructuredPostalAddress() {
	policy, err := (&DynamicPolicyBuilder{}).WithStructuredPostalAddress().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"structured_postal_address"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithNationality() {
	policy, err := (&DynamicPolicyBuilder{}).WithNationality().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"nationality"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExampleDynamicPolicyBuilder_WithPhoneNumber() {
	policy, err := (&DynamicPolicyBuilder{}).WithPhoneNumber().Build()
	if err != nil {
		return
	}
	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"phone_number"}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func TestDynamicPolicyBuilder_WithWantedAttributeByName_WithSourceConstraint(t *testing.T) {
	attributeName := "attributeName"
	builder := (&DynamicPolicyBuilder{})
	sourceConstraint, err := (&SourceConstraintBuilder{}).Build()
	assert.NilError(t, err)

	builder.WithWantedAttributeByName(
		attributeName,
		sourceConstraint,
	)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, attributeName)
	assert.Equal(t, len(policy.attributes[0].constraints), 1)
}

func TestDynamicPolicyBuilder_WithWantedAttributeByName_InvalidOptionsShouldPanic(t *testing.T) {
	attributeName := "attributeName"
	builder := (&DynamicPolicyBuilder{})
	invalidOption := "invalidOption"

	defer func() {
		r := recover().(string)
		assert.Check(t, strings.Contains(r, "Not a valid option type"))
	}()

	builder.WithWantedAttributeByName(
		attributeName,
		invalidOption,
	)

	t.Error("Expected Panic")

}

func TestDynamicPolicyBuilder_WithWantedAttributeByName_ShouldPropagateErrors(t *testing.T) {
	builder := (&DynamicPolicyBuilder{})

	builder.WithWantedAttributeByName("")
	builder.WithWantedAttributeByName("")

	_, err := builder.Build()

	assert.Error(t, err, "Wanted attribute names must not be empty, Wanted attribute names must not be empty")
	assert.Error(t, err.(yotierror.MultiError).Unwrap(), "Wanted attribute names must not be empty")
}

func TestDynamicPolicyBuilder_WithAgeDerivedAttribute_WithSourceConstraint(t *testing.T) {
	builder := (&DynamicPolicyBuilder{})
	sourceConstraint, err := (&SourceConstraintBuilder{}).Build()
	assert.NilError(t, err)

	builder.WithAgeDerivedAttribute(
		fmt.Sprintf(consts.AttrAgeOver, 18),
		sourceConstraint,
	)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrDateOfBirth)
	assert.Equal(t, len(policy.attributes[0].constraints), 1)
}

func TestDynamicPolicyBuilder_WithAgeDerivedAttribute_WithConstraintInterface(t *testing.T) {
	builder := (&DynamicPolicyBuilder{})
	var constraint constraintInterface
	sourceConstraint, err := (&SourceConstraintBuilder{}).Build()
	constraint = &sourceConstraint
	assert.NilError(t, err)

	builder.WithAgeDerivedAttribute(
		fmt.Sprintf(consts.AttrAgeOver, 18),
		constraint,
	)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrDateOfBirth)
	assert.Equal(t, len(policy.attributes[0].constraints), 1)
}

func TestDynamicPolicyBuilder_WithAgeDerivedAttribute_InvalidOptionsShouldPanic(t *testing.T) {
	builder := (&DynamicPolicyBuilder{})
	invalidOption := "invalidOption"

	defer func() {
		r := recover().(string)
		assert.Check(t, strings.Contains(r, "Not a valid option type"))
	}()

	builder.WithAgeDerivedAttribute(
		fmt.Sprintf(consts.AttrAgeOver, 18),
		invalidOption,
	)

	t.Error("Expected Panic")

}
