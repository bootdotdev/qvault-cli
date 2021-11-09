package main

import (
	"fmt"
	"os"

	"github.com/qvault/qvault-cli/internal/qvaultclient"
)

func runCheck(apiKey string) {
	projectUUID, err := getProjectUUID()
	if err != nil {
		fmt.Println("didn't find project id in local .qvault.json file. Please run 'qvault project [project id]'")
		os.Exit(1)
	}

	c := qvaultclient.NewClient(apiKey)

	currentStepResp, err := c.GetCurrentStep(projectUUID)
	if err != nil {
		fmt.Println("couldn't get current step.")
		os.Exit(1)
	}

	fmt.Printf("checking step: %v/%v\n", currentStepResp.StepNumber, currentStepResp.NumberOfSteps)
}
