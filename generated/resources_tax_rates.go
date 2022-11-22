package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_tax_rates() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			{
				v := d.Get("percentage")
				params["percentage"] = v
			}
			if v, exists := d.GetOk("description"); exists {
				params["description"] = v
			}
			if v, exists := d.GetOk("state"); exists {
				params["state"] = v
			}
			if v, exists := d.GetOk("tax_type"); exists {
				if inEnum(v.(string), []string{"gst", "hst", "jct", "pst", "qst", "rst", "sales_tax", "vat"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: gst, hst, jct, pst, qst, rst, sales_tax, vat"),
					})
				}
				params["tax_type"] = v
			}
			if v, exists := d.GetOk("country"); exists {
				params["country"] = v
			}
			{
				v := d.Get("display_name")
				params["display_name"] = v
			}
			{
				v := d.Get("inclusive")
				params["inclusive"] = v
			}
			if v, exists := d.GetOk("jurisdiction"); exists {
				params["jurisdiction"] = v
			}
			res, err := f.Post(ctx, "/v1/tax_rates", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new tax_rate: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/tax_rates/%v", d.Get("id")), map[string]any{"active": false})
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to set active to false on tax_rate %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/tax_rates/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up tax_rate %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("percentage", shared.Dig[any](res, "percentage"))
			d.Set("description", shared.Dig[any](res, "description"))
			d.Set("state", shared.Dig[any](res, "state"))
			d.Set("tax_type", shared.Dig[any](res, "tax_type"))
			d.Set("country", shared.Dig[any](res, "country"))
			d.Set("display_name", shared.Dig[any](res, "display_name"))
			d.Set("inclusive", shared.Dig[any](res, "inclusive"))
			d.Set("jurisdiction", shared.Dig[any](res, "jurisdiction"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"country": {
				Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"description": {
				Description: "An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"display_name": {
				Description: "The display name of the tax rate, which will be shown to users.",
				ForceNew:    false,
				Required:    true,
				Type:        schema.TypeString,
			},
			"inclusive": {
				Description: "This specifies if the tax rate is inclusive or exclusive.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeBool,
			},
			"jurisdiction": {
				Description: "The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"percentage": {
				Description: "This represents the tax rate percent out of 100.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeFloat,
			},
			"state": {
				Description: "[ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2:US), without country prefix. For example, \"NY\" for New York, United States.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"tax_type": {
				Description: "The high-level tax type, such as `vat` or `sales_tax`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("description") {
				v := d.Get("description")
				params["description"] = v
			}
			if d.HasChange("state") {
				v := d.Get("state")
				params["state"] = v
			}
			if d.HasChange("tax_type") {
				v := d.Get("tax_type")
				if inEnum(v.(string), []string{"gst", "hst", "jct", "pst", "qst", "rst", "sales_tax", "vat"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: gst, hst, jct, pst, qst, rst, sales_tax, vat"),
					})
				}
				params["tax_type"] = v
			}
			if d.HasChange("country") {
				v := d.Get("country")
				params["country"] = v
			}
			if d.HasChange("display_name") {
				v := d.Get("display_name")
				params["display_name"] = v
			}
			if d.HasChange("jurisdiction") {
				v := d.Get("jurisdiction")
				params["jurisdiction"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/tax_rates/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update tax_rate %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
