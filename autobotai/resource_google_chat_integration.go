package autobotai

import (
	"autobotAI/pkg"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGoogleChatIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGoogleChatIntegrationCreate,
		ReadContext:   resourceGoogleChatIntegrationRead,
		UpdateContext: resourceGoogleChatIntegrationUpdate,
		DeleteContext: resourceGoogleChatIntegrationDelete,
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
				Default:  "google_chat",
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGoogleChatIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.GoogleChatIntegration{}

	payload.Payload.Alias = d.Get("alias").(string)
	payload.Payload.Webhook = d.Get("webhook").(string)
	payload.Payload.Groups = d.Get("groups")
	payload.Payload.CspName = d.Get("cspname").(string)

	googleChatIntergrationId, err := client.CreateGoogleChatInetgration(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(googleChatIntergrationId)
	resourceGoogleChatIntegrationRead(ctx, d, m)
	return diags
}

func resourceGoogleChatIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	googleChatId := d.Id()

	googleChatDetails, err := client.GetGoogleChatInetgration(googleChatId)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("user_id", googleChatDetails.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", googleChatDetails.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cspname", googleChatDetails.CspName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", googleChatDetails.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", googleChatDetails.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("access_token", googleChatDetails.AccessToken); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", googleChatDetails.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", googleChatDetails.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("index_failures", googleChatDetails.IndexFailures); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_unauthorized", googleChatDetails.IsUnauthorized); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_used", googleChatDetails.LastUsed); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(googleChatId)

	return diags
}
func resourceGoogleChatIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceGoogleChatIntegrationDelete(ctx, d, m)
	resourceGoogleChatIntegrationCreate(ctx, d, m)

	return diags
}
func resourceGoogleChatIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	googleChatId := d.Id()

	_, err := client.DeleteGoogleChatIntegration(googleChatId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
