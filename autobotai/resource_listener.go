package autobotai

import (

	// "fmt"
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceListener() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceListenerCreate,
		ReadContext:   resourceListenerRead,
		UpdateContext: resourceListenerUpdate,
		DeleteContext: resourceListenerDelete,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"webhook_source_integration_type": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_schema": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_keys": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sample_input_json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceListenerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	// Warning or errors can be collected in a slice type

	var diags diag.Diagnostics
	payload := pkg.Listener{}

	payload.Description = d.Get("description").(string)
	payload.Name = d.Get("name").(string)
	payload.WebhookSourceIntegrationType = d.Get("webhook_source_integration_type").(string)

	listenerId, err := client.CreateListener(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(listenerId)
	resourceListenerRead(ctx, d, m)
	return diags
}
func resourceListenerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	listenerId := d.Id()

	listenerResponse, err := client.GetListener(listenerId)

	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", listenerResponse.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", listenerResponse.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", listenerResponse.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("arn", listenerResponse.Arn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("root_user_id", listenerResponse.RootUserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("webhook_source_integration_type", listenerResponse.WebhookSourceIntegrationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_id", listenerResponse.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", listenerResponse.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", listenerResponse.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret", listenerResponse.Secret); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("data_schema", listenerResponse.DataSchema); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("sample_input_json", listenerResponse.SampleInputJson); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", listenerResponse.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("data_keys", listenerResponse.DataKeys); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(listenerId)

	return diags
}
func resourceListenerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*pkg.Client)
	listenerId := d.Id()

	payload := pkg.Listener{}
	payload.Description = d.Get("description").(string)
	payload.WebhookSourceIntegrationType = d.Get("webhook_source_integration_type").(string)
	payload.Name = d.Get("name").(string)
	_, err := client.UpdateListener(listenerId, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceListenerRead(ctx, d, m)

	return diags
}
func resourceListenerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	listenerId := d.Id()

	err := client.DeleteListener(listenerId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
