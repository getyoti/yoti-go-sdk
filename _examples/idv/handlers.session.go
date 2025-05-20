package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/docscan"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

const (
	defaultBaseURL string = "https://api.yoti.com/idverify/v1"
)

var (
	sdkID               string
	key                 []byte
	client              *docscan.Client
	createSessionResult *create.SessionResult
)

func showIndexPage(c *gin.Context) {
	sessionSpec, err := buildSessionSpec()
	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error when building sessions spec",
				"ErrorMessage": err.Error()})
		return
	}
	pageFromSessionSpec(c, sessionSpec)
}

func showDBSPage(c *gin.Context) {
	sessionSpec, err := buildDBSSessionSpec()
	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error when building sessions spec",
				"ErrorMessage": err.Error()})
		return
	}
	pageFromSessionSpec(c, sessionSpec)
}

func showAdvancedIdentityProfilePage(c *gin.Context) {
	sessionSpec, err := buildAdvancedIdentityProfileSessionSpec()
	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error when building sessions spec",
				"ErrorMessage": err.Error()})
		return
	}
	pageFromSessionSpec(c, sessionSpec)
}

func pageFromSessionSpec(c *gin.Context, sessionSpec *create.SessionSpecification) {
	err := initialiseDocScanClient()
	if err != nil {
		c.HTML(
			http.StatusUnprocessableEntity,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error initialising Doc Scan (IDV) Client",
				"ErrorMessage": errors.Unwrap(err)})
		return
	}
	createSessionResult, err = client.CreateSession(sessionSpec)
	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error when creating Doc Scan (IDV) session",
				"ErrorMessage": err.Error()})
		return
	}

	c.SetCookie("session_id", createSessionResult.SessionID, 60*20, "/", "localhost", true, false)
	iFrameURL := getIframeURL(createSessionResult.SessionID, createSessionResult.ClientSessionToken)

	render(c, gin.H{
		"iframeURL": iFrameURL},
		"index.html")
	return
}

func getBaseURL() string {
	if value, exists := os.LookupEnv("YOTI_DOC_SCAN_API_URL"); exists && value != "" {
		return value
	}

	return defaultBaseURL
}

func getIframeURL(sessionID, sessionToken string) string {
	baseURL := getBaseURL()
	return fmt.Sprintf("%s/web/index.html?sessionID=%s&sessionToken=%s", baseURL, sessionID, sessionToken)
}

func showSuccessPage(c *gin.Context) {
	err := ensureDocScanClientInitialised()
	if err != nil {
		c.HTML(
			http.StatusUnprocessableEntity,
			"error.html",
			gin.H{
				"ErrorTitle":   "error setting the Doc Scan (IDV) Client",
				"ErrorMessage": err.Error()})
		return
	}

	sessionId, err := c.Cookie("session_id")
	if err != nil || sessionId == "" {
		c.HTML(
			http.StatusUnprocessableEntity,
			"error.html",
			gin.H{
				"ErrorTitle":   "Failed to get session ID",
				"ErrorMessage": err.Error()})
		return
	}

	c.Set("session_created", true)

	getSessionResult, err := client.GetSession(sessionId)
	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"error.html",
			gin.H{
				"ErrorTitle":   "Get Session Failed",
				"ErrorMessage": err.Error()})
		return
	}

	render(
		c,
		gin.H{
			"title":            "Success",
			"getSessionResult": getSessionResult,
			"add": func(a int, b int) int {
				return a + b
			},
		},
		"success.html",
	)
	return
}

func ensureDocScanClientInitialised() error {
	if client == nil {
		return initialiseDocScanClient()
	}
	return nil
}

func initialiseDocScanClient() error {
	var err error
	sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")
	keyFilePath := os.Getenv("YOTI_KEY_FILE_PATH")
	key, err = os.ReadFile(keyFilePath)
	if err != nil {
		return fmt.Errorf("failed to get key from YOTI_KEY_FILE_PATH :: %w", err)
	}

	client, err = docscan.NewClient(sdkID, key)
	if err != nil {
		return fmt.Errorf("failed to initialise Yoti Doc Scan (IDV) client :: %w", err)
	}

	return nil
}

func getMedia(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil || sessionId == "" {
		c.HTML(
			http.StatusInternalServerError,
			"error.html",
			gin.H{
				"ErrorTitle":   "Failed to get session ID",
				"ErrorMessage": err.Error()})
		return
	}

	mediaID := c.Query("mediaId")

	media, err := client.GetMediaContent(sessionId, mediaID)
	if err != nil || sessionId == "" {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Failed to get media",
				"ErrorMessage": err.Error()})
		return
	}

	if media == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.Header("Content-Type", media.MIME())
	c.Data(http.StatusOK, media.MIME(), media.Data())
	return
}

func showPrivacyPolicyPage(c *gin.Context) {
	render(c, gin.H{}, "privacy.html")
	return
}

func showErrorPage(c *gin.Context) {
	render(c, gin.H{
		"ErrorTitle":   "Error Code",
		"ErrorMessage": c.Request.URL.Query().Get("yotiErrorCode")},
		"error.html")
	return
}

func showFaceCaptureSessionPage(c *gin.Context) {
	err := initialiseDocScanClient()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"ErrorTitle":   "Error initializing Doc Scan client",
			"ErrorMessage": err.Error(),
		})
		return
	}

	sessionSpec, err := buildFaceCaptureSessionSpec()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"ErrorTitle":   "Failed to build session spec",
			"ErrorMessage": err.Error(),
		})
		return
	}

	sessionResult, err := client.CreateSession(sessionSpec)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"ErrorTitle":   "Session creation failed",
			"ErrorMessage": err.Error(),
		})
		return
	}

	sessionID := sessionResult.SessionID
	fmt.Printf(sessionResult.SessionID)
	sessionToken := sessionResult.ClientSessionToken
	c.SetCookie("session_id", sessionID, 60*20, "/", "localhost", true, false)

	err = client.AddFaceCaptureResourceToSession(sessionID)
	fmt.Printf(sessionResult.SessionID)
	fmt.Printf(sessionResult.ClientSessionToken)
	fmt.Printf("Error: %+v\n", err)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"ErrorTitle":   "Add Face Capture Resource Failed",
			"ErrorMessage": err.Error(),
		})
		return
	}

	iframeURL := getIframeURL(sessionID, sessionToken)
	fmt.Printf("Iframe: %s", iframeURL)

	render(c, gin.H{
		"iframeURL": iframeURL},
		"index.html")
}

func showSimpleFaceCapturePage(c *gin.Context) {
	showFaceCaptureSessionPage(c)
}
