package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const readableConfigFilePath = "$HOME/.qvault.json"

func getConfigFilePath() string {
	return fmt.Sprintf("%s/.qvault.json", os.Getenv("HOME"))
}

type configStructure struct {
	APIKey string `json:"apiKey"`
}

func runAuth() {
	if len(os.Args) < 3 {
		log.Fatal("An API key is required")
	}

	cfgStruct := configStructure{
		APIKey: os.Args[2],
	}
	dat, err := json.Marshal(cfgStruct)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(getConfigFilePath(), dat, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
}

func getAuth() (string, error) {
	dat, err := os.ReadFile(getConfigFilePath())
	if err != nil {
		return "", err
	}
	cfgStruct := configStructure{}
	err = json.Unmarshal(dat, &cfgStruct)
	if err != nil {
		return "", err
	}
	return cfgStruct.APIKey, nil
}
