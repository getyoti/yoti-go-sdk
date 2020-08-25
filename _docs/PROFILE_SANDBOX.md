# Yoti Go Profile Sandbox Module

This module contains the tools you need to test your Go back-end integration with the Yoti Sandbox service.

## Importing the Sandbox

You can reference the sandbox by adding the following import:

```Go
import "github.com/getyoti/yoti-go-sdk/v3/profile/sandbox"
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

The format of `privateKey` passed in to the client needs to be `*rsa.PrivateKey`. See the [sandboxexample_test.go](../_examples/profilesandbox/sandboxexample_test.go) to see how to easily create this struct.

## Examples

- [Profile Sandbox Example](../_examples/profilesandbox/)