package integration

import (
	"autobotai_integration/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"apikey": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ApiKey", nil),
			},
			"url": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("url", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"autobotai_azure_integration":             resourceAzureIntegration(),
			"autobotai_gcp_integration":               resourceGcpIntegration(),
			"autobotai_ms_teams_integration":          resourceMSTeamsIntegration(),
			"autobotai_workload_security_integration": resourceWorkloadSecurityIntegration(),
			"autobotai_conformity_integration":        resourceConformityIntegration(),
			"autobotai_git_integration":               resourceGitIntegration(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"autobotai_external_id": dataSourceExternalId(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("apikey").(string)
	url := d.Get("url").(string)
	var diags diag.Diagnostics

	c, err := pkg.NewClient(apiKey, url)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Autobot client",
			Detail:   err.Error(),
		})

		return nil, diags
	}

	return c, diags

}
