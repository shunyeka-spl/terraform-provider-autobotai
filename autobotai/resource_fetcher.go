package autobotai

import (

	// "fmt"
	"autobotAI/pkg"
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFetcher() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFetcherCreate,
		ReadContext:   resourceFetcherRead,
		UpdateContext: resourceFetcherUpdate,
		DeleteContext: resourceFetcherDelete,
		Schema: map[string]*schema.Schema{
			"integration_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_user_id": {
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
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_global": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"data_schema": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_keys": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"clients": {
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
func resourceFetcherCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.Fetcher{}

	payload.Integration_Type = d.Get("integration_type").(string)
	payload.Code = d.Get("code").(string)
	payload.Name = d.Get("name").(string)
	payload.Clients = d.Get("clients")

	fetcherID, err := client.CreateFetcher(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fetcherID)
	resourceFetcherRead(ctx, d, m)
	return diags
}
func resourceFetcherRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	fetcherId := d.Id()

	fetcherResponse, err := client.GetFetcher(fetcherId)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("integration_type", fetcherResponse.Integration_Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", fetcherResponse.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("arn", fetcherResponse.Arn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("root_user_id", fetcherResponse.RootUserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("clients", fetcherResponse.Clients); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_id", fetcherResponse.UserId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", fetcherResponse.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", fetcherResponse.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_global", fetcherResponse.IsGlobal); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("data_schema", fetcherResponse.DataSchema); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("code", fetcherResponse.Code); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", fetcherResponse.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("data_keys", fetcherResponse.DataKeys); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fetcherId)

	return diags
}
func resourceFetcherUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*pkg.Client)

	fetcherId := d.Id()

	payload := pkg.FetcherResponse{}
	payload.Code = d.Get("code").(string)
	payload.Integration_Type = d.Get("integration_type").(string)
	payload.Name = d.Get("name").(string)
	payload.Clients = d.Get("clients")
	payload.Arn = d.Get("arn").(string)
	payload.CreatedAt = d.Get("created_at").(string)
	payload.DataKeys = d.Get("data_keys").(string)
	payload.DataSchema = d.Get("data_schema").(string)
	payload.IsGlobal = d.Get("is_global").(bool)
	payload.ResourceType = d.Get("resource_type").(string)
	payload.RootUserId = d.Get("root_user_id").(string)
	payload.UpdatedAt = d.Get("updated_at").(string)
	payload.UserId = d.Get("user_id").(string)
	log.Println("[DEBUG] the payload is ", payload)
	_, err := client.UpdateFetcher(fetcherId, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceFetcherRead(ctx, d, m)
	return diags
}
func resourceFetcherDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	fetcherId := d.Id()

	err := client.DeleteFetcher(fetcherId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
