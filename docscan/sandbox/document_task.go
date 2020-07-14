package sandbox

type documentTask struct {
	task
	DocumentFilter documentFilter `json:"document_filter"`
}

type documentTaskBuilder struct {
	taskBuilder
	documentFilter documentFilter
	err            error
}

func (b *documentTaskBuilder) withDocumentFilter(filter documentFilter) {
	b.documentFilter = filter
}

func (b *documentTaskBuilder) build() (documentTask, error) {
	documentTask := documentTask{}

	task, err := b.taskBuilder.build()
	if err != nil {
		return documentTask, err
	}

	documentTask.task = task
	documentTask.DocumentFilter = b.documentFilter

	return documentTask, b.err
}
