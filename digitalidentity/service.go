package digitalidentity

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity/requests"
	"github.com/getyoti/yoti-go-sdk/v3/extra"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"google.golang.org/protobuf/proto"
)

const identitySessionCreationEndpoint = "/v2/sessions"
const identitySessionRetrieval = "/v2/sessions/%s"
const identitySessionQrCodeCreation = "/v2/sessions/%s/qr-codes"
const identitySessionQrCodeRetrieval = "/v2/qr-codes/%s"
const identitySessionReceiptRetrieval = "/v2/receipts/%s"
const identitySessionReceiptKeyRetrieval = "/v2/wrapped-item-keys/%s"

// CreateShareSession creates session using the supplied session specification
func CreateShareSession(httpClient requests.HttpClient, shareSessionRequest *ShareSessionRequest, clientSdkId, apiUrl string, key *rsa.PrivateKey) (*ShareSession, error) {
	endpoint := identitySessionCreationEndpoint

	payload, err := shareSessionRequest.MarshalJSON()
	if err != nil {
		return nil, err
	}

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    requests.AuthHeader(clientSdkId),
		Body:       payload,
		Params:     map[string]string{"sdkID": clientSdkId},
	}.Request()
	if err != nil {
		return nil, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	shareSession := &ShareSession{}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBytes, shareSession)
	return shareSession, err
}

// GetShareSession get session info using the supplied sessionID parameter
func GetShareSession(httpClient requests.HttpClient, sessionID string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (*ShareSession, error) {
	endpoint := fmt.Sprintf(identitySessionRetrieval, sessionID)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    requests.AuthHeader(clientSdkId),
		Params:     map[string]string{"sdkID": clientSdkId},
	}.Request()
	if err != nil {
		return nil, err
	}

	response, err := requests.Execute(httpClient, request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	shareSession := &ShareSession{}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBytes, shareSession)
	return shareSession, err
}

// CreateShareQrCode generates a sharing qr code using the supplied sessionID parameter
func CreateShareQrCode(httpClient requests.HttpClient, sessionID string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (*QrCode, error) {
	endpoint := fmt.Sprintf(identitySessionQrCodeCreation, sessionID)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    requests.AuthHeader(clientSdkId),
		Body:       nil,
		Params:     map[string]string{"sdkID": clientSdkId},
	}.Request()
	if err != nil {
		return nil, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	qrCode := &QrCode{}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBytes, qrCode)
	return qrCode, err
}

// GetShareSessionQrCode is used to fetch the qr code by  id.
func GetShareSessionQrCode(httpClient requests.HttpClient, qrCodeId string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (fetchedQrCode ShareSessionQrCode, err error) {
	endpoint := fmt.Sprintf(identitySessionQrCodeRetrieval, qrCodeId)
	headers := requests.AuthHeader(clientSdkId)
	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
	}.Request()
	if err != nil {
		return fetchedQrCode, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return fetchedQrCode, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fetchedQrCode, err
	}

	err = json.Unmarshal(responseBytes, &fetchedQrCode)

	return fetchedQrCode, err
}

// GetReceipt fetches receipt info using a receipt id.
func getReceipt(httpClient requests.HttpClient, receiptId string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (receipt ReceiptResponse, err error) {
	receiptUrl := requests.Base64ToBase64URL(receiptId)
	endpoint := fmt.Sprintf(identitySessionReceiptRetrieval, receiptUrl)

	headers := requests.AuthHeader(clientSdkId)
	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
	}.Request()
	if err != nil {
		return receipt, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return receipt, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return receipt, err
	}

	err = json.Unmarshal(responseBytes, &receipt)

	return receipt, err
}

// Get Receipt Item Key using the supplied receiptId Item Key ID
func getFetchReceiptItemKey(httpClient requests.HttpClient, fetchReceiptItemKeyId string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (receiptItemKey ReceiptItemKeyResponse, err error) {

	endpoint := fmt.Sprintf(identitySessionReceiptKeyRetrieval, fetchReceiptItemKeyId)
	headers := requests.AuthHeader(clientSdkId)
	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
	}.Request()
	if err != nil {
		return receiptItemKey, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return receiptItemKey, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return receiptItemKey, err
	}

	err = json.Unmarshal(responseBytes, &receiptItemKey)

	return receiptItemKey, err
}

func GetShareReceipt(httpClient requests.HttpClient, receiptId string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share SharedReceiptResponse, err error) {
	receiptResponse, err := getReceipt(httpClient, receiptId, clientSdkId, apiUrl, key)
	if err != nil {
		return share, err
	}

	itemKeyId := receiptResponse.WrappedItemKeyId

	encryptedItemKeyResponse, err := getFetchReceiptItemKey(httpClient, itemKeyId, clientSdkId, apiUrl, key)
	if err != nil {
		return share, err
	}

	receiptContentKey, err := cryptoutil.UnwrapReceiptKey(receiptResponse.WrappedKey, encryptedItemKeyResponse.Value, encryptedItemKeyResponse.Iv, key)
	if err != nil {
		return share, err
	}

	aattr, err := cryptoutil.DecryptReceiptContent([]byte(receiptResponse.Content.Profile), receiptContentKey)
	if err != nil {
		return share, err
	}

	aextra, err := cryptoutil.DecryptReceiptContent([]byte(receiptResponse.Content.ExtraData), receiptContentKey)
	if err != nil {
		return share, err
	}

	attrData := &yotiprotoattr.AttributeList{}
	if err := proto.Unmarshal(aattr, attrData); err != nil {
		return share, err
	}

	applicationProfile := newApplicationProfile(attrData)
	extraDataValue, err := extra.NewExtraData(aextra)
	if err != nil {
		return share, err
	}

	applicationContent := ApplicationContent{applicationProfile, extraDataValue}

	uattr, err := cryptoutil.DecryptReceiptContent([]byte(receiptResponse.OtherPartyContent.Profile), receiptContentKey)
	uextra, err := cryptoutil.DecryptReceiptContent([]byte(receiptResponse.OtherPartyContent.ExtraData), receiptContentKey)

	aattrData := &yotiprotoattr.AttributeList{}
	if err := proto.Unmarshal(uattr, aattrData); err != nil {
		return share, err
	}

	userProfile := newUserProfile(aattrData)
	userExtraDataValue, err := extra.NewExtraData(uextra)
	if err != nil {
		return share, err
	}

	userContent := UserContent{userProfile, userExtraDataValue}

	share.ApplicationContent = applicationContent
	share.UserContent = userContent
	share.ID = receiptResponse.ID
	share.SessionID = receiptResponse.ID
	share.UserContent = userContent

	return share, err
}
