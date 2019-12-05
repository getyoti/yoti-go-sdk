package sandbox

import (
	"fmt"
	"time"
)

func ExampleSourceAnchor() {
	time.Local = time.UTC
	source := SourceAnchor("subtype", time.Unix(1234567890, 0), "value")
	fmt.Println(source)
	// Output: {SOURCE value subtype 2009-02-13 23:31:30 +0000 UTC}
}

func ExampleVerifierAnchor() {
	time.Local = time.UTC
	verifier := VerifierAnchor("subtype", time.Unix(1234567890, 0), "value")
	fmt.Println(verifier)
	// Output: {VERIFIER value subtype 2009-02-13 23:31:30 +0000 UTC}
}
