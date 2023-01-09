package autobotai

import (
	"autobotAI/pkg"
	"context"

	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGetExternalId() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceExternalIdGet,
		ReadContext:   resourceExternalIdRead,
		UpdateContext: resourceExternalIdUpdate,
		DeleteContext: resourceExternalIdDelete,
		Schema: map[string]*schema.Schema{
			"alias": {
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
			"deploy_bots": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"success": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_principal": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceExternalIdGet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.AwsIntegration{}

	payload.Alias = d.Get("alias").(string)
	payload.Groups = d.Get("groups")
	payload.DeployedBots = d.Get("deploy_bots").(bool)

	awsIntergration, err := client.GetExternalId(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("success", awsIntergration.Success); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("external_id", awsIntergration.ExternalId); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("target_principal", awsIntergration.TargetPrincipal); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
func resourceExternalIdRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}
func resourceExternalIdUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}
func resourceExternalIdDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}
