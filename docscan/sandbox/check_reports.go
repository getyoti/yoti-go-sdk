package sandbox

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check"
)

type CheckReports struct {
	DocumentAuthenticityChecks []check.DocumentAuthenticityCheck `json:"ID_DOCUMENT_AUTHENTICITY"`
	DocumentTextDataChecks     []check.DocumentTextDataCheck     `json:"ID_DOCUMENT_TEXT_DATA_CHECK"`
	DocumentFaceMatchChecks    []check.DocumentFaceMatchCheck    `json:"ID_DOCUMENT_FACE_MATCH_CHECK"`
	LivenessChecks             []check.LivenessCheck             `json:"LIVENESS"`
	AsyncReportDelay           uint32                            `json:"async_report_delay"`
}

type checkReportsBuilder struct {
	documentAuthenticityChecks []check.DocumentAuthenticityCheck
	documentTextDataChecks     []check.DocumentTextDataCheck
	documentFaceMatchChecks    []check.DocumentFaceMatchCheck
	livenessChecks             []check.LivenessCheck
	asyncReportDelay           uint32
	err                        error
}

func NewCheckReportsBuilder() *checkReportsBuilder {
	return &checkReportsBuilder{}
}

func (b *checkReportsBuilder) WithDocumentAuthenticityCheck(documentAuthenticityCheck check.DocumentAuthenticityCheck) *checkReportsBuilder {
	b.documentAuthenticityChecks = append(b.documentAuthenticityChecks, documentAuthenticityCheck)
	return b
}

func (b *checkReportsBuilder) WithDocumentTextDataCheck(documentTextDataCheck check.DocumentTextDataCheck) *checkReportsBuilder {
	b.documentTextDataChecks = append(b.documentTextDataChecks, documentTextDataCheck)
	return b
}

func (b *checkReportsBuilder) WithDocumentFaceMatchCheck(documentFaceMatchCheck check.DocumentFaceMatchCheck) *checkReportsBuilder {
	b.documentFaceMatchChecks = append(b.documentFaceMatchChecks, documentFaceMatchCheck)
	return b
}

func (b *checkReportsBuilder) WithLivenessCheck(livenessCheck check.LivenessCheck) *checkReportsBuilder {
	b.livenessChecks = append(b.livenessChecks, livenessCheck)
	return b
}

func (b *checkReportsBuilder) WithAsyncReportDelay(asyncReportDelay uint32) *checkReportsBuilder {
	b.asyncReportDelay = asyncReportDelay
	return b
}

func (b *checkReportsBuilder) Build() (CheckReports, error) {
	return CheckReports{
		DocumentAuthenticityChecks: b.documentAuthenticityChecks,
		DocumentTextDataChecks:     b.documentTextDataChecks,
		DocumentFaceMatchChecks:    b.documentFaceMatchChecks,
		LivenessChecks:             b.livenessChecks,
		AsyncReportDelay:           b.asyncReportDelay,
	}, b.err
}
