package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
)

func receipt(w http.ResponseWriter, r *http.Request) {
	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		fmt.Fprintf(w, "Client could't be generated")
		return
	}
	receiptID := r.URL.Query().Get("ReceiptID")

	receiptValue, err := didClient.GetShareReceipt(receiptID)
	if err != nil {
		fmt.Fprintf(w, "failed to get share receipt: %v", err)
		return
	}

	if receiptValue.Error != "" {
		t, err := template.ParseFiles("error_receipt.html", "requirements_not_met_detail.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		templateVars := map[string]interface{}{
			"error":       receiptValue.Error,
			"errorReason": receiptValue.ErrorReason,
		}
		err = t.Execute(w, templateVars)
		if err != nil {
			errorPage(w, r.WithContext(context.WithValue(
				r.Context(),
				contextKey("yotiError"),
				fmt.Sprintf("Error applying the parsed error_receipt template. Error: `%s`", err),
			)))
			return
		}
		return
	}

	userProfile := receiptValue.UserContent.UserProfile

	selfie := userProfile.Selfie()

	var base64URL string
	if selfie != nil {
		base64URL = selfie.Value().Base64URL()

		decodedImage := decodeImage(selfie.Value().Data())
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
		"rememberMeID":    receiptValue.RememberMeID,
		"dateOfBirth":     dateOfBirthString,
	}

	var t *template.Template
	t, err = template.New("receipt.html").
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
			"jsonMarshalIndent": func(data interface{}) string {
				json, err := json.MarshalIndent(data, "", "\t")
				if err != nil {
					fmt.Println(err)
				}
				return string(json)
			},
		}).
		ParseFiles("receipt.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, templateVars)

	if err != nil {
		errorPage(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Error applying the parsed profile template. Error: `%s`", err),
		)))
		return
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
