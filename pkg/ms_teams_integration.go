package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateMSTeamsInetgration(MSTeamsIntegrationPayload MSTeamsIntegration) (string, error) {

	msTeamsIntergration := MSTeamsIntegration{}

	rb, err := json.Marshal(MSTeamsIntegrationPayload)
	if err != nil {
		return "", err
	}

	msTeamsIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &msTeamsIntergration)
	if err != nil {
		return "", err
	}
	msTeamsErr := json.Unmarshal(response, &msTeamsIntegrationResponse)
	if msTeamsErr != nil {
		return "", msTeamsErr
	}

	return msTeamsIntegrationResponse.AccountId, nil
}
func (c *Client) GetMSTeamsInetgration(msTeamsIntegrationId string) (*IntegrationsResponse, error) {
	msTeamsIntergrationResponse := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", msTeamsIntegrationId), nil, "", &msTeamsIntergrationResponse)

	if err != nil {
		return nil, err
	}

	return &msTeamsIntergrationResponse, nil
}

func (c *Client) DeleteMSTeamsIntegration(msTeamsIntegrationId string) (*DeleteIntegrationsResponse, error) {

	msTeamsIntergrationResponse := DeleteIntegrationsResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", msTeamsIntegrationId), nil, "", &msTeamsIntergrationResponse)
	if err != nil {
		return nil, err
	}

	return &msTeamsIntergrationResponse, nil
}
