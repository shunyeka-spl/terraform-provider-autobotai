package pkg

import (
	"encoding/json"
	"strings"
)

func (c *Client) CreateAgent(AgentPayload Agents) (string, error) {

	agent := Agents{}

	rb, err := json.Marshal(AgentPayload)
	if err != nil {
		return "", err
	}

	agentResponse := AgentsResponse{}
	response, err := c.ClientRequest(Post{}, "/agents", strings.NewReader(string(rb)), "", &agent)
	if err != nil {
		return "", err
	}

	agentErr := json.Unmarshal(response, &agentResponse)
	if agentErr != nil {
		return "", agentErr
	}
	return agentResponse.Id, nil
}
