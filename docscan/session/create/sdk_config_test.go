package create

import (
	"encoding/json"
	"fmt"
)

func ExampleSdkConfigBuilder_Build() {
	sdkConfig, err := NewSdkConfigBuilder().
		WithAllowsCamera().
		WithErrorUrl("https://errorurl.com").
		WithFontColour("#ff0000").
		WithLocale("fr_FR").
		WithPresetIssuingCountry("USA").
		WithPrimaryColour("#aa1111").
		WithSecondaryColour("#bb2222").
		WithSuccessUrl("https://successurl.com").
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(sdkConfig)
	fmt.Println(string(data))
	// Output: {"allowed_capture_methods":"CAMERA","primary_colour":"#aa1111","secondary_colour":"#bb2222","font_colour":"#ff0000","locale":"fr_FR","preset_issuing_country":"USA","success_url":"https://successurl.com","error_url":"https://errorurl.com"}
}
