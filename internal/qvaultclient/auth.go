package qvaultclient

import (
	"fmt"
	"net/http"
)

const qvaultSettingsURL = "https://api.qvault.io/dashboard/settings"

// CheckAPIKey -
func (c Client) CheckAPIKey(apiKey string) error {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/v1/apikey/validate", domain),
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Add("X-API-Key", c.APIKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid API key, please create a new one here: %s", qvaultSettingsURL)
	}
	return nil
}
