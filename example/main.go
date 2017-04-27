package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/getyoti/go"

	_ "github.com/joho/godotenv/autoload"
)

var sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")
var key, _ = ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
var client = yoti.YotiClient{
	SdkID: sdkID,
	Key:   key}

func home(w http.ResponseWriter, req *http.Request) {
	templateVars := map[string]interface{}{
		"yotiApplicationID": os.Getenv("YOTI_APPLICATION_ID")}

	t, _ := template.ParseFiles("login.html")
	t.Execute(w, templateVars)
}

func profile(w http.ResponseWriter, r *http.Request) {
	yotiToken := r.URL.Query().Get("token")
	profile, err := client.GetUserProfile(yotiToken)

	if err == nil {
		templateVars := map[string]interface{}{
			"profile": profile,
			"selfie":  template.URL(profile.Selfie.URL)}

		t, _ := template.ParseFiles("profile.html")
		t.Execute(w, templateVars)
	} else {
		fmt.Printf("%+v\n", err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/profile", profile)
	http.ListenAndServe(":8080", nil)
}
