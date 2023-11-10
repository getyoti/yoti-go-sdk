package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/dynamic"
)

func advancedIdentityProfile(w http.ResponseWriter, req *http.Request) {
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

	policy, err := (&dynamic.PolicyBuilder{}).
		WithAdvancedIdentityProfileRequirements(advancedIdentityProfile).
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

	pageFromScenario(w, req, "Advanced Identity Profile Example", scenario)
}
