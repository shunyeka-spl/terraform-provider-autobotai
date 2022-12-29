package integration

import (
	"autobotai_integration/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsIntegrationCreate,
		ReadContext:   resourceAwsIntegrationRead,
		UpdateContext: resourceAwsIntegrationUpdate,
		DeleteContext: resourceAwsIntegrationDelete,
		Schema: map[string]*schema.Schema{
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stack_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
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
				Optional: true,
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
			"access_token": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"alias": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"groups": {
				Type:     schema.TypeList,
				Computed: true,
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
func resourceAwsIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics

	accountId := d.Get("account_id").(string)
	stackId := d.Get("stack_id").(string)
	externalId := d.Get("external_id").(string)
	awsIntergration, err := client.GetAwsInetgration(accountId)

	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("stack_id", stackId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("external_id", externalId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_id", awsIntergration.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", awsIntergration.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", awsIntergration.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", awsIntergration.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", awsIntergration.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("access_token", awsIntergration.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", awsIntergration.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", awsIntergration.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("index_failures", awsIntergration.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_unauthorized", awsIntergration.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_used", awsIntergration.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(accountId)
	return diags
}
func resourceAwsIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}
func resourceAwsIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
func resourceAwsIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*pkg.Client)
	awsAccountId := d.Id()

	_, err := client.DeleteAwsIntegration(awsAccountId)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
