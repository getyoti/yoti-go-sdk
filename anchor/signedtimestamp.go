package anchor

import (
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotocom"
)

// SignedTimestamp is the object which contains a timestamp
type SignedTimestamp struct {
	version   int32
	timestamp *time.Time
}

func convertSignedTimestamp(protoSignedTimestamp yotiprotocom.SignedTimestamp) SignedTimestamp {
	uintTimestamp := protoSignedTimestamp.Timestamp
	intTimestamp := int64(uintTimestamp)
	unixTime := time.Unix(intTimestamp/1000000, 0)

	return SignedTimestamp{
		version:   protoSignedTimestamp.Version,
		timestamp: &unixTime,
	}
}

// Version indicates both the version of the protobuf message in use,
// as well as the specific hash algorithms.
func (s SignedTimestamp) Version() int32 {
	return s.version
}

// Timestamp is a point in time, to the nearest microsecond.
func (s SignedTimestamp) Timestamp() *time.Time {
	return s.timestamp
}
