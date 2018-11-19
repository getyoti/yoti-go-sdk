package yoti

//ActivityDetails represents the result of an activity between a user and the application
type ActivityDetails struct {
	UserProfile  Profile
	rememberMeID string
}

// RememberMeID is a unique identifier Yoti assigns to your user, but only for your app
// if the same user logs into your app again, you get the same id
// if she/he logs into another application, Yoti will assign a different id for that app
func (a ActivityDetails) RememberMeID() string {
	return a.rememberMeID
}
