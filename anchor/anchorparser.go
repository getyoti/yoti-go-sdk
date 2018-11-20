package anchor

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"log"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotocom"
	"github.com/golang/protobuf/proto"
)

type anchorExtension struct {
	Extension string `asn1:"tag:0,utf8"`
}

var (
	sourceOID   = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 47127, 1, 1, 1}
	verifierOID = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 47127, 1, 1, 2}
)

// ParseAnchors takes a slice of protobuf anchors, parses them, and returns a slice of Yoti SDK Anchors
func ParseAnchors(protoAnchors []*yotiprotoattr.Anchor) []*Anchor {
	var processedAnchors []*Anchor
	for _, protoAnchor := range protoAnchors {
		var extensions []string
		var (
			anchorType  = AnchorTypeUnknown
			parsedCerts = parseCertificates(protoAnchor.OriginServerCerts)
		)
		for _, cert := range parsedCerts {
			for _, ext := range cert.Extensions {
				var (
					value string
					err   error
				)
				anchorType, value, err = parseExtension(ext)
				if err != nil {
					log.Printf("error parsing anchor extension, %v", err)
					continue
				} else if anchorType == AnchorTypeUnknown {
					continue
				}
				extensions = append(extensions, value)
			}
		}

		processedAnchor := newAnchor(anchorType, parsedCerts, parseSignedTimestamp(protoAnchor.SignedTimeStamp), protoAnchor.SubType, extensions)

		processedAnchors = append(processedAnchors, processedAnchor)
	}

	return processedAnchors
}

func parseExtension(ext pkix.Extension) (anchorType Type, val string, err error) {
	anchorType = AnchorTypeUnknown

	switch {
	case ext.Id.Equal(sourceOID):
		anchorType = AnchorTypeSource
	case ext.Id.Equal(verifierOID):
		anchorType = AnchorTypeVerifier
	default:
		return anchorType, "", nil
	}

	var ae anchorExtension
	_, err = asn1.Unmarshal(ext.Value, &ae)
	switch {
	case err != nil:
		return anchorType, "", fmt.Errorf("unable to unmarshal extension: %v", err)
	case len(ae.Extension) == 0:
		return anchorType, "", errors.New("empty extension")
	default:
		val = ae.Extension
	}

	return anchorType, val, nil
}

func parseSignedTimestamp(rawBytes []byte) yotiprotocom.SignedTimestamp {
	signedTimestamp := &yotiprotocom.SignedTimestamp{}
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
