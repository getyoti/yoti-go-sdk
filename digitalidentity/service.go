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
const errorFailedToGetSignedRequest = "failed to get signed request: %v"
const errorFailedToExecuteRequest = "failed to execute request: %v"
const errorFailedToReadBody = "failed to read response body: %v"

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
		return nil, fmt.Errorf(errorFailedToGetSignedRequest, err)
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToExecuteRequest, err)
	}

	defer response.Body.Close()
	shareSession := &ShareSession{}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToReadBody, err)
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
		return nil, fmt.Errorf(errorFailedToGetSignedRequest, err)
	}

	response, err := requests.Execute(httpClient, request)

	if err != nil {
		return nil, fmt.Errorf(errorFailedToExecuteRequest, err)
	}
	defer response.Body.Close()
	shareSession := &ShareSession{}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToReadBody, err)
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
		return nil, fmt.Errorf(errorFailedToGetSignedRequest, err)
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToExecuteRequest, err)
	}

	defer response.Body.Close()
	qrCode := &QrCode{}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(errorFailedToReadBody, err)
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
		return fetchedQrCode, fmt.Errorf(errorFailedToGetSignedRequest, err)
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return fetchedQrCode, fmt.Errorf(errorFailedToExecuteRequest, err)
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fetchedQrCode, fmt.Errorf(errorFailedToReadBody, err)
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
		return receipt, fmt.Errorf(errorFailedToGetSignedRequest, err)
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return receipt, fmt.Errorf(errorFailedToExecuteRequest, err)
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return receipt, fmt.Errorf(errorFailedToReadBody, err)
	}

	err = json.Unmarshal(responseBytes, &receipt)

	return receipt, err
}

// GetReceiptItemKey retrieves the receipt item key for a receipt item key id.
func getReceiptItemKey(httpClient requests.HttpClient, receiptItemKeyId string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (receiptItemKey ReceiptItemKeyResponse, err error) {
	endpoint := fmt.Sprintf(identitySessionReceiptKeyRetrieval, receiptItemKeyId)
	headers := requests.AuthHeader(clientSdkId)
	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
	}.Request()
	if err != nil {
		return receiptItemKey, fmt.Errorf(errorFailedToGetSignedRequest, err)
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

func GetShareReceipt(httpClient requests.HttpClient, receiptId string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (receipt SharedReceiptResponse, err error) {
	receiptResponse, err := getReceipt(httpClient, receiptId, clientSdkId, apiUrl, key)
	if err != nil {
		return receipt, fmt.Errorf("failed to get receipt: %v", err)
	}

	itemKeyId := receiptResponse.WrappedItemKeyId

	encryptedItemKeyResponse, err := getReceiptItemKey(httpClient, itemKeyId, clientSdkId, apiUrl, key)
	if err != nil {
		return receipt, fmt.Errorf("failed to get receipt item key: %v", err)
	}

	receiptContentKey, err := cryptoutil.UnwrapReceiptKey(receiptResponse.WrappedKey, encryptedItemKeyResponse.Value, encryptedItemKeyResponse.Iv, key)
	if err != nil {
		return receipt, fmt.Errorf("failed to unwrap receipt content key: %v", err)
	}

	attrData, aextra, err := decryptReceiptContent(receiptResponse.Content, receiptContentKey)
	if err != nil {
		return receipt, fmt.Errorf("failed to decrypt receipt content: %v", err)
	}

	applicationProfile := newApplicationProfile(attrData)
	extraDataValue, err := extra.NewExtraData(aextra)
	if err != nil {
		return receipt, fmt.Errorf("failed to build application extra data: %v", err)
	}

	uattrData, uextra, err := decryptReceiptContent(receiptResponse.OtherPartyContent, receiptContentKey)
	if err != nil {
		return receipt, fmt.Errorf("failed to decrypt receipt other party content: %v", err)
	}

	userProfile := newUserProfile(uattrData)
	userExtraDataValue, err := extra.NewExtraData(uextra)
	if err != nil {
		return receipt, fmt.Errorf("failed to build other party extra data: %v", err)
	}

	return SharedReceiptResponse{
		ID:                 receiptResponse.ID,
		SessionID:          receiptResponse.SessionID,
		RememberMeID:       receiptResponse.RememberMeID,
		ParentRememberMeID: receiptResponse.ParentRememberMeID,
		Timestamp:          receiptResponse.Timestamp,
		UserContent: UserContent{
			UserProfile: userProfile,
			ExtraData:   userExtraDataValue,
		},
		ApplicationContent: ApplicationContent{
			ApplicationProfile: applicationProfile,
			ExtraData:          extraDataValue,
		},
		Error: receiptResponse.Error,
	}, nil
}

func decryptReceiptContent(content *Content, key []byte) (attrData *yotiprotoattr.AttributeList, aextra []byte, err error) {

	if content != nil {
		if len(content.Profile) > 0 {
			aattr, err := cryptoutil.DecryptReceiptContent(content.Profile, key)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to decrypt content profile: %v", err)
			}

			attrData = &yotiprotoattr.AttributeList{}
			if err := proto.Unmarshal(aattr, attrData); err != nil {
				return nil, nil, fmt.Errorf("failed to unmarshal attribute list: %v", err)
			}
		}

		if len(content.ExtraData) > 0 {
			aextra, err = cryptoutil.DecryptReceiptContent(content.ExtraData, key)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to decrypt receipt content extra data: %v", err)
			}
		}

	}

	return attrData, aextra, nil
}
