package dynamic

import (
	"fmt"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
	"gotest.tools/v3/assert"
)

func ExamplePolicyBuilder_WithFamilyName() {
	policy, err := (&PolicyBuilder{}).WithFamilyName().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"family_name","accept_self_asserted":false}
}

func ExamplePolicyBuilder_WithDocumentDetails() {
	policy, err := (&PolicyBuilder{}).WithDocumentDetails().Build()
	if err != nil {
		return
	}
	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"document_details","accept_self_asserted":false}
}

func ExamplePolicyBuilder_WithDocumentImages() {
	policy, err := (&PolicyBuilder{}).WithDocumentImages().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"document_images","accept_self_asserted":false}
}

func ExamplePolicyBuilder_WithSelfie() {
	policy, err := (&PolicyBuilder{}).WithSelfie().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"selfie","accept_self_asserted":false}
}

func ExamplePolicyBuilder_WithAgeOver() {
	constraint, err := (&SourceConstraintBuilder{}).WithDrivingLicence("").Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	policy, err := (&PolicyBuilder{}).WithAgeOver(18, constraint).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.attributes[0].MarshalJSON()
	fmt.Println(string(data))
	// Output: {"name":"date_of_birth","derivation":"age_over:18","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":false}}],"accept_self_asserted":false}
}

func ExamplePolicyBuilder_WithSelfieAuth() {
	policy, err := (&PolicyBuilder{}).WithSelfieAuth().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[1],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithWantedRememberMe() {
	policy, err := (&PolicyBuilder{}).WithWantedRememberMe().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[],"wanted_auth_types":[],"wanted_remember_me":true}
}

func ExamplePolicyBuilder_WithFullName() {
	constraint, err := (&SourceConstraintBuilder{}).WithPassport("").Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	policy, err := (&PolicyBuilder{}).WithFullName(&constraint).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	marshalledJSON, _ := policy.MarshalJSON()
	fmt.Println(string(marshalledJSON))
	// Output: {"wanted":[{"name":"full_name","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASSPORT","sub_type":""}],"soft_preference":false}}],"accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder() {
	policy, err := (&PolicyBuilder{}).WithFullName().
		WithPinAuth().WithWantedRememberMe().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"full_name","accept_self_asserted":false}],"wanted_auth_types":[2],"wanted_remember_me":true}
}

func ExamplePolicyBuilder_WithAgeUnder() {
	policy, err := (&PolicyBuilder{}).WithAgeUnder(18).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"date_of_birth","derivation":"age_under:18","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithGivenNames() {
	policy, err := (&PolicyBuilder{}).WithGivenNames().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"given_names","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithDateOfBirth() {
	policy, err := (&PolicyBuilder{}).WithDateOfBirth().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"date_of_birth","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithGender() {
	policy, err := (&PolicyBuilder{}).WithGender().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"gender","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithPostalAddress() {
	policy, err := (&PolicyBuilder{}).WithPostalAddress().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"postal_address","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithStructuredPostalAddress() {
	policy, err := (&PolicyBuilder{}).WithStructuredPostalAddress().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"structured_postal_address","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithNationality() {
	policy, err := (&PolicyBuilder{}).WithNationality().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"nationality","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func ExamplePolicyBuilder_WithPhoneNumber() {
	policy, err := (&PolicyBuilder{}).WithPhoneNumber().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := policy.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"wanted":[{"name":"phone_number","accept_self_asserted":false}],"wanted_auth_types":[],"wanted_remember_me":false}
}

func TestDynamicPolicyBuilder_WithWantedAttributeByName_WithSourceConstraint(t *testing.T) {
	attributeName := "attributeName"
	builder := &PolicyBuilder{}
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
	builder := &PolicyBuilder{}
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
	builder := &PolicyBuilder{}

	builder.WithWantedAttributeByName("")
	builder.WithWantedAttributeByName("")

	_, err := builder.Build()

	assert.Error(t, err, "Wanted attribute names must not be empty, Wanted attribute names must not be empty")
	assert.Error(t, err.(yotierror.MultiError).Unwrap(), "Wanted attribute names must not be empty")
}

func TestDynamicPolicyBuilder_WithAgeDerivedAttribute_WithSourceConstraint(t *testing.T) {
	builder := &PolicyBuilder{}
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
	builder := &PolicyBuilder{}
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
	builder := &PolicyBuilder{}
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
