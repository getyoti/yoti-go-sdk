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
	"strings"

	"github.com/getyoti/yoti-go-sdk"

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
		"yotiApplicationID": os.Getenv("YOTI_APPLICATION_ID")}

	t, _ := template.ParseFiles("login.html")
	t.Execute(w, templateVars)
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

	var base64URL string
	base64URL, err = userProfile.Selfie().Base64URL()

	if err != nil {
		log.Fatalf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: %q", err)
	}

	dob, err := userProfile.DateOfBirth()
	if err != nil {
		log.Fatalf("Error parsing Date of Birth attribute. Error %q", err)
	}

	templateVars := map[string]interface{}{
		"profile":         userProfile,
		"selfieBase64URL": template.URL(base64URL),
		"rememberMeID":    activityDetails.RememberMeID,
		"dateOfBirth":     dob,
	}

	decodedImage := decodeImage(userProfile.Selfie().Value)
	file := createImage()
	saveImage(decodedImage, file)

	var t *template.Template
	t, err = template.ParseFiles("profile.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, templateVars)
}

func main() {
	// Check if the cert files are available.
	certificatePresent := certificatePresenceCheck(selfSignedCertName, selfSignedKeyName)
	// If they are not available, generate new ones.
	if !certificatePresent {
		err := generateSelfSignedCertificate(selfSignedCertName, selfSignedKeyName, "127.0.0.1:"+portNumber)
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
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

	log.Printf("About to configure HTTPS redirection")
	go configureHTTPSRedirection()

	log.Printf("About to listen and serve on %[1]s. Go to https://localhost:%[1]s/", portNumber)
	http.ListenAndServeTLS(":"+portNumber, selfSignedCertName, selfSignedKeyName, nil)
}

func configureHTTPSRedirection() {
	err := http.ListenAndServe(":80", http.HandlerFunc(redirectHandler))
	if err != nil {
		panic("Error configuring HTTP redirection: " + err.Error())
	}
}

func redirectHandler(w http.ResponseWriter, req *http.Request) {
	hostParts := strings.Split(req.Host, ":")
	http.Redirect(
		w,
		req,
		fmt.Sprintf("https://%s%s", hostParts[0], req.RequestURI),
		http.StatusMovedPermanently)
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

func saveImage(img image.Image, file io.Writer) {
	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(file, img, &opt)
}
