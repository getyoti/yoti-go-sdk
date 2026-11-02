package retrieve

// CompanyProfileResponse contains company profile information returned in the session response
type CompanyProfileResponse struct {
	// CompanyName is the name of the company
	CompanyName string `json:"company_name"`
}

// GetCompanyName returns the company name
func (c *CompanyProfileResponse) GetCompanyName() string {
	return c.CompanyName
}
