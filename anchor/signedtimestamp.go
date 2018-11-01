package anchor

import (
	"time"

	"github.com/getyoti/yoti-go-sdk/yotiprotocom_v3"
)

// SignedTimestamp is the object which contains a timestamp
type SignedTimestamp struct {
	Version   int32
	Timestamp *time.Time
}

func convertSignedTimestamp(protoSignedTimestamp yotiprotocom_v3.SignedTimestamp) SignedTimestamp {
	uintTimestamp := protoSignedTimestamp.Timestamp
	intTimestamp := int64(uintTimestamp)
	unixTime := time.Unix(intTimestamp/1000000, 0)

	return SignedTimestamp{
		Version:   protoSignedTimestamp.Version,
		Timestamp: &unixTime,
	}
}
