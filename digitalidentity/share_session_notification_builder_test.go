package digitalidentity

import (
	"fmt"
)

func ExampleShareSessionNotificationBuilder() {
	shareSessionNotify, err := (&ShareSessionNotificationBuilder{}).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := shareSessionNotify.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"url":"","method":"","verifyTls":null,"headers":null}
}

func ExampleShareSessionNotificationBuilder_WithUrl() {
	shareSessionNotify, err := (&ShareSessionNotificationBuilder{}).WithUrl("Custom_Url").Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := shareSessionNotify.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"url":"Custom_Url","method":"","verifyTls":null,"headers":null}
}

func ExampleShareSessionNotificationBuilder_WithMethod() {
	shareSessionNotify, err := (&ShareSessionNotificationBuilder{}).WithMethod("CUSTOMMETHOD").Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := shareSessionNotify.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"url":"","method":"CUSTOMMETHOD","verifyTls":null,"headers":null}
}

func ExampleShareSessionNotificationBuilder_WithVerifyTls() {

	shareSessionNotify, err := (&ShareSessionNotificationBuilder{}).WithVerifyTls(true).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := shareSessionNotify.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"url":"","method":"","verifyTls":true,"headers":null}
}

func ExampleShareSessionNotificationBuilder_WithHeaders() {

	headers := make(map[string][]string)
	headers["key"] = append(headers["key"], "value")

	shareSessionNotify, err := (&ShareSessionNotificationBuilder{}).WithHeaders(headers).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := shareSessionNotify.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"url":"","method":"","verifyTls":null,"headers":{"key":["value"]}}
}
