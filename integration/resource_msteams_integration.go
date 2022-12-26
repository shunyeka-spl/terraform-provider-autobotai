package integration

import (
	"autobot_integration/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMSTeamsIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMSTeamsIntegrationCreate,
		ReadContext:   resourceMSTeamsIntegrationRead,
		UpdateContext: resourceMSTeamsIntegrationUpdate,
		DeleteContext: resourceMSTeamsIntegrationDelete,
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
			"webhook": {
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
				Default:  "ms_teams",
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceMSTeamsIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.MSTeamsIntegration{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.Webhook = d.Get("webhook").(string)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.CspName = d.Get("cspname").(string)

	msTeamsIntergrationId, err := client.CreateMSTeamsInetgration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(msTeamsIntergrationId)
	resourceMSTeamsIntegrationRead(ctx, d, m)
	return diags
}
func resourceMSTeamsIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	msTeamsId := d.Id()
	msTeamsDetails, err := client.GetMSTeamsInetgration(msTeamsId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("userid", msTeamsDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accountid", msTeamsDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", msTeamsDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", msTeamsDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", msTeamsDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accesstoken", msTeamsDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("createdat", msTeamsDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updatedat", msTeamsDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("indexfailures", msTeamsDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("isunauthorized", msTeamsDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lastused", msTeamsDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(msTeamsId)

	return diags
}
func resourceMSTeamsIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChange("alias") || d.HasChange("webhook") || d.HasChange("groups") || d.HasChange("cspname") {
		resourceMSTeamsIntegrationDelete(ctx, d, m)
		resourceMSTeamsIntegrationCreate(ctx, d, m)
	}
	var diags diag.Diagnostics

	return diags
}
func resourceMSTeamsIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*pkg.Client)
	msTeams := d.Id()

	_, err := client.DeleteMSTeamsIntegration(msTeams)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
