package check

// LivenessCheck represents a liveness check
type LivenessCheck struct {
	*check
	LivenessType string `json:"liveness_type"`
}

type livenessCheckBuilder struct {
	checkBuilder
	livenessType string
}

func (b *livenessCheckBuilder) withLivenessType(livenessType string) *livenessCheckBuilder {
	b.livenessType = livenessType
	return b
}

func (b *livenessCheckBuilder) build() *LivenessCheck {
	return &LivenessCheck{
		LivenessType: b.livenessType,
		check:        b.checkBuilder.build(),
	}
}
