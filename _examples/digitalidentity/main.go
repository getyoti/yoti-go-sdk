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
	errApplyingTheParsedTemplate = "Error applying the parsed template: "
	errParsingTheTemplate        = "Error parsing the template: "
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

func estimatedAgeExamplesPage(w http.ResponseWriter, req *http.Request) {
	templateVars := map[string]interface{}{
		"yotiClientSdkID": os.Getenv("YOTI_CLIENT_SDK_ID"),
	}

	t, err := template.ParseFiles("estimated_age_examples.html")
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Error parsing the template: %v", err),
		)))
		return
	}

	err = t.Execute(w, templateVars)
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Error applying the parsed template: %v", err),
		)))
		return
	}
}

func buildDigitalIdentitySessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	policy, err := (&digitalidentity.PolicyBuilder{}).WithFullName().WithEmail().WithPhoneNumber().WithSelfie().EstimatedAgeOver(18, 5).WithNationality().WithGender().WithDocumentDetails().WithDocumentImages().WithWantedRememberMe().Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build policy: %v", err)
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-examples"
	}`)

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).WithPolicy(policy).WithRedirectUri("https:/www.yoti.com").WithSubject(subject).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build create session request: %v", err)
	}
	return &sessionReq, nil
}

func generateSession(w http.ResponseWriter, r *http.Request) {
	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		fmt.Fprintf(w, "Client could't be generated: %v", err)
		return
	}

	sessionReq, err := buildDigitalIdentitySessionReq()
	if err != nil {
		fmt.Fprintf(w, "failed to build session request: %v", err)
		return
	}

	shareSession, err := didClient.CreateShareSession(sessionReq)
	if err != nil {
		fmt.Fprintf(w, "failed to create share session: %v", err)
		return
	}

	output, err := json.Marshal(shareSession)
	if err != nil {
		fmt.Fprintf(w, "failed to marshall share session: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(output))

}

func initialiseDigitalIdentityClient() (*yoti.DigitalIdentityClient, error) {
	var err error
	sdkID := os.Getenv("YOTI_CLIENT_SDK_ID")
	keyFilePath := os.Getenv("YOTI_KEY_FILE_PATH")
	key, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get key from YOTI_KEY_FILE_PATH :: %w", err)
	}

	didClient, err := yoti.NewDigitalIdentityClient(sdkID, key)
	if err != nil {
		return nil, fmt.Errorf("failed to initialise Share client :: %w", err)
	}

	return didClient, nil
}
func main() {
	// Check if the cert files are available.
	selfSignedCertName := "yotiSelfSignedCert.pem"
	selfSignedKeyName := "yotiSelfSignedKey.pem"
	certificatePresent := certificatePresenceCheck(selfSignedCertName, selfSignedKeyName)
	portNumber := "8080"
	// If they are not available, generate new ones.
	if !certificatePresent {
		err := generateSelfSignedCertificate(selfSignedCertName, selfSignedKeyName, "127.0.0.1:"+portNumber)
		if err != nil {
			panic("Error when creating https certs: " + err.Error())
		}
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/estimated-age-examples", estimatedAgeExamplesPage)
	http.HandleFunc("/v2/generate-share", generateSession)
	http.HandleFunc("/v2/generate-advanced-identity-share", generateAdvancedIdentitySession)
	http.HandleFunc("/v2/receipt-info", receipt)

	// Estimated Age Examples
	http.HandleFunc("/v2/generate-estimated-age-share", generateEstimatedAgeSession)
	http.HandleFunc("/v2/generate-estimated-age-over-share", generateEstimatedAgeOverSession)
	http.HandleFunc("/v2/generate-estimated-age-under-share", generateEstimatedAgeUnderSession)
	http.HandleFunc("/v2/generate-estimated-age-constrained-share", generateEstimatedAgeWithConstraintsSession)
	http.HandleFunc("/v2/estimated-age-receipt", estimatedAgeReceipt)
	http.HandleFunc("/v2/age-over-receipt", estimatedAgeReceipt)
	http.HandleFunc("/v2/age-under-receipt", estimatedAgeReceipt)
	http.HandleFunc("/v2/constrained-age-receipt", estimatedAgeReceipt)

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
