package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
)

var advancedIdentityProfile = []byte(`{
		"profiles": [
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

func buildAdvancedIdentitySessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	policy, err := (&digitalidentity.PolicyBuilder{}).WithAdvancedIdentityProfileRequirements(advancedIdentityProfile).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build Advanced Identity Requirements policy: %v", err)
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-examples"
	}`)

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).WithPolicy(policy).WithRedirectUri("https://localhost:8080/v2/receipt-info").WithSubject(subject).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build create session request: %v", err)
	}
	return &sessionReq, nil
}

func generateAdvancedIdentitySession(w http.ResponseWriter, r *http.Request) {
	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		fmt.Fprintf(w, "Client could't be generated: %v", err)
		return
	}

	sessionReq, err := buildAdvancedIdentitySessionReq()
	if err != nil {
		fmt.Fprintf(w, "failed to build session request: %v", err)
		return
	}

	shareSession, err := didClient.CreateShareSession(sessionReq)
	if err != nil {
		fmt.Fprintf(w, "failed to create share session: %v", err)
		return
	}

	output, err := json.Marshal(shareSession)
	if err != nil {
		fmt.Fprintf(w, "failed to marshall share session: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(output))

}
