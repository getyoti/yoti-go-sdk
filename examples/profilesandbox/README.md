# Yoti Go Sandbox Module

This module contains the tools you need to test your Go back-end integration with the Yoti Sandbox service.

## Importing the Sandbox

You can reference the sandbox by adding the following import:

```Go
import "github.com/getyoti/yoti-go-sdk/v2/profile/sandbox"
```

## Configuration
The sandbox is initialised in the following way:
```Go
sandboxClient := sandbox.Client{
		ClientSdkID: sandboxClientSdkId,
		Key:         privateKey,
	}
```
* `sandboxClientSdkId` is the Sandbox SDK identifier generated from the Sandbox section on Yoti Hub.
* `privateKey` is the PEM file for your Sandbox application downloaded from the Yoti Hub, in the Sandbox section.

Please do not open the PEM file, as this might corrupt the key, and you will need to redownload it.

The format of `privateKey` passed in to the client needs to be `*rsa.PrivateKey`. See the [sandboxexample_test.go](sandboxexample_test.go) to see how to easily create this struct.

## Example

See the [test_sandboxexample.go](test_sandboxexample.go) file for an example of how to use the Sandbox in your tests.
To run the example:
1. Copy the `.env.example` file and rename it to be `.env`.
1. Then modify the `SANDBOX_CLIENT_SDK_ID` and `YOTI_KEY_FILE_PATH` environment variables mentioned above.

To make this test runnable, the name needs to end in `_test.go`, so rename it accordingly.