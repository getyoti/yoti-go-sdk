package main

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/check"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/filter"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/objective"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/task"
)

func buildSessionSpec() (sessionSpec *create.SessionSpecification, err error) {
	var faceMatchCheck *check.RequestedFaceMatchCheck
	faceMatchCheck, err = check.NewRequestedFaceMatchCheckBuilder().
		WithManualCheckAlways().
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
		ForZoomLiveness().
		WithMaxRetries(5).
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
		WithManualCheckAlways().
		Build()
	if err != nil {
		return nil, err
	}

	var supplementaryDocTextExtractionTask *task.RequestedSupplementaryDocTextExtractionTask
	supplementaryDocTextExtractionTask, err = task.NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckAlways().
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
		Build()
	if err != nil {
		return nil, err
	}

	passportFilter, err := filter.NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithIncludedDocumentTypes(
			[]string{"PASSPORT"}).
		WithExpiredDocuments(true).
		Build()
	if err != nil {
		return nil, err
	}
	passportDoc, err := filter.NewRequiredIDDocumentBuilder().
		WithFilter(passportFilter).
		Build()
	if err != nil {
		return nil, err
	}

	idDoc, err := filter.NewRequiredIDDocumentBuilder().Build()
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
		WithResourcesTTL(90000).
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
		WithRequiredDocument(passportDoc).
		WithRequiredDocument(idDoc).
		WithRequiredDocument(supplementaryDoc).
		Build()

	fmt.Println("%v", sessionSpec)
	data, _ := json.Marshal(sessionSpec)
	fmt.Println(string(data))

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

	subject := []byte(`{
		"subject_id": "unique-user-id-for-examples"
	}`)

	sessionSpec, err = create.NewSessionSpecificationBuilder().
		WithClientSessionTokenTTL(6000).
		WithResourcesTTL(900000).
		WithUserTrackingID("some-tracking-id").
		WithSDKConfig(sdkConfig).
		WithIdentityProfileRequirements(identityProfile).
		WithCreateIdentityProfilePreview(true).
		WithSubject(subject).
		Build()

	if err != nil {
		return nil, err
	}
	return sessionSpec, nil
}
