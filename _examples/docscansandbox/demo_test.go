package docscansandbox

import "github.com/cucumber/godog"

func iAmOn(arg1 string) error {
	return godog.ErrPending
}

func iChoose(arg1 string) error {
	return godog.ErrPending
}

func iClickOnAddPhotoButton() error {
	return godog.ErrPending
}

func iClickOnFinishButton() error {
	return godog.ErrPending
}

func iConfigureTheSessionResponse() error {
	return godog.ErrPending
}

func iShouldSee(arg1 string) error {
	return godog.ErrPending
}

func iSwitchToTheIframe() error {
	return godog.ErrPending
}

func iUploadADocument() error {
	return godog.ErrPending
}

func iWaitSeconds(arg1 int) error {
	return godog.ErrPending
}

func theAuthenticityCheckBreakdownResultShouldBe(arg1 string) error {
	return godog.ErrPending
}

func theAuthenticityCheckBreakdownSubCheckShouldBe(arg1 string) error {
	return godog.ErrPending
}

func theTextDataCheckBreakdownResultShouldBe(arg1 string) error {
	return godog.ErrPending
}

func theTextDataCheckBreakdownSubCheckShouldBe(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I am on "([^"]*)"$`, iAmOn)
	s.Step(`^I choose "([^"]*)"$`, iChoose)
	s.Step(`^I click on add photo button$`, iClickOnAddPhotoButton)
	s.Step(`^I click on finish button$`, iClickOnFinishButton)
	s.Step(`^I configure the session response$`, iConfigureTheSessionResponse)
	s.Step(`^I should see "([^"]*)"$`, iShouldSee)
	s.Step(`^I switch to the iframe$`, iSwitchToTheIframe)
	s.Step(`^I upload a document$`, iUploadADocument)
	s.Step(`^I wait (\d+) seconds$`, iWaitSeconds)
	s.Step(`^the authenticity check breakdown result should be "([^"]*)"$`, theAuthenticityCheckBreakdownResultShouldBe)
	s.Step(`^the authenticity check breakdown sub check should be "([^"]*)"$`, theAuthenticityCheckBreakdownSubCheckShouldBe)
	s.Step(`^the text data check breakdown result should be "([^"]*)"$`, theTextDataCheckBreakdownResultShouldBe)
	s.Step(`^the text data check breakdown sub check should be "([^"]*)"$`, theTextDataCheckBreakdownSubCheckShouldBe)
}
