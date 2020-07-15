package sandbox

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
)

func Example_checkReportsBuilder() {
	breakdown, err := report.NewBreakdownBuilder().Build()
	if err != nil {
		return
	}

	recommendation, err := report.NewRecommendationBuilder().Build()
	if err != nil {
		return
	}

	authenticityCheck, err := check.NewDocumentAuthenticityCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		Build()
	if err != nil {
		return
	}

	faceMatchCheck, err := check.NewDocumentFaceMatchCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		Build()
	if err != nil {
		return
	}

	textDataCheck, err := check.NewDocumentTextDataCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		Build()
	if err != nil {
		return
	}

	zoomLivenessCheck, err := check.NewZoomLivenessCheckBuilder().Build()
	if err != nil {
		return
	}

	checkReports, err := NewCheckReportsBuilder().
		WithDocumentAuthenticityCheck(authenticityCheck).
		WithDocumentFaceMatchCheck(faceMatchCheck).
		WithDocumentTextDataCheck(textDataCheck).
		WithLivenessCheck(zoomLivenessCheck).
		WithAsyncReportDelay(10).
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(checkReports)
	fmt.Println(string(data))
	// Output: {"ID_DOCUMENT_AUTHENTICITY":[{"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]}},"document_filter":{"document_types":null,"country_codes":null}}],"ID_DOCUMENT_TEXT_DATA_CHECK":[{"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]},"document_fields":null},"document_filter":{"document_types":null,"country_codes":null}}],"ID_DOCUMENT_FACE_MATCH_CHECK":[{"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]}},"document_filter":{"document_types":null,"country_codes":null}}],"LIVENESS":[{"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":null}},"liveness_type":"ZOOM"}],"async_report_delay":10}
}

func Example_checkReportsBuilderMinimal() {
	checkReports, err := NewCheckReportsBuilder().Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(checkReports)
	fmt.Println(string(data))
	// Output: {"ID_DOCUMENT_AUTHENTICITY":[],"ID_DOCUMENT_TEXT_DATA_CHECK":[],"ID_DOCUMENT_FACE_MATCH_CHECK":[],"LIVENESS":[],"async_report_delay":0}
}
