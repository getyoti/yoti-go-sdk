package main

import (
	"log"
	"os"
	"strconv"

	"github.com/getyoti/yoti-go-sdk/v3"
	"github.com/getyoti/yoti-go-sdk/v3/aml"
	_ "github.com/joho/godotenv/autoload"
)

var (
	sdkID  string
	key    []byte
	client *yoti.Client
)

func main() {
	var err error
	key, err = os.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")

	if err != nil {
		log.Printf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: `%s`", err)
		return
	}

	client, err = yoti.NewClient(sdkID, key)
	if err != nil {
		log.Printf("Problem initialising client: Error: `%s`", err)
		return
	}

	givenNames := "Edward Richard George"
	familyName := "Heath"

	amlAddress := aml.Address{
		Country: "GBR"}

	amlProfile := aml.Profile{
		GivenNames: givenNames,
		FamilyName: familyName,
		Address:    amlAddress}

	var result aml.Result
	result, err = client.PerformAmlCheck(amlProfile)

	if err != nil {
		log.Printf(
			"Unable to retrieve AML result. Error: %s", err)
	} else {
		log.Printf(
			"AML Result for %s %s:",
			givenNames,
			familyName)
		log.Printf(
			"On PEP list: %s",
			strconv.FormatBool(result.OnPEPList))
		log.Printf(
			"On Fraud list: %s",
			strconv.FormatBool(result.OnFraudList))
		log.Printf(
			"On Watch list: %s",
			strconv.FormatBool(result.OnWatchList))
	}
}
