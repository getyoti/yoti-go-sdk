package retrieve_test

import (
	"encoding/json"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/retrieve"
	"gotest.tools/v3/assert"
)

func TestBreakdownResponse_ProcessAutomated(t *testing.T) {
	payload := `{"sub_check":"doc_number_validation","result":"PASS","details":[],"process":"AUTOMATED"}`

	var breakdown retrieve.BreakdownResponse
	err := json.Unmarshal([]byte(payload), &breakdown)
	assert.NilError(t, err)
	assert.Equal(t, "AUTOMATED", breakdown.Process)
}

func TestBreakdownResponse_ProcessExpertReview(t *testing.T) {
	payload := `{"sub_check":"doc_number_validation","result":"PASS","details":[],"process":"EXPERT_REVIEW"}`

	var breakdown retrieve.BreakdownResponse
	err := json.Unmarshal([]byte(payload), &breakdown)
	assert.NilError(t, err)
	assert.Equal(t, "EXPERT_REVIEW", breakdown.Process)
}

func TestBreakdownResponse_ProcessAbsent(t *testing.T) {
	payload := `{"sub_check":"doc_number_validation","result":"PASS","details":[]}`

	var breakdown retrieve.BreakdownResponse
	err := json.Unmarshal([]byte(payload), &breakdown)
	assert.NilError(t, err)
	assert.Equal(t, "", breakdown.Process)
}
