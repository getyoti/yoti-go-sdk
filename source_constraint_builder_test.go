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
