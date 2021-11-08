package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const configFileName = ".qvault.json"

func getConfigFilePath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dirname, configFileName), nil
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
	path, err := getConfigFilePath()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, dat, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
}

func getAuth() (string, error) {
	path, err := getConfigFilePath()
	if err != nil {
		log.Fatal(err)
	}
	dat, err := os.ReadFile(path)
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
