package check_test

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/check"
)

func ExampleNewRequestedWatchlistAdvancedCACheckYotiAccountBuilder() {
	advancedCAYotiAccountCheck, err := check.NewRequestedWatchlistAdvancedCACheckYotiAccountBuilder().
		WithRemoveDeceased(true).
		WithShareURL(true).
		WithSources(check.RequestedTypeListSources{
			Types: []string{"pep", "fitness-probity", "warning"}}).
		WithMatchingStrategy(check.RequestedFuzzyMatchingStrategy{Fuzziness: 0.5}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := advancedCAYotiAccountCheck.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"WATCHLIST_ADVANCED_CA","config":{"type":"WITH_YOTI_ACCOUNT","remove_deceased":true,"share_url":true,"sources":{"type":"TYPE_LIST","types":["pep","fitness-probity","warning"]},"matching_strategy":{"type":"FUZZY","fuzziness":0.5}}}
}
