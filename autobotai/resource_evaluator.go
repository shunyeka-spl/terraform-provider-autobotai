package autobotai

import (
	"autobotAI/pkg"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEvaluator() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEvaluatorCreate,
		ReadContext:   resourceEvaluatorRead,
		UpdateContext: resourceEvaluatorUpdate,
		DeleteContext: resourceEvaluatorDelete,
		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"integration_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"preference": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"qb_rules_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"combinator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"not": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"qb_rules_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"rules_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value_source": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"combinator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"not": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"rules": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"qb_rules_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"rules_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"field": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"operator": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value_source": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceEvaluatorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)

	var diags diag.Diagnostics
	payload := pkg.Evaluator{}

	payload.Name = d.Get("name").(string)
	payload.EvalDetails.Code = d.Get("code").(string)
	payload.EvalDetails.Preference = d.Get("preference").(string)
	payload.EvalDetails.QbRules.ID = d.Get("qb_rules_id").(string)
	payload.EvalDetails.QbRules.Combinator = d.Get("combinator").(string)
	payload.EvalDetails.QbRules.Not = d.Get("not").(bool)
	if v, ok := d.GetOk("rules"); ok && len(v.(*schema.Set).List()) > 0 {
		processRules(&payload, d)
	}
	evaluatorId, err := client.CreateEvaluator(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(evaluatorId)
	resourceEvaluatorRead(ctx, d, m)
	return diags

}
func processRules(payload *pkg.Evaluator, d *schema.ResourceData) {

	rulesIn := d.Get("rules").(*schema.Set).List()
	rulesOut := make([]pkg.Rules, len(rulesIn))
	for i, rules := range rulesIn {
		m := rules.(map[string]interface{})
		obj := pkg.Rules{}
		if fmt.Sprintf("%s", m["rules_id"]) != "" {
			obj.ID = m["rules_id"].(string)
		}
		if m["field"].(string) != "" {
			obj.Field = m["field"].(string)
		}
		if m["operator"].(string) != "" {
			obj.Operator = m["operator"].(string)
		}
		if m["value"].(string) != "" {
			obj.Value = m["value"].(string)
		}
		if m["value_source"].(string) != "" {
			obj.ValueSource = m["value_source"].(string)
		}
		if m["combinator"].(string) != "" {
			obj.Combinator = m["combinator"].(string)
		}
		if m["not"].(bool) {
			obj.Not = m["not"].(bool)
		}
		obj.Rules = processInputRules(m["rules"].(*schema.Set).List())
		rulesOut[i] = obj
	}
	payload.EvalDetails.QbRules.Rules = rulesOut

}
func processInputRules(rulesIn []interface{}) []pkg.Rules {

	rulesOut := make([]pkg.Rules, len(rulesIn))
	for i, rules := range rulesIn {
		m := rules.(map[string]interface{})
		obj := pkg.Rules{}
		if m["qb_rules_id"].(string) != "" {
			obj.ID = m["qb_rules_id"].(string)
			rulesOut[i] = obj
		}
		obj.ID = fmt.Sprintf("%s", m["rules_id"])
		obj.Field = fmt.Sprintf("%s", m["field"])

		obj.Operator = fmt.Sprintf("%s", m["operator"])
		obj.Value = fmt.Sprintf("%s", m["value"])
		obj.ValueSource = fmt.Sprintf("%s", m["value_source"])
		rulesOut = append(rulesOut, obj)

	}

	return rulesOut
}
func resourceEvaluatorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	var diags diag.Diagnostics
	evaluatorId := d.Id()

	evaluatorResponse, err := client.GetEvaluator(evaluatorId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", evaluatorResponse.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("code", evaluatorResponse.EvalDetails.Code); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("preference", evaluatorResponse.EvalDetails.Preference); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("qb_rules_id", evaluatorResponse.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("combinator", evaluatorResponse.EvalDetails.QbRules.Combinator); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("not", evaluatorResponse.EvalDetails.QbRules.Not); err != nil {
		return diag.FromErr(err)

	}
	if err := d.Set("arn", evaluatorResponse.Arn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", evaluatorResponse.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("root_user_id", evaluatorResponse.RootUserID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("integration_type", evaluatorResponse.IntegrationType); err != nil {
		return diag.FromErr(err)
	}
	rules := flattenRules(evaluatorResponse)
	if err := d.Set("rules", rules); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(evaluatorId)

	return diags
}
func flattenRules(rules *pkg.EvaluatorResponse) []interface{} {
	rulesOut := make([]interface{}, len(rules.EvalDetails.QbRules.Rules))
	for i, rules := range rules.EvalDetails.QbRules.Rules {
		m := make(map[string]interface{})
		m["rules_id"] = rules.ID
		m["field"] = rules.Field
		m["operator"] = rules.Operator
		m["value_source"] = rules.ValueSource
		m["rules"] = flattenInnerRules(rules.Rules)
		rulesOut[i] = m
	}
	return rulesOut
}

func flattenInnerRules(rulesIn []pkg.Rules) []interface{} {
	rulesOut := make([]interface{}, len(rulesIn))
	for i, rules := range rulesIn {
		m := make(map[string]interface{})

		m["rules_id"] = rules.ID
		m["field"] = rules.Field
		m["operator"] = rules.Operator
		m["value_source"] = rules.ValueSource
		rulesOut[i] = m
	}
	return rulesOut
}
func resourceEvaluatorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*pkg.Client)
	evaluatorId := d.Id()
	payload := pkg.Evaluator{}

	payload.Name = d.Get("name").(string)
	payload.EvaluatorId = evaluatorId
	payload.EvalDetails.Code = d.Get("code").(string)
	payload.EvalDetails.Preference = d.Get("preference").(string)
	payload.EvalDetails.QbRules.ID = d.Get("qb_rules_id").(string)
	payload.EvalDetails.QbRules.Combinator = d.Get("combinator").(string)
	payload.EvalDetails.QbRules.Not = d.Get("not").(bool)
	if v, ok := d.GetOk("rules"); ok && len(v.(*schema.Set).List()) > 0 {
		processRules(&payload, d)
	}

	_, err := client.UpdateEvaluator(evaluatorId, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceEvaluatorRead(ctx, d, m)
}
func resourceEvaluatorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*pkg.Client)
	evaluatorId := d.Id()

	err := client.DeleteEvaluator(evaluatorId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
