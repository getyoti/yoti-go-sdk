package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/getyoti/yoti-go-sdk/v3"
	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
	_ "github.com/joho/godotenv/autoload"
)

type contextKey string

var (
	sdkID                        string
	key                          []byte
	client                       *yoti.Client
	selfSignedCertName           = "yotiSelfSignedCert.pem"
	selfSignedKeyName            = "yotiSelfSignedKey.pem"
	portNumber                   = "8080"
	errApplyingTheParsedTemplate = "Error applying the parsed template: "
	errParsingTheTemplate        = "Error parsing the template: "
	profileEndpoint              = "/profile"
	scenarioBuilderErr           = "Scenario Builder Error: `%s`"
	didClient                    *yoti.DigitalIdentityClient
)

func home(w http.ResponseWriter, req *http.Request) {
	templateVars := map[string]interface{}{
		"yotiScenarioID":  os.Getenv("YOTI_SCENARIO_ID"),
		"yotiClientSdkID": os.Getenv("YOTI_CLIENT_SDK_ID")}
	t, err := template.ParseFiles("login.html")

	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(errParsingTheTemplate+err.Error()),
		)))
		return
	}

	err = t.Execute(w, templateVars)
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(errApplyingTheParsedTemplate+err.Error()),
		)))
		return
	}
}
func buildDigitalIdentitySessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	policy, err := (&digitalidentity.PolicyBuilder{}).WithFullName().WithEmail().WithPhoneNumber().WithSelfie().WithAgeOver(18).WithNationality().WithGender().WithDocumentDetails().WithDocumentImages().WithWantedRememberMe().Build()

	if err != nil {
		return nil, err
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-examples"
	}`)

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).WithPolicy(policy).WithRedirectUri("https:/www.yoti.com").WithSubject(subject).Build()
	if err != nil {
		return nil, err
	}
	return &sessionReq, nil
}

func generateSession(w http.ResponseWriter, r *http.Request) {

	initialiseDigitalIdentityClient()
	sessionReq, err := buildDigitalIdentitySessionReq()
	if err != nil {
		fmt.Fprintf(w, string(""))
	}

	shareSession, err := didClient.CreateShareSession(sessionReq)
	if err != nil {
		fmt.Fprintf(w, string(""))
	}

	output, err := json.Marshal(shareSession)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(output))
}

func getReceipt(w http.ResponseWriter, r *http.Request) {
	initialiseDigitalIdentityClient()
	receiptID := r.URL.Query().Get("ReceiptID")

	//urlQUery, err := url.QueryUnescape(receiptID)
	//if err != nil {
	//	fmt.Fprintf(w, string(""))
	//}

	receiptValue, err := didClient.GetShareReceipt(receiptID)
	if err != nil {
		fmt.Fprintf(w, string("Receipt not found"))
	}
	output, err := json.Marshal(receiptValue)
	if err != nil {
		fmt.Fprintf(w, string("Receipt value cannot be marshalled"))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(output))
}

func initialiseDigitalIdentityClient() error {
	var err error
	sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")
	keyFilePath := os.Getenv("YOTI_KEY_FILE_PATH")
	key, err = os.ReadFile(keyFilePath)
	if err != nil {
		return fmt.Errorf("failed to get key from YOTI_KEY_FILE_PATH :: %w", err)
	}

	didClient, err = yoti.NewDigitalIdentityClient(sdkID, key)
	if err != nil {
		return fmt.Errorf("failed to initialise Share client :: %w", err)
	}
	didClient.OverrideAPIURL("https://api.yoti.com/share")

	return nil
}
func main() {
	// Check if the cert files are available.
	certificatePresent := certificatePresenceCheck(selfSignedCertName, selfSignedKeyName)
	// If they are not available, generate new ones.
	if !certificatePresent {
		err := generateSelfSignedCertificate(selfSignedCertName, selfSignedKeyName, "127.0.0.1:"+portNumber)
		if err != nil {
			panic("Error when creating https certs: " + err.Error())
		}
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/v2/generateShare", generateSession)
	http.HandleFunc("/v2/receiptInfo", getReceipt)

	rootdir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error: Couldn't get current working directory")
	}
	http.Handle("/images/", http.StripPrefix("/images",
		http.FileServer(http.Dir(path.Join(rootdir, "images/")))))
	http.Handle("/static/", http.StripPrefix("/static",
		http.FileServer(http.Dir(path.Join(rootdir, "static/")))))

	log.Printf("About to listen and serve on %[1]s. Go to https://localhost:%[1]s/", portNumber)
	err = http.ListenAndServeTLS(":"+portNumber, selfSignedCertName, selfSignedKeyName, nil)

	if err != nil {
		panic("Error when calling `ListenAndServeTLS`: " + err.Error())
	}
}
