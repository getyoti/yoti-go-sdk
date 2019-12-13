package yoti

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/requests"
	"github.com/getyoti/yoti-go-sdk/v2/share"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

const (
	apiDefaultURL = "https://api.yoti.com/api/v1"

	defaultUnknownErrorMessageConst = "Unknown HTTP Error: %[1]d: %[2]s"
)

var (
	// DefaultHTTPErrorMessages maps HTTP error status codes to format strings
	// to create useful error messages. -1 is used to specify a default message
	// that can be used if an error code is not explicitly defined
	DefaultHTTPErrorMessages = map[int]string{
		-1: defaultUnknownErrorMessageConst,
	}
)

// ClientInterface defines the interface required to Mock the YotiClient for
// testing
type clientInterface interface {
	makeRequest(string, string, []byte, bool, ...map[int]string) (string, error)
	GetSdkID() string
}

// Client represents a client that can communicate with yoti and return information about Yoti users.
type Client struct {
	// SdkID represents the SDK ID and NOT the App ID. This can be found in the integration section of your
	// application hub at https://hub.yoti.com/
	SdkID string

	// Key should be the security key given to you by yoti (see: security keys section of
	// https://hub.yoti.com) for more information about how to load your key from a file see:
	// https://github.com/getyoti/yoti-go-sdk/blob/master/README.md
	Key []byte

	apiURL     string
	HTTPClient httpClient // Mockable HTTP Client Interface
}

func (client *Client) doRequest(request *http.Request) (*http.Response, error) {
	if client.HTTPClient == nil {
		client.HTTPClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}
	return client.HTTPClient.Do(request)
}

// OverrideAPIURL overrides the default API URL for this Yoti Client to permit
// testing
func (client *Client) OverrideAPIURL(apiURL string) {
	client.apiURL = apiURL
}

func (client *Client) getAPIURL() string {
	if client.apiURL != "" {
		return client.apiURL
	}
	return apiDefaultURL
}

// GetSdkID gets the Client SDK ID attached to this client instance
func (client *Client) GetSdkID() string {
	return client.SdkID
}

// GetActivityDetails requests information about a Yoti user using the one time use token generated by the Yoti login process.
// It returns the outcome of the request. If the request was successful it will include the users details, otherwise
// it will specify a reason the request failed.
func (client *Client) GetActivityDetails(token string) (ActivityDetails, []string) {
	activity, errStrings := client.getActivityDetails(token)
	return activity, errStrings
}

func (client *Client) getActivityDetails(token string) (activity ActivityDetails, errStrings []string) {

	httpMethod := http.MethodGet
	key, err := loadRsaKey(client.Key)
	if err != nil {
		errStrings = append(errStrings, fmt.Sprintf("Invalid Key: %s", err.Error()))
		return
	}
	token, err = decryptToken(token, key)
	if err != nil {
		errStrings = append(errStrings, fmt.Sprintf("Invalid Key: %s", err.Error()))
		return
	}
	endpoint := getProfileEndpoint(token, client.GetSdkID())

	response, err := client.makeRequest(
		httpMethod,
		endpoint,
		nil,
		true,
		map[int]string{404: "Profile Not Found%[2]s"},
		DefaultHTTPErrorMessages,
	)
	if err != nil {
		errStrings = append(errStrings, err.Error())
		return
	}
	return handleSuccessfulResponse(response, key)
}

func handleHTTPError(response *http.Response, errorMessages ...map[int]string) error {
	var body []byte
	if response.Body != nil {
		body, _ = ioutil.ReadAll(response.Body)
	} else {
		body = make([]byte, 0)
	}
	for _, handler := range errorMessages {
		for code, message := range handler {
			if code == response.StatusCode {
				return fmt.Errorf(
					message,
					response.StatusCode,
					body,
				)
			}

		}
		if defaultMessage, ok := handler[-1]; ok {
			return fmt.Errorf(
				defaultMessage,
				response.StatusCode,
				body,
			)
		}

	}
	return fmt.Errorf(
		defaultUnknownErrorMessageConst,
		response.StatusCode,
		body,
	)
}

func (client *Client) makeRequest(httpMethod, endpoint string, payload []byte, includeKey bool, httpErrorMessages ...map[int]string) (responseData string, err error) {
	key, err := loadRsaKey(client.Key)
	if err != nil {
		return
	}

	var headers map[string][]string
	if includeKey {
		headers = requests.AuthKeyHeader(&key.PublicKey)
	}

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: httpMethod,
		BaseURL:    client.getAPIURL(),
		Endpoint:   endpoint,
		Headers:    headers,
		Body:       payload,
	}.Request()

	if err != nil {
		return
	}

	var response *http.Response
	if response, err = client.doRequest(request); err != nil {
		return
	}

	if response.StatusCode < 300 && response.StatusCode >= 200 {
		var tmp []byte
		if response.Body != nil {
			tmp, err = ioutil.ReadAll(response.Body)
		} else {
			tmp = make([]byte, 0)
		}
		responseData = string(tmp)
		return
	}
	err = handleHTTPError(response, httpErrorMessages...)
	return
}

func getProtobufAttribute(profile Profile, key string) *yotiprotoattr.Attribute {
	for _, v := range profile.attributeSlice {
		if v.Name == AttrConstStructuredPostalAddress {
			return v
		}
	}

	return nil
}

func handleSuccessfulResponse(responseContent string, key *rsa.PrivateKey) (activityDetails ActivityDetails, errStrings []string) {
	var parsedResponse = profileDO{}
	var err error

	if err = json.Unmarshal([]byte(responseContent), &parsedResponse); err != nil {
		errStrings = append(errStrings, err.Error())
		return
	}

	if parsedResponse.Receipt.SharingOutcome != "SUCCESS" {
		err = ErrSharingFailure
		errStrings = append(errStrings, err.Error())
	} else {
		var userAttributeList, applicationAttributeList *yotiprotoattr.AttributeList
		if userAttributeList, err = parseUserProfile(&parsedResponse.Receipt, key); err != nil {
			errStrings = append(errStrings, err.Error())
			return
		}
		if applicationAttributeList, err = parseApplicationProfile(&parsedResponse.Receipt, key); err != nil {
			errStrings = append(errStrings, err.Error())
			return
		}
		id := parsedResponse.Receipt.RememberMeID

		profile := Profile{
			baseProfile{
				attributeSlice: createAttributeSlice(userAttributeList),
			},
		}
		appProfile := ApplicationProfile{
			baseProfile{
				attributeSlice: createAttributeSlice(applicationAttributeList),
			},
		}

		var formattedAddress string
		formattedAddress, err = ensureAddressProfile(profile)
		if err != nil {
			log.Printf("Unable to get 'Formatted Address' from 'Structured Postal Address'. Error: %q", err)
		} else if formattedAddress != "" {
			if _, err = profile.StructuredPostalAddress(); err != nil {
				errStrings = append(errStrings, err.Error())
				return
			}

			protoStructuredPostalAddress := getProtobufAttribute(profile, AttrConstStructuredPostalAddress)

			addressAttribute := &yotiprotoattr.Attribute{
				Name:        AttrConstAddress,
				Value:       []byte(formattedAddress),
				ContentType: yotiprotoattr.ContentType_STRING,
				Anchors:     protoStructuredPostalAddress.Anchors,
			}

			profile.attributeSlice = append(profile.attributeSlice, addressAttribute)
		}

		decryptedExtraData, err := parseExtraData(&parsedResponse.Receipt, key)
		if err != nil {
			log.Printf("Unable to decrypt ExtraData from the receipt. Error: %q", err)
			errStrings = append(errStrings, err.Error())
		}

		extraData, err := share.NewExtraData(decryptedExtraData)

		if err != nil {
			log.Printf("Unable to parse ExtraData from the receipt. Error: %q", err)
			errStrings = append(errStrings, err.Error())
		}

		activityDetails = ActivityDetails{
			UserProfile:        profile,
			rememberMeID:       id,
			parentRememberMeID: parsedResponse.Receipt.ParentRememberMeID,
			timestamp:          parsedResponse.Receipt.Timestamp,
			receiptID:          parsedResponse.Receipt.ReceiptID,
			ApplicationProfile: appProfile,
			extraData:          extraData,
		}
	}

	return activityDetails, errStrings
}

func createAttributeSlice(protoAttributeList *yotiprotoattr.AttributeList) (result []*yotiprotoattr.Attribute) {
	if protoAttributeList != nil {
		result = append(result, protoAttributeList.Attributes...)
	}

	return result
}

func ensureAddressProfile(profile Profile) (address string, err error) {
	if profile.Address() == nil {
		var structuredPostalAddress *attribute.JSONAttribute
		if structuredPostalAddress, err = profile.StructuredPostalAddress(); err == nil {
			if (structuredPostalAddress != nil && !reflect.DeepEqual(structuredPostalAddress, attribute.JSONAttribute{})) {
				var formattedAddress string
				formattedAddress, err = retrieveFormattedAddressFromStructuredPostalAddress(structuredPostalAddress.Value())
				if err == nil {
					return formattedAddress, nil
				}
			}
		}
	}

	return "", err
}

func retrieveFormattedAddressFromStructuredPostalAddress(structuredPostalAddress interface{}) (address string, err error) {
	parsedStructuredAddressMap := structuredPostalAddress.(map[string]interface{})
	if formattedAddress, ok := parsedStructuredAddressMap["formatted_address"]; ok {
		return formattedAddress.(string), nil
	}
	return
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

// PerformAmlCheck performs an Anti Money Laundering Check (AML) for a particular user.
// Returns three boolean values: 'OnPEPList', 'OnWatchList' and 'OnFraudList'.
func (client *Client) PerformAmlCheck(amlProfile AmlProfile) (amlResult AmlResult, err error) {
	var httpMethod = http.MethodPost
	endpoint := getAMLEndpoint(client.GetSdkID())
	content, err := json.Marshal(amlProfile)
	if err != nil {
		return
	}
	amlErrorMessages := make(map[int]string)
	amlErrorMessages[-1] = "AML Check was unsuccessful, status code: '%[1]d', content '%[2]s'"

	response, err := client.makeRequest(httpMethod, endpoint, content, false, amlErrorMessages)
	if err != nil {
		return
	}

	amlResult, err = GetAmlResult([]byte(response))
	return
}
