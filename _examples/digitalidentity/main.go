package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/getyoti/yoti-go-sdk/v3"
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
	http.HandleFunc(profileEndpoint, profile)
	http.HandleFunc("/share", dynamicShare)
	http.HandleFunc("/source-constraints", sourceConstraints)
	http.HandleFunc("/dbs-check", dbsCheck)

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
