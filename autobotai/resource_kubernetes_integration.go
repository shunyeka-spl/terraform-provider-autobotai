package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKubernetesIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKubernetesIntegrationCreate,
		ReadContext:   resourceKubernetesIntegrationRead,
		UpdateContext: resourceKubernetesIntegrationUpdate,
		DeleteContext: resourceKubernetesIntegrationDelete,
		Schema: map[string]*schema.Schema{
			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deploy_bots": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceKubernetesIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.KubernetesIntegration{}
	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.DeployedBots = d.Get("deploy_bots").(bool)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.CspName = "kubernetes"

	kubernetesId, err := client.CreateKubernetesIntegration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(kubernetesId)
	resourceKubernetesIntegrationRead(ctx, d, m)
	return diags
}
func resourceKubernetesIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}
func resourceKubernetesIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	resourceKubernetesIntegrationDelete(ctx, d, m)
	resourceKubernetesIntegrationCreate(ctx, d, m)
	return diags
}

func resourceKubernetesIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	kubernetesId := d.Id()

	err := client.DeleteKubernetesIntegration(kubernetesId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
