package sandbox

const (
	zoom = "ZOOM"
)

type zoomLivenessCheck struct {
	livenessCheck
}

type zoomLivenessCheckBuilder struct {
	livenessCheckBuilder
	err error
}

func NewZoomLivenessCheckBuilder() *zoomLivenessCheckBuilder {
	return &zoomLivenessCheckBuilder{}
}

func (b *zoomLivenessCheckBuilder) WithRecommendation(recommendation recommendation) *zoomLivenessCheckBuilder {
	b.livenessCheckBuilder.withRecommendation(recommendation)
	return b
}

func (b *zoomLivenessCheckBuilder) WithBreakdown(breakdown breakdown) *zoomLivenessCheckBuilder {
	b.livenessCheckBuilder.withBreakdown(breakdown)
	return b
}

func (b *zoomLivenessCheckBuilder) Build() (zoomLivenessCheck, error) {
	zoomLivenessCheck := zoomLivenessCheck{}

	livenessCheck, err := b.livenessCheckBuilder.build()
	if err != nil {
		return zoomLivenessCheck, err
	}

	zoomLivenessCheck.livenessCheck = livenessCheck
	zoomLivenessCheck.LivenessType = zoom

	return zoomLivenessCheck, b.err
}
