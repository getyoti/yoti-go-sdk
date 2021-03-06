package main

import (
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

	var textExtractionTask *task.RequestedTextExtractionTask
	textExtractionTask, err = task.NewRequestedTextExtractionTaskBuilder().
		WithManualCheckAlways().
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

	var supplementaryDocTextExtractionTask *task.RequestedSupplementaryDocTextExtractionTask
	supplementaryDocTextExtractionTask, err = task.NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckAlways().
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

	passportFilter, err := filter.NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithIncludedDocumentTypes(
			[]string{"PASSPORT"}).
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
		WithRequestedTask(textExtractionTask).
		WithRequestedTask(supplementaryDocTextExtractionTask).
		WithSDKConfig(sdkConfig).
		WithRequiredDocument(passportDoc).
		WithRequiredDocument(idDoc).
		WithRequiredDocument(supplementaryDoc).
		Build()

	if err != nil {
		return nil, err
	}
	return sessionSpec, nil
}
