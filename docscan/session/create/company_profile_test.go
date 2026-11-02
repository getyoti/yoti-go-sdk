package create

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewCompanyProfile(t *testing.T) {
	cp := NewCompanyProfile("Acme Corp")
	if cp == nil {
		t.Fatal("expected non-nil CompanyProfile")
	}
	if cp.GetCompanyName() != "Acme Corp" {
		t.Errorf("expected company name 'Acme Corp', got '%s'", cp.GetCompanyName())
	}
}

func TestCompanyProfile_JSONMarshalling(t *testing.T) {
	cp := NewCompanyProfile("Acme Corp")
	data, err := json.Marshal(cp)
	if err != nil {
		t.Fatalf("failed to marshal CompanyProfile: %v", err)
	}
	expected := `{"company_name":"Acme Corp"}`
	if string(data) != expected {
		t.Errorf("expected JSON %s, got %s", expected, string(data))
	}
}

func ExampleSessionSpecificationBuilder_Build_withCompanyProfile() {
	companyProfile := NewCompanyProfile("Acme Corp")

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithCompanyProfile(companyProfile).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"company_profile":{"company_name":"Acme Corp"}}
}

func TestSessionSpecification_CompanyProfile_Omitted_When_Not_Set(t *testing.T) {
	sessionSpecification, err := NewSessionSpecificationBuilder().
		Build()
	if err != nil {
		t.Fatalf("error building session spec: %v", err)
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		t.Fatalf("error marshalling session spec: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("error unmarshalling result: %v", err)
	}

	if _, exists := m["company_profile"]; exists {
		t.Error("expected 'company_profile' to be omitted when not set")
	}
}

func TestSessionSpecification_CompanyProfile_IncludedInJSON_When_Set(t *testing.T) {
	companyProfile := NewCompanyProfile("Test Company")

	sessionSpecification, err := NewSessionSpecificationBuilder().
		WithCompanyProfile(companyProfile).
		Build()
	if err != nil {
		t.Fatalf("error building session spec: %v", err)
	}

	data, err := json.Marshal(sessionSpecification)
	if err != nil {
		t.Fatalf("error marshalling session spec: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("error unmarshalling result: %v", err)
	}

	cp, exists := m["company_profile"]
	if !exists {
		t.Fatal("expected 'company_profile' to be present in JSON")
	}

	cpMap, ok := cp.(map[string]interface{})
	if !ok {
		t.Fatal("expected 'company_profile' to be a JSON object")
	}

	if cpMap["company_name"] != "Test Company" {
		t.Errorf("expected company_name 'Test Company', got '%v'", cpMap["company_name"])
	}
}
