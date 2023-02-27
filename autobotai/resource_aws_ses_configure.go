package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsSesConfigureIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsSesConfigureIntegrationCreate,
		ReadContext:   resourceAwsSesConfigureIntegrationRead,
		UpdateContext: resourceAwsSesConfigureIntegrationUpdate,
		DeleteContext: resourceAwsSesConfigureIntegrationDelete,
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
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dependson": {
				Type:     schema.TypeString,
				Required: true,
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
				Default:  "aws_ses",
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsSesConfigureIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.AwsSesConfigure{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.Region = d.Get("region").(string)

	payload.Payload.CspName = d.Get("cspname").(string)
	payload.DependsOn = d.Get("dependson").(string)
	payload.Payload.IntegrationID = payload.DependsOn

	awsSesConfigureIntegrationId, err := client.CreateAwsSesConfigureIntegration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(awsSesConfigureIntegrationId)
	resourceAwsSesConfigureIntegrationRead(ctx, d, m)
	return diags
}

func resourceAwsSesConfigureIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	awsSesId := d.Id()

	awsSesConfigureDetails, err := client.GetAwsSesConfigureIntegration(awsSesId)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("user_id", awsSesConfigureDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", awsSesConfigureDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", awsSesConfigureDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", awsSesConfigureDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", awsSesConfigureDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("access_token", awsSesConfigureDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", awsSesConfigureDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", awsSesConfigureDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("index_failures", awsSesConfigureDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_unauthorized", awsSesConfigureDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_used", awsSesConfigureDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(awsSesId)

	return diags
}
func resourceAwsSesConfigureIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceAwsSesConfigureIntegrationDelete(ctx, d, m)
	resourceAwsSesConfigureIntegrationCreate(ctx, d, m)

	return diags
}
func resourceAwsSesConfigureIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	awsSesId := d.Id()

	_, err := client.DeleteAwsSesConfigureIntegration(awsSesId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
