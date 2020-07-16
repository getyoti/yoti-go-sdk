package aml

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/web"
)

func getAMLEndpoint(sdkID string) string {
	return fmt.Sprintf("/aml-check?appId=%s", sdkID)
}

// PerformAmlCheck performs an Anti Money Laundering Check (AML) for a particular user.
// Returns three boolean values: 'OnPEPList', 'OnWatchList' and 'OnFraudList'.
func PerformAmlCheck(httpClient web.HttpClient, amlProfile AmlProfile, clientSdkId, apiUrl string, key *rsa.PrivateKey) (amlResult AmlResult, err error) {
	payload, err := json.Marshal(amlProfile)
	if err != nil {
		return
	}
	headers := requests.AuthKeyHeader(&key.PublicKey)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   getAMLEndpoint(clientSdkId),
		Headers:    headers,
		Body:       payload,
	}.Request()
	if err != nil {
		return
	}

	amlErrorMessages := make(map[int]string)
	amlErrorMessages[-1] = "AML Check was unsuccessful, status code: '%[1]d', content '%[2]s'"

	response, err := web.MakeRequest(httpClient, request, amlErrorMessages)
	if err != nil {
		return
	}

	amlResult, err = GetAmlResult([]byte(response))
	return
}
