package main

import (
	bytes "bytes"
	"context"
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
	dynamic "github.com/getyoti/yoti-go-sdk/v2/dynamic_sharing_service"
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
	ScenarioBuilderErr           = "Scenario Builder Error: `%s`"
)

func home(w http.ResponseWriter, req *http.Request) {
	templateVars := map[string]interface{}{
		"yotiScenarioID":  os.Getenv("YOTI_SCENARIO_ID"),
		"yotiClientSdkID": os.Getenv("YOTI_CLIENT_SDK_ID")}

	t, err := template.ParseFiles("login.html")

	if err != nil {
		panic(errParsingTheTemplate + err.Error())
	}

	err = t.Execute(w, templateVars)

	if err != nil {
		panic(errApplyingTheParsedTemplate + err.Error())
	}
}

func sourceConstraints(w http.ResponseWriter, req *http.Request) {
	constraint, err := (&dynamic.SourceConstraintBuilder{}).WithDrivingLicence("").WithPassport("").Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Constraint Builder Error: `%s`", err),
		)))
		return
	}

	policy, err := (&dynamic.DynamicPolicyBuilder{}).WithFullName(constraint).WithStructuredPostalAddress(constraint).Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Policy Builder Error: `%s`", err),
		)))
		return
	}

	scenario, err := (&dynamic.DynamicScenarioBuilder{}).WithPolicy(policy).
		WithCallbackEndpoint(profileEndpoint).Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(ScenarioBuilderErr, err),
		)))
		return
	}

	pageFromScenario(w, req, "Source Constraint example", scenario)
}

func dynamicShare(w http.ResponseWriter, req *http.Request) {
	policy, err := (&dynamic.DynamicPolicyBuilder{}).WithFullName().WithEmail().Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(ScenarioBuilderErr, err),
		)))
		return
	}
	scenario, err := (&dynamic.DynamicScenarioBuilder{}).WithPolicy(policy).WithCallbackEndpoint(profileEndpoint).Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(ScenarioBuilderErr, err),
		)))
		return
	}

	pageFromScenario(w, req, "Dynamic Share example", scenario)
}

func pageFromScenario(w http.ResponseWriter, req *http.Request, title string, scenario dynamic.DynamicScenario) {
	sdkID := os.Getenv("YOTI_CLIENT_SDK_ID")

	key, err := ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: `%s`", err.Error()),
		)))
		return
	}

	client, err := yoti.NewClient(sdkID, key)
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("%s", err),
		)))
	}

	share, err := yoti.CreateShareURL(client, &scenario)
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("%s", err.Error()),
		)))
		return
	}

	templateVars := map[string]interface{}{
		"pageTitle":       title,
		"yotiClientSdkID": sdkID,
		"yotiShareURL":    share.ShareURL,
	}

	t, err := template.ParseFiles("dynamic-share.html")
	if err != nil {
		panic("Error parsing template: " + err.Error())
	}

	err = t.Execute(w, templateVars)
	if err != nil {
		panic("Error applying the parsed template: " + err.Error())
	}
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	templateVars := map[string]interface{}{
		"yotiError": r.Context().Value(contextKey("yotiError")).(string),
	}
	log.Printf("%s", templateVars["yotiError"])
	t, err := template.ParseFiles("error.html")
	if err != nil {
		panic(errParsingTheTemplate + err.Error())
	}

	err = t.Execute(w, templateVars)
	if err != nil {
		panic(errApplyingTheParsedTemplate + err.Error())
	}

}

func profile(w http.ResponseWriter, r *http.Request) {
	var err error
	key, err = ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")

	if err != nil {
		errorPage(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: `%s`", err.Error()),
		)))
		return
	}

	client, err = yoti.NewClient(sdkID, key)
	if err != nil {
		errorPage(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("%s", err),
		)))
	}

	yotiOneTimeUseToken := r.URL.Query().Get("token")

	activityDetails, err := client.GetActivityDetails(yotiOneTimeUseToken)
	if err != nil {
		errorPage(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey("yotiError"),
			err.Error(),
		)))
		return
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
		errorPage(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Error parsing Date of Birth attribute. Error %q", err.Error()),
		)))
		return
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
	t, err = template.New("profile.html").
		Funcs(template.FuncMap{
			"escapeURL": func(s string) template.URL {
				return template.URL(s)
			},
			"marshalAttribute": func(name string, icon string, property interface{}, prevalue string) interface{} {
				return struct {
					Name     string
					Icon     string
					Prop     interface{}
					Prevalue string
				}{
					name,
					icon,
					property,
					prevalue,
				}
			},
		}).
		ParseFiles("profile.html")
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
	http.HandleFunc(profileEndpoint, profile)
	http.HandleFunc("/dynamic-share", dynamicShare)
	http.HandleFunc("/source-constraints", sourceConstraints)

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
