// test_session_json.go
// Basit bir test programı - Session sonucunu raw JSON olarak gösterir
// Kullanım: go run test_session_json.go <SESSION_ID>

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/docscan"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run test_session_json.go <SESSION_ID>")
		fmt.Println("Örnek: go run test_session_json.go 12345678-1234-1234-1234-123456789012")
		os.Exit(1)
	}

	sessionID := os.Args[1]

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	// Initialize DocScan client
	sdkID := os.Getenv("YOTI_CLIENT_SDK_ID")
	key, err := os.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	if err != nil {
		fmt.Printf("Error reading key file: %v\n", err)
		os.Exit(1)
	}

	client, err := docscan.NewClient(sdkID, key)
	if err != nil {
		fmt.Printf("Error creating DocScan client: %v\n", err)
		os.Exit(1)
	}

	// Get session result
	sessionResult, err := client.GetSession(sessionID)
	if err != nil {
		fmt.Printf("Error getting session: %v\n", err)
		os.Exit(1)
	}

	// Output raw JSON only
	jsonData, err := json.MarshalIndent(sessionResult, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
