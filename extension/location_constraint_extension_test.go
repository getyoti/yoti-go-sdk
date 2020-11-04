package extension

import (
	"fmt"
)

func ExampleLocationConstraintExtension() {
	extension, err := (&LocationConstraintExtensionBuilder{}).
		WithLatitude(51.511831).
		WithLongitude(-0.081446).
		WithRadius(0.001).
		WithMaxUncertainty(0.001).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := extension.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"LOCATION_CONSTRAINT","content":{"expected_device_location":{"latitude":51.511831,"longitude":-0.081446,"radius":0.001,"max_uncertainty_radius":0.001}}}
}
