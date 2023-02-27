package pkg

import (
	"fmt"
)

func (c *Client) GetAwsInetgration(awsAccountId string) (*IntegrationsResponse, error) {

	awsGetIntergration := IntegrationsResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/integrations/id/%s", awsAccountId), nil, "", &awsGetIntergration)

	if err != nil {
		return nil, err
	}

	return &awsGetIntergration, nil
}
func (c *Client) DeleteAwsIntegration(AwsAccountId string) (*DeleteResponse, error) {

	awsDeleteIntergration := DeleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/integrations/%s", AwsAccountId), nil, "", &awsDeleteIntergration)

	if err != nil {
		return nil, err
	}

	return &awsDeleteIntergration, nil
}
