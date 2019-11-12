package yoti

import (
	"crypto/rsa"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotocom"
	"github.com/golang/protobuf/proto"
)

func parseApplicationProfile(receipt *receiptDO, key *rsa.PrivateKey) (result *yotiprotoattr.AttributeList, err error) {
	var unwrappedKey []byte
	if unwrappedKey, err = unwrapKey(receipt.WrappedReceiptKey, key); err != nil {
		return
	}

	if receipt.ProfileContent == "" {
		return
	}

	var profileContentBytes []byte
	if profileContentBytes, err = base64ToBytes(receipt.ProfileContent); err != nil {
		return
	}

	encryptedData := &yotiprotocom.EncryptedData{}
	if err = proto.Unmarshal(profileContentBytes, encryptedData); err != nil {
		return nil, err
	}

	var decipheredBytes []byte
	if decipheredBytes, err = decipherAes(unwrappedKey, encryptedData.Iv, encryptedData.CipherText); err != nil {
		return nil, err
	}

	attributeList := &yotiprotoattr.AttributeList{}
	if err := proto.Unmarshal(decipheredBytes, attributeList); err != nil {
		return nil, err
	}

	return attributeList, nil
}

func parseUserProfile(receipt *receiptDO, key *rsa.PrivateKey) (result *yotiprotoattr.AttributeList, err error) {
	decipheredBytes, err := parseEncryptedProto(receipt, receipt.OtherPartyProfileContent, key)
	if err != nil {
		return
	}

	attributeList := &yotiprotoattr.AttributeList{}
	if err := proto.Unmarshal(decipheredBytes, attributeList); err != nil {
		return nil, err
	}

	return attributeList, nil
}

func parseExtraData(receipt *receiptDO, key *rsa.PrivateKey) (result string, err error) {
	bytes, err := parseEncryptedProto(receipt, receipt.ExtraDataContent, key)
	return string(bytes), err
}

func parseEncryptedProto(receipt *receiptDO, encryptedBase64 string, key *rsa.PrivateKey) (result []byte, err error) {
	unwrappedKey, err := unwrapKey(receipt.WrappedReceiptKey, key)
	if err != nil {
		return
	}
	encryptedBytes, err := base64ToBytes(encryptedBase64)
	if err != nil || len(encryptedBytes) == 0 {
		return
	}
	encryptedData := &yotiprotocom.EncryptedData{}
	if err = proto.Unmarshal(encryptedBytes, encryptedData); err != nil {
		return
	}

	return decipherAes(unwrappedKey, encryptedData.Iv, encryptedData.CipherText)
}
