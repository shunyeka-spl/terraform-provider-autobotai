package integration

import (
	"autobot_integration/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGitIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGitIntegrationCreate,
		ReadContext:   resourceGitIntegrationRead,
		UpdateContext: resourceGitIntegrationUpdate,
		DeleteContext: resourceGitIntegrationDelete,
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
			"host": {
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
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cspname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "git",
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceGitIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.Git{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.Host = d.Get("host").(string)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.Token = d.Get("token").(string)
	payload.Payload.CspName = d.Get("cspname").(string)

	gitId, err := client.CreateGitInetgration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(gitId)
	resourceGitIntegrationRead(ctx, d, m)
	return diags
}
func resourceGitIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	gitId := d.Id()
	gitDetails, err := client.GetGitInetgration(gitId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("userid", gitDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accountid", gitDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", gitDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", gitDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", gitDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accesstoken", gitDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("createdat", gitDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updatedat", gitDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("indexfailures", gitDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("isunauthorized", gitDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lastused", gitDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(gitId)

	return diags
}
func resourceGitIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if d.HasChange("alias") || d.HasChange("host") || d.HasChange("token") || d.HasChange("groups") || d.HasChange("cspname") {
		resourceGitIntegrationDelete(ctx, d, m)
		resourceGitIntegrationCreate(ctx, d, m)
	}

	return diags
}
func resourceGitIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*pkg.Client)
	gitId := d.Id()

	_, err := client.DeleteGitIntegration(gitId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
