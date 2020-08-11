package constants

type CaptureMethod string

const (
	Camera          CaptureMethod = "CAMERA"
	CameraAndUpload CaptureMethod = "CAMERA_AND_UPLOAD"
)

type Topic string

const (
	ResourceUpdate    Topic = "RESOURCE_UPDATE"
	TaskCompletion    Topic = "TASK_COMPLETION"
	CheckCompletion   Topic = "CHECK_COMPLETION"
	SessionCompletion Topic = "SESSION_COMPLETION"
)

type ManualCheck string

const (
	Always   ManualCheck = "ALWAYS"
	Fallback ManualCheck = "FALLBACK"
	Never    ManualCheck = "NEVER"
)

type IDDocument string

const (
	Authenticity       IDDocument = "ID_DOCUMENT_AUTHENTICITY"
	TextDataCheck      IDDocument = "ID_DOCUMENT_TEXT_DATA_CHECK"
	TextDataExtraction IDDocument = "ID_DOCUMENT_TEXT_DATA_EXTRACTION"
	FaceMatch          IDDocument = "ID_DOCUMENT_FACE_MATCH"
)
