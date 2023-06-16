package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInventorySchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInventoryScheduleEnable,
		ReadContext:   resourceInventoryScheduleRead,
		UpdateContext: resourceInventoryScheduleUpdate,
		DeleteContext: resourceInventoryScheduleDelete,
		Schema: map[string]*schema.Schema{
			"integration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cron_expression": {
				Type:     schema.TypeString,
				Required: true,
			},
			"integration_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceInventoryScheduleEnable(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.InventorySchedule{}

	payload.IntegrationId = d.Get("integration_id").(string)
	payload.IntegrationType = d.Get("integration_type").(string)
	payload.CronExpression = d.Get("cron_expression").(string)

	enventoryId, err := client.EnableInventorySchedule(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(enventoryId)

	return diags
}
func resourceInventoryScheduleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}
func resourceInventoryScheduleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}

func resourceInventoryScheduleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	inventoryId := d.Id()

	err := client.DisableInventorySchedule(inventoryId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
