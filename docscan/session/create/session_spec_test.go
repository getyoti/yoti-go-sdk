package create

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/check"
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
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(sessionSpecification)
	fmt.Println(string(data))
	// Output: {"client_session_token_ttl":789,"resources_ttl":456,"user_tracking_id":"some-tracking-id","notifications":{"topics":["some-topic"]},"requested_checks":[{"type":"ID_DOCUMENT_FACE_MATCH","config":{"manual_check":"NEVER"}},{"type":"ID_DOCUMENT_AUTHENTICITY","config":{}},{"type":"LIVENESS","config":{"max_retries":5,"liveness_type":"LIVENESSTYPE"}}],"requested_tasks":[{"type":"ID_DOCUMENT_TEXT_DATA_EXTRACTION","config":{"manual_check":"FALLBACK"}}],"sdk_config":{"allowed_capture_methods":"CAMERA"}}
}
