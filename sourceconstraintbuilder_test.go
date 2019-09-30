package yoti

import (
	"fmt"
)

func ExampleSourceConstraint() {
	drivingLicence := (&WantedAnchorBuilder{}).New().WithValue("DRIVING_LICENCE").Build()
	sourceConstraint := (&SourceConstraintBuilder{}).New().
		WithAnchor(drivingLicence).
		WithSoftPreference(true).
		Build()

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println("SourceConstraint:", string(json))
	// Output: SourceConstraint: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":true}}
}

func ExampleSourceConstraintBuilder_WithPassport() {
	sourceConstraint := (&SourceConstraintBuilder{}).New().
		WithPassport("").
		Build()

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASSPORT","sub_type":""}],"soft_preference":false}}
}

func ExampleSourceConstraintBuilder_WithDrivingLicence() {
	sourceConstraint := (&SourceConstraintBuilder{}).New().
		WithDrivingLicence("").
		Build()

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"DRIVING_LICENCE","sub_type":""}],"soft_preference":false}}
}

func ExampleSourceConstraintBuilder_WithNationalID() {
	sourceConstraint := (&SourceConstraintBuilder{}).New().
		WithNationalID("").
		Build()

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"NATIONAL_ID","sub_type":""}],"soft_preference":false}}
}

func ExampleSourceConstraintBuilder_WithPasscard() {
	sourceConstraint := (&SourceConstraintBuilder{}).New().
		WithPasscard("").
		Build()

	json, _ := sourceConstraint.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"type":"SOURCE","preferred_sources":{"anchors":[{"name":"PASS_CARD","sub_type":""}],"soft_preference":false}}
}
