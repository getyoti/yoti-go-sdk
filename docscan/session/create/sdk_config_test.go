package create

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
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

func ExampleSdkConfigBuilder_WithEarlyBiometricConsentFlow() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithEarlyBiometricConsentFlow().
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
	// Output: {"biometric_consent_flow":"EARLY"}
}

func ExampleSdkConfigBuilder_WithJustInTimeBiometricConsentFlow() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithJustInTimeBiometricConsentFlow().
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
	// Output: {"biometric_consent_flow":"JUST_IN_TIME"}
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

func ExampleSdkConfigBuilder_WithAllowHandOff_false() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithAllowHandOff(false).
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
	// Output: {"allow_handoff":false}
}

func ExampleSdkConfigBuilder_Build_allowHandOff_omittedWhenNotSet() {
	sdkConfig, err := NewSdkConfigBuilder().
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
	// Output: {}
}

func ExampleSdkConfigBuilder_WithEnforceHandOff() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithEnforceHandOff(true).
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
	// Output: {"enforce_handoff":true}
}

func ExampleSdkConfigBuilder_WithEnforceHandOff_false() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithEnforceHandOff(false).
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
	// Output: {"enforce_handoff":false}
}

func ExampleSdkConfigBuilder_Build_enforceHandOff_omittedWhenNotSet() {
	sdkConfig, err := NewSdkConfigBuilder().
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
	// Output: {}
}

func ExampleSdkConfigBuilder_WithAllowHandOff_andEnforceHandOff() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithAllowHandOff(true).
		WithEnforceHandOff(true).
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
	// Output: {"allow_handoff":true,"enforce_handoff":true}
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

func ExampleSdkConfigBuilder_WithBrandId() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithBrandId("some_brand_id").
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
	// Output: {"brand_id":"some_brand_id"}
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

func ExampleSdkConfigBuilder_WithPrimaryColourDarkMode() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithPrimaryColourDarkMode("SOME_COLOUR").
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
	// Output: {"primary_colour_dark_mode":"SOME_COLOUR"}
}

func ExampleSdkConfigBuilder_WithSuppressedScreens() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithSuppressedScreens([]string{
			constants.IDDocumentEducationScreen,
			constants.IDDocumentRequirementsScreen,
			constants.SupplementaryDocumentEducationScreen,
			constants.ZoomLivenessEducationScreen,
			constants.StaticLivenessEducationScreen,
			constants.FaceCaptureEducationScreen,
			constants.FlowCompletionScreen,
		}).
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
	// Output: {"suppressed_screens":["ID_DOCUMENT_EDUCATION","ID_DOCUMENT_REQUIREMENTS","SUPPLEMENTARY_DOCUMENT_EDUCATION","ZOOM_LIVENESS_EDUCATION","STATIC_LIVENESS_EDUCATION","FACE_CAPTURE_EDUCATION","FLOW_COMPLETION"]}
}

func ExampleSdkConfigBuilder_WithSuppressedScreens_replacesPreviousValue() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithSuppressedScreens([]string{constants.FlowCompletionScreen}).
		WithSuppressedScreens([]string{constants.IDDocumentEducationScreen}).
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
	// Output: {"suppressed_screens":["ID_DOCUMENT_EDUCATION"]}
}

func ExampleSdkConfigBuilder_WithSuppressedScreen() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithSuppressedScreen(constants.FlowCompletionScreen).
		WithSuppressedScreen(constants.IDDocumentEducationScreen).
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
	// Output: {"suppressed_screens":["FLOW_COMPLETION","ID_DOCUMENT_EDUCATION"]}
}

func ExampleSdkConfigBuilder_Build_suppressedScreensOmittedWhenNotSet() {
	sdkConfig, err := NewSdkConfigBuilder().
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
	// Output: {}
}
