Go Yoti App Integration
=============================

1) [An Architectural View](#an-architectural-view) -
High level overview of integration

1) [Profile Retrieval](#profile-retrieval) -
How to retrieve a Yoti profile using the one time use token

1) [Running the example](#running-the-example) -
Running the profile example

1) [API Coverage](#api-coverage) -
Attributes defined

## An Architectural View

To integrate your application with Yoti, your back-end must expose a GET endpoint that Yoti will use to forward tokens.
The endpoint can be configured in Yoti Hub when you create/update your application.

The image below shows how your application back-end and Yoti integrate in the context of a Login flow.
Yoti SDK carries out for you steps 6, 7 ,8 and the profile decryption in step 9.

![alt text](login_flow.png "Login flow")

Yoti also allows you to enable user details verification from your mobile app by means of the [Android](https://github.com/getyoti/android-sdk-button) and [iOS](https://github.com/getyoti/ios-sdk-button) SDKs. In that scenario, your Yoti-enabled mobile app is playing both the role of the browser and the Yoti app. By the way, your back-end doesn't need to handle these cases in a significantly different way. You might just decide to handle the `User-Agent` header in order to provide different responses for web and mobile clients.


## Profile Retrieval

When your application receives a one time use token via the exposed endpoint (it will be assigned to a query string parameter named `token`), you can easily retrieve the activity details by adding the following to your endpoint handler:

```Go
activityDetails, err := client.GetActivityDetails(yotiOneTimeUseToken)
if err != nil {
  // handle unhappy path
}
```

### Handling Errors
If a network error occurs that can be handled by resending the request,
the error returned by the SDK will implement the temporary error interface.
This can be tested for using either `errors.Is` or a type assertion, and resent.

```Go
while true {
  activityDetails, err := client.GetActivityDetails(token)
  var temp interface{ Temporary() bool }
  if !errors.Is(err, &temp) {
    break
  }
  // Log the temporary error as a warning
}
```

### Retrieving the user profile

You can then get the user profile from the activityDetails struct:

```Go
var rememberMeID string = activityDetails.RememberMeID()
var parentRememberMeID string = activityDetails.ParentRememberMeID()
var userProfile yoti.Profile = activityDetails.UserProfile

var selfie = userProfile.Selfie().Value()
var givenNames string = userProfile.GivenNames().Value()
var familyName string = userProfile.FamilyName().Value()
var fullName string = userProfile.FullName().Value()
var mobileNumber string = userProfile.MobileNumber().Value()
var emailAddress string = userProfile.EmailAddress().Value()
var address string = userProfile.Address().Value()
var gender string = userProfile.Gender().Value()
var nationality string = userProfile.Nationality().Value()
var dateOfBirth *time.Time
dobAttr, err := userProfile.DateOfBirth()
if err != nil {
    // handle error
} else {
    dateOfBirth = dobAttr.Value()
}
var structuredPostalAddress map[string]interface{}
structuredPostalAddressAttribute, err := userProfile.StructuredPostalAddress()
if err != nil {
    // handle error
} else {
    structuredPostalAddress := structuredPostalAddressAttribute.Value().(map[string]interface{})
}
```

If you have chosen "Verify Condition" on the Yoti Hub with the age condition of "Over 18", you can retrieve the user information with the generic .GetAttribute method, which requires the result to be cast to the original type:

```Go
userProfile.GetAttribute("age_over:18").Value().(string)
```

GetAttribute returns an interface, the value can be acquired through a type assertion.

### Anchors, Sources and Verifiers

An `Anchor` represents how a given Attribute has been _sourced_ or _verified_.  These values are created and signed whenever a Profile Attribute is created, or verified with an external party.

For example, an attribute value that was _sourced_ from a Passport might have the following values:

`Anchor` property | Example value
-----|------
type | SOURCE
value | PASSPORT
subType | OCR
signedTimestamp | 2017-10-31, 19:45:59.123789

Similarly, an attribute _verified_ against the data held by an external party will have an `Anchor` of type _VERIFIER_, naming the party that verified it.

From each attribute you can retrieve the `Anchors`, and subsets `Sources` and `Verifiers` (all as `[]*anchor.Anchor`) as follows:

```Go
givenNamesAnchors := userProfile.GivenNames().Anchors()
givenNamesSources := userProfile.GivenNames().Sources()
givenNamesVerifiers := userProfile.GivenNames().Verifiers()
```

You can also retrieve further properties from these respective anchors in the following way:

```Go
var givenNamesFirstAnchor *anchor.Anchor = givenNamesAnchors[0]

var anchorType anchor.Type = givenNamesFirstAnchor.Type()
var signedTimestamp *time.Time = givenNamesFirstAnchor.SignedTimestamp().Timestamp()
var subType string = givenNamesFirstAnchor.SubType()
var value string = givenNamesFirstAnchor.Value()
```

## Running the Example

Follow the below link for instructions on how to run the example project:

1) [Profile example](../_examples/profile/README.md)

## API Coverage

* [X] Activity Details
  * [X] Remember Me ID `RememberMeID()`
  * [X] Parent Remember Me ID `ParentRememberMeID()`
  * [X] User Profile `UserProfile`
    * [X] Selfie `Selfie()`
    * [X] Selfie Base64 URL `Selfie().Value().Base64URL()`
    * [X] Given Names `GivenNames()`
    * [X] Family Name `FamilyName()`
    * [X] Full Name `FullName()`
    * [X] Mobile Number `MobileNumber()`
    * [X] Email Address `EmailAddress()`
    * [X] Date of Birth `DateOfBirth()`
    * [X] Postal Address `Address()`
    * [X] Structured Postal Address `StructuredPostalAddress()`
    * [X] Gender `Gender()`
    * [X] Nationality `Nationality()`