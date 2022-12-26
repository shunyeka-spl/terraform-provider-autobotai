package pkg

import (
	"encoding/json"
	"strings"
)

func (c *Client) GetExternalId(AwsIntegrationPayload AwsIntegration) (*AwsIntegrationResponse, error) {

	awsIntergration := AwsIntegration{}

	rb, err := json.Marshal(AwsIntegrationPayload)
	if err != nil {
		return nil, err
	}

	awsIntegrationResponse := AwsIntegrationResponse{}
	response, err := c.ClientRequest(Post{}, "/integrations/aws/setup", strings.NewReader(string(rb)), "", &awsIntergration)
	if err != nil {
		return nil, err
	}

	awsErr := json.Unmarshal(response, &awsIntegrationResponse)
	if awsErr != nil {
		return nil, awsErr
	}

	return &awsIntegrationResponse, nil
}
