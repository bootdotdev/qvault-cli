package main

import (
	"fmt"
	"log"
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

	if len(os.Args) > 1 && os.Args[1] == "project" {
		setProjectUUID()
		return
	}

	apiKey, err := getAuth()
	if err != nil {
		path, err := getConfigFilePath()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("API key not found in", path)
		fmt.Println("Please run 'qvault auth [token]'")
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "check" {
		runCheck(apiKey)
		return
	}
}
