package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_terminal_configurations() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("bbpos_wisepos_e"); exists {
				params["bbpos_wisepos_e"] = v
			}
			if v, exists := d.GetOk("tipping"); exists {
				params["tipping"] = v
			}
			if v, exists := d.GetOk("verifone_p400"); exists {
				params["verifone_p400"] = v
			}
			res, err := f.Post(ctx, "/v1/terminal/configurations", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new terminal.configuration: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/terminal/configurations/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete terminal.configuration %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/terminal/configurations/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up terminal.configuration %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("bbpos_wisepos_e", shared.Dig[any](res, "bbpos_wisepos_e"))
			d.Set("tipping", shared.Dig[any](res, "tipping"))
			d.Set("verifone_p400", shared.Dig[any](res, "verifone_p400"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"bbpos_wisepos_e": {
				Description: "An object containing device type specific settings for BBPOS WisePOS E readers",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"splashscreen": {
					Description: "",
					Required:    false,
					Type:        schema.TypeString,
				}}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"tipping": {
				Description: "Tipping configurations for readers supporting on-reader tips",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"aud": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"cad": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"chf": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"czk": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"dkk": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"eur": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"gbp": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"hkd": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"myr": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"nok": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"nzd": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"sek": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"sgd": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"usd": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"fixed_amounts": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"percentages": {
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeInt},
								Required:    false,
								Type:        schema.TypeList,
							},
							"smart_tip_threshold": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"verifone_p400": {
				Description: "An object containing device type specific settings for Verifone P400 readers",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"splashscreen": {
					Description: "",
					Required:    false,
					Type:        schema.TypeString,
				}}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("bbpos_wisepos_e") {
				v := d.Get("bbpos_wisepos_e")
				params["bbpos_wisepos_e"] = v
			}
			if d.HasChange("tipping") {
				v := d.Get("tipping")
				params["tipping"] = v
			}
			if d.HasChange("verifone_p400") {
				v := d.Get("verifone_p400")
				params["verifone_p400"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/terminal/configurations/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update terminal.configuration %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
