# Yoti Go IDV (Doc Scan) Sandbox Module

This module contains the tools you need to test your Go back-end integration with the IDV Sandbox service.

## Importing the Sandbox

You can reference the sandbox by adding the following import:

```Go
import "github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox"
```

## Configuration
The sandbox is initialised in the following way:
```Go
sandboxClient := sandbox.NewClient(sandboxClientSdkId, privateKey)
```
* `sandboxClientSdkId` is the Sandbox SDK identifier generated from the Sandbox section on Yoti Hub.
* `privateKey` is the PEM file for your Sandbox application downloaded from the Yoti Hub, in the Sandbox section.

Please do not open the PEM file, as this might corrupt the key, and you will need to redownload it.

## Examples

- [IDV Sandbox WebDriver Example](../_examples/docscansandbox/)