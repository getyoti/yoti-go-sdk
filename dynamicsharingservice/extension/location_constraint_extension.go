package extension

import (
	"encoding/json"
)

const (
	locationConstraintExtensionTypeConst = "LOCATION_CONSTRAINT"
)

// LocationConstraintExtensionBuilder is used to construct a LocationConstraintExtension
type LocationConstraintExtensionBuilder struct {
	extension LocationConstraintExtension
}

// LocationConstraintExtension is an extension representing a geographic constraint
type LocationConstraintExtension struct {
	latitude    float64
	longtitude  float64
	radius      float64
	uncertainty float64
}

// New initializes the builder
func (builder *LocationConstraintExtensionBuilder) New() *LocationConstraintExtensionBuilder {
	builder.extension.latitude = 0
	builder.extension.longtitude = 0
	builder.extension.radius = 0
	builder.extension.uncertainty = 0
	return builder
}

// WithLatitude sets the latitude of the location constraint
func (builder *LocationConstraintExtensionBuilder) WithLatitude(latitude float64) *LocationConstraintExtensionBuilder {
	builder.extension.latitude = latitude
	return builder
}

// WithLongtitude sets the longtitude of the location constraint
func (builder *LocationConstraintExtensionBuilder) WithLongtitude(longtitude float64) *LocationConstraintExtensionBuilder {
	builder.extension.longtitude = longtitude
	return builder
}

// WithRadius sets the radius within which the location constraint will be satisfied
func (builder *LocationConstraintExtensionBuilder) WithRadius(radius float64) *LocationConstraintExtensionBuilder {
	builder.extension.radius = radius
	return builder
}

// WithUncertainty sets the max uncertainty allowed by the location constraint extension
func (builder *LocationConstraintExtensionBuilder) WithUncertainty(uncertainty float64) *LocationConstraintExtensionBuilder {
	builder.extension.uncertainty = uncertainty
	return builder
}

// Build constructs a LocationConstraintExtension from the builder
func (builder *LocationConstraintExtensionBuilder) Build() LocationConstraintExtension {
	return builder.extension
}

// MarshalJSON ...
func (extension LocationConstraintExtension) MarshalJSON() ([]byte, error) {
	type content struct {
		Latitude       float64 `json:"latitude"`
		Longtitude     float64 `json:"longtitude"`
		Radius         float64 `json:"radius"`
		MaxUncertainty float64 `json:"max_uncertainty_radius"`
	}
	return json.Marshal(&struct {
		Type    string  `json:"type"`
		Content content `json:"content"`
	}{
		Type: locationConstraintExtensionTypeConst,
		Content: content{
			Latitude:       extension.latitude,
			Longtitude:     extension.longtitude,
			Radius:         extension.radius,
			MaxUncertainty: extension.uncertainty,
		},
	})
}
