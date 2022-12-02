package create

import "github.com/getyoti/yoti-go-sdk/v3/docscan/constants"

// SDKConfig provides configuration properties for the the web/native clients
type SDKConfig struct {
	AllowedCaptureMethods string                 `json:"allowed_capture_methods,omitempty"`
	PrimaryColour         string                 `json:"primary_colour,omitempty"`
	SecondaryColour       string                 `json:"secondary_colour,omitempty"`
	FontColour            string                 `json:"font_colour,omitempty"`
	Locale                string                 `json:"locale,omitempty"`
	PresetIssuingCountry  string                 `json:"preset_issuing_country,omitempty"`
	SuccessUrl            string                 `json:"success_url,omitempty"`
	ErrorUrl              string                 `json:"error_url,omitempty"`
	PrivacyPolicyUrl      string                 `json:"privacy_policy_url,omitempty"`
	AttemptsConfiguration *AttemptsConfiguration `json:"attempts_configuration,omitempty"`
}

type AttemptsConfiguration struct {
	IdDocumentTextDataExtraction map[string]int `json:"ID_DOCUMENT_TEXT_DATA_EXTRACTION,omitempty"`
}

// NewSdkConfigBuilder creates a new SdkConfigBuilder
func NewSdkConfigBuilder() *SdkConfigBuilder {
	return &SdkConfigBuilder{}
}

// SdkConfigBuilder builds the SDKConfig struct
type SdkConfigBuilder struct {
	allowedCaptureMethods                string
	primaryColour                        string
	secondaryColour                      string
	fontColour                           string
	locale                               string
	presetIssuingCountry                 string
	successUrl                           string
	errorUrl                             string
	privacyPolicyUrl                     string
	idDocumentTextDataExtractionAttempts map[string]int
}

// WithAllowedCaptureMethods sets the allowed capture methods on the builder
func (b *SdkConfigBuilder) WithAllowedCaptureMethods(captureMethods string) *SdkConfigBuilder {
	b.allowedCaptureMethods = captureMethods
	return b
}

// WithAllowsCamera sets the allowed capture method to "CAMERA"
func (b *SdkConfigBuilder) WithAllowsCamera() *SdkConfigBuilder {
	return b.WithAllowedCaptureMethods(constants.Camera)
}

// WithAllowsCameraAndUpload sets the allowed capture method to "CAMERA_AND_UPLOAD"
func (b *SdkConfigBuilder) WithAllowsCameraAndUpload() *SdkConfigBuilder {
	return b.WithAllowedCaptureMethods(constants.CameraAndUpload)
}

// WithPrimaryColour sets the primary colour to be used by the web/native client, hexadecimal value e.g. #ff0000
func (b *SdkConfigBuilder) WithPrimaryColour(colour string) *SdkConfigBuilder {
	b.primaryColour = colour
	return b
}

// WithSecondaryColour sets the secondary colour to be used by the web/native client (used on the button), hexadecimal value e.g. #ff0000
func (b *SdkConfigBuilder) WithSecondaryColour(colour string) *SdkConfigBuilder {
	b.secondaryColour = colour
	return b
}

// WithFontColour the font colour to be used by the web/native client (used on the button), hexadecimal value e.g. #ff0000
func (b *SdkConfigBuilder) WithFontColour(colour string) *SdkConfigBuilder {
	b.fontColour = colour
	return b
}

// WithLocale sets the language locale use by the web/native client
func (b *SdkConfigBuilder) WithLocale(locale string) *SdkConfigBuilder {
	b.locale = locale
	return b
}

// WithPresetIssuingCountry sets the preset issuing country used by the web/native client
func (b *SdkConfigBuilder) WithPresetIssuingCountry(country string) *SdkConfigBuilder {
	b.presetIssuingCountry = country
	return b
}

// WithSuccessUrl sets the success URL for the redirect that follows the web/native client uploading documents successfully
func (b *SdkConfigBuilder) WithSuccessUrl(url string) *SdkConfigBuilder {
	b.successUrl = url
	return b
}

// WithErrorUrl sets the error URL for the redirect that follows the web/native client uploading documents unsuccessfully
func (b *SdkConfigBuilder) WithErrorUrl(url string) *SdkConfigBuilder {
	b.errorUrl = url
	return b
}

// WithPrivacyPolicyUrl sets the privacy policy URL
func (b *SdkConfigBuilder) WithPrivacyPolicyUrl(url string) *SdkConfigBuilder {
	b.privacyPolicyUrl = url
	return b
}

func (b *SdkConfigBuilder) WithIdDocumentTextExtractionCategoryAttempts(category string, attempts int) *SdkConfigBuilder {
	if b.idDocumentTextDataExtractionAttempts == nil {
		b.idDocumentTextDataExtractionAttempts = make(map[string]int)
	}
	b.idDocumentTextDataExtractionAttempts[category] = attempts
	return b
}

func (b *SdkConfigBuilder) WithIdDocumentTextExtractionReclassificationAttempts(attempts int) *SdkConfigBuilder {
	return b.WithIdDocumentTextExtractionCategoryAttempts(reclassification, attempts)
}

func (b *SdkConfigBuilder) WithIdDocumentTextExtractionGenericAttempts(attempts int) *SdkConfigBuilder {
	return b.WithIdDocumentTextExtractionCategoryAttempts(generic, attempts)
}

// Build builds the SDKConfig struct using the supplied values
func (b *SdkConfigBuilder) Build() (*SDKConfig, error) {
	sdkConf := &SDKConfig{
		b.allowedCaptureMethods,
		b.primaryColour,
		b.secondaryColour,
		b.fontColour,
		b.locale,
		b.presetIssuingCountry,
		b.successUrl,
		b.errorUrl,
		b.privacyPolicyUrl,
		nil,
	}

	if b.idDocumentTextDataExtractionAttempts != nil {
		sdkConf.AttemptsConfiguration = &AttemptsConfiguration{
			IdDocumentTextDataExtraction: b.idDocumentTextDataExtractionAttempts,
		}
	}
	return sdkConf, nil
}
