package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3"
	"github.com/getyoti/yoti-go-sdk/v3/dynamic"
)

func dynamicShare(w http.ResponseWriter, req *http.Request) {
	policy, err := (&dynamic.PolicyBuilder{}).WithFullName().WithEmail().Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(scenarioBuilderErr, err),
		)))
		return
	}
	scenario, err := (&dynamic.ScenarioBuilder{}).WithPolicy(policy).WithCallbackEndpoint(profileEndpoint).Build()
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf(scenarioBuilderErr, err),
		)))
		return
	}

	pageFromScenario(w, req, "Dynamic Share example", scenario)
}

func pageFromScenario(w http.ResponseWriter, req *http.Request, title string, scenario dynamic.Scenario) {
	sdkID := os.Getenv("YOTI_CLIENT_SDK_ID")

	key, err := os.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
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

	share, err := client.CreateShareURL(&scenario)
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

	var t *template.Template
	t, err = template.ParseFiles("dynamic-share.html")
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("error parsing template: "+err.Error()),
		)))
		return
	}

	err = t.Execute(w, templateVars)
	if err != nil {
		errorPage(w, req.WithContext(context.WithValue(
			req.Context(),
			contextKey("yotiError"),
			fmt.Sprintf("error applying the parsed template: "+err.Error()),
		)))
		return
	}
}
