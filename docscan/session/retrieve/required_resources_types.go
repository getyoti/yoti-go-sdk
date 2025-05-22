package retrieve

type RequiredIdDocumentResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredIdDocumentResourceResponse) String() string {
	return "ID Document Resource - " + r.BaseRequiredResource.String()
}

type RequiredSupplementaryDocumentResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredSupplementaryDocumentResourceResponse) String() string {
	return "Supplementary Document Resource - " + r.BaseRequiredResource.String()
}

type RequiredZoomLivenessResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredZoomLivenessResourceResponse) String() string {
	return "Zoom Liveness Resource - " + r.BaseRequiredResource.String()
}

type RequiredLivenessResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredLivenessResourceResponse) String() string {
	return "Liveness Resource - " + r.BaseRequiredResource.String()
}

type RequiredStaticLivenessResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredStaticLivenessResourceResponse) String() string {
	return "Static Liveness Resource - " + r.BaseRequiredResource.String()
}

type RequiredFaceCaptureResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredFaceCaptureResourceResponse) String() string {
	return "Face Capture Resource - " + r.BaseRequiredResource.String()
}

type UnknownRequiredResourceResponse struct {
	BaseRequiredResource
}

func (r *UnknownRequiredResourceResponse) String() string {
	return "Unknown Resource Type - " + r.BaseRequiredResource.String()
}
