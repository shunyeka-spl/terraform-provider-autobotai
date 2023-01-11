package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAutomationCreate,
		ReadContext:   resourceAutomationRead,
		UpdateContext: resourceAutomationUpdate,
		DeleteContext: resourceAutomationDelete,
		Schema: map[string]*schema.Schema{
			"approval_required": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"clients": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"integration_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
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
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhook_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"required_payload": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceAutomationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	

	var diags diag.Diagnostics
	payload := pkg.Automation{}

	payload.ApprovalRequired = d.Get("approval_required").(bool)
	payload.Name = d.Get("name").(string)
	payload.Code = d.Get("code").(string)
	payload.Clients = d.Get("clients")
	payload.Integration_Type = d.Get("integration_type").(string)
	payload.Title = d.Get("title").(string)
	payload.Type = d.Get("type").(string)

	automationId, err := client.CreateAutomation(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(automationId)
	resourceAutomationRead(ctx, d, m)
	return diags
}
func resourceAutomationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	automationId := d.Id()

	automationResponse, err := client.GetAutomation(automationId)

	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", automationResponse.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", automationResponse.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("arn", automationResponse.Arn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("root_user_id", automationResponse.RootUserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_id", automationResponse.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("title", automationResponse.Title); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("code", automationResponse.Code); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", automationResponse.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("integration_type", automationResponse.Integration_Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("clients", automationResponse.Clients); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("webhook_url", automationResponse.WebhookUrl); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required_payload", automationResponse.Requiredpayload); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("approval_required", automationResponse.ApprovalRequired); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(automationId)

	return diags
}
func resourceAutomationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*pkg.Client)
	automationId := d.Id()
	if d.HasChange("approval_required") || d.HasChange("name") || d.HasChange("code") || d.HasChange("clients") || d.HasChange("integration_type") || d.HasChange("title") || d.HasChange("type") {
		payload := pkg.AutomationResponse{}
		payload.ApprovalRequired = d.Get("approval_required").(bool)
		payload.Arn = d.Get("arn").(string)
		payload.Clients = d.Get("clients")
		payload.Code = d.Get("code").(string)
		payload.Integration_Type = d.Get("integration_type").(string)
		payload.Name = d.Get("name").(string)

		payload.Requiredpayload = d.Get("required_payload")
		payload.ResourceType = d.Get("resource_type").(string)
		payload.RootUserId = d.Get("root_user_id").(string)
		payload.Title = d.Get("title").(string)
		payload.Type = d.Get("type").(string)
		payload.UserId = d.Get("user_id").(string)
		payload.WebhookUrl = d.Get("webhook_url").(string)

		_, err := client.UpdateAutomation(automationId, payload)
		if err != nil {
			return diag.FromErr(err)
		}

	}
	resourceAutomationRead(ctx, d, m)
	return diags
}
func resourceAutomationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	automationId := d.Id()

	err := client.DeleteAutomation(automationId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
