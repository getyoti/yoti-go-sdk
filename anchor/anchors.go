package anchor

import (
	"crypto/x509"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotocom"
)

// Anchor is the metadata associated with an attribute. It describes how an attribute has been provided
// to Yoti (SOURCE Anchor) and how it has been verified (VERIFIER Anchor).
// If an attribute has only one SOURCE Anchor with the value set to
// "USER_PROVIDED" and zero VERIFIER Anchors, then the attribute
// is a self-certified one.
type Anchor struct {
	anchorType        Type
	originServerCerts []*x509.Certificate
	signedTimestamp   SignedTimestamp
	subtype           string
	value             []string
}

func newAnchor(anchorType Type, originServerCerts []*x509.Certificate, signedTimestamp yotiprotocom.SignedTimestamp, subtype string, value []string) *Anchor {
	return &Anchor{
		anchorType:        anchorType,
		originServerCerts: originServerCerts,
		signedTimestamp:   convertSignedTimestamp(signedTimestamp),
		subtype:           subtype,
		value:             value,
	}
}

// Type Anchor type, based on the Object Identifier (OID)
type Type int

const (
	// AnchorTypeUnknown - default value
	AnchorTypeUnknown Type = 1 + iota
	// AnchorTypeSource - how the anchor has been sourced
	AnchorTypeSource
	// AnchorTypeVerifier - how the anchor has been verified
	AnchorTypeVerifier
)

// Type of the Anchor - most likely either SOURCE or VERIFIER, but it's
// possible that new Anchor types will be added in future.
func (a Anchor) Type() Type {
	return a.anchorType
}

// OriginServerCerts are the X.509 certificate chain(DER-encoded ASN.1)
// from the service that assigned the attribute.
//
// The first certificate in the chain holds the public key that can be
// used to verify the Signature field; any following entries (zero or
// more) are for intermediate certificate authorities (in order).
//
// The last certificate in the chain must be verified against the Yoti root
// CA certificate. An extension in the first certificate holds the main artifact type,
// e.g. “PASSPORT”, which can be retrieved with .Value().
func (a Anchor) OriginServerCerts() []*x509.Certificate {
	return a.originServerCerts
}

// SignedTimestamp is the time at which the signature was created. The
// message associated with the timestamp is the marshaled form of
// AttributeSigning (i.e. the same message that is signed in the
// Signature field). This method returns the SignedTimestamp
// object, the actual timestamp as a *time.Time can be called with
// .Timestamp() on the result of this function.
func (a Anchor) SignedTimestamp() SignedTimestamp {
	return a.signedTimestamp
}

// SubType is an indicator of any specific processing method, or
// subcategory, pertaining to an artifact. For example, for a passport, this would be
// either "NFC" or "OCR".
func (a Anchor) SubType() string {
	return a.subtype
}

// Value identifies the provider that either sourced or verified the attribute value.
// The range of possible values is not limited. For a SOURCE anchor, expect values like
// PASSPORT, DRIVING_LICENSE. For a VERIFIER anchor expect valuues like YOTI_ADMIN.
func (a Anchor) Value() []string {
	return a.value
}

// GetSources returns the anchors which identify how and when an attribute value was acquired.
func GetSources(anchors []*Anchor) (sources []*Anchor) {
	return filterAnchors(anchors, AnchorTypeSource)
}

// GetVerifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func GetVerifiers(anchors []*Anchor) (sources []*Anchor) {
	return filterAnchors(anchors, AnchorTypeVerifier)
}

func filterAnchors(anchors []*Anchor, anchorType Type) (result []*Anchor) {
	for _, v := range anchors {
		if v.anchorType == anchorType {
			result = append(result, v)
		}
	}
	return result
}
