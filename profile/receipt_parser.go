package profile

import (
	"crypto/rsa"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/util"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotocom"
	"github.com/golang/protobuf/proto"
)

func parseApplicationProfile(receipt *receiptDO, key *rsa.PrivateKey) (result *yotiprotoattr.AttributeList, err error) {
	var unwrappedKey []byte
	if unwrappedKey, err = cryptoutil.UnwrapKey(receipt.WrappedReceiptKey, key); err != nil {
		return
	}

	if receipt.ProfileContent == "" {
		return
	}

	var profileContentBytes []byte
	if profileContentBytes, err = util.Base64ToBytes(receipt.ProfileContent); err != nil {
		return
	}

	encryptedData := &yotiprotocom.EncryptedData{}
	if err = proto.Unmarshal(profileContentBytes, encryptedData); err != nil {
		return nil, err
	}

	var decipheredBytes []byte
	if decipheredBytes, err = cryptoutil.DecipherAes(unwrappedKey, encryptedData.Iv, encryptedData.CipherText); err != nil {
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

func decryptExtraData(receipt *receiptDO, key *rsa.PrivateKey) (result []byte, err error) {
	bytes, err := parseEncryptedProto(receipt, receipt.ExtraDataContent, key)
	return bytes, err
}

func parseEncryptedProto(receipt *receiptDO, encryptedBase64 string, key *rsa.PrivateKey) (result []byte, err error) {
	unwrappedKey, err := cryptoutil.UnwrapKey(receipt.WrappedReceiptKey, key)
	if err != nil {
		return
	}
	encryptedBytes, err := util.Base64ToBytes(encryptedBase64)
	if err != nil || len(encryptedBytes) == 0 {
		return
	}
	encryptedData := &yotiprotocom.EncryptedData{}
	if err = proto.Unmarshal(encryptedBytes, encryptedData); err != nil {
		return
	}

	return cryptoutil.DecipherAes(unwrappedKey, encryptedData.Iv, encryptedData.CipherText)
}
