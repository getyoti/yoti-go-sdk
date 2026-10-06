package main

import (
	"fmt"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// TestAnalyzeAgeVerificationSources demonstrates how the anchor analysis works
func TestAnalyzeAgeVerificationSources(t *testing.T) {
	// Create a mock age verification without anchors for basic testing
	// In real scenarios, anchors would be populated from the API response
	ageOverAttr := &yotiprotoattr.Attribute{
		Name:    "age_over:18",
		Value:   []byte("true"),
		Anchors: []*yotiprotoattr.Anchor{},
	}

	ageVerification, err := attribute.NewAgeVerification(ageOverAttr)
	if err != nil {
		t.Fatalf("Failed to create age verification: %v", err)
	}

	// Analyze the source
	details := analyzeAgeVerificationSources([]attribute.AgeVerification{ageVerification})

	if len(details) != 1 {
		t.Fatalf("Expected 1 detail, got %d", len(details))
	}

	detail := details[0]

	fmt.Printf("Age Verification: %s:%d\n", detail.Verification.CheckType, detail.Verification.Age)
	fmt.Printf("Result: %v\n", detail.Verification.Result)
	fmt.Printf("Source Type: %s\n", detail.SourceType)
	fmt.Printf("Used Estimated Age: %v\n", detail.UsedEstimatedAge)
	fmt.Printf("Used Fallback: %v\n", detail.UsedFallback)
	fmt.Printf("Source Anchors: %v\n", detail.SourceAnchors)

	// Verify the age verification was parsed correctly
	if detail.Verification.CheckType != "age_over" {
		t.Errorf("Expected CheckType 'age_over', got '%s'", detail.Verification.CheckType)
	}

	if detail.Verification.Age != 18 {
		t.Errorf("Expected Age 18, got %d", detail.Verification.Age)
	}

	if !detail.Verification.Result {
		t.Error("Expected Result to be true")
	}
}

// TestAnalyzeMultipleAgeVerifications tests handling of multiple verifications
func TestAnalyzeMultipleAgeVerifications(t *testing.T) {
	// Create multiple age verifications
	ageOver18 := &yotiprotoattr.Attribute{
		Name:    "age_over:18",
		Value:   []byte("true"),
		Anchors: []*yotiprotoattr.Anchor{},
	}

	ageUnder25 := &yotiprotoattr.Attribute{
		Name:    "age_under:25",
		Value:   []byte("false"),
		Anchors: []*yotiprotoattr.Anchor{},
	}

	verification1, _ := attribute.NewAgeVerification(ageOver18)
	verification2, _ := attribute.NewAgeVerification(ageUnder25)

	// Analyze the sources
	details := analyzeAgeVerificationSources([]attribute.AgeVerification{verification1, verification2})

	if len(details) != 2 {
		t.Fatalf("Expected 2 details, got %d", len(details))
	}

	fmt.Printf("\nMultiple Verifications:\n")
	for i, detail := range details {
		fmt.Printf("Verification %d: %s:%d = %v\n", i+1, detail.Verification.CheckType, detail.Verification.Age, detail.Verification.Result)
		fmt.Printf("  Source Type: %s\n", detail.SourceType)
		fmt.Printf("  Source Anchors: %v\n", detail.SourceAnchors)
	}

	// Verify both were parsed correctly
	if details[0].Verification.Age != 18 {
		t.Errorf("Expected first verification Age 18, got %d", details[0].Verification.Age)
	}

	if details[1].Verification.Age != 25 {
		t.Errorf("Expected second verification Age 25, got %d", details[1].Verification.Age)
	}
}

// TestAnalyzeAgeVerificationSourcesWithNoAnchors tests the case when no anchors are present
func TestAnalyzeAgeVerificationSourcesWithNoAnchors(t *testing.T) {
	// Create a mock age verification without anchors
	ageOverAttr := &yotiprotoattr.Attribute{
		Name:    "age_over:25",
		Value:   []byte("false"),
		Anchors: []*yotiprotoattr.Anchor{},
	}

	ageVerification, err := attribute.NewAgeVerification(ageOverAttr)
	if err != nil {
		t.Fatalf("Failed to create age verification: %v", err)
	}

	// Analyze the source
	details := analyzeAgeVerificationSources([]attribute.AgeVerification{ageVerification})

	if len(details) != 1 {
		t.Fatalf("Expected 1 detail, got %d", len(details))
	}

	detail := details[0]

	fmt.Printf("\nAge Verification (no anchors): %s:%d\n", detail.Verification.CheckType, detail.Verification.Age)
	fmt.Printf("Result: %v\n", detail.Verification.Result)
	fmt.Printf("Source Type: %s\n", detail.SourceType)
	fmt.Printf("Source Anchors: %v\n", detail.SourceAnchors)

	// Verify expectations for no anchors case
	if detail.SourceType != "UNKNOWN" {
		t.Errorf("Expected SourceType to be 'UNKNOWN', got '%s'", detail.SourceType)
	}

	if len(detail.SourceAnchors) != 1 || detail.SourceAnchors[0] != "No source anchors found" {
		t.Errorf("Expected 'No source anchors found' message, got %v", detail.SourceAnchors)
	}

	if detail.UsedEstimatedAge {
		t.Error("Expected UsedEstimatedAge to be false when no anchors")
	}

	if detail.UsedFallback {
		t.Error("Expected UsedFallback to be false when no anchors")
	}
}

// Note: Testing with actual anchor data would require:
// 1. Real certificates from Yoti API responses
// 2. Proper DER-encoded X.509 certificate structures
// 3. Protobuf anchor objects with correct OIDs
//
// The above tests verify the basic functionality of the analyzeAgeVerificationSources
// function. To test with real anchor data, you would need to:
// - Capture actual API responses from Yoti
// - Parse the certificates and anchors
// - Verify the source detection logic works correctly
//
// Example of what real anchor data looks like:
// - Anchor Type: SOURCE or VERIFIER (identified by OID)
// - Origin Server Certs: X.509 certificates with extensions
// - Extension Value: e.g., "PASSPORT", "DRIVING_LICENCE"
// - Signed Timestamp: When the attribute was signed
// - SubType: e.g., "OCR", "NFC"
