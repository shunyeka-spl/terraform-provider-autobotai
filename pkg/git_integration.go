package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateGitInetgration(gitIntegrationPayload Git) (string, error) {

	gitIntergration := Git{}

	rb, err := json.Marshal(gitIntegrationPayload)
	if err != nil {
		return "", err
	}

	gitIntegrationResponse := IntegrationsResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations", strings.NewReader(string(rb)), "", &gitIntergration)
	if err != nil {
		return "", err
	}

	gitErr := json.Unmarshal(response, &gitIntegrationResponse)
	if gitErr != nil {
		return "", gitErr
	}

	return gitIntegrationResponse.AccountId, nil
}
func (c *Client) GetGitInetgration(gitId string) (*IntegrationsResponse, error) {

	gitGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", gitId), nil, "", &gitGetIntergration)

	if err != nil {
		return nil, err
	}

	return &gitGetIntergration, nil
}

func (c *Client) DeleteGitIntegration(gitId string) (*DeleteIntegrationsResponse, error) {

	gitDeleteIntergration := DeleteIntegrationsResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", gitId), nil, "", &gitDeleteIntergration)
	if err != nil {
		return nil, err
	}

	return &gitDeleteIntergration, nil
}
