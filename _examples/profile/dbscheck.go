package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/dynamic"
)

func dbsCheck(w http.ResponseWriter, req *http.Request) {
	identityProfile := []byte(`{
		"trust_framework": "UK_TFIDA",
		"scheme": {
			"type":      "DBS",
			"objective": "BASIC"
		}
	}`)

	policy, err := (&dynamic.PolicyBuilder{}).
		WithIdentityProfileRequirements(identityProfile).
		Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(scenarioBuilderErr, err),
		)))
		return
	}

	subject := []byte(`{
		"subject_id": "my_subject_id"
	}`)
	scenario, err := (&dynamic.ScenarioBuilder{}).
		WithPolicy(policy).
		WithSubject(subject).
		WithCallbackEndpoint(profileEndpoint).Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(scenarioBuilderErr, err),
		)))
		return
	}

	pageFromScenario(w, req, "DBS Check Example", scenario)
}
