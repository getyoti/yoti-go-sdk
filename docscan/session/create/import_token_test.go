package create

import (
	"encoding/json"
	"fmt"
	"time"
)

func ExampleImportTokenBuilder_Build() {
	ttl := time.Hour * 24 * 30
	it, err := NewImportTokenBuilder().
		WithTTL(int(ttl.Seconds())).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(it)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"ttl":2592000}
}
