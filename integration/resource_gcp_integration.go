package integration

import (
	"autobot_integration/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGcpIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGcpIntegrationCreate,
		ReadContext:   resourceGcpIntegrationRead,
		UpdateContext: resourceGcpIntegrationUpdate,
		DeleteContext: resourceGcpIntegrationDelete,
		Schema: map[string]*schema.Schema{
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
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_provider_x_cert_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_x_cert_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cspname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "gcp",
			},
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
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceGcpIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.GCPIntegration{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.Type = d.Get("type").(string)
	payload.Payload.ProjectId = d.Get("project_id").(string)
	payload.Payload.PrivateKeyId = d.Get("private_key_id").(string)
	payload.Payload.PrivateKey = d.Get("private_key").(string)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.CspName = d.Get("cspname").(string)
	payload.Payload.ClientEmail = d.Get("client_email").(string)
	payload.Payload.ClientId = d.Get("client_id").(string)
	payload.Payload.AuthUri = d.Get("auth_uri").(string)
	payload.Payload.TokenUri = d.Get("token_uri").(string)
	payload.Payload.AuthProviderX509CertUrl = d.Get("auth_provider_x_cert_url").(string)
	payload.Payload.ClientC509CertUrl = d.Get("client_x_cert_url").(string)

	gcpIntergrationId, err := client.CreateGcpInetgration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(gcpIntergrationId)
	resourceGcpIntegrationRead(ctx, d, m)
	return diags
}

func resourceGcpIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	GcpId := d.Id()
	gcpDetails, err := client.GetGcpInetgration(GcpId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("userid", gcpDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accountid", gcpDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", gcpDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", gcpDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", gcpDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("accesstoken", gcpDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("createdat", gcpDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updatedat", gcpDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("indexfailures", gcpDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("isunauthorized", gcpDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lastused", gcpDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(GcpId)

	return diags
}
func resourceGcpIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasChange("alias") || d.HasChange("type") || d.HasChange("project_id") || d.HasChange("private_key_id") || d.HasChange("private_key") || d.HasChange("groups") || d.HasChange("cspname") || d.HasChange("client_email") || d.HasChange("client_id") || d.HasChange("auth_uri") || d.HasChange("token_uri") || d.HasChange("auth_provider_x_cert_url") || d.HasChange("client_x_cert_url") {

		resourceGcpIntegrationDelete(ctx, d, m)
		resourceGcpIntegrationCreate(ctx, d, m)

	}

	return diags
}
func resourceGcpIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	gcpId := d.Id()

	_, err := client.DeleteGcpIntegration(gcpId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
