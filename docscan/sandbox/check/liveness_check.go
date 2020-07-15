package check

type LivenessCheck struct {
	check
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

func (b *livenessCheckBuilder) build() (LivenessCheck, error) {
	livenessCheck := LivenessCheck{
		LivenessType: b.livenessType,
	}

	check, err := b.checkBuilder.build()
	if err != nil {
		return livenessCheck, err
	}

	livenessCheck.check = check

	return livenessCheck, nil
}
