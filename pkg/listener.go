package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateListener(listenerPayload Listener) (string, error) {

	listener := Listener{}

	rb, err := json.Marshal(listenerPayload)
	if err != nil {
		return "", err
	}

	listenerResponse := ListenerResponse{}
	response, err := c.ClientRequest(Post{}, "/listeners", strings.NewReader(string(rb)), "", &listener)
	if err != nil {
		return "", err
	}

	listenerErr := json.Unmarshal(response, &listenerResponse)
	if listenerErr != nil {
		return "", listenerErr
	}

	return listenerResponse.ID, nil
}
func (c *Client) GetListener(listenerId string) (*ListenerResponse, error) {

	listenerResponse := ListenerResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/listeners/%s", listenerId), nil, "", &listenerResponse)

	if err != nil {
		return nil, err
	}

	return &listenerResponse, nil
}
func (c *Client) UpdateListener(listenerId string, updatedListenerPayload Listener) (string, error) {

	listener := Listener{}
	rb, err := json.Marshal(updatedListenerPayload)
	if err != nil {
		return "", err
	}

	response, err := c.ClientRequest(Put{}, fmt.Sprintf("/listeners/%s", listenerId), strings.NewReader(string(rb)), "", &listener)

	if err != nil {
		return "", err
	}
	listenerResponse := ListenerResponse{}
	listenerErr := json.Unmarshal(response, &listenerResponse)
	if listenerErr != nil {
		return "", listenerErr
	}
	return listenerResponse.ID, nil
}
func (c *Client) DeleteListener(listenerId string) error {

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/listeners/%s", listenerId), nil, "", "")
	if err != nil {
		return err
	}

	return nil
}
