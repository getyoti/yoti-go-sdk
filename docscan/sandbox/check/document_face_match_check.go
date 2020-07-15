package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/filter"
)

type DocumentFaceMatchCheck struct {
	documentCheck
}

type documentFaceMatchCheckBuilder struct {
	documentCheckBuilder
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

func (b *documentFaceMatchCheckBuilder) WithDocumentFilter(filter filter.DocumentFilter) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentFaceMatchCheckBuilder) Build() (DocumentFaceMatchCheck, error) {
	documentFaceMatchCheck := DocumentFaceMatchCheck{}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentFaceMatchCheck, err
	}

	documentFaceMatchCheck.documentCheck = documentCheck

	return documentFaceMatchCheck, nil
}
