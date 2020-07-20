package profile

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/extra"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

func getProfileEndpoint(token, sdkID string) string {
	return fmt.Sprintf("/profile/%s?appId=%s", token, sdkID)
}

// GetActivityDetails requests information about a Yoti user using the one time
// use token generated by the Yoti login process. Don't call this directly, use yoti.GetActivityDetails
func GetActivityDetails(httpClient requests.HttpClient, token, clientSdkId, apiUrl string, key *rsa.PrivateKey) (activity ActivityDetails, err error) {
	if len(token) < 1 {
		return activity, errors.New("Invalid Token")
	}

	var decryptedToken string
	decryptedToken, err = cryptoutil.DecryptToken(token, key)
	if err != nil {
		return activity, errors.New("Unable to decrypt token")
	}

	headers := requests.AuthKeyHeader(&key.PublicKey)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   getProfileEndpoint(decryptedToken, clientSdkId),
		Headers:    headers,
		Body:       nil,
	}.Request()
	if err != nil {
		return
	}

	response, err := requests.Execute(httpClient, request, map[int]string{404: "Profile Not Found"}, requests.DefaultHTTPErrorMessages)
	if err != nil {
		return
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return handleSuccessfulResponse(responseBytes, key)
}

func handleSuccessfulResponse(responseBytes []byte, key *rsa.PrivateKey) (activityDetails ActivityDetails, err error) {
	var parsedResponse = profileDO{}

	if err = json.Unmarshal(responseBytes, &parsedResponse); err != nil {
		return
	}

	if parsedResponse.Receipt.SharingOutcome != "SUCCESS" {
		err = ErrSharingFailure
	} else {
		var userAttributeList, applicationAttributeList *yotiprotoattr.AttributeList
		if userAttributeList, err = parseUserProfile(&parsedResponse.Receipt, key); err != nil {
			return
		}
		if applicationAttributeList, err = parseApplicationProfile(&parsedResponse.Receipt, key); err != nil {
			return
		}
		id := parsedResponse.Receipt.RememberMeID

		userProfile := newUserProfile(userAttributeList)
		applicationProfile := NewApplicationProfile(applicationAttributeList)

		var extraData *extra.ExtraData
		extraData, err = parseExtraData(&parsedResponse.Receipt, key, err)

		timestamp, timestampErr := time.Parse(time.RFC3339Nano, parsedResponse.Receipt.Timestamp)
		if timestampErr != nil {
			err = yotierror.MultiError{This: errors.New("Unable to read timestamp. Error: " + timestampErr.Error()), Next: err}
		}

		activityDetails = ActivityDetails{
			UserProfile:        userProfile,
			rememberMeID:       id,
			parentRememberMeID: parsedResponse.Receipt.ParentRememberMeID,
			timestamp:          timestamp,
			receiptID:          parsedResponse.Receipt.ReceiptID,
			ApplicationProfile: applicationProfile,
			extraData:          extraData,
		}
	}

	return activityDetails, err
}

func parseExtraData(receipt *receiptDO, key *rsa.PrivateKey, err error) (*extra.ExtraData, error) {
	decryptedExtraData, decryptErr := decryptExtraData(receipt, key)
	if decryptErr != nil {
		err = yotierror.MultiError{This: errors.New("Unable to decrypt ExtraData from the receipt. Error: " + decryptErr.Error()), Next: err}
	}

	extraData, parseErr := extra.NewExtraData(decryptedExtraData)
	if parseErr != nil {
		err = yotierror.MultiError{This: errors.New("Unable to parse ExtraData from the receipt. Error: " + parseErr.Error()), Next: err}
	}
	return extraData, err
}

func parseIsAgeVerifiedValue(byteValue []byte) (result *bool, err error) {
	stringValue := string(byteValue)

	var parseResult bool
	parseResult, err = strconv.ParseBool(stringValue)

	if err != nil {
		return nil, err
	}

	result = &parseResult

	return
}
