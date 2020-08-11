package filter

type documentType string

const identity documentType = "ID_DOCUMENT"

// type restrictions string

// const (
// 	orthogonal restrictions = "ORTHOGONAL_RESTRICTIONS"
// 	document   restrictions = "DOCUMENT_RESTRICTIONS"
// )

type inclusionType string

const (
	includeList inclusionType = "WHITELIST"
	excludeList inclusionType = "BLACKLIST"
)
