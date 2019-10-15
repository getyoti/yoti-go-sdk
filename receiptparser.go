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
	var unwrappedKey []byte
	if unwrappedKey, err = unwrapKey(receipt.WrappedReceiptKey, key); err != nil {
		return
	}

	if receipt.OtherPartyProfileContent == "" {
		return
	}

	var otherPartyProfileContentBytes []byte
	if otherPartyProfileContentBytes, err = base64ToBytes(receipt.OtherPartyProfileContent); err != nil {
		return
	}

	encryptedData := &yotiprotocom.EncryptedData{}
	if err = proto.Unmarshal(otherPartyProfileContentBytes, encryptedData); err != nil {
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
