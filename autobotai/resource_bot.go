package autobotai

import (
	"autobotAI/pkg"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBotCreate,
		ReadContext:   resourceBotRead,
		UpdateContext: resourceBotUpdate,
		DeleteContext: resourceBotDelete,
		Schema: map[string]*schema.Schema{

			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cron_expression": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"subject": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"links": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"examples": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"permissions": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"listener_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"run_at": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"fleet_in_sync": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"topic": {
				Type:     schema.TypeString,
				Required: true,
			},
			"category": {
				Type:     schema.TypeString,
				Required: true,
			},

			"importance": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fetcher_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"integration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"integration_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"evaluator_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"actions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bot_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"batch_size": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"integration_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"integration": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"automation_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"params": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"values": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"required": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"action_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceBotCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.Bot{}

	payload.Name = d.Get("name").(string)
	payload.Topic = d.Get("topic").(string)
	payload.Category = d.Get("category").(string)
	payload.Importance = d.Get("importance").(string)
	payload.FetcherID = d.Get("fetcher_id").(string)
	payload.IntegrationID = d.Get("integration_id").(string)
	payload.IntegrationType = d.Get("integration_type").(string)
	payload.IntegrationID = d.Get("integration_id").(string)
	payload.IntegrationID = d.Get("integration_id").(string)
	payload.EvaluatorID = d.Get("evaluator_id").(string)

	if v, ok := d.GetOk("actions"); ok && len(v.(*schema.Set).List()) > 0 {
		processActions(&payload, d)
	}

	botId, err := client.CreateBot(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(botId)
	resourceBotRead(ctx, d, m)
	return diags
}

func processActions(payload *pkg.Bot, d *schema.ResourceData) {

	actionIn := d.Get("actions").(*schema.Set).List()
	actionOut := make([]pkg.Action, len(actionIn))
	for i, action := range actionIn {

		m := action.(map[string]interface{})
		obj := pkg.Action{}
		if fmt.Sprintf("%s", m["type"]) != "" {
			obj.Type = fmt.Sprintf("%s", m["type"])
		}
		if m["params"] != nil {
			obj.Params = processParams(m["params"].(*schema.Set).List())
		}
		if m["required"].(bool) {
			obj.ApprovalDetails.Required = m["required"].(bool)
		}
		if fmt.Sprintf("%s", m["action_id"]) != "" {
			obj.ActionID = fmt.Sprintf("%s", m["action_id"])
		}
		if fmt.Sprintf("%s", m["automation_id"]) != "" {
			obj.AutomationID = fmt.Sprintf("%s", m["automation_id"])
		}
		actionOut[i] = obj
	}
	payload.Actions = actionOut

}

func processParams(paramsIn []interface{}) []pkg.Params {

	paramsOut := make([]pkg.Params, len(paramsIn))
	for i, rules := range paramsIn {
		m := rules.(map[string]interface{})
		obj := pkg.Params{}
		if fmt.Sprintf("%s", m["description"]) != "" {
			obj.Description = fmt.Sprintf("%s", m["description"])
		}
		if fmt.Sprintf("%s", m["type"]) != "" {
			obj.Type = fmt.Sprintf("%s", m["type"])
		}
		if fmt.Sprintf("%s", m["name"]) != "" {
			obj.Name = fmt.Sprintf("%s", m["name"])
		}
		if m["values"] != nil {
			obj.Values = m["values"]
		}

		paramsOut[i] = obj

	}

	return paramsOut
}
func resourceBotRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	botId := d.Id()

	botResponse, err := client.GetBot(botId)

	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", botResponse.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", botResponse.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("arn", botResponse.Arn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("root_user_id", botResponse.RootUserID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_id", botResponse.UserID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cron_expression", botResponse.CronExpression); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("topic", botResponse.Topic); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subject", botResponse.Subject); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", botResponse.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("links", botResponse.Links); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("category", botResponse.Category); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("importance", botResponse.Importance); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("examples", botResponse.Examples); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("permissions", botResponse.Permissions); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fetcher_id", botResponse.FetcherID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("listener_id", botResponse.ListenerID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status", botResponse.Status); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fleet_in_sync", botResponse.FleetInSync); err != nil {
		return diag.FromErr(err)
	}

	if d.Get("fleet_id") != nil && d.Get("fleet_id") != "" {
		if err := d.Set("fleet_id", botResponse.FleetID); err != nil {
			return diag.FromErr(err)
		}
	}

	action := processReadAction(botResponse)

	if err := d.Set("actions", action); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("integration_id", botResponse.IntegrationID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("integration_type", botResponse.IntegrationType); err != nil {
		return diag.FromErr(err)
	}

	if d.Get("run_at") != nil && d.Get("run_at") != "" {
		if err := d.Set("run_at", botResponse.RunAt); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(botId)

	return diags
}

func processReadAction(action *pkg.BotResponse) []interface{} {

	actionOut := make([]interface{}, len(action.Action))
	for i, action := range action.Action {
		m := make(map[string]interface{})
		m["type"] = action.Type
		m["required"] = action.ApprovalDetails.Required
		m["action_id"] = action.ActionID
		m["automation_id"] = action.AutomationID
		m["params"] = processReadParams(action.Params)
		actionOut[i] = m
	}
	return actionOut

}
func processReadParams(paramsIn []pkg.Params) []interface{} {
	paramsOut := make([]interface{}, len(paramsIn))
	for i, params := range paramsIn {
		m := make(map[string]interface{})
		m["description"] = params.Description
		m["type"] = params.Type
		m["name"] = params.Name
		m["values"] = params.Values
		paramsOut[i] = m
	}
	return paramsOut
}

func resourceBotUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	botId := d.Id()
	payload := pkg.BotResponse{}
	payload.ResourceType = d.Get("resource_type").(string)
	payload.Name = d.Get("name").(string)
	payload.Arn = d.Get("arn").(string)
	payload.RootUserID = d.Get("root_user_id").(string)
	payload.UserID = d.Get("user_id").(string)
	payload.CronExpression = d.Get("cron_expression")
	payload.Topic = d.Get("topic").(string)
	payload.Subject = d.Get("subject")
	payload.Description = d.Get("description")

	payload.Links = d.Get("links")
	payload.Category = d.Get("category").(string)
	payload.Importance = d.Get("importance").(string)
	payload.Examples = d.Get("examples")
	payload.Permissions = d.Get("permissions")
	payload.FetcherID = d.Get("fetcher_id").(string)
	payload.ListenerID = d.Get("listener_id").(string)
	payload.IntegrationID = d.Get("integration_id").(string)
	payload.IntegrationType = d.Get("integration_type").(string)

	payload.Status = d.Get("status").(string)
	payload.FleetInSync = d.Get("fleet_in_sync").(bool)

	if d.Get("run_at") != nil && d.Get("run_at") != "" {
		payload.RunAt = d.Get("run_at")
	}
	if d.Get("fleet_id") != nil && d.Get("fleet_id") != "" {
		payload.FleetID = d.Get("fleet_id")
	}
	if v, ok := d.GetOk("actions"); ok && len(v.(*schema.Set).List()) > 0 {
		processUpdateActions(&payload, d)
	}

	_, err := client.UpdateBot(botId, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(botId)

	resourceBotRead(ctx, d, m)
	return diags
}

func processUpdateActions(payload *pkg.BotResponse, d *schema.ResourceData) {

	actionIn := d.Get("actions").(*schema.Set).List()
	actionOut := make([]pkg.Actions, len(actionIn))
	for i, actions := range actionIn {

		m := actions.(map[string]interface{})
		obj := pkg.Actions{}

		if fmt.Sprintf("%s", m["type"]) != "" {
			obj.Type = fmt.Sprintf("%s", m["type"])
		}

		if m["params"] != nil {
			obj.Params = processParams(m["params"].(*schema.Set).List())
		}
		if m["bot_id"] != nil {
			obj.BotID = m["bot_id"]
		}
		if m["batch_size"] != nil {
			obj.BatchSize = m["batch_size"]
		}
		if m["integration_id"] != nil {
			obj.IntegrationID = m["integration_id"]
		}
		if m["integration"] != nil {
			obj.Integration = m["integration"]
		}
		obj.ApprovalDetails.Required = m["required"].(bool)
		if fmt.Sprintf("%s", m["action_id"]) != "" {
			obj.ActionID = fmt.Sprintf("%s", m["action_id"])
		}
		if fmt.Sprintf("%s", m["automation_id"]) != "" {
			obj.AutomationID = fmt.Sprintf("%s", m["automation_id"])
		}

		actionOut[i] = obj
	}
	payload.Action = actionOut

}
func resourceBotDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	botId := d.Id()

	err := client.DeleteBot(botId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
