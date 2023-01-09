package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateFetcher(fetcherPayload Fetcher) (string, error) {

	fetcher := Fetcher{}

	rb, err := json.Marshal(fetcherPayload)
	if err != nil {
		return "", err
	}

	fetcherResponse := FetcherResponse{}
	response, err := c.ClientRequest(Post{}, "/fetchers", strings.NewReader(string(rb)), "", &fetcher)
	if err != nil {
		return "", err
	}

	fetcherErr := json.Unmarshal(response, &fetcherResponse)
	if fetcherErr != nil {
		return "", fetcherErr
	}

	return fetcherResponse.ID, nil
}
func (c *Client) GetFetcher(fetcherId string) (*FetcherResponse, error) {

	fetcherResponse := FetcherResponse{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/fetchers/%s", fetcherId), nil, "", &fetcherResponse)

	if err != nil {
		return nil, err
	}

	return &fetcherResponse, nil
}
func (c *Client) UpdateFetcher(fetcherId string, updateFetcherPayload FetcherResponse) (string, error) {

	fetcherResponse := FetcherResponse{}
	rb, err := json.Marshal(updateFetcherPayload)
	if err != nil {
		return "", err
	}

	response, err := c.ClientRequest(Put{}, fmt.Sprintf("/fetchers/%s", fetcherId), strings.NewReader(string(rb)), "", &fetcherResponse)

	if err != nil {
		return "", err
	}
	fetcherErr := json.Unmarshal(response, &fetcherResponse)
	if fetcherErr != nil {
		return "", fetcherErr
	}

	return fetcherResponse.ID, nil
}
func (c *Client) DeleteFetcher(fetcherId string) error {

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/fetchers/%s", fetcherId), nil, "", "")
	if err != nil {
		return err
	}

	return nil
}
