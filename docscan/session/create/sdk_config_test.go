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
		WithIdDocumentTextExtractionCategoryRetries("test_category", 3).
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
	// Output: {"allowed_capture_methods":"CAMERA","primary_colour":"#aa1111","secondary_colour":"#bb2222","font_colour":"#ff0000","locale":"fr_FR","preset_issuing_country":"USA","success_url":"https://example.com/success","error_url":"https://example.com/error","privacy_policy_url":"https://example.com/privacy","attempts_configuration":{"id_document_text_data_extraction":{"test_category":3}}}
}

func ExampleSdkConfigBuilder_Build_repeatedCallWithIdDocumentTextExtractionCategoryRetries() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithIdDocumentTextExtractionCategoryRetries("test_category", 3).
		WithIdDocumentTextExtractionCategoryRetries("test_category", 2).
		WithIdDocumentTextExtractionCategoryRetries("test_category", 1).
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
	// Output: {"attempts_configuration":{"id_document_text_data_extraction":{"test_category":1}}}
}

func ExampleSdkConfigBuilder_Build_multipleCategoriesWithIdDocumentTextExtractionCategoryRetries() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithIdDocumentTextExtractionGenericRetries(3).
		WithIdDocumentTextExtractionCategoryRetries("test_category", 2).
		WithIdDocumentTextExtractionReclassificationRetries(1).
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
	// Output: {"attempts_configuration":{"id_document_text_data_extraction":{"GENERIC":3,"RECLASSIFICATION":1,"test_category":2}}}
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
