package qvaultclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// GetCurrentStep -
func (c Client) GetCurrentStep(projectUUID string) (GetCurrentStepResponseBody, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/v1/project_steps/current?project_uuid=%s", domain, projectUUID),
		nil,
	)
	if err != nil {
		return GetCurrentStepResponseBody{}, err
	}
	req.Header.Add("X-API-Key", c.APIKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return GetCurrentStepResponseBody{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return GetCurrentStepResponseBody{}, errors.New("unexpected stastus code")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetCurrentStepResponseBody{}, err
	}

	respBody := GetCurrentStepResponseBody{}
	err = json.Unmarshal(dat, &respBody)
	return respBody, err
}
