package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_test_helpers_test_clocks() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			{
				v := d.Get("frozen_time")
				params["frozen_time"] = v
			}
			if v, exists := d.GetOk("name"); exists {
				params["name"] = v
			}
			res, err := f.Post(ctx, "/v1/test_helpers/test_clocks", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new test_helpers.test_clock: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/test_helpers/test_clocks/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete test_helpers.test_clock %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/test_helpers/test_clocks/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up test_helpers.test_clock %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("frozen_time", shared.Dig[any](res, "frozen_time"))
			d.Set("name", shared.Dig[any](res, "name"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"frozen_time": {
				Description: "The initial frozen time for this test clock.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeInt,
			},
			"name": {
				Description: "The name for this test clock.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/test_helpers/test_clocks/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update test_helpers.test_clock %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
