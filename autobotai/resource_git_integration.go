package autobotai

import (
	"autobotAI/pkg"
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

	if err := d.Set("user_id", gitDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", gitDetails.AccountId); err != nil {
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
	if err := d.Set("access_token", gitDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", gitDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", gitDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("index_failures", gitDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_unauthorized", gitDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_used", gitDetails.LastUsed); err != nil {
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
