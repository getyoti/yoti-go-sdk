package retrieve

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestResource_StringMethods(t *testing.T) {
	resources := []RequiredResourceResponse{
		&RequiredIdDocumentResourceResponse{BaseRequiredResource{"ID_DOCUMENT", "id1", "state1"}},
		&RequiredSupplementaryDocumentResourceResponse{BaseRequiredResource{"SUPPLEMENTARY_DOCUMENT", "id2", "state2"}},
		&RequiredZoomLivenessResourceResponse{BaseRequiredResource{"LIVENESS", "id3", "state3"}},
		&RequiredFaceCaptureResourceResponse{BaseRequiredResource{"FACE_CAPTURE", "id4", "state4"}},
		&UnknownRequiredResourceResponse{BaseRequiredResource{"UNKNOWN", "id5", "state5"}},
	}

	for _, r := range resources {
		str := r.String()
		assert.Assert(t, str != "", "String method should return non-empty string")
	}
}
