package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_terminal_locations() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			{
				v := d.Get("address")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("country")
						v := shared.Dig[any](outerV, "country")
						if v != nil {
						} else {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       "Field is missing but required.",
							})
						}
					}
				}
				params["address"] = v
			}
			if v, exists := d.GetOk("configuration_overrides"); exists {
				params["configuration_overrides"] = v
			}
			{
				v := d.Get("display_name")
				params["display_name"] = v
			}
			res, err := f.Post(ctx, "/v1/terminal/locations", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new terminal.location: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/terminal/locations/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete terminal.location %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/terminal/locations/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up terminal.location %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("address", shared.Dig[any](res, "address"))
			d.Set("configuration_overrides", shared.Dig[any](res, "configuration_overrides"))
			d.Set("display_name", shared.Dig[any](res, "display_name"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"address": {
				Description: "The full address of the location.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"city": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"country": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"line1": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"line2": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"postal_code": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"state": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: false,
				Required: true,
				Type:     schema.TypeMap,
			},
			"configuration_overrides": {
				Description: "The ID of a configuration that will be used to customize all readers in this location.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"display_name": {
				Description: "A name for the location.",
				ForceNew:    false,
				Required:    true,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("address") {
				v := d.Get("address")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("country")
						v := shared.Dig[any](outerV, "country")
						if v != nil {
						} else {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       "Field is missing but required.",
							})
						}
					}
				}
				params["address"] = v
			}
			if d.HasChange("configuration_overrides") {
				v := d.Get("configuration_overrides")
				params["configuration_overrides"] = v
			}
			if d.HasChange("display_name") {
				v := d.Get("display_name")
				params["display_name"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/terminal/locations/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update terminal.location %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
