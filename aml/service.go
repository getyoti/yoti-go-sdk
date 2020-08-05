package aml

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/requests"
)

func getCheckEndpoint(sdkID string) string {
	return fmt.Sprintf("/aml-check?appId=%s", sdkID)
}

// PerformCheck performs an Anti Money Laundering Check (AML) for a particular user.
// Returns three boolean values: 'OnPEPList', 'OnWatchList' and 'OnFraudList'.
func PerformCheck(httpClient requests.HttpClient, profile Profile, clientSdkId, apiUrl string, key *rsa.PrivateKey) (result Result, err error) {
	payload, err := json.Marshal(profile)
	if err != nil {
		return
	}
	headers := requests.AuthKeyHeader(&key.PublicKey)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   getCheckEndpoint(clientSdkId),
		Headers:    headers,
		Body:       payload,
	}.Request()
	if err != nil {
		return
	}

	httpErrorMessages := make(map[int]string)
	httpErrorMessages[-1] = "AML Check was unsuccessful"

	response, err := requests.Execute(httpClient, request, httpErrorMessages)
	if err != nil {
		return result, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	result, err = GetResult(responseBytes)
	return
}
