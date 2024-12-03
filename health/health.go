package health

import (
	"fmt"
	"health-check/pkg/httpclient"
	"log"
)

// CheckHealth sends a GET request to the provided URL and checks the health of the service
func CheckHealth(url string, client httpclient.Client) error {
	// Use client to send GET request
	resp, err := client.SendGetRequest(url)
	if err != nil {
		// Log the error and return a wrapped error
		log.Printf("Health check request failed: %v", err)
		return fmt.Errorf("error sending request: %v", err)
	}

	// Check if the response status code is in the 2xx range (success)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Return error if status code is not in 2xx range
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Optionally, you can check the response body for "OK" or other criteria (but not mandatory)
	if resp.Body != "OK" {
		return fmt.Errorf("unexpected response body: %s", resp.Body)
	}

	// Return nil if the health check passes
	return nil
}
