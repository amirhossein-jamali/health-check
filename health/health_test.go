package health

import (
	"errors"
	"health-check/pkg/httpclient"
	"health-check/pkg/httpclient/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckHealth(t *testing.T) {
	tests := []struct {
		name       string // Name of the test case
		statusCode int    // HTTP status code to simulate
		body       string // Response body to simulate
		err        error  // Error to simulate from SendGetRequest
		expected   bool   // Expected result: true if no error expected, false otherwise
	}{
		{"Valid Response", http.StatusOK, "OK", nil, true},               // Test with correct response (OK)
		{"Invalid Body", http.StatusOK, "NOT OK", nil, false},            // Test with incorrect response body
		{"Server Error", http.StatusInternalServerError, "", nil, false}, // Test with server error
		{"Network Error", 0, "", errors.New("network error"), false},     // Simulate network error
		{"Empty Response Body", http.StatusOK, "", nil, false},           // Test with empty response body
		{"Large Response Body", http.StatusOK, "OKOKOKOK", nil, false},   // Test with unexpected large body
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new mock client
			mockClient := new(mocks.Client)

			// Setup mock behavior
			mockClient.On("SendGetRequest", "https://dev-panel-api.demo.fedshi.ice.global/dev/readiness/").
				Return(&httpclient.Response{StatusCode: tt.statusCode, Body: tt.body}, tt.err)

			// Call CheckHealth with the mock client
			err := CheckHealth("https://dev-panel-api.demo.fedshi.ice.global/dev/readiness/", mockClient)

			// Assert the result matches the expected outcome
			if tt.expected {
				assert.NoError(t, err, "Expected no error, but got: %v", err)
			} else {
				assert.Error(t, err, "Expected an error but got none")
			}

			// Assert that the mock expectations were met
			mockClient.AssertExpectations(t)
		})
	}
}
