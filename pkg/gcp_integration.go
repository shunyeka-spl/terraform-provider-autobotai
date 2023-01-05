package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateGcpInetgration(GcpIntegrationPayload GCPIntegration) (string, error) {

	gcpIntergration := GCPIntegration{}

	rb, err := json.Marshal(GcpIntegrationPayload)
	if err != nil {
		return "", err
	}

	gcpIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &gcpIntergration)
	if err != nil {
		return "", err
	}
	gcpErr := json.Unmarshal(response, &gcpIntegrationResponse)
	if gcpErr != nil {
		return "", gcpErr
	}

	return gcpIntegrationResponse.AccountId, nil
}
func (c *Client) GetGcpInetgration(GcpIntegrationId string) (*IntegrationsResponse, error) {
	GcpGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", GcpIntegrationId), nil, "", &GcpGetIntergration)

	if err != nil {
		return nil, err
	}

	return &GcpGetIntergration, nil
}

func (c *Client) DeleteGcpIntegration(GcpIntegrationId string) (*DeleteIntegrationsResponse, error) {

	gcpDeleteIntergration := DeleteIntegrationsResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", GcpIntegrationId), nil, "", &gcpDeleteIntergration)
	if err != nil {
		return nil, err
	}

	return &gcpDeleteIntergration, nil
}
