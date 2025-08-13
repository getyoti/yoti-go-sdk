package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
)

// buildEstimatedAgeOverSessionReq creates a Digital Identity session request that
// verifies the user is over a specific age using estimated_age with date_of_birth fallback
func buildEstimatedAgeOverSessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	policy, err := (&digitalidentity.PolicyBuilder{}).
		WithFullName().
		WithEmail().
		EstimatedAgeOver(18, 5). // Estimated age checks for 23, date_of_birth fallback checks for 18
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build age over policy: %v", err)
	}

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).
		WithPolicy(policy).
		WithRedirectUri("https://localhost:8080/v2/age-over-receipt").
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build create session request: %v", err)
	}
	return &sessionReq, nil
}

// buildEstimatedAgeWithConstraintsSessionReq creates a Digital Identity session request that
// requests estimated age with source constraints (e.g., passport verification)
func buildEstimatedAgeWithConstraintsSessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	// Create a source constraint requiring passport verification
	constraint, err := (&digitalidentity.SourceConstraintBuilder{}).
		WithPassport("").
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build constraint: %v", err)
	}

	policy, err := (&digitalidentity.PolicyBuilder{}).
		WithFullName().
		WithEmail().
		EstimatedAgeOver(18, 5, &constraint). // Estimated age checks for 23, date_of_birth fallback checks for 18
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build constrained age policy: %v", err)
	}

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).
		WithPolicy(policy).
		WithRedirectUri("https://localhost:8080/v2/constrained-age-receipt").
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build create session request: %v", err)
	}
	return &sessionReq, nil
}

// generateEstimatedAgeOverSession handles requests to create an age over verification session
func generateEstimatedAgeOverSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		http.Error(w, `{"error": "Client couldn't be generated"}`, http.StatusInternalServerError)
		return
	}

	sessionReq, err := buildEstimatedAgeOverSessionReq()
	if err != nil {
		http.Error(w, `{"error": "failed to build session request"}`, http.StatusInternalServerError)
		return
	}

	shareSession, err := didClient.CreateShareSession(sessionReq)
	if err != nil {
		http.Error(w, `{"error": "failed to create share session"}`, http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(shareSession)
	if err != nil {
		http.Error(w, `{"error": "failed to marshall share session"}`, http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(output))
}

// generateEstimatedAgeSession handles requests to create a basic estimated age verification session
func generateEstimatedAgeSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		http.Error(w, `{"error": "Client couldn't be generated"}`, http.StatusInternalServerError)
		return
	}

	sessionReq, err := buildEstimatedAgeOverSessionReq()
	if err != nil {
		http.Error(w, `{"error": "failed to build session request"}`, http.StatusInternalServerError)
		return
	}

	shareSession, err := didClient.CreateShareSession(sessionReq)
	if err != nil {
		http.Error(w, `{"error": "failed to create share session"}`, http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(shareSession)
	if err != nil {
		http.Error(w, `{"error": "failed to marshall share session"}`, http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(output))
}

// generateEstimatedAgeWithConstraintsSession handles requests to create an age verification session with constraints
func generateEstimatedAgeWithConstraintsSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		http.Error(w, `{"error": "Client couldn't be generated"}`, http.StatusInternalServerError)
		return
	}

	sessionReq, err := buildEstimatedAgeWithConstraintsSessionReq()
	if err != nil {
		http.Error(w, `{"error": "failed to build session request"}`, http.StatusInternalServerError)
		return
	}

	shareSession, err := didClient.CreateShareSession(sessionReq)
	if err != nil {
		http.Error(w, `{"error": "failed to create share session"}`, http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(shareSession)
	if err != nil {
		http.Error(w, `{"error": "failed to marshall share session"}`, http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(output))
}
