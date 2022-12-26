package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateWorkloadSecurityInetgration(workloadSecurityIntegrationPayload WorkloadSecurity) (string, error) {

	workloadSecurityIntergration := WorkloadSecurity{}

	rb, err := json.Marshal(workloadSecurityIntegrationPayload)
	if err != nil {
		return "", err
	}

	workloadSecurityIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &workloadSecurityIntergration)
	if err != nil {
		return "", err
	}

	workloadSecurityErr := json.Unmarshal(response, &workloadSecurityIntegrationResponse)
	if workloadSecurityErr != nil {
		return "", workloadSecurityErr
	}

	return workloadSecurityIntegrationResponse.AccountId, nil
}
func (c *Client) GetWorkloadSecurityInetgration(workloadSecurityId string) (*IntegrationsResponse, error) {
	workloadSecurityGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", workloadSecurityId), nil, "", &workloadSecurityGetIntergration)

	if err != nil {
		return nil, err
	}

	return &workloadSecurityGetIntergration, nil
}

func (c *Client) DeleteWorkloadSecurityIntegration(workloadSecurityId string) (*DeleteIntegrationsResponse, error) {

	workloadSecurityDeleteIntergration := DeleteIntegrationsResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", workloadSecurityId), nil, "", &workloadSecurityDeleteIntergration)
	if err != nil {
		return nil, err
	}

	return &workloadSecurityDeleteIntergration, nil
}
