package main

import (
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/check"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/filter"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/objective"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/task"
)

func buildSessionSpec() (sessionSpec *create.SessionSpecification, err error) {
	var faceMatchCheck *check.RequestedFaceMatchCheck
	faceMatchCheck, err = check.NewRequestedFaceMatchCheckBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		return nil, err
	}

	var documentAuthenticityCheck *check.RequestedDocumentAuthenticityCheck
	documentAuthenticityCheck, err = check.NewRequestedDocumentAuthenticityCheckBuilder().
		Build()
	if err != nil {
		return nil, err
	}

	var livenessCheck *check.RequestedLivenessCheck
	livenessCheck, err = check.NewRequestedLivenessCheckBuilder().
		ForStaticLiveness().
		WithMaxRetries(3).
		Build()
	if err != nil {
		return nil, err
	}

	var idDocsComparisonCheck *check.RequestedIDDocumentComparisonCheck
	idDocsComparisonCheck, err = check.NewRequestedIDDocumentComparisonCheckBuilder().
		Build()
	if err != nil {
		return nil, err
	}

	var thirdPartyCheck *check.RequestedThirdPartyIdentityCheck
	thirdPartyCheck, err = check.NewRequestedThirdPartyIdentityCheckBuilder().
		Build()
	if err != nil {
		return nil, err
	}

	var watchlistScreeningCheck *check.RequestedWatchlistScreeningCheck
	watchlistScreeningCheck, err = check.NewRequestedWatchlistScreeningCheckBuilder().
		WithAdverseMediaCategory().
		WithSanctionsCategory().
		Build()
	if err != nil {
		return nil, err
	}

	yotiAccountWatchlistAdvancedCACheck, err := check.NewRequestedWatchlistAdvancedCACheckYotiAccountBuilder().
		WithRemoveDeceased(true).
		WithShareURL(true).
		WithSources(check.RequestedTypeListSources{
			Types: []string{"pep", "fitness-probity", "warning"}}).
		WithMatchingStrategy(check.RequestedFuzzyMatchingStrategy{Fuzziness: 0.5}).
		Build()
	if err != nil {
		return nil, err
	}

	var textExtractionTask *task.RequestedTextExtractionTask
	textExtractionTask, err = task.NewRequestedTextExtractionTaskBuilder().
		WithManualCheckFallback().
		WithExpandedDocumentFields(true).
		Build()
	if err != nil {
		return nil, err
	}

	var supplementaryDocTextExtractionTask *task.RequestedSupplementaryDocTextExtractionTask
	supplementaryDocTextExtractionTask, err = task.NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		return nil, err
	}

	var sdkConfig *create.SDKConfig
	sdkConfig, err = create.NewSdkConfigBuilder().
		WithAllowsCameraAndUpload().
		WithPrimaryColour("#2d9fff").
		WithSecondaryColour("#FFFFFF").
		WithFontColour("#FFFFFF").
		WithLocale("en-GB").
		WithPresetIssuingCountry("GBR").
		WithSuccessUrl("https://localhost:8080/success").
		WithErrorUrl("https://localhost:8080/error").
		WithPrivacyPolicyUrl("https://localhost:8080/privacy-policy").
		WithIdDocumentTextExtractionGenericAttempts(2).
		WithAllowHandOff(true).
		WithDarkModeOn().
		WithEarlyBiometricConsentFlow().
		//WithBrandId("some_brand_id").
		Build()
	if err != nil {
		return nil, err
	}

	//This section is used for Orthogonal Restriction
	/*passportFilter, err := filter.NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithIncludedDocumentTypes(
			[]string{"PASSPORT"}).
		WithNonLatinDocuments(true).
		WithExpiredDocuments(false).
		Build()
	if err != nil {
		return nil, err
	}*/

	docRestriction, err := filter.NewRequestedDocumentRestrictionBuilder().
		WithDocumentTypes([]string{"PASSPORT"}).
		WithCountryCodes([]string{"GBR"}).
		Build()
	if err != nil {
		return nil, err
	}

	/*passportDoc, err := filter.NewRequiredIDDocumentBuilder().
		WithFilter(passportFilter).
		Build()
	if err != nil {
		return nil, err
	}*/

	docFilter, err := filter.NewRequestedDocumentRestrictionsFilterBuilder().
		ForIncludeList().
		WithDocumentRestriction(docRestriction).
		WithAllowNonLatinDocuments(true).
		WithExpiredDocuments(false).
		Build()
	if err != nil {
		return nil, err
	}

	idDoc, err := filter.NewRequiredIDDocumentBuilder().WithFilter(docFilter).Build()
	if err != nil {
		return nil, err
	}

	proofOfAddressObjective, err := objective.NewProofOfAddressObjectiveBuilder().Build()
	if err != nil {
		return nil, err
	}

	supplementaryDoc, err := filter.NewRequiredSupplementaryDocumentBuilder().
		WithObjective(proofOfAddressObjective).
		Build()
	if err != nil {
		return nil, err
	}

	sessionSpec, err = create.NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(600).
		WithResourcesTTL(87000).
		WithUserTrackingID("some-tracking-id").
		WithRequestedCheck(faceMatchCheck).
		WithRequestedCheck(documentAuthenticityCheck).
		WithRequestedCheck(livenessCheck).
		WithRequestedCheck(idDocsComparisonCheck).
		WithRequestedCheck(thirdPartyCheck).
		WithRequestedCheck(watchlistScreeningCheck).
		WithRequestedCheck(yotiAccountWatchlistAdvancedCACheck).
		WithRequestedTask(textExtractionTask).
		WithRequestedTask(supplementaryDocTextExtractionTask).
		WithSDKConfig(sdkConfig).
		//Below line will be enabled when orthogonal Restriction is Needed
		//WithRequiredDocument(passportDoc).
		WithRequiredDocument(idDoc).
		WithRequiredDocument(supplementaryDoc).
		Build()

	if err != nil {
		return nil, err
	}

	return sessionSpec, nil
}

func buildDBSSessionSpec() (sessionSpec *create.SessionSpecification, err error) {
	var sdkConfig *create.SDKConfig
	sdkConfig, err = create.NewSdkConfigBuilder().
		WithAllowsCameraAndUpload().
		WithPrimaryColour("#2d9fff").
		WithSecondaryColour("#FFFFFF").
		WithFontColour("#FFFFFF").
		WithLocale("en-GB").
		WithPresetIssuingCountry("GBR").
		WithSuccessUrl("https://localhost:8080/success").
		WithErrorUrl("https://localhost:8080/error").
		WithPrivacyPolicyUrl("https://localhost:8080/privacy-policy").
		Build()
	if err != nil {
		return nil, err
	}

	identityProfile := []byte(`{
		"trust_framework": "UK_TFIDA",
		"scheme": {
    		"type": "DBS",
    		"objective": "BASIC"
		}
	}`)

	ttl := time.Hour * 24 * 30
	importToken, err := create.NewImportTokenBuilder().
		WithTTL(int(ttl.Seconds())).
		Build()
	if err != nil {
		return nil, err
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-examples"
	}`)

	sessionSpec, err = create.NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(600).
		WithResourcesTTL(87000).
		WithUserTrackingID("some-tracking-id").
		WithSDKConfig(sdkConfig).
		WithIdentityProfileRequirements(identityProfile).
		WithCreateIdentityProfilePreview(true).
		WithSubject(subject).
		WithImportToken(importToken).
		Build()

	return sessionSpec, nil
}

func buildAdvancedIdentityProfileSessionSpec() (sessionSpec *create.SessionSpecification, err error) {
	var sdkConfig *create.SDKConfig
	sdkConfig, err = create.NewSdkConfigBuilder().
		WithAllowsCameraAndUpload().
		WithPrimaryColour("#2d9fff").
		WithSecondaryColour("#FFFFFF").
		WithFontColour("#FFFFFF").
		WithLocale("en-GB").
		WithPresetIssuingCountry("GBR").
		WithSuccessUrl("https://localhost:8080/success").
		WithErrorUrl("https://localhost:8080/error").
		WithPrivacyPolicyUrl("https://localhost:8080/privacy-policy").
		Build()
	if err != nil {
		return nil, err
	}

	advancedIdentityProfile := []byte(`{
		"profiles": [
			{
				"trust_framework": "UK_TFIDA",
				"schemes": [
					{
						"label": "LB912",
						"type": "RTW"
					}
				]
			},
			{
				"trust_framework": "YOTI_GLOBAL",
				"schemes": [
					{
						"label": "LB321",
						"type": "IDENTITY",
						"objective": "AL_L1"
					}
				]
			}
		]
	}`)

	subject := []byte(`{
		"subject_id": "unique-user-id-for-examples"
	}`)

	return create.NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(6000).
		WithResourcesTTL(900000).
		WithUserTrackingID("some-tracking-id").
		WithSDKConfig(sdkConfig).
		WithAdvancedIdentityProfileRequirements(advancedIdentityProfile).
		WithCreateIdentityProfilePreview(true).
		WithSubject(subject).
		Build()
}

func buildFaceCaptureSessionSpec() (*create.SessionSpecification, error) {
	sdkConfig, err := create.NewSdkConfigBuilder().
		WithSuccessUrl("https://localhost:8080/success").
		WithErrorUrl("https://localhost:8080/error").
		WithPrivacyPolicyUrl("https://localhost:8080/privacy-policy").
		WithJustInTimeBiometricConsentFlow().
		Build()
	if err != nil {
		return nil, err
	}

	faceComparisonCheck, err := check.NewRequestedFaceComparisonCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		return nil, err
	}

	livenessCheck, err := check.NewRequestedLivenessCheckBuilder().
		ForStaticLiveness().
		WithMaxRetries(1).
		WithManualCheckNever().
		Build()
	if err != nil {
		return nil, err
	}

	sessionSpec, err := create.NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(600).
		WithResourcesTTL(90000).
		WithUserTrackingID("simple-face-capture-tracking-id").
		WithRequestedCheck(livenessCheck).
		WithRequestedCheck(faceComparisonCheck).
		WithSDKConfig(sdkConfig).
		Build()
	if err != nil {
		return nil, err
	}

	return sessionSpec, nil
}

// buildSessionSpecWithSuppressedScreens demonstrates how to create a session specification with suppressed screens
func buildSessionSpecWithSuppressedScreens() (sessionSpec *create.SessionSpecification, err error) {
	var faceMatchCheck *check.RequestedFaceMatchCheck
	faceMatchCheck, err = check.NewRequestedFaceMatchCheckBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		return nil, err
	}

	var documentAuthenticityCheck *check.RequestedDocumentAuthenticityCheck
	documentAuthenticityCheck, err = check.NewRequestedDocumentAuthenticityCheckBuilder().
		Build()
	if err != nil {
		return nil, err
	}

	var livenessCheck *check.RequestedLivenessCheck
	livenessCheck, err = check.NewRequestedLivenessCheckBuilder().
		ForStaticLiveness().
		WithMaxRetries(3).
		Build()
	if err != nil {
		return nil, err
	}

	var idDocsComparisonCheck *check.RequestedIDDocumentComparisonCheck
	idDocsComparisonCheck, err = check.NewRequestedIDDocumentComparisonCheckBuilder().
		Build()
	if err != nil {
		return nil, err
	}

	var textDataExtractionTask *task.RequestedTextExtractionTask
	textDataExtractionTask, err = task.NewRequestedTextExtractionTaskBuilder().
		WithManualCheckFallback().
		WithChipDataDesired().
		Build()
	if err != nil {
		return nil, err
	}

	// Create SDK configuration with suppressed screens for a shortened flow
	var sdkConfig *create.SDKConfig
	sdkConfig, err = create.NewSdkConfigBuilder().
		WithAllowsCameraAndUpload().
		WithPrimaryColour("#2d9fff").
		WithSecondaryColour("#FFFFFF").
		WithFontColour("#FFFFFF").
		WithLocale("en-GB").
		WithPresetIssuingCountry("GBR").
		WithSuccessUrl("https://localhost:8080/success").
		WithErrorUrl("https://localhost:8080/error").
		WithPrivacyPolicyUrl("https://localhost:8080/privacy-policy").
		WithIdDocumentTextExtractionGenericAttempts(2).
		WithAllowHandOff(true).
		WithDarkModeOn().
		WithEarlyBiometricConsentFlow().
		// Use the shortened flow convenience method
		WithShortenedFlow().
		Build()
	if err != nil {
		return nil, err
	}

	sessionSpec, err = create.NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(600).
		WithResourcesTTL(90000).
		WithUserTrackingID("shortened-flow-tracking-id").
		WithRequestedCheck(faceMatchCheck).
		WithRequestedCheck(documentAuthenticityCheck).
		WithRequestedCheck(livenessCheck).
		WithRequestedCheck(idDocsComparisonCheck).
		WithRequestedTask(textDataExtractionTask).
		WithSDKConfig(sdkConfig).
		Build()
	if err != nil {
		return nil, err
	}

	return sessionSpec, nil
}
