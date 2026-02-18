package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
)

func estimatedAgeReceipt(w http.ResponseWriter, r *http.Request) {
	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		fmt.Fprintf(w, "Client couldn't be generated")
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
	}

	// Get estimated age with fallback logic
	estimatedAge := userProfile.EstimatedAge()
	result, isEstimatedAge := userProfile.EstimatedAgeWithFallback()

	var estimatedAgeString string
	var usedEstimatedAge bool
	var usedFallback bool

	if result != nil {
		if isEstimatedAge {
			// estimated_age was returned
			usedEstimatedAge = true
			if estimatedAge != nil {
				estimatedAgeString = estimatedAge.Value()
			}
		} else {
			// date_of_birth was returned as fallback
			usedFallback = true
		}
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

	// Get age verifications (e.g., age_over:18, age_under:21)
	ageVerifications, err := userProfile.AgeVerifications()
	if err != nil {
		errorPage(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Error parsing Age Verifications. Error %q", err.Error()),
		)))
		return
	}

	// Analyze which source was used for each age verification
	ageVerificationDetails := analyzeAgeVerificationSources(ageVerifications)

	templateVars := map[string]interface{}{
		"profile":                userProfile,
		"selfieBase64URL":        template.URL(base64URL),
		"rememberMeID":           receiptValue.RememberMeID,
		"dateOfBirth":            dateOfBirthString,
		"estimatedAge":           estimatedAgeString,
		"usedEstimatedAge":       usedEstimatedAge,
		"usedFallback":           usedFallback,
		"hasEstimatedAge":        estimatedAge != nil,
		"hasDateOfBirth":         dob != nil,
		"ageVerifications":       ageVerifications,
		"ageVerificationDetails": ageVerificationDetails,
		"fullReceipt":            receiptValue, // Add full receipt for JSON display
		"profileAttributes": map[string]interface{}{
			"fullName":     userProfile.FullName(),
			"givenNames":   userProfile.GivenNames(),
			"familyName":   userProfile.FamilyName(),
			"emailAddress": userProfile.EmailAddress(),
			"mobileNumber": userProfile.MobileNumber(),
			"nationality":  userProfile.Nationality(),
			"address":      userProfile.Address(),
			"selfie":       userProfile.Selfie(),
			"estimatedAge": userProfile.EstimatedAge(),
			"dateOfBirth":  dob,
		},
	}

	var t *template.Template
	t, err = template.New("estimated_age_receipt.html").
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
				if data == nil {
					return "null"
				}
				json, err := json.MarshalIndent(data, "", "  ")
				if err != nil {
					return fmt.Sprintf("Error marshaling JSON: %v\nData type: %T", err, data)
				}
				return string(json)
			},
		}).
		ParseFiles("estimated_age_receipt.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, templateVars)

	if err != nil {
		errorPage(w, r.WithContext(context.WithValue(
			r.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Error applying the parsed estimated age template. Error: `%s`", err),
		)))
		return
	}
}

// AgeVerificationDetail contains information about an age verification and its source
type AgeVerificationDetail struct {
	Verification     attribute.AgeVerification
	SourceType       string   // "ESTIMATED_AGE", "DATE_OF_BIRTH", or "UNKNOWN"
	SourceAnchors    []string // List of source anchor values
	UsedEstimatedAge bool     // True if derived from estimated_age
	UsedFallback     bool     // True if derived from date_of_birth fallback
}

// analyzeAgeVerificationSources examines the anchors of age verifications to determine
// which source (estimated_age or date_of_birth) was actually used for each check
func analyzeAgeVerificationSources(verifications []attribute.AgeVerification) []AgeVerificationDetail {
	details := make([]AgeVerificationDetail, 0, len(verifications))

	for _, verification := range verifications {
		detail := AgeVerificationDetail{
			Verification:  verification,
			SourceType:    "UNKNOWN",
			SourceAnchors: []string{},
		}

		// Parse anchors from the protobuf attribute
		if verification.Attribute != nil && len(verification.Attribute.Anchors) > 0 {
			anchors := anchor.ParseAnchors(verification.Attribute.Anchors)

			// Get SOURCE anchors only
			sources := anchor.GetSources(anchors)

			for _, src := range sources {
				sourceValue := src.Value()
				detail.SourceAnchors = append(detail.SourceAnchors, sourceValue)

				// Check if the source indicates estimated_age or date_of_birth
				upperValue := strings.ToUpper(sourceValue)
				if strings.Contains(upperValue, "ESTIMATED") || strings.Contains(upperValue, "AGE_ESTIMATE") {
					detail.SourceType = "ESTIMATED_AGE"
					detail.UsedEstimatedAge = true
				} else if strings.Contains(upperValue, "DATE_OF_BIRTH") || strings.Contains(upperValue, "DOB") {
					detail.SourceType = "DATE_OF_BIRTH"
					detail.UsedFallback = true
				}
			}
		}

		// If we couldn't determine from anchors, mark as unknown
		if len(detail.SourceAnchors) == 0 {
			detail.SourceAnchors = []string{"No source anchors found"}
		}

		details = append(details, detail)
	}

	return details
}
