package main

import "fmt"

func runCurrent(apiKey string) {
	projectName := "HTTP Server in Go"
	fmt.Println("Current project: ", projectName)

	stepName := "Step 2"
	fmt.Println("Current step: ", stepName)

	stepURL := "https://qvault.io"
	fmt.Println("Instructions:", stepURL)
}
