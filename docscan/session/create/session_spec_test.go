package create

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/check"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/filter"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/task"
)

func ExampleSessionSpecificationBuilder_Build() {
	notifications, err := NewNotificationConfigBuilder().
		WithTopic("some-topic").
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	faceMatchCheck, err := check.NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	documentAuthenticityCheck, err := check.NewRequestedDocumentAuthenticityCheckBuilder().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	livenessCheck, err := check.NewRequestedLivenessCheckBuilder().
		ForLivenessType("LIVENESSTYPE").
		WithMaxRetries(5).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	textExtractionTask, err := task.NewRequestedTextExtractionTaskBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	sdkConfig, err := NewSdkConfigBuilder().
		WithAllowsCamera().
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	requiredIDDocument, err := filter.NewRequiredIDDocumentBuilder().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(789).
		WithResourcesTTL(456).
		WithUserTrackingID("some-tracking-id").
		WithNotifications(notifications).
		WithRequestedCheck(faceMatchCheck).
		WithRequestedCheck(documentAuthenticityCheck).
		WithRequestedCheck(livenessCheck).
		WithRequestedTask(textExtractionTask).
		WithSDKConfig(sdkConfig).
		WithRequiredDocument(requiredIDDocument).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"client_session_token_ttl":789,"resources_ttl":456,"user_tracking_id":"some-tracking-id","notifications":{"topics":["some-topic"]},"requested_checks":[{"type":"ID_DOCUMENT_FACE_MATCH","config":{"manual_check":"NEVER"}},{"type":"ID_DOCUMENT_AUTHENTICITY","config":{}},{"type":"LIVENESS","config":{"max_retries":5,"liveness_type":"LIVENESSTYPE"}}],"requested_tasks":[{"type":"ID_DOCUMENT_TEXT_DATA_EXTRACTION","config":{"manual_check":"FALLBACK"}}],"sdk_config":{"allowed_capture_methods":"CAMERA"},"required_documents":[{"type":"ID_DOCUMENT"}]}
}

func ExampleSessionSpecificationBuilder_Build_withBlockBiometricConsentTrue() {
	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithBlockBiometricConsent(true).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"block_biometric_consent":true}
}

func ExampleSessionSpecificationBuilder_Build_withBlockBiometricConsentFalse() {
	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithBlockBiometricConsent(false).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"block_biometric_consent":false}
}

func ExampleSessionSpecificationBuilder_WithRequiredDocument_supplementary() {
	requiredSupplementaryDocument, err := filter.NewRequiredSupplementaryDocumentBuilder().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithRequiredDocument(requiredSupplementaryDocument).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"required_documents":[{"type":"SUPPLEMENTARY_DOCUMENT"}]}
}

func ExampleSessionSpecificationBuilder_Build_withIdentityProfileRequirements() {
	identityProfile := []byte(`{
		"trust_framework": "UK_TFIDA",
		"scheme": {
			"type":      "DBS",
			"objective": "STANDARD"
		}
	}`)

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithIdentityProfileRequirements(identityProfile).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"identity_profile_requirements":{"trust_framework":"UK_TFIDA","scheme":{"type":"DBS","objective":"STANDARD"}}}
}

func TestExampleSessionSpecificationBuilder_Build_WithIdentityProfileRequirements_InvalidJSON(t *testing.T) {
	identityProfile := []byte(`{
		"trust_framework": UK_TFIDA",
		,
	}`)

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithIdentityProfileRequirements(identityProfile).
		Build()

	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}

	_, err = json.Marshal(sessionSpecification)
	if err == nil {
		t.Error("expected an error")
		return
	}
	var marshallerErr *json.MarshalerError
	if !errors.As(err, &marshallerErr) {
		t.Errorf("wanted err to be of type '%v', got: '%v'", reflect.TypeOf(marshallerErr), reflect.TypeOf(err))
	}
}

func ExampleSessionSpecificationBuilder_Build_withAdvancedIdentityProfileRequirements() {
	advancedIdentityProfile := []byte(`{
		"profiles": [
			{
				"trust_framework": "UK_TFIDA",
				"schemes": [
					{
						"label": "LB912",
						"type": "RTW"
					},
					{
						"label": "LB777",
						"type": "DBS",
						"objective": "BASIC"
					}
				]
			},
			{
				"trust_framework": "YOTI_GLOBAL",
				"schemes": [
					{
						"label": "LB321",
						"type": "IDENTITY",
						"objective": "AL_L1",
						"config": {}
					}
				]
			}
		]
	}`)

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithAdvancedIdentityProfileRequirements(advancedIdentityProfile).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"advanced_identity_profile_requirements":{"profiles":[{"trust_framework":"UK_TFIDA","schemes":[{"label":"LB912","type":"RTW"},{"label":"LB777","type":"DBS","objective":"BASIC"}]},{"trust_framework":"YOTI_GLOBAL","schemes":[{"label":"LB321","type":"IDENTITY","objective":"AL_L1","config":{}}]}]}}
}

func TestExampleSessionSpecificationBuilder_Build_WithAdvancedIdentityProfileRequirements_InvalidJSON(t *testing.T) {
	advancedIdentityProfile := []byte(`{
		"trust_framework": UK_TFIDA",
		,
	}`)

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithAdvancedIdentityProfileRequirements(advancedIdentityProfile).
		Build()

	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}

	_, err = json.Marshal(sessionSpecification)
	if err == nil {
		t.Error("expected an error")
		return
	}
	var marshallerErr *json.MarshalerError
	if !errors.As(err, &marshallerErr) {
		t.Errorf("wanted err to be of type '%v', got: '%v'", reflect.TypeOf(marshallerErr), reflect.TypeOf(err))
	}
}

func ExampleSessionSpecificationBuilder_Build_withSubject() {
	subject := []byte(`{
		"subject_id": "Original subject ID"
	}`)

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithSubject(subject).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"subject":{"subject_id":"Original subject ID"}}
}

func TestExampleSessionSpecificationBuilder_Build_WithSubject_InvalidJSON(t *testing.T) {
	subject := []byte(`{
		"Original subject ID"
	}`)

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithSubject(subject).
		Build()

	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}

	_, err = json.Marshal(sessionSpecification)
	if err == nil {
		t.Error("expected an error")
		return
	}
	var marshallerErr *json.MarshalerError
	if !errors.As(err, &marshallerErr) {
		t.Errorf("wanted err to be of type '%v', got: '%v'", reflect.TypeOf(marshallerErr), reflect.TypeOf(err))
	}
}

func ExampleSessionSpecificationBuilder_Build_withCreateIdentityProfilePreview() {

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithCreateIdentityProfilePreview(true).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"create_identity_profile_preview":true}
}
