package constants

//TODO: should these be exported?
// If we export, they should be commented :S
const (
	IDDocumentAuthenticity       = "ID_DOCUMENT_AUTHENTICITY"
	IDDocumentTextDataCheck      = "ID_DOCUMENT_TEXT_DATA_CHECK"
	IDDocumentTextDataExtraction = "ID_DOCUMENT_TEXT_DATA_EXTRACTION"
	IDDocumentFaceMatch          = "ID_DOCUMENT_FACE_MATCH"
	Liveness                     = "LIVENESS"
	Zoom                         = "ZOOM"

	Camera          = "CAMERA"
	CameraAndUpload = "CAMERA_AND_UPLOAD"

	ResourceUpdate    = "RESOURCE_UPDATE"
	TaskCompletion    = "TASK_COMPLETION"
	CheckCompletion   = "CHECK_COMPLETION"
	SessionCompletion = "SESSION_COMPLETION"

	IDDocument             = "ID_DOCUMENT"
	OrthogonalRestrictions = "ORTHOGONAL_RESTRICTIONS"
	DocumentRestrictions   = "DOCUMENT_RESTRICTIONS"
	Includelist            = "WHITELIST"
	Excludelist            = "BLACKLIST"

	Always   = "ALWAYS"
	Fallback = "FALLBACK"
	Never    = "NEVER"

	Desired = "DESIRED"
	Ignore  = "IGNORE"
)
