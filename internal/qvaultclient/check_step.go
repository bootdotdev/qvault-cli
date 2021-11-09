package qvaultclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// CheckCurrentStep -
func (c Client) CheckCurrentStep(stepUUID string, currentCheckRequestBody []byte) (CheckCurrentStepResponseBody, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/v1/project_steps/check?step_uuid=%s", domain, stepUUID),
		bytes.NewReader(currentCheckRequestBody),
	)
	if err != nil {
		return CheckCurrentStepResponseBody{}, err
	}
	req.Header.Add("X-API-Key", c.APIKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CheckCurrentStepResponseBody{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return CheckCurrentStepResponseBody{}, errors.New("unexpected stastus code")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CheckCurrentStepResponseBody{}, err
	}

	respBody := CheckCurrentStepResponseBody{}
	err = json.Unmarshal(dat, &respBody)
	return respBody, err
}
