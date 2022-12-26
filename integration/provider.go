package integration

import (
	"autobot_integration/pkg"
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
				DefaultFunc: schema.EnvDefaultFunc("APIKEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"autobot_azure_integration":             resourceAzureIntegration(),
			"autobot_gcp_integration":               resourceGcpIntegration(),
			"autobot_ms_teams_integration":          resourceMSTeamsIntegration(),
			"autobot_workload_security_integration": resourceWorkloadSecurityIntegration(),
			"autobot_conformity_integration":        resourceConformityIntegration(),
			"autobot_git_integration":               resourceGitIntegration(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"autobot_external_id": dataSourceExternalId(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("apikey").(string)
	var diags diag.Diagnostics

	c, err := pkg.NewClient(apiKey)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Conformity client",
			Detail:   err.Error(),
		})

		return nil, diags
	}

	return c, diags

}
