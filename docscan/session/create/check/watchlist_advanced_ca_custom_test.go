package check_test

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/check"
)

func ExampleRequestedWatchlistAdvancedCACheckCustomAccountBuilder_Build() {
	advancedCACustomAccountCheck, err := check.NewRequestedWatchlistAdvancedCACheckCustomAccountBuilder().
		WithAPIKey("api-key").
		WithMonitoring(true).
		WithTags(map[string]string{
			"tag_name": "value",
		}).
		WithClientRef("client-ref").
		WithMatchingStrategy(check.RequestedExactMatchingStrategy{ExactMatch: true}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := advancedCACustomAccountCheck.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"WATCHLIST_ADVANCED_CA","config":{"type":"WITH_CUSTOM_ACCOUNT","matching_strategy":{"type":"EXACT","exact_match":true},"api_key":"api-key","monitoring":true,"tags":{"tag_name":"value"},"client_ref":"client-ref"}}
}
