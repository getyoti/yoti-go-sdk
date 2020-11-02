package create

import (
	"encoding/json"
	"fmt"
)

func ExampleNotificationConfigBuilder_Build() {
	notifications, err := NewNotificationConfigBuilder().
		WithAuthToken("auth-token").
		WithEndpoint("/endpoint").
		WithTopic("SOME_TOPIC").
		ForCheckCompletion().
		ForResourceUpdate().
		ForSessionCompletion().
		ForTaskCompletion().
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(notifications)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"auth_token":"auth-token","endpoint":"/endpoint","topics":["SOME_TOPIC","CHECK_COMPLETION","RESOURCE_UPDATE","SESSION_COMPLETION","TASK_COMPLETION"]}
}
