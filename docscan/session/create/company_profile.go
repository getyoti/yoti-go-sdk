package create

// CompanyProfile contains the company profile information for the session
type CompanyProfile struct {
	// CompanyName is the name of the company
	CompanyName string `json:"company_name"`
}

// NewCompanyProfile creates a new CompanyProfile with the given company name
func NewCompanyProfile(companyName string) *CompanyProfile {
	return &CompanyProfile{
		CompanyName: companyName,
	}
}

// GetCompanyName returns the company name
func (c *CompanyProfile) GetCompanyName() string {
	return c.CompanyName
}
