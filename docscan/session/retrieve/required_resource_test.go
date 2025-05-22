package retrieve

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestResource_StringMethods(t *testing.T) {
	resources := []RequiredResourceResponse{
		&RequiredIdDocumentResourceResponse{
			BaseRequiredResource{
				Type:  "ID_DOCUMENT",
				ID:    "id1",
				State: "state1",
			},
		},
		&RequiredSupplementaryDocumentResourceResponse{
			BaseRequiredResource{
				Type:  "SUPPLEMENTARY_DOCUMENT",
				ID:    "id2",
				State: "state2",
			},
		},
		&RequiredZoomLivenessResourceResponse{
			BaseRequiredResource{
				Type:         "LIVENESS",
				ID:           "id3",
				State:        "state3",
				LivenessType: "ZOOM",
			},
		},
		&RequiredStaticLivenessResourceResponse{
			BaseRequiredResource{
				Type:         "LIVENESS",
				ID:           "id3",
				State:        "state3",
				LivenessType: "STATIC",
			},
		},
		&RequiredFaceCaptureResourceResponse{
			BaseRequiredResource{
				Type:  "FACE_CAPTURE",
				ID:    "id4",
				State: "state4",
			},
		},
		&UnknownRequiredResourceResponse{
			BaseRequiredResource{
				Type:  "UNKNOWN",
				ID:    "id5",
				State: "state5",
			},
		},
	}

	for _, r := range resources {
		str := r.String()
		assert.Assert(t, str != "", "String method should return non-empty string")
	}
}
