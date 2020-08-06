package sandbox

import (
	"fmt"
	"time"
)

func ExampleSourceAnchor() {
	time.Local = time.UTC
	source := SourceAnchor("subtype", time.Unix(1234567890, 0), "value")
	marshalled, _ := source.MarshalJSON()

	fmt.Println(string(marshalled))
	// Output: {"type":"SOURCE","value":"value","sub_type":"subtype","timestamp":1234567890000}
}

func ExampleVerifierAnchor() {
	time.Local = time.UTC
	verifier := VerifierAnchor("subtype", time.Unix(1234567890, 0), "value")
	marshalled, _ := verifier.MarshalJSON()

	fmt.Println(string(marshalled))
	// Output: {"type":"VERIFIER","value":"value","sub_type":"subtype","timestamp":1234567890000}
}
