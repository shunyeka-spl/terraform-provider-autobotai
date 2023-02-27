package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateAwsSesConfigureIntegration(AwsSesPayload AwsSesConfigure) (string, error) {

	awsSesConfigureIntegration := AwsSesConfigure{}

	rb, err := json.Marshal(AwsSesPayload)
	if err != nil {
		return "", err
	}

	awsSesIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &awsSesConfigureIntegration)
	if err != nil {
		return "", err
	}

	awsSesErr := json.Unmarshal(response, &awsSesIntegrationResponse)
	if awsSesErr != nil {
		return "", awsSesErr
	}

	return awsSesIntegrationResponse.AccountId, nil
}
func (c *Client) GetAwsSesConfigureIntegration(AzureIntegrationId string) (*IntegrationsResponse, error) {

	awsSesConfigureGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", AzureIntegrationId), nil, "", &awsSesConfigureGetIntergration)

	if err != nil {
		return nil, err
	}

	return &awsSesConfigureGetIntergration, nil
}
func (c *Client) DeleteAwsSesConfigureIntegration(AzureIntegrationId string) (*DeleteResponse, error) {

	awsSesConfigureDeleteIntergration := DeleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", AzureIntegrationId), nil, "", &awsSesConfigureDeleteIntergration)
	if err != nil {
		return nil, err
	}

	return &awsSesConfigureDeleteIntergration, nil
}
