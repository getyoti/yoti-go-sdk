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
		WithIdDocumentTextExtractionCategoryAttempts("test_category", 3).
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
	// Output: {"allowed_capture_methods":"CAMERA","primary_colour":"#aa1111","secondary_colour":"#bb2222","font_colour":"#ff0000","locale":"fr_FR","preset_issuing_country":"USA","success_url":"https://example.com/success","error_url":"https://example.com/error","privacy_policy_url":"https://example.com/privacy","attempts_configuration":{"ID_DOCUMENT_TEXT_DATA_EXTRACTION":{"test_category":3}}}
}

func ExampleSdkConfigBuilder_Build_repeatedCallWithIdDocumentTextExtractionCategoryAttempts() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithIdDocumentTextExtractionCategoryAttempts("test_category", 3).
		WithIdDocumentTextExtractionCategoryAttempts("test_category", 2).
		WithIdDocumentTextExtractionCategoryAttempts("test_category", 1).
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
	// Output: {"attempts_configuration":{"ID_DOCUMENT_TEXT_DATA_EXTRACTION":{"test_category":1}}}
}

func ExampleSdkConfigBuilder_Build_multipleCategoriesWithIdDocumentTextExtractionCategoryAttempts() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithIdDocumentTextExtractionGenericAttempts(3).
		WithIdDocumentTextExtractionCategoryAttempts("test_category", 2).
		WithIdDocumentTextExtractionReclassificationAttempts(1).
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
	// Output: {"attempts_configuration":{"ID_DOCUMENT_TEXT_DATA_EXTRACTION":{"GENERIC":3,"RECLASSIFICATION":1,"test_category":2}}}
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

func ExampleSdkConfigBuilder_WithAllowHandOff() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithAllowHandOff(true).
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
	// Output: {"allow_handoff":true}
}

func ExampleSdkConfigBuilder_WithDarkMode() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithDarkMode("ON").
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
	// Output: {"dark_mode":"ON"}
}

func ExampleSdkConfigBuilder_WithDarkModeOff() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithDarkModeOff().
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
	// Output: {"dark_mode":"OFF"}
}

func ExampleSdkConfigBuilder_WithDarkModeAuto() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithDarkModeAuto().
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
	// Output: {"dark_mode":"AUTO"}
}
