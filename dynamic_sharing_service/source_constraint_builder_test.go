package dynamic_sharing_service

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func ExampleSourceConstraint() {
	drivingLicence, err := (&WantedAnchorBuilder{}).WithValue("DRIVING_LICENCE").Build()
	if err != nil {
		return
	}
	sourceConstraint, err := (&SourceConstraintBuilder{}).
		WithAnchor(drivingLicence).
		WithSoftPreference(true).
		Build()
	if err != nil {
		return
	}

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println("SourceConstraint:", string(json))
	// Output: SourceConstraint: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":true}}
}

func ExampleSourceConstraintBuilder_WithPassport() {
	sourceConstraint, err := (&SourceConstraintBuilder{}).
		WithPassport("").
		Build()
	if err != nil {
		return
	}

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASSPORT","sub_type":""}],"soft_preference":false}}
}

func ExampleSourceConstraintBuilder_WithDrivingLicence() {
	sourceConstraint, err := (&SourceConstraintBuilder{}).
		WithDrivingLicence("").
		Build()
	if err != nil {
		return
	}

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":false}}
}

func ExampleSourceConstraintBuilder_WithNationalID() {
	sourceConstraint, err := (&SourceConstraintBuilder{}).
		WithNationalID("").
		Build()
	if err != nil {
		return
	}

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"NATIONAL_ID","sub_type":""}],"soft_preference":false}}
}

func ExampleSourceConstraintBuilder_WithPasscard() {
	sourceConstraint, err := (&SourceConstraintBuilder{}).
		WithPasscard("").
		Build()
	if err != nil {
		return
	}

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASS_CARD","sub_type":""}],"soft_preference":false}}
}

func TestSourceConstraint_isConstraintImplemented(t *testing.T) {
	constraint := &SourceConstraint{}
	assert.Check(t, constraint.isConstraint())
}
