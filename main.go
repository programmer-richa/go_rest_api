package main

import (
	"fmt"
	"net/http"

	"github.com/programmer-richa/go_rest_api/routes"
	"github.com/programmer-richa/go_rest_api/settings"
)

func main() {
	routes.Routes()
	// Listener
	err := http.ListenAndServe(settings.Port, nil)

	if err != nil {
		fmt.Printf("Could not start application on port 8080. %v", err)
	}
}
