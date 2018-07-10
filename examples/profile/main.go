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
	var sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")
	var key, err = ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))

	if err != nil {
		log.Printf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: `%s`", err)
		return
	}

	var client = yoti.Client{
		SdkID: sdkID,
		Key:   key}

	yotiOneTimeUseToken := r.URL.Query().Get("token")
	profile, err := client.GetUserProfile(yotiOneTimeUseToken)

	if err == nil {
		templateVars := map[string]interface{}{
			"profile":         profile,
			"selfieBase64URL": template.URL(profile.Selfie.URL())}

		decodedImage := decodeImage(profile.Selfie.Data)
		file := createImage()
		saveImage(decodedImage, file)

		t, err := template.ParseFiles("profile.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		t.Execute(w, templateVars)
	} else {
		fmt.Println(err)
	}
}

func main() {
	// Check if the cert files are available.
	certificatePresent := certificatePresenceCheck(selfSignedCertName, selfSignedKeyName)
	// If they are not available, generate new ones.
	if certificatePresent == false {
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

func saveImage(img image.Image, file *os.File) {
	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(file, img, &opt)
}
