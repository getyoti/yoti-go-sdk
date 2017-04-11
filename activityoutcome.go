package yoti

type ActivityOutcome int

const (
	ActivityOutcome_ProfileNotFound = "ProfileNotFound"
	ActivityOutcome_Failure         = "Failure"
	ActivityOutcome_SharingFailure  = "SharingFailure"
)
