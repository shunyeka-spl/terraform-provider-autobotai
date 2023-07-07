package pkg

import (
	"encoding/json"
	"strings"
)

func (c *Client) CreateApiKey(ApiKeyPayload ApiKeys) (*ApiKeysResponse, error) {

	apikey := ApiKeys{}

	rb, err := json.Marshal(ApiKeyPayload)
	if err != nil {
		return nil, err
	}
	apikeysResponse := ApiKeysResponse{}
	response, err := c.ClientRequest(Post{}, "/apikeys", strings.NewReader(string(rb)), "", &apikey)
	if err != nil {
		return nil, err
	}

	apikeyErr := json.Unmarshal(response, &apikeysResponse)
	if apikeyErr != nil {
		return nil, apikeyErr
	}
	return &apikeysResponse, nil
}
