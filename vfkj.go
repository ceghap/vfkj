package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Function to make a POST request and return whether it was successful
func makePostRequest(uuid, answer, identityType string) bool {
	// Prepare form data
	data := url.Values{}
	data.Set("uuid", uuid)
	data.Set("answers[]", answer)
	data.Set("identity_type", identityType)

	// Convert form data to a string reader for the body
	bodyData := strings.NewReader(data.Encode())

	// Create a new POST request
	req, err := http.NewRequest("POST", "https://vote-client.mediaprima.com.my/web/vote", bodyData)
	if err != nil {
		fmt.Printf("Error creating POST request: %v\n", err)
		return false
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Length", "109")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://www.abpbh.com.my")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://www.abpbh.com.my/")
	req.Header.Set("Sec-CH-UA", `"Chromium";v="128", "Not;A=Brand";v="24", "Google Chrome";v="128"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	// Send the request using http.DefaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return false
	}

	// Check if the response body contains "Job Queued."
	return resp.StatusCode == http.StatusOK && strings.Contains(string(body), "Job Queued.")
}

func main() {
	// UUID, answer, and identityType values
	uuid := "aab0e78f-8025-4ecd-a6c3-e3b1a77bc1fe"
	answer := "a2b301ea-8807-4934-ad79-3396d638c15a"
	identityType := "ip"

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Infinite loop to submit votes every 1 second
	successCount := 0
	for {
		// Make the POST request and count successes
		if makePostRequest(uuid, answer, identityType) {
			successCount++
			fmt.Printf("Vote %d: Success\n", successCount)
		} else {
			fmt.Printf("Vote %d: Failed\n", successCount+1)
		}

		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}
}
