package qvaultclient

import (
	"net/http"
	"time"
)

const domain = "https://qvault-classroom.wm.r.appspot.com"

// Client -
type Client struct {
	httpClient *http.Client
	APIKey     string
}

// NewClient -
func NewClient(apiKey string) Client {
	return Client{
		APIKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 60,
		},
	}
}
