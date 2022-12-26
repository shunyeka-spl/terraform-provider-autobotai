package integration

import (
	"autobot_integration/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAzureIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAzureIntegrationCreate,
		ReadContext:   resourceAzureIntegrationRead,
		UpdateContext: resourceAzureIntegrationUpdate,
		DeleteContext: resourceAzureIntegrationDelete,
		Schema: map[string]*schema.Schema{
			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"userid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"accountid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"accesstoken": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"createdat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updatedat": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"indexfailures": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"isunauthorized": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lastused": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cspname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "azure",
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAzureIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.AzureIntegrations{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.ClientId = d.Get("client_id").(string)
	payload.Payload.TenantId = d.Get("tenant_id").(string)
	payload.Payload.SecretKey = d.Get("secret_key").(string)
	payload.Payload.SubscriptionId = d.Get("subscription_id").(string)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.CspName = d.Get("cspname").(string)

	azureIntergrationId, err := client.CreateAzureInetgration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(azureIntergrationId)
	resourceAzureIntegrationRead(ctx, d, m)
	return diags
}

func resourceAzureIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	azureId := d.Id()

	azureDetails, err := client.GetAzureInetgration(azureId)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("userid", azureDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accountid", azureDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", azureDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", azureDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", azureDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accesstoken", azureDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("createdat", azureDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updatedat", azureDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("indexfailures", azureDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("isunauthorized", azureDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lastused", azureDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(azureId)

	return diags
}
func resourceAzureIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if d.HasChange("alias") || d.HasChange("client_id") || d.HasChange("tenant_id") || d.HasChange("subscription_id") || d.HasChange("groups") || d.HasChange("cspname") {
		resourceAzureIntegrationDelete(ctx, d, m)
		resourceAzureIntegrationCreate(ctx, d, m)
	}

	return diags
}
func resourceAzureIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	azureId := d.Id()

	_, err := client.DeleteAzureIntegration(azureId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
