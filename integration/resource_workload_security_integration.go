package integration

import (
	"autobot_integration/pkg"
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
			"api_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secretkey": {
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
	payload.Payload.SecretKey = d.Get("secretkey").(string)
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

	if err := d.Set("userid", workloadSecurityDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accountid", workloadSecurityDetails.AccountId); err != nil {
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
	if err := d.Set("accesstoken", workloadSecurityDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("createdat", workloadSecurityDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updatedat", workloadSecurityDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("indexfailures", workloadSecurityDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("isunauthorized", workloadSecurityDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lastused", workloadSecurityDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(workloadSecurityId)

	return diags
}
func resourceWorkloadSecurityIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasChange("alias") || d.HasChange("url") || d.HasChange("secretkey") || d.HasChange("api_version") || d.HasChange("groups") || d.HasChange("cspname") {
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
