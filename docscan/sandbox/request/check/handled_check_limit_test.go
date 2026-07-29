package check

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// ExampleDocumentAuthenticityCheckBuilder_withHandledCheckLimit shows handled_check_limit serialization
func ExampleDocumentAuthenticityCheckBuilder_withHandledCheckLimit() {
	recommendation, err := report.NewRecommendationBuilder().
		WithValue("some_value").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	check, err := NewDocumentAuthenticityCheckBuilder().
		WithRecommendation(recommendation).
		WithDocumentFilter(docFilter).
		WithHandledCheckLimit(3).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(check)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"some_value"}}},"handled_check_limit":3,"document_filter":{"document_types":[],"country_codes":[]}}
}

func TestHandledCheckLimit_IncludedWhenSet(t *testing.T) {
	recommendation, err := report.NewRecommendationBuilder().
		WithValue("APPROVE").
		Build()
	if err != nil {
		t.Fatalf("unexpected error building recommendation: %v", err)
	}

	c, err := NewThirdPartyIdentityCheckBuilder().
		WithRecommendation(recommendation).
		WithHandledCheckLimit(5).
		Build()
	if err != nil {
		t.Fatalf("unexpected error building check: %v", err)
	}

	assertHandledCheckLimitInJSON(t, c, 5)

	// Zero is a valid limit and must still be serialized (should not be omitted).
	cZero, err := NewThirdPartyIdentityCheckBuilder().
		WithRecommendation(recommendation).
		WithHandledCheckLimit(0).
		Build()
	if err != nil {
		t.Fatalf("unexpected error building check (zero limit): %v", err)
	}

	assertHandledCheckLimitInJSON(t, cZero, 0)
}

func TestHandledCheckLimit_OmittedWhenNotSet(t *testing.T) {
	recommendation, err := report.NewRecommendationBuilder().
		WithValue("APPROVE").
		Build()
	if err != nil {
		t.Fatalf("unexpected error building recommendation: %v", err)
	}

	c, err := NewThirdPartyIdentityCheckBuilder().
		WithRecommendation(recommendation).
		Build()
	if err != nil {
		t.Fatalf("unexpected error building check: %v", err)
	}

	assertHandledCheckLimitAbsentFromJSON(t, c)
}

func TestHandledCheckLimit_DocumentAuthenticityCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewDocumentAuthenticityCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewDocumentAuthenticityCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func TestHandledCheckLimit_DocumentFaceMatchCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewDocumentFaceMatchCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewDocumentFaceMatchCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func TestHandledCheckLimit_DocumentTextDataCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewDocumentTextDataCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewDocumentTextDataCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func TestHandledCheckLimit_IDDocumentComparisonCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewIDDocumentComparisonCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewIDDocumentComparisonCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func TestHandledCheckLimit_SupplementaryDocumentTextDataCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewSupplementaryDocumentTextDataCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewSupplementaryDocumentTextDataCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func TestHandledCheckLimit_ThirdPartyIdentityCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewThirdPartyIdentityCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewThirdPartyIdentityCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func TestHandledCheckLimit_ZoomLivenessCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewZoomLivenessCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewZoomLivenessCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func TestHandledCheckLimit_StaticLivenessCheck(t *testing.T) {
	rec := mustBuildRecommendation(t)
	c, err := NewStaticLivenessCheckBuilder().WithRecommendation(rec).WithHandledCheckLimit(2).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitInJSON(t, c, 2)
	c2, err := NewStaticLivenessCheckBuilder().WithRecommendation(rec).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertHandledCheckLimitAbsentFromJSON(t, c2)
}

func mustBuildRecommendation(t *testing.T) *report.Recommendation {
	t.Helper()
	rec, err := report.NewRecommendationBuilder().WithValue("APPROVE").Build()
	if err != nil {
		t.Fatalf("unexpected error building recommendation: %v", err)
	}
	return rec
}

func assertHandledCheckLimitInJSON(t *testing.T, v interface{}, limit int) {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("unexpected error marshaling: %v", err)
	}
	jsonStr := string(data)
	expected := fmt.Sprintf(`"handled_check_limit":%d`, limit)
	if !strings.Contains(jsonStr, expected) {
		t.Errorf("expected JSON to contain %s, got: %s", expected, jsonStr)
	}
}

func assertHandledCheckLimitAbsentFromJSON(t *testing.T, v interface{}) {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("unexpected error marshaling: %v", err)
	}
	jsonStr := string(data)
	if strings.Contains(jsonStr, "handled_check_limit") {
		t.Errorf("expected JSON to omit handled_check_limit, got: %s", jsonStr)
	}
}
