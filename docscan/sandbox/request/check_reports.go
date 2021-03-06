package request

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check"
)

// CheckReports represents check reports
type CheckReports struct {
	DocumentAuthenticityChecks          []*check.DocumentAuthenticityCheck          `json:"ID_DOCUMENT_AUTHENTICITY"`
	DocumentTextDataChecks              []*check.DocumentTextDataCheck              `json:"ID_DOCUMENT_TEXT_DATA_CHECK"`
	DocumentFaceMatchChecks             []*check.DocumentFaceMatchCheck             `json:"ID_DOCUMENT_FACE_MATCH"`
	LivenessChecks                      []*check.LivenessCheck                      `json:"LIVENESS"`
	IDDocumentComparisonChecks          []*check.IDDocumentComparisonCheck          `json:"ID_DOCUMENT_COMPARISON"`
	SupplementaryDocumentTextDataChecks []*check.SupplementaryDocumentTextDataCheck `json:"SUPPLEMENTARY_DOCUMENT_TEXT_DATA_CHECK"`
	AsyncReportDelay                    uint32                                      `json:"async_report_delay,omitempty"`
	ThirdPartyIdentityCheck             *check.ThirdPartyIdentityCheck              `json:"THIRD_PARTY_IDENTITY,omitempty"`
}

// CheckReportsBuilder builds CheckReports
type CheckReportsBuilder struct {
	documentAuthenticityChecks          []*check.DocumentAuthenticityCheck
	documentTextDataChecks              []*check.DocumentTextDataCheck
	documentFaceMatchChecks             []*check.DocumentFaceMatchCheck
	livenessChecks                      []*check.LivenessCheck
	idDocumentComparisonChecks          []*check.IDDocumentComparisonCheck
	supplementaryDocumentTextDataChecks []*check.SupplementaryDocumentTextDataCheck
	thirdPartyIdentityCheck             *check.ThirdPartyIdentityCheck
	asyncReportDelay                    uint32
}

// NewCheckReportsBuilder creates a new CheckReportsBuilder
func NewCheckReportsBuilder() *CheckReportsBuilder {
	return &CheckReportsBuilder{
		documentAuthenticityChecks:          []*check.DocumentAuthenticityCheck{},
		documentTextDataChecks:              []*check.DocumentTextDataCheck{},
		documentFaceMatchChecks:             []*check.DocumentFaceMatchCheck{},
		livenessChecks:                      []*check.LivenessCheck{},
		idDocumentComparisonChecks:          []*check.IDDocumentComparisonCheck{},
		supplementaryDocumentTextDataChecks: []*check.SupplementaryDocumentTextDataCheck{},
	}
}

// WithDocumentAuthenticityCheck adds a document authenticity check
func (b *CheckReportsBuilder) WithDocumentAuthenticityCheck(documentAuthenticityCheck *check.DocumentAuthenticityCheck) *CheckReportsBuilder {
	b.documentAuthenticityChecks = append(b.documentAuthenticityChecks, documentAuthenticityCheck)
	return b
}

// WithDocumentTextDataCheck adds a document text data check
func (b *CheckReportsBuilder) WithDocumentTextDataCheck(documentTextDataCheck *check.DocumentTextDataCheck) *CheckReportsBuilder {
	b.documentTextDataChecks = append(b.documentTextDataChecks, documentTextDataCheck)
	return b
}

// WithDocumentTextDataCheck adds a supplementary document text data check
func (b *CheckReportsBuilder) WithSupplementaryDocumentTextDataCheck(supplementaryDocumentTextDataCheck *check.SupplementaryDocumentTextDataCheck) *CheckReportsBuilder {
	b.supplementaryDocumentTextDataChecks = append(
		b.supplementaryDocumentTextDataChecks,
		supplementaryDocumentTextDataCheck,
	)
	return b
}

// WithDocumentFaceMatchCheck adds a document face match check
func (b *CheckReportsBuilder) WithDocumentFaceMatchCheck(documentFaceMatchCheck *check.DocumentFaceMatchCheck) *CheckReportsBuilder {
	b.documentFaceMatchChecks = append(b.documentFaceMatchChecks, documentFaceMatchCheck)
	return b
}

// WithLivenessCheck adds a liveness check
func (b *CheckReportsBuilder) WithLivenessCheck(livenessCheck *check.LivenessCheck) *CheckReportsBuilder {
	b.livenessChecks = append(b.livenessChecks, livenessCheck)
	return b
}

// WithIDDocumentComparisonCheck adds an identity document comparison check
func (b *CheckReportsBuilder) WithIDDocumentComparisonCheck(idDocumentComparisonCheck *check.IDDocumentComparisonCheck) *CheckReportsBuilder {
	b.idDocumentComparisonChecks = append(b.idDocumentComparisonChecks, idDocumentComparisonCheck)
	return b
}

// WithAsyncReportDelay sets the async report delay
func (b *CheckReportsBuilder) WithAsyncReportDelay(asyncReportDelay uint32) *CheckReportsBuilder {
	b.asyncReportDelay = asyncReportDelay
	return b
}

func (b *CheckReportsBuilder) WithThirdPartyIdentityCheck(thirdPartyIdentityCheck *check.ThirdPartyIdentityCheck) *CheckReportsBuilder {
	b.thirdPartyIdentityCheck = thirdPartyIdentityCheck
	return b
}

// Build creates CheckReports
func (b *CheckReportsBuilder) Build() (CheckReports, error) {
	return CheckReports{
		DocumentAuthenticityChecks:          b.documentAuthenticityChecks,
		DocumentTextDataChecks:              b.documentTextDataChecks,
		DocumentFaceMatchChecks:             b.documentFaceMatchChecks,
		LivenessChecks:                      b.livenessChecks,
		IDDocumentComparisonChecks:          b.idDocumentComparisonChecks,
		SupplementaryDocumentTextDataChecks: b.supplementaryDocumentTextDataChecks,
		AsyncReportDelay:                    b.asyncReportDelay,
		ThirdPartyIdentityCheck:             b.thirdPartyIdentityCheck,
	}, nil
}
