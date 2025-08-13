package digitalidentity

import (
	"fmt"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"gotest.tools/v3/assert"
)

func ExamplePolicyBuilder_EstimatedAgeOver() {
	policy, err := (&PolicyBuilder{}).EstimatedAgeOver(18, 5).Build()
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
	// Output: {"name":"estimated_age","derivation":"age_over:18:5","accept_self_asserted":false,"alternative_names":["date_of_birth"]}
}

func TestPolicyBuilderEstimatedAgeSetsCorrectAttributes(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.EstimatedAgeOver(18, 5)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, len(policy.attributes[0].alternativeNames), 1)
	assert.Equal(t, policy.attributes[0].alternativeNames[0], consts.AttrDateOfBirth)
	assert.Equal(t, policy.attributes[0].derivation, "age_over:18:5")
	assert.Equal(t, policy.attributes[0].acceptSelfAsserted, false)
	assert.Equal(t, len(policy.attributes[0].constraints), 0)
}

func TestPolicyBuilderEstimatedAgeWithSourceConstraint(t *testing.T) {
	builder := &PolicyBuilder{}
	sourceConstraint, err := (&SourceConstraintBuilder{}).Build()
	assert.NilError(t, err)

	builder.EstimatedAgeOver(18, 5, &sourceConstraint)

	policy, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, len(policy.attributes), 1)
	assert.Equal(t, policy.attributes[0].name, consts.AttrEstimatedAge)
	assert.Equal(t, policy.attributes[0].derivation, "age_over:18:5")
	assert.Equal(t, len(policy.attributes[0].constraints), 1)
}

func TestPolicyBuilderCombiningEstimatedAgeWithOtherAttributes(t *testing.T) {
	builder := &PolicyBuilder{}
	builder.WithFullName().
		EstimatedAgeOver(18, 5).
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
	assert.Equal(t, estimatedAgeAttr.derivation, "age_over:18:5")
	assert.Equal(t, len(estimatedAgeAttr.alternativeNames), 1)
	assert.Equal(t, estimatedAgeAttr.alternativeNames[0], consts.AttrDateOfBirth)
}
