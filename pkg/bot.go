package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateBot(botPayload Bot) (string, error) {

	bot := Bot{}

	rb, err := json.Marshal(botPayload)
	if err != nil {
		return "", err
	}

	response, err := c.ClientRequest(Post{}, "/bots", strings.NewReader(string(rb)), "", &bot)
	if err != nil {
		return "", err
	}
	var botDetails interface{}
	botErr := json.Unmarshal(response, &botDetails)
	botResponse := botDetails.(map[string]interface{})
	if botErr != nil {
		return "", botErr
	}
	return botResponse["_id"].(string), nil
}
func (c *Client) GetBot(botId string) (*BotResponse, error) {

	botResponse := BotResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/bots/%s", botId), nil, "", &botResponse)

	if err != nil {
		return nil, err
	}

	return &botResponse, nil
}
func (c *Client) UpdateBot(botId string, updateBotPayload BotResponse) (string, error) {

	botResponse := BotResponse{}
	rb, err := json.Marshal(updateBotPayload)
	if err != nil {
		return "", err
	}

	response, err := c.ClientRequest(Put{}, fmt.Sprintf("/bots/%s", botId), strings.NewReader(string(rb)), "", &botResponse)

	if err != nil {
		return "", err
	}
	var botDetails interface{}
	botErr := json.Unmarshal(response, &botDetails)
	bot := botDetails.(map[string]interface{})
	if botErr != nil {
		return "", botErr
	}
	return bot["_id"].(string), nil
}
func (c *Client) DeleteBot(botId string) error {

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/bots/%s", botId), nil, "", "")
	if err != nil {
		return err
	}

	return nil
}
