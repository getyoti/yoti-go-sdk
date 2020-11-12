package create

import (
	"encoding/json"
	"fmt"
)

func ExampleSdkConfigBuilder_Build() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithAllowsCamera().
		WithErrorUrl("https://example.com/error").
		WithFontColour("#ff0000").
		WithLocale("fr_FR").
		WithPresetIssuingCountry("USA").
		WithPrimaryColour("#aa1111").
		WithSecondaryColour("#bb2222").
		WithSuccessUrl("https://example.com/success").
		WithPrivacyPolicyUrl("https://example.com/privacy").
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sdkConfig)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"allowed_capture_methods":"CAMERA","primary_colour":"#aa1111","secondary_colour":"#bb2222","font_colour":"#ff0000","locale":"fr_FR","preset_issuing_country":"USA","success_url":"https://example.com/success","error_url":"https://example.com/error","privacy_policy_url":"https://example.com/privacy"}
}

func ExampleSdkConfigBuilder_WithAllowsCameraAndUpload() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithAllowsCameraAndUpload().
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sdkConfig)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"allowed_capture_methods":"CAMERA_AND_UPLOAD"}
}
