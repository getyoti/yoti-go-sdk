package digitalidentity

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateShareQrCodeBuilder_Build(t *testing.T) {
	expectedID := "SOME_ID"
	expectedURI := "https://example.com/qr"

	builder := CreateShareQrCodeBuilder{
		createShareQrCode: CreateShareQrCodeResult{
			id:  expectedID,
			uri: expectedURI,
		},
	}

	result, err := builder.Build()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result.id != expectedID {
		t.Errorf("Expected ID: %s, but got: %s", expectedID, result.id)
	}
	if result.uri != expectedURI {
		t.Errorf("Expected URI: %s, but got: %s", expectedURI, result.uri)
	}
}

func TestCreateShareQrCode_MarshalJSON(t *testing.T) {
	createShareQrCode := CreateShareQrCodeResult{
		id:  "SOME_ID",
		uri: "https://example.com/qr",
	}

	expectedJSON := `{"id":"SOME_ID","uri":"https://example.com/qr"}`

	resultJSON, err := json.Marshal(createShareQrCode)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if string(resultJSON) != expectedJSON {
		t.Errorf("Expected JSON:\n%s\nbut got:\n%s\n", expectedJSON, string(resultJSON))
	}
}

func ExampleCreateShareQrCodeBuilder_Build() {
	builder := CreateShareQrCodeBuilder{
		createShareQrCode: CreateShareQrCodeResult{
			id:  "SOME_ID",
			uri: "https://example.com/qr",
		},
	}

	result, err := builder.Build()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(string(data))
	// Output: {"id":"SOME_ID","uri":"https://example.com/qr"}
}
