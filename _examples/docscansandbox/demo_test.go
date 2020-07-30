package docscansandbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/task"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	appBaseUrl        = "https://app:3000"
	documentImagePath = "/usr/src/resources/image.jpg"
)

var pool = sync.Pool{
	New: func() interface{} {
		return startWebDriver()
	},
}

func startWebDriver() selenium.WebDriver {
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{
		Args: []string{
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-gpu",
			"--window-size=1280,2000",
			"--ignore-certificate-errors",
		},
	})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
	if err != nil {
		panic(err)
	}

	// Set timeouts
	var timeout time.Duration = 5000 * time.Millisecond
	wd.SetPageLoadTimeout(timeout)
	wd.SetImplicitWaitTimeout(timeout)
	wd.SetAsyncScriptTimeout(timeout)

	return wd
}

func newSandboxClient() *sandbox.Client {
	key, err := ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	if err != nil {
		panic(err)
	}
	sdkID := os.Getenv("YOTI_SANDBOX_CLIENT_SDK_ID")

	client, err := sandbox.NewClient(sdkID, key)
	if err != nil {
		panic(err)
	}
	return client
}

type webContext struct {
	wd     selenium.WebDriver
	client *sandbox.Client
}

func (c *webContext) getIFrameSessionID() (sessionId string, err error) {
	var iFrame selenium.WebElement
	iFrame, err = c.wd.FindElement(selenium.ByTagName, "iframe")
	if err != nil {
		return
	}

	iFrameURL, err := iFrame.GetAttribute("src")
	if err != nil {
		return
	}

	parsedURL, err := url.Parse(iFrameURL)
	if err != nil {
		return
	}

	query, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return
	}

	return query.Get("sessionID"), nil
}

func (c *webContext) iConfigureTheSessionResponse() error {
	documentFilter, err := filter.NewDocumentFilterBuilder().
		WithDocumentType("PASSPORT").
		Build()
	if err != nil {
		return err
	}

	authenticityBreakdown, err := report.NewBreakdownBuilder().
		WithSubCheck("security_features").
		WithResult("NOT_AVAILABLE").
		WithDetail("some_detail", "some_detail_value").
		Build()
	if err != nil {
		return err
	}

	authenticityRecommendation, err := report.NewRecommendationBuilder().
		WithValue("NOT_AVAILABLE").
		WithReason("PICTURE_TOO_DARK").
		WithRecoverySuggestion("BETTER_LIGHTING").
		Build()
	if err != nil {
		return err
	}

	documentAuthenticityCheck, err := check.NewDocumentAuthenticityCheckBuilder().
		WithBreakdown(authenticityBreakdown).
		WithRecommendation(authenticityRecommendation).
		WithDocumentFilter(documentFilter).
		Build()
	if err != nil {
		return err
	}

	textDataCheckBreakdown, err := report.NewBreakdownBuilder().
		WithSubCheck("document_in_date").
		WithResult("PASS").
		Build()
	if err != nil {
		return err
	}

	textDataCheckRecommendation, err := report.NewRecommendationBuilder().
		WithValue("APPROVE").
		Build()

	textDataCheck, err := check.NewDocumentTextDataCheckBuilder().
		WithBreakdown(textDataCheckBreakdown).
		WithRecommendation(textDataCheckRecommendation).
		WithDocumentFields(map[string]string{
			"full_name":       "John Doe",
			"nationality":     "GBR",
			"date_of_birth":   "1986-06-01",
			"document_number": "123456789",
		}).
		WithDocumentFilter(documentFilter).
		Build()
	if err != nil {
		return err
	}

	checkReports, err := request.NewCheckReportsBuilder().
		WithAsyncReportDelay(5).
		WithDocumentAuthenticityCheck(documentAuthenticityCheck).
		WithDocumentTextDataCheck(textDataCheck).
		Build()
	if err != nil {
		return err
	}

	textExtractionTask, err := task.NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentFields(map[string]string{
			"full_name":       "John Doe",
			"nationality":     "GBR",
			"date_of_birth":   "1986-06-01",
			"document_number": "123456789",
		}).
		Build()
	if err != nil {
		return err
	}

	taskResults, err := request.NewTaskResultsBuilder().
		WithDocumentTextDataExtractionTask(textExtractionTask).
		Build()
	if err != nil {
		return err
	}

	responseConfig, err := request.NewResponseConfigBuilder().
		WithCheckReports(checkReports).
		WithTaskResults(taskResults).
		Build()
	if err != nil {
		return err
	}

	var sessionId string
	sessionId, err = c.getIFrameSessionID()
	if err != nil {
		return err
	}

	configErr := c.client.ConfigureSessionResponse(sessionId, responseConfig)

	if configErr != nil {
		request, _ := json.Marshal(responseConfig)
		return errors.New(string(request) + configErr.Error())
	}

	return nil
}

func (c *webContext) iAmOn(path string) error {
	return c.wd.Get(fmt.Sprintf("%s%s", appBaseUrl, path))
}

func (c *webContext) clickOn(selector string) error {
	elem, err := c.wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		return err
	}

	return elem.Click()
}

func (c *webContext) iChoose(value string) error {
	return c.clickOn(fmt.Sprintf("input[value='%s']", value))
}

func (c *webContext) iClickOnAddPhotoButton() error {
	return c.clickOn(fmt.Sprintf("*[data-qa='addPhotoButton']"))
}

func (c *webContext) iClickOnFinishButton() error {
	return c.clickOn(fmt.Sprintf("*[data-qa='finish-button']"))
}

func (c *webContext) iShouldSee(text string) error {
	source, err := c.wd.PageSource()
	if err != nil {
		return err
	}
	if !strings.Contains(source, text) {
		return fmt.Errorf("Page source does not contain \"%s\"", text)
	}
	return nil
}

func (c *webContext) iSwitchToTheIframe() error {
	iFrame, err := c.wd.FindElement(selenium.ByTagName, "iframe")
	if err != nil {
		return err
	}
	return c.wd.SwitchFrame(iFrame)
}

func (c *webContext) iUploadADocument() error {
	uploadElement, err := c.wd.FindElement(selenium.ByCSSSelector, "input[data-qa='change-photo']")
	if err != nil {
		return err
	}
	return uploadElement.SendKeys(documentImagePath)
}

func (c *webContext) iWaitSeconds(seconds int) error {
	time.Sleep(time.Duration(seconds) * time.Second)
	return nil
}

func (c *webContext) elementContains(selector string, text string) error {
	elem, err := c.wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		return err
	}
	elemText, err := elem.Text()
	if err != nil {
		return err
	}
	if !strings.Contains(elemText, text) {
		return fmt.Errorf("\"%s\" does not contain \"%s\"", selector, text)
	}
	return nil
}

func (c *webContext) theAuthenticityCheckBreakdownResultShouldBe(result string) error {
	return c.elementContains(
		"*[data-qa='authenticity-checks'] *[data-qa='result']",
		result,
	)
}

func (c *webContext) theAuthenticityCheckBreakdownSubCheckShouldBe(subCheck string) error {
	return c.elementContains(
		"*[data-qa='authenticity-checks'] *[data-qa='sub-check']",
		subCheck,
	)
}

func (c *webContext) theTextDataCheckBreakdownResultShouldBe(result string) error {
	return c.elementContains(
		"*[data-qa='text-data-checks'] *[data-qa='result']",
		result,
	)
}

func (c *webContext) theTextDataCheckBreakdownSubCheckShouldBe(subCheck string) error {
	return c.elementContains(
		"*[data-qa='text-data-checks'] *[data-qa='sub-check']",
		subCheck,
	)
}

func FeatureContext(s *godog.Suite) {
	context := &webContext{}

	s.BeforeScenario(func(*messages.Pickle) {
		context.client = newSandboxClient()
		context.wd = pool.Get().(selenium.WebDriver)
	})
	s.AfterScenario(func(*messages.Pickle, error) {
		context.wd.Quit()
	})

	s.Step(`^I am on "([^"]*)"$`, context.iAmOn)
	s.Step(`^I choose "([^"]*)"$`, context.iChoose)
	s.Step(`^I click on add photo button$`, context.iClickOnAddPhotoButton)
	s.Step(`^I click on finish button$`, context.iClickOnFinishButton)
	s.Step(`^I configure the session response$`, context.iConfigureTheSessionResponse)
	s.Step(`^I should see "([^"]*)"$`, context.iShouldSee)
	s.Step(`^I switch to the iframe$`, context.iSwitchToTheIframe)
	s.Step(`^I upload a document$`, context.iUploadADocument)
	s.Step(`^I wait (\d+) seconds$`, context.iWaitSeconds)
	s.Step(`^the authenticity check breakdown result should be "([^"]*)"$`, context.theAuthenticityCheckBreakdownResultShouldBe)
	s.Step(`^the authenticity check breakdown sub check should be "([^"]*)"$`, context.theAuthenticityCheckBreakdownSubCheckShouldBe)
	s.Step(`^the text data check breakdown result should be "([^"]*)"$`, context.theTextDataCheckBreakdownResultShouldBe)
	s.Step(`^the text data check breakdown sub check should be "([^"]*)"$`, context.theTextDataCheckBreakdownSubCheckShouldBe)
}
