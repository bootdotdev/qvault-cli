package main

import "fmt"

func runExec(apiKey string) {
	stepName := "my step"
	fmt.Println("Executing step: ", stepName)
	fmt.Println("API key: ", apiKey)
}
