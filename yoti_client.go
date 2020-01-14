package yoti

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/consts"
	"github.com/getyoti/yoti-go-sdk/v2/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v2/requests"
	"github.com/getyoti/yoti-go-sdk/v2/share"
	"github.com/getyoti/yoti-go-sdk/v2/yotierror"
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
	Key *rsa.PrivateKey

	apiURL     string
	HTTPClient httpClient // Mockable HTTP Client Interface
}

// NewClient constructs a Client object
func NewClient(sdkID string, key []byte) (*Client, error) {
	decodedKey, err := loadRsaKey(key)
	return &Client{
		SdkID: sdkID,
		Key:   decodedKey,
	}, err
}

func (client *Client) doRequest(request *http.Request) (*http.Response, error) {
	if client.HTTPClient == nil {
		client.HTTPClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}
	return client.HTTPClient.Do(request)
}

// OverrideAPIURL overrides the default API URL for this Yoti Client
func (client *Client) OverrideAPIURL(apiURL string) {
	client.apiURL = apiURL
}

func (client *Client) getAPIURL() string {
	if client.apiURL != "" {
		return client.apiURL
	}

	if value, exists := os.LookupEnv("YOTI_API_URL"); exists && value != "" {
		return value
	}

	return apiDefaultURL
}

// GetSdkID gets the Client SDK ID attached to this client instance
func (client *Client) GetSdkID() string {
	return client.SdkID
}

// GetActivityDetails requests information about a Yoti user using the one time
// use token generated by the Yoti login process. It returns the outcome of the
// request. If the request was successful it will include the user's details,
// otherwise an error will be returned, which will specify the reason the
// request failed. If the function call can be reattempted with the same token
// the error will implement interface{ Temporary() bool }.
func (client *Client) GetActivityDetails(token string) (ActivityDetails, error) {
	activity, err := client.getActivityDetails(token)
	return activity, err
}

func (client *Client) getActivityDetails(token string) (activity ActivityDetails, err error) {

	httpMethod := http.MethodGet
	key, err := cryptoutil.ParseRSAKey(client.Key)
	if err != nil {
		err = fmt.Errorf("Invalid Token: %s", err.Error())
		return
	}
	endpoint := getProfileEndpoint(token, client.GetSdkID())

	response, err := client.makeRequest(
		httpMethod,
		endpoint,
		nil,
		true,
		map[int]string{404: "Profile Not Found %[1]d"},
		DefaultHTTPErrorMessages,
	)
	if err != nil {
		return
	}
	return handleSuccessfulResponse(response, client.Key)
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
	key, err := cryptoutil.ParseRSAKey(&client.Key.PublicKey)
	if err != nil {
		return
	}

	var headers map[string][]string
	if includeKey {
		headers = requests.AuthKeyHeader(key)
	}

	request, err := requests.SignedRequest{
		Key:        client.Key,
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
	if response.StatusCode >= 500 {
		err = yotierror.NewTemporary(err)
	}
	return
}

func getProtobufAttribute(profile Profile, key string) *yotiprotoattr.Attribute {
	for _, v := range profile.attributeSlice {
		if v.Name == key {
			return v
		}
	}

	return nil
}

func handleSuccessfulResponse(responseContent string, key *rsa.PrivateKey) (activityDetails ActivityDetails, err error) {
	var parsedResponse = profileDO{}

	if err = json.Unmarshal([]byte(responseContent), &parsedResponse); err != nil {
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

		ensureAddressProfile(&profile)

		decryptedExtraData, errTemp := parseExtraData(&parsedResponse.Receipt, key)
		if errTemp != nil {
			log.Printf("Unable to decrypt ExtraData from the receipt. Error: %q", err)
			err = yotierror.MultiError{This: errTemp, Next: err}
		}

		extraData, errTemp := share.NewExtraData(decryptedExtraData)

		if errTemp != nil {
			log.Printf("Unable to parse ExtraData from the receipt. Error: %q", err)
			err = yotierror.MultiError{This: errTemp, Next: err}
		}

		timestamp, err := time.Parse(time.RFC3339Nano, parsedResponse.Receipt.Timestamp)
		if err != nil {
			log.Printf("Unable to read timestamp. Error: %q", err)
		}

		activityDetails = ActivityDetails{
			UserProfile:        profile,
			rememberMeID:       id,
			parentRememberMeID: parsedResponse.Receipt.ParentRememberMeID,
			timestamp:          timestamp,
			receiptID:          parsedResponse.Receipt.ReceiptID,
			ApplicationProfile: appProfile,
			extraData:          extraData,
		}
	}

	return activityDetails, err
}

func createAttributeSlice(protoAttributeList *yotiprotoattr.AttributeList) (result []*yotiprotoattr.Attribute) {
	if protoAttributeList != nil {
		result = append(result, protoAttributeList.Attributes...)
	}

	return result
}

func setFormattedAddress(profile *Profile, formattedAddress string) {
	proto := getProtobufAttribute(*profile, consts.AttrStructuredPostalAddress)

	addressAttribute := &yotiprotoattr.Attribute{
		Name:        consts.AttrAddress,
		Value:       []byte(formattedAddress),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     proto.Anchors,
	}
	profile.attributeSlice = append(profile.attributeSlice, addressAttribute)
}

func ensureAddressProfile(profile *Profile) {
	if profile.Address() == nil {
		if structuredPostalAddress, err := profile.StructuredPostalAddress(); err == nil {
			if (structuredPostalAddress != nil && !reflect.DeepEqual(structuredPostalAddress, attribute.JSONAttribute{})) {
				var formattedAddress string
				formattedAddress, err = retrieveFormattedAddressFromStructuredPostalAddress(structuredPostalAddress.Value())
				if err == nil && formattedAddress != "" {
					setFormattedAddress(profile, formattedAddress)
					return
				}
			}
		}
	}
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