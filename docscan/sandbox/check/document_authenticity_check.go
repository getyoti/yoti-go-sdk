package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/filter"
)

type DocumentAuthenticityCheck struct {
	documentCheck
}

type documentAuthenticityCheckBuilder struct {
	documentCheckBuilder
	err error
}

func NewDocumentAuthenticityCheckBuilder() *documentAuthenticityCheckBuilder {
	return &documentAuthenticityCheckBuilder{}
}

func (b *documentAuthenticityCheckBuilder) WithRecommendation(recommendation report.Recommendation) *documentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentAuthenticityCheckBuilder) WithBreakdown(breakdown report.Breakdown) *documentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentAuthenticityCheckBuilder) WithDocumentFilter(filter filter.DocumentFilter) *documentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentAuthenticityCheckBuilder) Build() (DocumentAuthenticityCheck, error) {
	documentAuthenticityCheck := DocumentAuthenticityCheck{}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentAuthenticityCheck, err
	}

	documentAuthenticityCheck.documentCheck = documentCheck

	return documentAuthenticityCheck, b.err
}
