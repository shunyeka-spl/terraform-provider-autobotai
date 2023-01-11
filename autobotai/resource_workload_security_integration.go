package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkloadSecurityIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkloadSecurityIntegrationCreate,
		ReadContext:   resourceWorkloadSecurityIntegrationRead,
		UpdateContext: resourceWorkloadSecurityIntegrationUpdate,
		DeleteContext: resourceWorkloadSecurityIntegrationDelete,
		Schema: map[string]*schema.Schema{
			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"access_token": {
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
			"index_failures": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_unauthorized": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_used": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
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
				Default:  "workload_security",
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceWorkloadSecurityIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.WorkloadSecurity{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.Url = d.Get("url").(string)
	payload.Payload.ApiVersion = d.Get("api_version").(string)
	payload.Payload.SecretKey = d.Get("secret_key").(string)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.CspName = d.Get("cspname").(string)

	workloadSecurityIntergrationId, err := client.CreateWorkloadSecurityInetgration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(workloadSecurityIntergrationId)
	resourceWorkloadSecurityIntegrationRead(ctx, d, m)
	return diags
}
func resourceWorkloadSecurityIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	workloadSecurityId := d.Id()
	workloadSecurityDetails, err := client.GetWorkloadSecurityInetgration(workloadSecurityId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("user_id", workloadSecurityDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", workloadSecurityDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", workloadSecurityDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", workloadSecurityDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", workloadSecurityDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("access_token", workloadSecurityDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", workloadSecurityDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", workloadSecurityDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("index_failures", workloadSecurityDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_unauthorized", workloadSecurityDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_used", workloadSecurityDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(workloadSecurityId)

	return diags
}
func resourceWorkloadSecurityIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasChange("alias") || d.HasChange("url") || d.HasChange("secret_key") || d.HasChange("api_version") || d.HasChange("groups") || d.HasChange("cspname") {
		resourceWorkloadSecurityIntegrationDelete(ctx, d, m)
		resourceWorkloadSecurityIntegrationCreate(ctx, d, m)
	}

	return diags
}
func resourceWorkloadSecurityIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*pkg.Client)
	workloadSecurityId := d.Id()

	_, err := client.DeleteWorkloadSecurityIntegration(workloadSecurityId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
