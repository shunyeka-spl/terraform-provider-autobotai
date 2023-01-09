package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateEvaluator(evluatorPayload Evaluator) (string, error) {

	evaluator := Evaluator{}

	rb, err := json.Marshal(evluatorPayload)
	if err != nil {
		return "", err
	}

	evaluatorResponse := EvaluatorResponse{}
	response, err := c.ClientRequest(Post{}, "/evaluators", strings.NewReader(string(rb)), "", &evaluator)
	if err != nil {
		return "", err
	}

	evluatorErr := json.Unmarshal(response, &evaluatorResponse)
	if evluatorErr != nil {
		return "", evluatorErr
	}

	return evaluatorResponse.ID, nil
}
func (c *Client) GetEvaluator(evaluatorId string) (*EvaluatorResponse, error) {

	evaluatorResponse := EvaluatorResponse{}

	response, err := c.ClientRequest(Get{}, fmt.Sprintf("/evaluators/%s", evaluatorId), nil, "", &evaluatorResponse)

	if err != nil {
		return nil, err
	}

	evaluatorErr := json.Unmarshal(response, &evaluatorResponse)
	if evaluatorErr != nil {
		return nil, evaluatorErr
	}

	return &evaluatorResponse, nil
}
func (c *Client) UpdateEvaluator(evaluatorId string, updateEvaluatorPayload Evaluator) (string, error) {

	evaluator := Evaluator{}
	rb, err := json.Marshal(updateEvaluatorPayload)
	if err != nil {
		return "", err
	}

	response, err := c.ClientRequest(Put{}, fmt.Sprintf("/evaluators/%s", evaluatorId), strings.NewReader(string(rb)), "", &evaluator)

	if err != nil {
		return "", err
	}
	evaluatorResponse := EvaluatorResponse{}
	evaluatorErr := json.Unmarshal(response, &evaluatorResponse)
	if evaluatorErr != nil {
		return "", evaluatorErr
	}

	return evaluatorResponse.ID, nil
}
func (c *Client) DeleteEvaluator(evaluatorId string) error {

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/evaluators/%s", evaluatorId), nil, "", "")
	if err != nil {
		return err
	}

	return nil
}
