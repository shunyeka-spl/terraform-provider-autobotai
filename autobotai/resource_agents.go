package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAgent() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAgentCreate,
		ReadContext:   resourceAgentRead,
		UpdateContext: resourceAgentUpdate,
		DeleteContext: resourceAgentDelete,
		Schema: map[string]*schema.Schema{
			"integration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"integration_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAgentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.Agents{}

	payload.IntegrationId = d.Get("integration_id").(string)
	payload.IntegrationType = d.Get("integration_type").(string)
	payload.Name = d.Get("name").(string)

	agentId, err := client.CreateAgent(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(agentId)

	return diags
}
func resourceAgentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}
func resourceAgentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}

func resourceAgentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
