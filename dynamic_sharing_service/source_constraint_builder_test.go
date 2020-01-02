package dynamic_sharing_service

import (
	"fmt"
)

func ExampleSourceConstraint() {
	drivingLicence, err := (&WantedAnchorBuilder{}).New().WithValue("DRIVING_LICENCE").Build()
	if err != nil {
		return
	}
	sourceConstraint, err := (&SourceConstraintBuilder{}).New().
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
	sourceConstraint, err := (&SourceConstraintBuilder{}).New().
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
	sourceConstraint, err := (&SourceConstraintBuilder{}).New().
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
	sourceConstraint, err := (&SourceConstraintBuilder{}).New().
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
	sourceConstraint, err := (&SourceConstraintBuilder{}).New().
		WithPasscard("").
		Build()
	if err != nil {
		return
	}

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASS_CARD","sub_type":""}],"soft_preference":false}}
}
