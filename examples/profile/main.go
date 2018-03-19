package main

import (
	bytes "bytes"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/getyoti/yoti-go-sdk"

	_ "github.com/joho/godotenv/autoload"
)

var (
	sdkID  string
	key    []byte
	client *yoti.Client
)

func home(w http.ResponseWriter, req *http.Request) {
	templateVars := map[string]interface{}{
		"yotiApplicationID": os.Getenv("YOTI_APPLICATION_ID")}

	t, _ := template.ParseFiles("login.html")
	t.Execute(w, templateVars)
}

func profile(w http.ResponseWriter, r *http.Request) {
	var sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")
	var key, err = ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))

	if err != nil {
		log.Printf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: `%s`", err)
		return
	}

	var client = yoti.Client{
		SdkID: sdkID,
		Key:   key}

	yotiToken := r.URL.Query().Get("token")
	profile, err := client.GetUserProfile(yotiToken)

	if err == nil {
		templateVars := map[string]interface{}{
			"profile":         profile,
			"selfieBase64URL": template.URL(profile.Selfie.URL())}

		// This key uses the  format: age_[over|under]:[1-999] and is dynamically
		// generated based on the dashboard attribute Age / Verify Condition
		templateVars["AgeVerified"] = string(profile.OtherAttributes["age_over:18"].Value)
		decodedImage := decodeImage(profile.Selfie.Data)
		file := createImage()
		saveImage(decodedImage, file)

		t, _ := template.ParseFiles("profile.html")
		t.Execute(w, templateVars)
	} else {
		fmt.Printf("%+v\n", err)
	}
}

func main() {
	rootdir, _ := os.Getwd()
	http.HandleFunc("/", home)
	http.HandleFunc("/profile", profile)
	http.Handle("/images/", http.StripPrefix("/images",
		http.FileServer(http.Dir(path.Join(rootdir, "images/")))))
	log.Printf("About to listen and serve on 8080. Go to 127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}

func decodeImage(imageBytes []byte) (decodedImage image.Image) {
	decodedImage, _, _ = image.Decode(bytes.NewReader(imageBytes))
	return
}

func createImage() (file *os.File) {
	file, err := os.Create("./images/YotiSelfie.jpeg")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func saveImage(img image.Image, file *os.File) {
	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(file, img, &opt)
}
