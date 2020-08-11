package create

import (
	"encoding/json"
	"fmt"
	"os"
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
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(notifications)
	fmt.Println(string(data))
	// Output: {"auth_token":"auth-token","endpoint":"/endpoint","topics":["SOME_TOPIC","CHECK_COMPLETION","RESOURCE_UPDATE","SESSION_COMPLETION","TASK_COMPLETION"]}
}
