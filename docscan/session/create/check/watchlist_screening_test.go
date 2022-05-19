package check

import (
	"encoding/json"
	"fmt"
)

func ExampleNewRequestedWatchlistScreeningCheckBuilder() {
	watchlistScreeningCheck, err := NewRequestedWatchlistScreeningCheckBuilder().
		WithAdverseMediaCategory().
		WithSanctionsCategory().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(watchlistScreeningCheck)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"WATCHLIST_SCREENING","config":{"categories":["ADVERSE-MEDIA","SANCTIONS"]}}
}
