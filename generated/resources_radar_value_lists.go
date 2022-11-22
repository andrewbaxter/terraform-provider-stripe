package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_radar_value_lists() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			{
				v := d.Get("name")
				params["name"] = v
			}
			{
				v := d.Get("alias")
				params["alias"] = v
			}
			if v, exists := d.GetOk("item_type"); exists {
				if inEnum(v.(string), []string{"card_bin", "card_fingerprint", "case_sensitive_string", "country", "customer_id", "email", "ip_address", "string"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: card_bin, card_fingerprint, case_sensitive_string, country, customer_id, email, ip_address, string"),
					})
				}
				params["item_type"] = v
			}
			res, err := f.Post(ctx, "/v1/radar/value_lists", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new radar.value_list: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/radar/value_lists/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete radar.value_list %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/radar/value_lists/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up radar.value_list %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("name", shared.Dig[any](res, "name"))
			d.Set("alias", shared.Dig[any](res, "alias"))
			d.Set("item_type", shared.Dig[any](res, "item_type"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"alias": {
				Description: "The name of the value list for use in rules.",
				ForceNew:    false,
				Required:    true,
				Type:        schema.TypeString,
			},
			"item_type": {
				Description: "Type of the items in the value list. One of `card_fingerprint`, `card_bin`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, or `customer_id`. Use `string` if the item type is unknown or mixed.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"name": {
				Description: "The human-readable name of the value list.",
				ForceNew:    false,
				Required:    true,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("name") {
				v := d.Get("name")
				params["name"] = v
			}
			if d.HasChange("alias") {
				v := d.Get("alias")
				params["alias"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/radar/value_lists/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update radar.value_list %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
