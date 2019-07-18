package main

import (
	bytes "bytes"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	yoti "github.com/getyoti/yoti-go-sdk/v2"
	_ "github.com/joho/godotenv/autoload"
)

var (
	sdkID              string
	key                []byte
	client             *yoti.Client
	selfSignedCertName = "yotiSelfSignedCert.pem"
	selfSignedKeyName  = "yotiSelfSignedKey.pem"
	portNumber         = "8080"
)

func home(w http.ResponseWriter, req *http.Request) {
	templateVars := map[string]interface{}{
		"yotiScenarioID":    os.Getenv("YOTI_SCENARIO_ID"),

	t, err := template.ParseFiles("login.html")

	if err != nil {
		panic("Error parsing the template: " + err.Error())
	}

	err = t.Execute(w, templateVars)

	if err != nil {
		panic("Error applying the parsed template: " + err.Error())
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	var err error
	key, err = ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")

	if err != nil {
		log.Fatalf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: `%s`", err)
	}

	client = &yoti.Client{
		SdkID: sdkID,
		Key:   key}

	yotiOneTimeUseToken := r.URL.Query().Get("token")

	activityDetails, errStrings := client.GetActivityDetails(yotiOneTimeUseToken)
	if len(errStrings) != 0 {
		log.Fatalf("Errors: %v", errStrings)
	}

	userProfile := activityDetails.UserProfile

	selfie := userProfile.Selfie()
	var base64URL string
	if selfie != nil {
		base64URL = selfie.Value().Base64URL()

		decodedImage := decodeImage(selfie.Value().Data)
		file := createImage()
		saveImage(decodedImage, file)
	}

	dob, err := userProfile.DateOfBirth()
	if err != nil {
		log.Fatalf("Error parsing Date of Birth attribute. Error %q", err)
	}

	var dateOfBirthString string
	if dob != nil {
		dateOfBirthString = dob.Value().String()
	}

	templateVars := map[string]interface{}{
		"profile":         userProfile,
		"selfieBase64URL": template.URL(base64URL),
		"rememberMeID":    activityDetails.RememberMeID(),
		"dateOfBirth":     dateOfBirthString,
	}

	var t *template.Template
	t, err = template.ParseFiles("profile.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, templateVars)

	if err != nil {
		panic("Error applying the parsed profile template. Error: " + err.Error())
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
	http.HandleFunc("/profile", profile)

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

func decodeImage(imageBytes []byte) image.Image {
	decodedImage, _, err := image.Decode(bytes.NewReader(imageBytes))

	if err != nil {
		panic("Error when decoding the image: " + err.Error())
	}

	return decodedImage
}

func createImage() (file *os.File) {
	file, err := os.Create("./images/YotiSelfie.jpeg")

	if err != nil {
		panic("Error when creating the image: " + err.Error())
	}
	return
}

func saveImage(img image.Image, file io.Writer) {
	var opt jpeg.Options
	opt.Quality = 100

	err := jpeg.Encode(file, img, &opt)

	if err != nil {
		panic("Error when saving the image: " + err.Error())
	}
}
