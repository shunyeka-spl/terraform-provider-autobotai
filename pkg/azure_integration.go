package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateAzureInetgration(AzureIntegrationPayload AzureIntegrations) (string, error) {

	azureIntergration := AzureIntegrations{}

	rb, err := json.Marshal(AzureIntegrationPayload)
	if err != nil {
		return "", err
	}

	azureIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &azureIntergration)
	if err != nil {
		return "", err
	}

	azureErr := json.Unmarshal(response, &azureIntegrationResponse)
	if azureErr != nil {
		return "", azureErr
	}

	return azureIntegrationResponse.AccountId, nil
}
func (c *Client) GetAzureInetgration(AzureIntegrationId string) (*IntegrationsResponse, error) {

	azureGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", AzureIntegrationId), nil, "", &azureGetIntergration)

	if err != nil {
		return nil, err
	}

	return &azureGetIntergration, nil
}
func (c *Client) DeleteAzureIntegration(AzureIntegrationId string) (*DeleteIntegrationsResponse, error) {

	azureDeleteIntergration := DeleteIntegrationsResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", AzureIntegrationId), nil, "", &azureDeleteIntergration)
	if err != nil {
		return nil, err
	}

	return &azureDeleteIntergration, nil
}
