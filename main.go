package main

import (
	"fmt"
	"health-check/health"
	"health-check/pkg/httpclient"
	"log"
)

func main() {
	// URL of the health check service
	url := "https://dev-panel-api.demo.fedshi.ice.global/dev/readiness/"

	// Create an instance of DefaultClient (or mock client in tests)
	client := &httpclient.DefaultClient{}

	// Calling the CheckHealth function with both URL and client as arguments
	err := health.CheckHealth(url, client)
	if err != nil {
		// Log error if health check fails
		log.Printf("Error: %v", err)
	} else {
		// Print success message if the health check is successful
		fmt.Println("Service health check passed!")
	}
}
