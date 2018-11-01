package anchor

import (
	"crypto/x509"
	"encoding/asn1"
	"log"

	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
	"github.com/getyoti/yoti-go-sdk/yotiprotocom_v3"
	"github.com/golang/protobuf/proto"
)

type anchorExtension struct {
	Extension string `asn1:"tag:0,utf8"`
}

var sourceOID = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 47127, 1, 1, 1}
var verifierOID = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 47127, 1, 1, 2}

//ParseAnchors takes a slice of protobuf anchors, parses them, and returns a slice of Yoti SDK Anchors
func ParseAnchors(protoAnchors []*yotiprotoattr_v3.Anchor) []*Anchor {
	var processedAnchors []*Anchor
	for _, protoAnchor := range protoAnchors {
		var extensions []string
		var anchorType = AnchorTypeUnknown

		var parsedCerts = parseCertificates(protoAnchor.OriginServerCerts)

		for _, cert := range parsedCerts {
			for _, ext := range cert.Extensions {
				var ae anchorExtension
				if sourceOID.Equal(ext.Id) {
					anchorType = AnchorTypeSource
				} else if verifierOID.Equal(ext.Id) {
					anchorType = AnchorTypeVerifier
				}

				_, err := asn1.Unmarshal(ext.Value, &ae)
				if err == nil {
					extensions = append(extensions, ae.Extension)
				} else {
					log.Printf("Error unmarshalling anchor extension: %q", err)
				}
			}
		}

		processedAnchor := newAnchor(anchorType, parsedCerts, parseSignedTimestamp(protoAnchor.SignedTimeStamp), protoAnchor.SubType, extensions)

		processedAnchors = append(processedAnchors, processedAnchor)
	}

	return processedAnchors
}

func parseSignedTimestamp(rawBytes []byte) yotiprotocom_v3.SignedTimestamp {
	signedTimestamp := &yotiprotocom_v3.SignedTimestamp{}
	if err := proto.Unmarshal(rawBytes, signedTimestamp); err != nil {
		signedTimestamp = nil
	}

	return *signedTimestamp
}

func parseCertificates(rawCerts [][]byte) (result []*x509.Certificate) {
	for _, cert := range rawCerts {
		parsedCertificate, _ := x509.ParseCertificate(cert)

		result = append(result, parsedCertificate)
	}

	return result
}
