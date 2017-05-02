Yoti Golang SDK
===============

Welcome to the Yoti Golang SDK. This repo contains the tools and step by step instructions you need to quickly integrate your Golang back-end with Yoti so that your users can share their identity details with your application in a secure and trusted way.

## Table of Contents

1) [An Architectural view](#an-architectural-view) -
High level overview of integration

2) [References](#references)-
Guides before you start

3) [Requirements](#requirements)-
Everything you need to get started

4) [Installing the SDK](#installing-the-sdk)-
How to install our SDK

5) [SDK Project import](#sdk-project-import)-
How to install the SDK to your project

6) [Profile Retrieval](#profile-retrieval)-
How to retrieve a Yoti profile using the token

7) [Handling users](#handling-users)-
How to manage users

8) [API Coverage](#api-coverage)-
Attributes defined

9) [Running the tests](running-the-tests)-
Attributes defined

10) [Support](#support)-
Please feel free to reach out

## An Architectural view

Before you start your integration, here is a bit of background on how the integration works. To integrate your application with Yoti, your back-end must expose a GET endpoint that Yoti will use to forward tokens.
The endpoint can be configured in the Yoti Dashboard when you create/update your application. For more information on how to create an application please check our [developer page](https://www.yoti.com/developers/documentation/#login-button-setup).

The image below shows how your application back-end and Yoti integrate into the context of a Login flow.
Yoti SDK carries out for you steps 6, 7 and the profile decryption in step 8.

![alt text](login_flow.png "Login flow")


Yoti also allows you to enable user details verification from your mobile app by means of the Android (TBA) and iOS (TBA) SDKs. In that scenario, your Yoti-enabled mobile app is playing both the role of the browser and the Yoti app. Your back-end doesn't need to handle these cases in a significantly different way. You might just decide to handle the `User-Agent` header in order to provide different responses for desktop and mobile clients.

## References

* [AES-256 symmetric encryption][]
* [RSA pkcs asymmetric encryption][]
* [Protocol buffers][]
* [Base64 data][]

[AES-256 symmetric encryption]:   https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
[RSA pkcs asymmetric encryption]: https://en.wikipedia.org/wiki/RSA_(cryptosystem)
[Protocol buffers]:               https://en.wikipedia.org/wiki/Protocol_Buffers
[Base64 data]:                    https://en.wikipedia.org/wiki/Base64


## Requirements

TBC

## Installing the SDK

To import the Yoti SDK inside your project, simply run the following command from your terminal:

```
go get "github.com/getyoti/go"
```

## SDK Project import

You can reference the project URL by adding the following import:
```golang
import "github.com/getyoti/go"
```

## Configuration

The YotiClient is the SDK entry point. To initialise it you need include the following snippet inside your endpoint initialisation section:
```golang
sdkID := "your-sdk-id";
key, err := ioutil.ReadFile("path/to/your-application-pem-file.pem")
if err != nil {
	// handle key load error
}

client := yoti.YotiClient{
	SdkID: sdkID,
	Key: key}
```
Where:
- `sdkID` is the SDK identifier generated by Yoti Dashboard in the Key tab when you create your app. Note this is not your Application Identifier which is needed by your client-side code.

- `path/to/your-application-pem-file.pem` is the path to the application pem file. It can be downloaded only once from the Keys tab in your Yoti Dashboard.

Please do not open the pem file as this might corrupt the key and you will need to create a new application.

Keeping your settings and access keys outside your repository is highly recommended. You can use gems like [godotenv](https://github.com/joho/godotenv) to manage environment variables more easily.

## Profile Retrieval

When your application receives a token via the exposed endpoint (it will be assigned to a query string parameter named `token`), you can easily retrieve the user profile by adding the following to your endpoint handler:

```golang
profile, err := client.GetUserProfile(yotiToken)
```

Before you inspect the user profile, you might want to check whether the user validation was successful.
This is done as follows:

```golang
profile, err := client.GetUserProfile(yotiToken)
if err != nil {
  // handle unhappy path
}
```

## Handling users

When you retrieve the user profile, you receive a userId generated by Yoti exclusively for your application.
This means that if the same individual logs into another app, Yoti will assign her/him a different ID.
You can use such ID to verify whether the retrieved profile identifies a new or an existing user.
Here is an example of how this works:

```golang
profile, err := client.GetUserProfile(yotiToken)
if err == nil {
	user := YourUserSearchFunction(profile.ID)
	if user != nil {
		// handle login
	} else {
      // handle registration
    }
} else {
    // handle unhappy path
}
```
Where `yourUserSearchFunction` is a piece of logic in your app that is supposed to find a user, given a userID.
No matter if the user is a new or an existing one, Yoti will always provide her/his profile, so you don't necessarily need to store it.

The `profile` object provides a set of attributes corresponding to user attributes. Whether the attributes are present or not depends on the settings you have applied to your app on Yoti Dashboard.

## Running the tests

You can run the unit tests for this project by executing the following command inside the repository folder
```
go test
```

## API Coverage

* Activity Details
    * [X] User ID `user_id`
    * [X] Profile
        * [X] Photo `selfie`
        * [X] Given Names `given_names`
        * [X] Family Name `family_name`
        * [X] Mobile Number `phone_number`
        * [X] Email address `email_address`
        * [X] Date of Birth `date_of_birth`
        * [X] Address `postal_address`
        * [X] Gender `gender`
        * [X] Nationality `nationality`

## Support

For any questions or support please email [sdksupport@yoti.com](mailto:sdksupport@yoti.com).
Please provide the following the get you up and working as quick as possible:

- Computer Type
- OS Version
- Version of Go being used
- Screenshot
