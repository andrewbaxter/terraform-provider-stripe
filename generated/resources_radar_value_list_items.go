package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_radar_value_list_items() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			{
				v := d.Get("value")
				params["value"] = v
			}
			{
				v := d.Get("value_list")
				params["value_list"] = v
			}
			res, err := f.Post(ctx, "/v1/radar/value_list_items", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new radar.value_list_item: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/radar/value_list_items/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete radar.value_list_item %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/radar/value_list_items/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up radar.value_list_item %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("value", shared.Dig[any](res, "value"))
			d.Set("value_list", shared.Dig[any](res, "value_list"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"value": {
				Description: "The value of the item (whose type must match the type of the parent value list).",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"value_list": {
				Description: "The identifier of the value list which the created item will be added to.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/radar/value_list_items/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update radar.value_list_item %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
