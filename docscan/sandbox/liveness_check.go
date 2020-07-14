package sandbox

type livenessCheck struct {
	check
	LivenessType string `json:"liveness_type"`
}

type livenessCheckBuilder struct {
	checkBuilder
	err error
}

func (b *livenessCheckBuilder) build() (livenessCheck, error) {
	livenessCheck := livenessCheck{}

	check, err := b.checkBuilder.build()
	if err != nil {
		return livenessCheck, err
	}

	livenessCheck.check = check

	return livenessCheck, b.err
}
