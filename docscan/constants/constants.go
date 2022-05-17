package constants

const (
	IDDocumentAuthenticity                  string = "ID_DOCUMENT_AUTHENTICITY"
	IDDocumentComparison                    string = "ID_DOCUMENT_COMPARISON"
	IDDocumentTextDataCheck                 string = "ID_DOCUMENT_TEXT_DATA_CHECK"
	IDDocumentTextDataExtraction            string = "ID_DOCUMENT_TEXT_DATA_EXTRACTION"
	IDDocumentFaceMatch                     string = "ID_DOCUMENT_FACE_MATCH"
	SupplementaryDocumentTextDataCheck      string = "SUPPLEMENTARY_DOCUMENT_TEXT_DATA_CHECK"
	ThirdPartyIdentityCheck                 string = "THIRD_PARTY_IDENTITY"
	SupplementaryDocumentTextDataExtraction string = "SUPPLEMENTARY_DOCUMENT_TEXT_DATA_EXTRACTION"
	WatchlistScreening                      string = "WATCHLIST_SCREENING"
	WatchlistAdvancedCA                     string = "WATCHLIST_ADVANCED_CA"

	WithYotiAccounts  = "WITH_YOTI_ACCOUNT"
	WithCustomAccount = "WITH_CUSTOM_ACCOUNT"
	TypeList          = "TYPE_LIST"
	Profiles          = "PROFILE"
	Exact             = "EXACT"
	Fuzzy             = "FUZZY"

	Liveness string = "LIVENESS"
	Zoom     string = "ZOOM"

	Camera          string = "CAMERA"
	CameraAndUpload string = "CAMERA_AND_UPLOAD"

	ResourceUpdate    string = "RESOURCE_UPDATE"
	TaskCompletion    string = "TASK_COMPLETION"
	CheckCompletion   string = "CHECK_COMPLETION"
	SessionCompletion string = "SESSION_COMPLETION"

	Sanctions    string = "SANCTIONS"
	AdverseMedia string = "ADVERSE-MEDIA"

	Always   string = "ALWAYS"
	Fallback string = "FALLBACK"
	Never    string = "NEVER"

	ProofOfAddress string = "PROOF_OF_ADDRESS"
)
