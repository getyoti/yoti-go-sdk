package create

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/check"
)

func ExampleSessionSpecificationBuilder_Build() {
	notifications, err := NewNotificationConfigBuilder().
		WithTopic("some-topic").
		Build()

	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	faceMatchCheck, err := check.NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()

	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(789).
		WithResourcesTtl(456).
		WithUserTrackingID("some-tracking-id").
		WithNotifications(notifications).
		WithRequestedCheck(faceMatchCheck).
		// WithRequestedTasks().
		// WithSdkConfig().
		Build()

	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(sessionSpecification)
	fmt.Println(string(data))
	// Output: {"client_session_token_ttl":789,"resources_ttl":456,"user_tracking_id":"some-tracking-id","notifications":{"topics":["some-topic"]},"requested_checks":[{"type":"ID_DOCUMENT_FACE_MATCH","config":{"manual_check":"NEVER"}}]}
}
