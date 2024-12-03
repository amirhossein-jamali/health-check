package httpclient

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Response struct holds the HTTP response data
type Response struct {
	StatusCode int    // HTTP status code
	Body       string // Response body
}

// Client defines the interface for sending HTTP requests
type Client interface {
	SendGetRequest(url string) (*Response, error)
}

// DefaultClient is the default implementation of the Client interface
type DefaultClient struct{}

// SendGetRequest sends a GET request to the specified URL and returns the response
func (c *DefaultClient) SendGetRequest(url string) (*Response, error) {
	// Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		// Return error if request fails
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	// Ensure the response body gets closed once the function completes
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Return error if reading the body fails
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Return the response with status code and body
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}, nil
}
