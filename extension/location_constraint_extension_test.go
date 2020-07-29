package extension

import (
	"fmt"
	"os"
)

func ExampleLocationConstraintExtension() {
	extension, err := (&LocationConstraintExtensionBuilder{}).
		WithLatitude(51.511831).
		WithLongitude(-0.081446).
		WithRadius(0.001).
		WithUncertainty(0.001).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"LOCATION_CONSTRAINT","content":{"expected_device_location":{"latitude":51.511831,"longitude":-0.081446,"radius":0.001,"max_uncertainty_radius":0.001}}}
}
