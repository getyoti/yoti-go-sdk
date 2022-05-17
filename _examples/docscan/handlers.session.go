package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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
	initialiseDocScanClient(c)

	sessionSpec, err := buildSessionSpec()
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error when building sessions spec",
				"ErrorMessage": err.Error()})
	}

	createSessionResult, err = client.CreateSession(sessionSpec)
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error when creating Doc Scan session",
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
	ensureDocScanClientInitialised(c)

	sessionId, err := c.Cookie("session_id")
	if err != nil || sessionId == "" {
		c.HTML(
			http.StatusBadRequest,
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
			http.StatusBadRequest,
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
			"stringsJoin": strings.Join,
		},
		"success.html",
	)
	return
}

func ensureDocScanClientInitialised(c *gin.Context) {
	if client == nil {
		initialiseDocScanClient(c)
	}
	return
}

func initialiseDocScanClient(c *gin.Context) {
	var err error
	sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")
	keyFilePath := os.Getenv("YOTI_KEY_FILE_PATH")
	key, err = ioutil.ReadFile(keyFilePath)
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Failed to get key from YOTI_KEY_FILE_PATH",
				"ErrorMessage": err.Error()})
	}

	client, err = docscan.NewClient(sdkID, key)
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Failed to initialise Yoti Doc Scan client",
				"ErrorMessage": err.Error()})
	}
	return
}

func getMedia(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil || sessionId == "" {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Failed to get session ID",
				"ErrorMessage": err.Error()})
		return
	}

	mediaID := c.Query("mediaId")

	media, err := client.GetMediaContent(sessionId, mediaID)

	if media == nil && err == nil {
		c.Status(http.StatusNoContent)
		return
	}

	if err != nil || sessionId == "" {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Failed to get media",
				"ErrorMessage": err.Error()})
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
