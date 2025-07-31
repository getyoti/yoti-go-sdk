package digitalidentity

import (
	"fmt"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"gotest.tools/v3/assert"
)

func ExamplePolicyBuilder_WithEstimatedAge() {
	policy, err := (&PolicyBuilder{}).WithEstimatedAge().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := policy.attributes[0].MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"name":"estimated_age","accept_self_asserted":false,"alternative_names":["date_of_birth"]}
}

func ExamplePolicyBuilder_WithEstimatedAgeOver() {
	policy, err := (&PolicyBuilder{}).WithEstimatedAgeOverSimple(18).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := policy.attributes[0].MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"name":"estimated_age","derivation":"age_over:18","accept_self_asserted":false,"alternative_names":["date_of_birth"]}
}

func ExamplePolicyBuilder_WithEstimatedAgeUnder() {
	policy, err := (&PolicyBuilder{}).WithEstimatedAgeUnderSimple(21).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := policy.attributes[0].MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"name":"estimated_age","derivation":"age_under:21","accept_self_asserted":false,"alternative_names":["date_of_birth"]}
}

func TestPolicyBuilder_WithEstimatedAge_SetsCorrectAttributes(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.WithEstimatedAge()

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, len(policy.attributes[0].alternativeNames), 1)
	assert.Equal(t, policy.attributes[0].alternativeNames[0], consts.AttrDateOfBirth)
	assert.Equal(t, policy.attributes[0].derivation, "")
	assert.Equal(t, policy.attributes[0].acceptSelfAsserted, false)
	assert.Equal(t, len(policy.attributes[0].constraints), 0)
}

func TestPolicyBuilder_WithEstimatedAgeOver_SetsCorrectAttributes(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.WithEstimatedAgeOverSimple(18)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, len(policy.attributes[0].alternativeNames), 1)
	assert.Equal(t, policy.attributes[0].alternativeNames[0], consts.AttrDateOfBirth)
	assert.Equal(t, policy.attributes[0].derivation, "age_over:18")
	assert.Equal(t, policy.attributes[0].acceptSelfAsserted, false)
	assert.Equal(t, len(policy.attributes[0].constraints), 0)
}

func TestPolicyBuilder_WithEstimatedAgeUnder_SetsCorrectAttributes(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.WithEstimatedAgeUnderSimple(21)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, len(policy.attributes[0].alternativeNames), 1)
	assert.Equal(t, policy.attributes[0].alternativeNames[0], consts.AttrDateOfBirth)
	assert.Equal(t, policy.attributes[0].derivation, "age_under:21")
	assert.Equal(t, policy.attributes[0].acceptSelfAsserted, false)
	assert.Equal(t, len(policy.attributes[0].constraints), 0)
}

func TestPolicyBuilder_WithEstimatedAgeOver_WithSourceConstraint(t *testing.T) {
	builder := &PolicyBuilder{}
	sourceConstraint, err := (&SourceConstraintBuilder{}).Build()
	assert.NilError(t, err)

	builder.WithEstimatedAgeOverSimple(18, &sourceConstraint)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, policy.attributes[0].derivation, "age_over:18")
	assert.Equal(t, len(policy.attributes[0].constraints), 1)
}

func TestPolicyBuilder_WithEstimatedAge_WithSourceConstraint(t *testing.T) {
	builder := &PolicyBuilder{}
	sourceConstraint, err := (&SourceConstraintBuilder{}).Build()
	assert.NilError(t, err)

	builder.WithEstimatedAge(&sourceConstraint)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, len(policy.attributes[0].constraints), 1)
}

func TestPolicyBuilder_CombiningEstimatedAgeWithOtherAttributes(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.WithFullName().
		WithEstimatedAgeOverSimple(18).
		WithEmail()

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 3)

	// Find the estimated age attribute
	var estimatedAgeAttr WantedAttribute
	for _, attr := range policy.attributes {
		if attr.name == consts.AttrEstimatedAge {
			estimatedAgeAttr = attr
			break
		}
	}

	assert.Equal(t, estimatedAgeAttr.name, consts.AttrEstimatedAge)
	assert.Equal(t, estimatedAgeAttr.derivation, "age_over:18")
	assert.Equal(t, len(estimatedAgeAttr.alternativeNames), 1)
	assert.Equal(t, estimatedAgeAttr.alternativeNames[0], consts.AttrDateOfBirth)
}

func TestPolicyBuilderWithEstimatedAgeOverWithBuffer(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.WithEstimatedAgeOver(18, 5)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, policy.attributes[0].derivation, "age_over:18:5")
	assert.Equal(t, len(policy.attributes[0].alternativeNames), 1)
	assert.Equal(t, policy.attributes[0].alternativeNames[0], consts.AttrDateOfBirth)
}

func TestPolicyBuilderWithEstimatedAgeUnderWithBuffer(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.WithEstimatedAgeUnder(21, 5)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, policy.attributes[0].derivation, "age_under:21:5")
	assert.Equal(t, len(policy.attributes[0].alternativeNames), 1)
	assert.Equal(t, policy.attributes[0].alternativeNames[0], consts.AttrDateOfBirth)
}
