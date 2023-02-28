package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateGoogleChatInetgration(googleChatIntegration GoogleChatIntegration) (string, error) {

	googleChat := GoogleChatIntegration{}

	rb, err := json.Marshal(googleChatIntegration)
	if err != nil {
		return "", err
	}
	googleChatIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &googleChat)
	if err != nil {
		return "", err
	}

	googleChatErr := json.Unmarshal(response, &googleChatIntegrationResponse)
	if googleChatErr != nil {
		return "", googleChatErr
	}

	return googleChatIntegrationResponse.AccountId, nil
}
func (c *Client) GetGoogleChatInetgration(googleChatIntegrationId string) (*IntegrationsResponse, error) {

	googleChatGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", googleChatIntegrationId), nil, "", &googleChatGetIntergration)

	if err != nil {
		return nil, err
	}

	return &googleChatGetIntergration, nil
}
func (c *Client) DeleteGoogleChatIntegration(googleChatIntegrationId string) (*DeleteIntegrationsResponse, error) {

	googleChatDeleteIntergration := DeleteIntegrationsResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", googleChatIntegrationId), nil, "", &googleChatDeleteIntergration)
	if err != nil {
		return nil, err
	}

	return &googleChatDeleteIntergration, nil
}
