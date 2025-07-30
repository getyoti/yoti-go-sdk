package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
)

// buildEstimatedAgeSessionReq creates a Digital Identity session request that
// requests the estimated_age attribute with automatic fallback to date_of_birth
func buildEstimatedAgeSessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	policy, err := (&digitalidentity.PolicyBuilder{}).
		WithFullName().
		WithEmail().
		WithEstimatedAge(). // This will request estimated_age with date_of_birth fallback
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build estimated age policy: %v", err)
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-estimated-age-example"
	}`)

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).
		WithPolicy(policy).
		WithRedirectUri("https://localhost:8080/v2/estimated-age-receipt").
		WithSubject(subject).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build create session request: %v", err)
	}
	return &sessionReq, nil
}

// buildEstimatedAgeOverSessionReq creates a Digital Identity session request that
// verifies the user is over a specific age using estimated_age with date_of_birth fallback
func buildEstimatedAgeOverSessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	policy, err := (&digitalidentity.PolicyBuilder{}).
		WithFullName().
		WithEmail().
		WithEstimatedAgeOver(18). // Age verification with fallback
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build age over policy: %v", err)
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-age-over-example"
	}`)

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).
		WithPolicy(policy).
		WithRedirectUri("https://localhost:8080/v2/age-over-receipt").
		WithSubject(subject).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build create session request: %v", err)
	}
	return &sessionReq, nil
}

// buildEstimatedAgeUnderSessionReq creates a Digital Identity session request that
// verifies the user is under a specific age using estimated_age with date_of_birth fallback
func buildEstimatedAgeUnderSessionReq() (sessionSpec *digitalidentity.ShareSessionRequest, err error) {
	policy, err := (&digitalidentity.PolicyBuilder{}).
		WithFullName().
		WithEmail().
		WithEstimatedAgeUnder(21). // Age verification with fallback
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build age under policy: %v", err)
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-age-under-example"
	}`)

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).
		WithPolicy(policy).
		WithRedirectUri("https://localhost:8080/v2/age-under-receipt").
		WithSubject(subject).
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
		WithEstimatedAgeOver(18, &constraint). // Age verification with constraint and fallback
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build constrained age policy: %v", err)
	}

	subject := []byte(`{
		"subject_id": "unique-user-id-for-constrained-age-example"
	}`)

	sessionReq, err := (&digitalidentity.ShareSessionRequestBuilder{}).
		WithPolicy(policy).
		WithRedirectUri("https://localhost:8080/v2/constrained-age-receipt").
		WithSubject(subject).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build create session request: %v", err)
	}
	return &sessionReq, nil
}

// generateEstimatedAgeSession handles requests to create a basic estimated age session
func generateEstimatedAgeSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		http.Error(w, `{"error": "Client couldn't be generated"}`, http.StatusInternalServerError)
		return
	}

	sessionReq, err := buildEstimatedAgeSessionReq()
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

// generateEstimatedAgeUnderSession handles requests to create an age under verification session
func generateEstimatedAgeUnderSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	didClient, err := initialiseDigitalIdentityClient()
	if err != nil {
		http.Error(w, `{"error": "Client couldn't be generated"}`, http.StatusInternalServerError)
		return
	}

	sessionReq, err := buildEstimatedAgeUnderSessionReq()
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

// generateEstimatedAgeWithConstraintsSession handles requests to create a constrained age verification session
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
