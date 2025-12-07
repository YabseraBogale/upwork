package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://effoysira.com/page/1"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the response body
	fmt.Printf("Response Body:\n%s\n", string(body))
}
