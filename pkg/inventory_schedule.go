package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) EnableInventorySchedule(InventorySchedulePayload InventorySchedule) (string, error) {

	inventorySchedule := InventorySchedule{}

	rb, err := json.Marshal(InventorySchedulePayload)
	if err != nil {
		return "", err
	}

	scheduleresponse := InventoryScheduleResponse{}
	response, err := c.ClientRequest(Post{}, "/inventory_schedules/", strings.NewReader(string(rb)), "", &inventorySchedule)
	if err != nil {
		return "", err
	}

	inventoryErr := json.Unmarshal(response, &scheduleresponse)
	if inventoryErr != nil {
		return "", inventoryErr
	}
	return scheduleresponse.Id, nil
}

func (c *Client) DisableInventorySchedule(InventoryScheduleId string) error {

	InventoryScheduleDisable := InventoryDeleteresponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/inventory_schedules/%s", InventoryScheduleId), nil, "", &InventoryScheduleDisable)
	if err != nil {
		return nil
	}

	return nil
}
