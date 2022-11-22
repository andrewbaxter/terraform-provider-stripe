package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_ephemeral_keys() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("customer"); exists {
				params["customer"] = v
			}
			if v, exists := d.GetOk("issuing_card"); exists {
				params["issuing_card"] = v
			}
			res, err := f.Post(ctx, "/v1/ephemeral_keys", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new ephemeral_key: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/ephemeral_keys/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete ephemeral_key %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/ephemeral_keys/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up ephemeral_key %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("customer", shared.Dig[any](res, "customer"))
			d.Set("issuing_card", shared.Dig[any](res, "issuing_card"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"customer": {
				Description: "The ID of the Customer you'd like to modify using the resulting ephemeral key.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"issuing_card": {
				Description: "The ID of the Issuing Card you'd like to access using the resulting ephemeral key.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/ephemeral_keys/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update ephemeral_key %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
