package extension

import (
	"fmt"
)

func ExampleLocationConstraintExtension() {
	extension := (&LocationConstraintExtensionBuilder{}).New().
		WithLatitude(51.511831).
		WithLongtitude(-0.081446).
		WithRadius(0.001).
		WithUncertainty(0.001).
		Build()

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"LOCATION_CONSTRAINT","content":{"latitude":51.511831,"longtitude":-0.081446,"radius":0.001,"max_uncertainty_radius":0.001}}
}
