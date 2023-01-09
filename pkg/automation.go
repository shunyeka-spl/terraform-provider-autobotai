package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateAutomation(AutomationPayload Automation) (string, error) {

	automation := Automation{}

	rb, err := json.Marshal(AutomationPayload)
	if err != nil {
		return "", err
	}

	automationResponse := AutomationResponse{}
	response, err := c.ClientRequest(Post{}, "/automations", strings.NewReader(string(rb)), "", &automation)
	if err != nil {
		return "", err
	}

	automationErr := json.Unmarshal(response, &automationResponse)
	if automationErr != nil {
		return "", automationErr
	}

	return automationResponse.ID, nil
}
func (c *Client) GetAutomation(automationId string) (*AutomationResponse, error) {

	automationResponse := AutomationResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/automations/%s", automationId), nil, "", &automationResponse)

	if err != nil {
		return nil, err
	}

	return &automationResponse, nil
}
func (c *Client) UpdateAutomation(automationId string, updateAutomationPayload AutomationResponse) (string, error) {

	autoationResponse := AutomationResponse{}
	rb, err := json.Marshal(updateAutomationPayload)
	if err != nil {
		return "", err
	}

	response, err := c.ClientRequest(Put{}, fmt.Sprintf("/automations/%s", automationId), strings.NewReader(string(rb)), "", &autoationResponse)

	if err != nil {
		return "", err
	}
	automationErr := json.Unmarshal(response, &autoationResponse)
	if automationErr != nil {
		return "", automationErr
	}

	return autoationResponse.ID, nil
}
func (c *Client) DeleteAutomation(automationId string) error {

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/automations/%s", automationId), nil, "", "")
	if err != nil {
		return err
	}

	return nil
}
