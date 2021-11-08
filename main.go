package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "help" {
		runHelp()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "auth" {
		runAuth()
		return
	}

	apiKey, err := getAuth()
	if err != nil {
		fmt.Println("API key not found in", readableConfigFilePath)
		fmt.Println("Please run 'qvault auth [token]")
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "current" {
		runCurrent(apiKey)
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "exec" {
		runExec(apiKey)
		return
	}
}
