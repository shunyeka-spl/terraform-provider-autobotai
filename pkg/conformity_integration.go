package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateConformityInetgration(conformityIntegrationPayload Conformity) (string, error) {

	conformityIntergration := Conformity{}

	rb, err := json.Marshal(conformityIntegrationPayload)
	if err != nil {
		return "", err
	}

	conformityIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &conformityIntergration)
	if err != nil {
		return "", err
	}

	conformityErr := json.Unmarshal(response, &conformityIntegrationResponse)
	if conformityErr != nil {
		return "", conformityErr
	}

	return conformityIntegrationResponse.AccountId, nil
}
func (c *Client) GetConformityInetgration(conformityId string) (*IntegrationsResponse, error) {
	conformityGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", conformityId), nil, "", &conformityGetIntergration)

	if err != nil {
		return nil, err
	}

	return &conformityGetIntergration, nil
}

func (c *Client) DeleteConformityIntegration(conformityId string) (*DeleteIntegrationsResponse, error) {

	conformityDeleteIntergration := DeleteIntegrationsResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", conformityId), nil, "", &conformityDeleteIntergration)
	if err != nil {
		return nil, err
	}

	return &conformityDeleteIntergration, nil
}
