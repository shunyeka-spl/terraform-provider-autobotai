package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateKubernetesIntegration(KuberenetesIntegrationPayload KubernetesIntegration) (string, error) {

	kubernetesIntegration := KubernetesIntegration{}

	rb, err := json.Marshal(KuberenetesIntegrationPayload)
	if err != nil {
		return "", err
	}
	kubernetesResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &kubernetesIntegration)
	if err != nil {
		return "", err
	}
	kubernetesErr := json.Unmarshal(response, &kubernetesResponse)
	if kubernetesErr != nil {
		return "", kubernetesErr
	}

	return kubernetesResponse.AccountId, nil
}

func (c *Client) DeleteKubernetesIntegration(kubernetesId string) error {

	kubernetesIntegrationId := DeleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", kubernetesId), nil, "", &kubernetesIntegrationId)
	if err != nil {
		return nil
	}

	return nil
}
