package create

import "time"

type ImportToken struct {
	Ttl int `json:"ttl"`
}

// defaultImportTokenTTL is 12 Months
const defaultImportTokenTTL = time.Hour * 24 * 365

// NewImportTokenBuilder creates a new ImportTokenBuilder
func NewImportTokenBuilder() *ImportTokenBuilder {
	return &ImportTokenBuilder{
		Ttl: int(defaultImportTokenTTL.Seconds()),
	}
}

// ImportTokenBuilder builds the ImportToken struct
type ImportTokenBuilder struct {
	Ttl int
}

// WithTTL sets the TTL of the import-token (in seconds)
func (b *ImportTokenBuilder) WithTTL(ttl int) *ImportTokenBuilder {
	b.Ttl = ttl
	return b
}

// Build builds the ImportToken struct using the supplied values
func (b *ImportTokenBuilder) Build() (*ImportToken, error) {
	return &ImportToken{
		Ttl: b.Ttl,
	}, nil
}
