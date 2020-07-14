package sandbox

type documentTextDataExtractionTask struct {
	Result documentTextDataExtractionTaskResult `json:"result"`
	documentTask
}

type documentTextDataExtractionTaskBuilder struct {
	documentTaskBuilder
	documentFields map[string]string
	err            error
}

type documentTextDataExtractionTaskResult struct {
	taskResult
	DocumentFields map[string]string `json:"document_fields"`
}

func NewDocumentTextDataExtractionTaskBuilder() *documentTextDataExtractionTaskBuilder {
	return &documentTextDataExtractionTaskBuilder{}
}

func (b *documentTextDataExtractionTaskBuilder) WithRecommendation(recommendation recommendation) *documentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentTextDataExtractionTaskBuilder) WithBreakdown(breakdown breakdown) *documentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentTextDataExtractionTaskBuilder) WithDocumentFilter(filter documentFilter) *documentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentTextDataExtractionTaskBuilder) WithDocumentField(key string, value string) *documentTextDataExtractionTaskBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]string)
	}
	b.documentFields[key] = value
	return b
}

func (b *documentTextDataExtractionTaskBuilder) Build() (documentTextDataExtractionTask, error) {
	documentTextDataExtractionTask := documentTextDataExtractionTask{}

	documentTask, err := b.documentTaskBuilder.build()
	if err != nil {
		return documentTextDataExtractionTask, err
	}

	documentTextDataExtractionTask.documentTask = documentTask
	documentTextDataExtractionTask.Result = documentTextDataExtractionTaskResult{
		taskResult:     documentTask.Result,
		DocumentFields: b.documentFields,
	}

	return documentTextDataExtractionTask, b.err
}
