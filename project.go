package main

import (
	"encoding/json"
	"log"
	"os"
)

const localConfigFileName = ".qvault.json"

type localConfigStructure struct {
	ProjectUUID string `json:"projectUUID"`
}

func setProjectUUID() {
	if len(os.Args) < 3 {
		log.Fatal("a project id is required")
	}

	cfgStruct := localConfigStructure{
		ProjectUUID: os.Args[2],
	}
	dat, err := json.Marshal(cfgStruct)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(localConfigFileName, dat, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
}

func getProjectUUID() (string, error) {
	dat, err := os.ReadFile(localConfigFileName)
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
