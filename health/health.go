package health

import (
	"fmt"
	"health-check/pkg/httpclient"
	"strings"
)

// CheckHealth sends a GET request to the provided URL and checks the health of the service
func CheckHealth(url string, client httpclient.Client) error {
	// Use client to send GET request
	resp, err := client.SendGetRequest(url)
	if err != nil {
		// Log the error and return a wrapped error
		return fmt.Errorf("error sending request: %v", err)
	}

	// Check if the response status code is in the 2xx range (success)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Trim spaces and remove quotes before comparing
	cleanedBody := strings.Trim(resp.Body, "\"") // Remove surrounding quotes

	// Compare using strings.Compare for a more reliable comparison
	if strings.Compare(cleanedBody, "OK") != 0 {
		return fmt.Errorf("unexpected response body: %q", cleanedBody)
	}

	// Return nil if the health check passes
	return nil
}
