package yoti

import (
	"fmt"
)

type yotiClientMock struct {
	mockGetSdkID    func() string
	mockMakeRequest func(string, string, []byte, ...map[int]string) (string, error)
}

func (mock *yotiClientMock) GetSdkID() string {
	if mock.mockGetSdkID != nil {
		return mock.mockGetSdkID()
	}
	panic("Mock undefined")
}

func (mock *yotiClientMock) makeRequest(httpMethod, endpoint string, payload []byte, httpErrorMessages ...map[int]string) (string, error) {
	if mock.mockMakeRequest != nil {
		return mock.mockMakeRequest(httpMethod, endpoint, payload, httpErrorMessages...)
	}
	panic("Mock undefined")
}

func ExampleCreateShareURL() {
	mockYoti := yotiClientMock{
		mockGetSdkID: func() string { return "0000-0000-0000-0000" },
		mockMakeRequest: func(string, string, []byte, ...map[int]string) (string, error) {
			return "{\"qrcode\":\"https://code.yoti.com/CAEaJDQzNzllZDc0LTU0YjItNDkxMy04OTE4LTExYzM2ZDU2OTU3ZDAC\",\"ref_id\":\"0\"}", nil
		},
	}

	policy := (&DynamicPolicyBuilder{}).New().WithFullName().WithWantedRememberMe().Build()
	scenario := (&DynamicScenarioBuilder{}).New().WithPolicy(policy).Build()

	share, _ := CreateShareURL(&mockYoti, &scenario)

	fmt.Printf("QR code: %s", share.ShareURL)
	// Output: QR code: https://code.yoti.com/CAEaJDQzNzllZDc0LTU0YjItNDkxMy04OTE4LTExYzM2ZDU2OTU3ZDAC

}
