package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConformityIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityIntegrationCreate,
		ReadContext:   resourceConformityIntegrationRead,
		UpdateContext: resourceConformityIntegrationUpdate,
		DeleteContext: resourceConformityIntegrationDelete,
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
			"api_key": {
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
			"region": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cspname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "conformity",
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceConformityIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.Conformity{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.ApiKey = d.Get("api_key").(string)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.Region = d.Get("region")
	payload.Payload.CspName = d.Get("cspname").(string)

	conformityId, err := client.CreateConformityInetgration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(conformityId)
	resourceConformityIntegrationRead(ctx, d, m)
	return diags
}
func resourceConformityIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	conformityId := d.Id()

	ConformityDetails, err := client.GetConformityInetgration(conformityId)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("user_id", ConformityDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", ConformityDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", ConformityDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", ConformityDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", ConformityDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("access_token", ConformityDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", ConformityDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", ConformityDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("index_failures", ConformityDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_unauthorized", ConformityDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_used", ConformityDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(conformityId)

	return diags
}
func resourceConformityIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if d.HasChange("alias") || d.HasChange("api_key") || d.HasChange("groups") || d.HasChange("region") || d.HasChange("cspname") {
		resourceConformityIntegrationDelete(ctx, d, m)
		resourceConformityIntegrationCreate(ctx, d, m)
	}

	return diags
}
func resourceConformityIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	conformityId := d.Id()

	_, err := client.DeleteConformityIntegration(conformityId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
