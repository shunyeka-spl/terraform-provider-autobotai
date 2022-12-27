package pkg

import (
	"net/http"
	"time"
)

type Client struct {
	Apikey     string
	Url        string
	HttpClient *http.Client
}

// create a client with region and apiKey
func NewClient(apikey string, url string) (*Client, error) {

	client := Client{Apikey: apikey, Url: url, HttpClient: &http.Client{
		Timeout: time.Second * 60,
	}}

	_, err := client.validateApiKey()
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// Validate ApiKey by sending API request using the API key provided
func (c *Client) validateApiKey() (*ApiList, error) {

	apiKeyListResult := ApiList{}
	_, err := c.ClientRequest(Get{}, "", nil, "", &apiKeyListResult)

	if err != nil {

		return nil, err
	}
	return &apiKeyListResult, nil
}

// func getUrl() string {
// 	return ("https://api-test.autobot.live")
// }
