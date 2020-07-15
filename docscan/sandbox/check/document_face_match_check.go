package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
)

type documentFaceMatchCheck struct {
	documentCheck
}

type documentFaceMatchCheckBuilder struct {
	documentCheckBuilder
	err error
}

func NewDocumentFaceMatchCheckBuilder() *documentFaceMatchCheckBuilder {
	return &documentFaceMatchCheckBuilder{}
}

func (b *documentFaceMatchCheckBuilder) WithRecommendation(recommendation report.Recommendation) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentFaceMatchCheckBuilder) WithBreakdown(breakdown report.Breakdown) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentFaceMatchCheckBuilder) WithDocumentFilter(filter sandbox.DocumentFilter) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentFaceMatchCheckBuilder) Build() (documentFaceMatchCheck, error) {
	documentFaceMatchCheck := documentFaceMatchCheck{}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentFaceMatchCheck, err
	}

	documentFaceMatchCheck.documentCheck = documentCheck

	return documentFaceMatchCheck, b.err
}
