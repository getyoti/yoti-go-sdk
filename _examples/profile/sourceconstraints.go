package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/dynamic"
)

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

	policy, err := (&dynamic.PolicyBuilder{}).WithFullName(constraint).WithStructuredPostalAddress(constraint).Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("Policy Builder Error: `%s`", err),
		)))
		return
	}

	scenario, err := (&dynamic.ScenarioBuilder{}).WithPolicy(policy).
		WithCallbackEndpoint(profileEndpoint).Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(scenarioBuilderErr, err),
		)))
		return
	}

	pageFromScenario(w, req, "Source Constraint example", scenario)
}
