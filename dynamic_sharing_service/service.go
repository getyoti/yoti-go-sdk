package dynamic_sharing_service

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/web"
)

// GetDynamicShareEndpoint gets the Dynamic Share Endpoint URI
func GetDynamicShareEndpoint(client web.ClientInterface) (string, error) {
	return fmt.Sprintf(
		"/qrcodes/apps/%s",
		client.GetSdkID(),
	), nil
}
