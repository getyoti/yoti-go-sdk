package filter

// TypeRestriction is a restriction of the type of document required
type TypeRestriction struct {
	Inclusion     string   `json:"inclusion"`
	DocumentTypes []string `json:"document_types"`
}

// CountryRestriction is a restriction of the country in which a document pertains to
type CountryRestriction struct {
	Inclusion    string   `json:"inclusion"`
	CountryCodes []string `json:"country_codes"`
}

// DigitalIDProvider is a digital identity provider allowed to satisfy a required document
type DigitalIDProvider struct {
	Name string `json:"name"`
}
